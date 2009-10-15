// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package rand

/*
 * Uniform distribution
 *
 * algorithm by
 * DP Mitchell and JA Reeds
 */

const (
	_LEN	= 607;
	_TAP	= 273;
	_MAX	= 1<<63;
	_MASK	= _MAX-1;
	_A	= 48271;
	_M	= (1<<31)-1;
	_Q	= 44488;
	_R	= 3399;
)

var (
	// cooked random numbers
	// the state of the rng
	// after 780e10 iterations
	rng_cooked [_LEN]int64 = [...]int64{
		5041579894721019882, 4646389086726545243, 1395769623340756751, 5333664234075297259,
		2875692520355975054, 9033628115061424579, 7143218595135194537, 4812947590706362721,
		7937252194349799378, 5307299880338848416, 8209348851763925077, 2115741599318814044,
		4593015457530856296, 8140875735541888011, 3319429241265089026, 8619815648190321034,
		1727074043483619500, 113108499721038619, 4569519971459345583, 5062833859075314731,
		2387618771259064424, 2716131344356686112, 6559392774825876886, 7650093201692370310,
		7684323884043752161, 257867835996031390, 6593456519409015164, 271327514973697897,
		2789386447340118284, 1065192797246149621, 3344507881999356393, 4459797941780066633,
		7465081662728599889, 1014950805555097187, 4449440729345990775, 3481109366438502643,
		2418672789110888383, 5796562887576294778, 4484266064449540171, 3738982361971787048,
		4523597184512354423, 10530508058128498, 8633833783282346118, 2625309929628791628,
		8660405965245884302, 10162832508971942, 6540714680961817391, 7031802312784620857,
		6240911277345944669, 831864355460801054, 8004434137542152891, 2116287251661052151,
		2202309800992166967, 9161020366945053561, 4069299552407763864, 4936383537992622449,
		457351505131524928, 342195045928179354, 2847771682816600509, 2068020115986376518,
		4368649989588021065, 887231587095185257, 5563591506886576496, 6816225200251950296,
		5616972787034086048, 8471809303394836566, 1686575021641186857, 4045484338074262002,
		4244156215201778923, 7848217333783577387, 5632136521049761902, 833283142057835272,
		9029726508369077193, 3243583134664087292, 4316371101804477087, 8937849979965997980,
		6446940406810434101, 1679342092332374735, 6050638460742422078, 6993520719509581582,
		7640877852514293609, 5881353426285907985, 812786550756860885, 4541845584483343330,
		2725470216277009086, 4980675660146853729, 5210769080603236061, 8894283318990530821,
		6326442804750084282, 1495812843684243920, 7069751578799128019, 7370257291860230865,
		6756929275356942261, 4706794511633873654, 7824520467827898663, 8549875090542453214,
		33650829478596156, 1328918435751322643, 7297902601803624459, 1011190183918857495,
		2238025036817854944, 5147159997473910359, 896512091560522982, 2659470849286379941,
		6097729358393448602, 1731725986304753684, 4106255841983812711, 8327155210721535508,
		8477511620686074402, 5803876044675762232, 8435417780860221662, 5988852856651071244,
		4715837297103951910, 7566171971264485114, 505808562678895611, 5070098180695063370,
		842110666775871513, 572156825025677802, 1791881013492340891, 3393267094866038768,
		3778721850472236509, 2352769483186201278, 1292459583847367458, 8897907043675088419,
		5781809037144163536, 2733958794029492513, 5092019688680754699, 8996124554772526841,
		4234737173186232084, 5027558287275472836, 4635198586344772304, 8687338893267139351,
		5907508150730407386, 784756255473944452, 972392927514829904, 5422057694808175112,
		5158420642969283891, 9048531678558643225, 2407211146698877100, 7583282216521099569,
		3940796514530962282, 3341174631045206375, 3095313889586102949, 7405321895688238710,
		5832080132947175283, 7890064875145919662, 8184139210799583195, 1149859861409226130,
		1464597243840211302, 4641648007187991873, 3516491885471466898, 956288521791657692,
		6657089965014657519, 5220884358887979358, 1796677326474620641, 5340761970648932916,
		1147977171614181568, 5066037465548252321, 2574765911837859848, 1085848279845204775,
		3350107529868390359, 6116438694366558490, 2107701075971293812, 1803294065921269267,
		2469478054175558874, 7368243281019965984, 3791908367843677526, 185046971116456637,
		2257095756513439648, 7217693971077460129, 909049953079504259, 7196649268545224266,
		5637660345400869599, 3955544945427965183, 8057528650917418961, 4139268440301127643,
		6621926588513568059, 1373361136802681441, 6527366231383600011, 3507654575162700890,
		9202058512774729859, 1954818376891585542, 6640380907130175705, 8299563319178235687,
		3901867355218954373, 7046310742295574065, 6847195391333990232, 1572638100518868053,
		8850422670118399721, 3631909142291992901, 5158881091950831288, 2882958317343121593,
		4763258931815816403, 6280052734341785344, 4243789408204964850, 2043464728020827976,
		6545300466022085465, 4562580375758598164, 5495451168795427352, 1738312861590151095,
		553004618757816492, 6895160632757959823, 8233623922264685171, 7139506338801360852,
		8550891222387991669, 5535668688139305547, 2430933853350256242, 5401941257863201076,
		8159640039107728799, 6157493831600770366, 7632066283658143750, 6308328381617103346,
		3681878764086140361, 3289686137190109749, 6587997200611086848, 244714774258135476,
		4079788377417136100, 8090302575944624335, 2945117363431356361, 864324395848741045,
		3009039260312620700, 8430027460082534031, 401084700045993341, 7254622446438694921,
		4707864159563588614, 5640248530963493951, 5982507712689997893, 3315098242282210105,
		5503847578771918426, 3941971367175193882, 8118566580304798074, 3839261274019871296,
		7062410411742090847, 741381002980207668, 6027994129690250817, 2497829994150063930,
		6251390334426228834, 1368930247903518833, 8809096399316380241, 6492004350391900708,
		2462145737463489636, 404828418920299174, 4153026434231690595, 261785715255475940,
		5464715384600071357, 592710404378763017, 6764129236657751224, 8513655718539357449,
		5820343663801914208, 385298524683789911, 5224135003438199467, 6303131641338802145,
		7150122561309371392, 368107899140673753, 3115186834558311558, 2915636353584281051,
		4782583894627718279, 6718292300699989587, 8387085186914375220, 3387513132024756289,
		4654329375432538231, 8930667561363381602, 5374373436876319273, 7623042350483453954,
		7725442901813263321, 9186225467561587250, 4091027289597503355, 2357631606492579800,
		2530936820058611833, 1636551876240043639, 5564664674334965799, 1452244145334316253,
		2061642381019690829, 1279580266495294036, 9108481583171221009, 6023278686734049809,
		5007630032676973346, 2153168792952589781, 6720334534964750538, 6041546491134794105,
		3433922409283786309, 2285479922797300912, 3110614940896576130, 6366559590722842893,
		5418791419666136509, 7163298419643543757, 4891138053923696990, 580618510277907015,
		1684034065251686769, 4429514767357295841, 330346578555450005, 1119637995812174675,
		7177515271653460134, 4589042248470800257, 7693288629059004563, 143607045258444228,
		246994305896273627, 866417324803099287, 6473547110565816071, 3092379936208876896,
		2058427839513754051, 5133784708526867938, 8785882556301281247, 6149332666841167611,
		8585842181454472135, 6137678347805511274, 2070447184436970006, 5708223427705576541,
		5999657892458244504, 4358391411789012426, 325123008708389849, 6837621693887290924,
		4843721905315627004, 6010651222149276415, 5398352198963874652, 4602025990114250980,
		1044646352569048800, 9106614159853161675, 829256115228593269, 4919284369102997000,
		2681532557646850893, 3681559472488511871, 5307999518958214035, 6334130388442829274,
		2658708232916537604, 1163313865052186287, 581945337509520675, 3648778920718647903,
		4423673246306544414, 1620799783996955743, 220828013409515943, 8150384699999389761,
		4287360518296753003, 4590000184845883843, 5513660857261085186, 6964829100392774275,
		478991688350776035, 8746140185685648781, 228500091334420247, 1356187007457302238,
		3019253992034194581, 3152601605678500003, 430152752706002213, 5559581553696971176,
		4916432985369275664, 663574931734554391, 3420773838927732076, 2868348622579915573,
		1999319134044418520, 3328689518636282723, 2587672709781371173, 1517255313529399333,
		3092343956317362483, 3662252519007064108, 972445599196498113, 7664865435875959367,
		1708913533482282562, 6917817162668868494, 3217629022545312900, 2570043027221707107,
		8739788839543624613, 2488075924621352812, 4694002395387436668, 4559628481798514356,
		2997203966153298104, 1282559373026354493, 240113143146674385, 8665713329246516443,
		628141331766346752, 4571950817186770476, 1472811188152235408, 7596648026010355826,
		6091219417754424743, 7834161864828164065, 7103445518877254909, 4390861237357459201,
		4442653864240571734, 8903482404847331368, 622261699494173647, 6037261250297213248,
		504404948065709118, 7275215526217113061, 1011176780856001400, 2194750105623461063,
		2623071828615234808, 5157313728073836108, 3738405111966602044, 2539767524076729570,
		2467284396349269342, 5256026990536851868, 7841086888628396109, 6640857538655893162,
		1202087339038317498, 2113514992440715978, 7534350895342931403, 4925284734898484745,
		5145623771477493805, 8225140880134972332, 2719520354384050532, 9132346697815513771,
		4332154495710163773, 7137789594094346916, 6994721091344268833, 6667228574869048934,
		655440045726677499, 59934747298466858, 6124974028078036405, 8957774780655365418,
		2332206071942466437, 1701056712286369627, 3154897383618636503, 1637766181387607527,
		2460521277767576533, 197309393502684135, 643677854385267315, 2543179307861934850,
		4350769010207485119, 4754652089410667672, 2015595502641514512, 7999059458976458608,
		4287946071480840813, 8362686366770308971, 6486469209321732151, 3617727845841796026,
		7554353525834302244, 4450022655153542367, 1605195740213535749, 5327014565305508387,
		4626575813550328320, 2692222020597705149, 241045573717249868, 5098046974627094010,
		7916882295460730264, 884817090297530579, 5329160409530630596, 7790979528857726136,
		4955070238059373407, 4918537275422674302, 3008076183950404629, 3007769226071157901,
		2470346235617803020, 8928702772696731736, 7856187920214445904, 4474874585391974885,
		7900176660600710914, 2140571127916226672, 2425445057265199971, 2486055153341847830,
		4186670094382025798, 1883939007446035042, 8808666044074867985, 3734134241178479257,
		4065968871360089196, 6953124200385847784, 1305686814738899057, 1637739099014457647,
		3656125660947993209, 3966759634633167020, 3106378204088556331, 6328899822778449810,
		4565385105440252958, 1979884289539493806, 2331793186920865425, 3783206694208922581,
		8464961209802336085, 2843963751609577687, 3030678195484896323, 4793717574095772604,
		4459239494808162889, 402587895800087237, 8057891408711167515, 4541888170938985079,
		1042662272908816815, 5557303057122568958, 2647678726283249984, 2144477441549833761,
		5806352215355387087, 7117771003473903623, 5916597177708541638, 462597715452321361,
		8833658097025758785, 5970273481425315300, 563813119381731307, 2768349550652697015,
		1598828206250873866, 5206393647403558110, 6235043485709261823, 3152217402014639496,
		8469693267274066490, 125672920241807416, 5311079624024060938, 6663754932310491587,
		8736848295048751716, 4488039774992061878, 5923302823487327109, 140891791083103236,
		7414942793393574290, 7990420780896957397, 4317817392807076702, 3625184369705367340,
		2740722765288122703, 5743100009702758344, 5997898640509039159, 8854493341352484163,
		5242208035432907801, 701338899890987198, 7609280429197514109, 3020985755112334161,
		6651322707055512866, 2635195723621160615, 5144520864246028816, 1035086515727829828,
		1567242097116389047, 8172389260191636581, 6337820351429292273, 2163012566996458925,
		2743190902890262681, 1906367633221323427, 6011544915663598137, 5932255307352610768,
		2241128460406315459, 895504896216695588, 3094483003111372717, 4583857460292963101,
		9079887171656594975, 8839289181930711403, 5762740387243057873, 4225072055348026230,
		1838220598389033063, 3801620336801580414, 8823526620080073856, 1776617605585100335,
		7899055018877642622, 5421679761463003041, 5521102963086275121, 4248279443559365898,
		8735487530905098534, 1760527091573692978, 7142485049657745894, 8222656872927218123,
		4969531564923704323, 3394475942196872480, 6424174453260338141, 359248545074932887,
		3273651282831730598, 6797106199797138596, 3030918217665093212, 145600834617314036,
		6036575856065626233, 740416251634527158, 7080427635449935582, 6951781370868335478,
		399922722363687927, 294902314447253185, 7844950936339178523, 880320858634709042,
		6192655680808675579, 411604686384710388, 9026808440365124461, 6440783557497587732,
		4615674634722404292, 539897290441580544, 2096238225866883852, 8751955639408182687,
		1907224908052289603, 7381039757301768559, 6157238513393239656, 7749994231914157575,
		8629571604380892756, 5280433031239081479, 7101611890139813254, 2479018537985767835,
		7169176924412769570, 7942066497793203302, 1357759729055557688, 2278447439451174845,
		3625338785743880657, 6477479539006708521, 8976185375579272206, 5511371554711836120,
		1326024180520890843, 7537449876596048829, 5464680203499696154, 3189671183162196045,
		6346751753565857109, 241159987320630307, 3095793449658682053, 8978332846736310159,
		2902794662273147216, 7208698530190629697, 7276901792339343736, 1732385229314443140,
		4133292154170828382, 2918308698224194548, 1519461397937144458, 5293934712616591764,
		4922828954023452664, 2879211533496425641, 5896236396443472108, 8465043815351752425,
		7329020396871624740, 8915471717014488588, 2944902635677463047, 7052079073493465134,
		8382142935188824023, 9103922860780351547, 4152330101494654406,
	};
)

type rngSource struct {
	tap	int;		// index into vec
	feed	int;		// index into vec
	vec	[_LEN]int64;	// current feedback register
}

// seed rng x[n+1] = 48271 * x[n] mod (2**31 - 1)
func seedrand(x int32) int32 {
	hi := x/_Q;
	lo := x%_Q;
	x = _A*lo - _R*hi;
	if x < 0 {
		x += _M;
	}
	return x;
}

// Seed uses the provided seed value to initialize the generator to a deterministic state.
func (rng *rngSource) Seed(seed int64) {
	rng.tap = 0;
	rng.feed = _LEN-_TAP;

	seed = seed%_M;
	if seed < 0 {
		seed += _M;
	}
	if seed == 0 {
		seed = 89482311;
	}

	x := int32(seed);
	for i := -20; i < _LEN; i++ {
		x = seedrand(x);
		if i >= 0 {
			var u int64;
			u = int64(x)<<40;
			x = seedrand(x);
			u ^= int64(x)<<20;
			x = seedrand(x);
			u ^= int64(x);
			u ^= rng_cooked[i];
			rng.vec[i] = u&_MASK;
		}
	}
}

// Int63 returns a non-negative pseudo-random 63-bit integer as an int64.
func (rng *rngSource) Int63() int64 {
	rng.tap--;
	if rng.tap < 0 {
		rng.tap += _LEN;
	}

	rng.feed--;
	if rng.feed < 0 {
		rng.feed += _LEN;
	}

	x := (rng.vec[rng.feed] + rng.vec[rng.tap])&_MASK;
	rng.vec[rng.feed] = x;
	return x;
}
