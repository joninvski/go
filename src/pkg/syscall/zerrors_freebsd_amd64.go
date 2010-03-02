// mkerrors.sh
// MACHINE GENERATED BY THE COMMAND ABOVE; DO NOT EDIT

// godefs -gsyscall _const.c

// MACHINE GENERATED - DO NOT EDIT.

package syscall

// Constants
const (
	AF_APPLETALK              = 0x10
	AF_ARP                    = 0x23
	AF_ATM                    = 0x1e
	AF_BLUETOOTH              = 0x24
	AF_CCITT                  = 0xa
	AF_CHAOS                  = 0x5
	AF_CNT                    = 0x15
	AF_COIP                   = 0x14
	AF_DATAKIT                = 0x9
	AF_DECnet                 = 0xc
	AF_DLI                    = 0xd
	AF_E164                   = 0x1a
	AF_ECMA                   = 0x8
	AF_HYLINK                 = 0xf
	AF_IEEE80211              = 0x25
	AF_IMPLINK                = 0x3
	AF_INET                   = 0x2
	AF_INET6                  = 0x1c
	AF_IPX                    = 0x17
	AF_ISDN                   = 0x1a
	AF_ISO                    = 0x7
	AF_LAT                    = 0xe
	AF_LINK                   = 0x12
	AF_LOCAL                  = 0x1
	AF_MAX                    = 0x26
	AF_NATM                   = 0x1d
	AF_NETBIOS                = 0x6
	AF_NETGRAPH               = 0x20
	AF_OSI                    = 0x7
	AF_PUP                    = 0x4
	AF_ROUTE                  = 0x11
	AF_SCLUSTER               = 0x22
	AF_SIP                    = 0x18
	AF_SLOW                   = 0x21
	AF_SNA                    = 0xb
	AF_UNIX                   = 0x1
	AF_UNSPEC                 = 0
	AF_VENDOR00               = 0x27
	AF_VENDOR01               = 0x29
	AF_VENDOR02               = 0x2b
	AF_VENDOR03               = 0x2d
	AF_VENDOR04               = 0x2f
	AF_VENDOR05               = 0x31
	AF_VENDOR06               = 0x33
	AF_VENDOR07               = 0x35
	AF_VENDOR08               = 0x37
	AF_VENDOR09               = 0x39
	AF_VENDOR10               = 0x3b
	AF_VENDOR11               = 0x3d
	AF_VENDOR12               = 0x3f
	AF_VENDOR13               = 0x41
	AF_VENDOR14               = 0x43
	AF_VENDOR15               = 0x45
	AF_VENDOR16               = 0x47
	AF_VENDOR17               = 0x49
	AF_VENDOR18               = 0x4b
	AF_VENDOR19               = 0x4d
	AF_VENDOR20               = 0x4f
	AF_VENDOR21               = 0x51
	AF_VENDOR22               = 0x53
	AF_VENDOR23               = 0x55
	AF_VENDOR24               = 0x57
	AF_VENDOR25               = 0x59
	AF_VENDOR26               = 0x5b
	AF_VENDOR27               = 0x5d
	AF_VENDOR28               = 0x5f
	AF_VENDOR29               = 0x61
	AF_VENDOR30               = 0x63
	AF_VENDOR31               = 0x65
	AF_VENDOR32               = 0x67
	AF_VENDOR33               = 0x69
	AF_VENDOR34               = 0x6b
	AF_VENDOR35               = 0x6d
	AF_VENDOR36               = 0x6f
	AF_VENDOR37               = 0x71
	AF_VENDOR38               = 0x73
	AF_VENDOR39               = 0x75
	AF_VENDOR40               = 0x77
	AF_VENDOR41               = 0x79
	AF_VENDOR42               = 0x7b
	AF_VENDOR43               = 0x7d
	AF_VENDOR44               = 0x7f
	AF_VENDOR45               = 0x81
	AF_VENDOR46               = 0x83
	AF_VENDOR47               = 0x85
	E2BIG                     = 0x7
	EACCES                    = 0xd
	EADDRINUSE                = 0x30
	EADDRNOTAVAIL             = 0x31
	EAFNOSUPPORT              = 0x2f
	EAGAIN                    = 0x23
	EALREADY                  = 0x25
	EAUTH                     = 0x50
	EBADF                     = 0x9
	EBADMSG                   = 0x59
	EBADRPC                   = 0x48
	EBUSY                     = 0x10
	ECANCELED                 = 0x55
	ECHILD                    = 0xa
	ECONNABORTED              = 0x35
	ECONNREFUSED              = 0x3d
	ECONNRESET                = 0x36
	EDEADLK                   = 0xb
	EDESTADDRREQ              = 0x27
	EDOM                      = 0x21
	EDOOFUS                   = 0x58
	EDQUOT                    = 0x45
	EEXIST                    = 0x11
	EFAULT                    = 0xe
	EFBIG                     = 0x1b
	EFTYPE                    = 0x4f
	EHOSTDOWN                 = 0x40
	EHOSTUNREACH              = 0x41
	EIDRM                     = 0x52
	EILSEQ                    = 0x56
	EINPROGRESS               = 0x24
	EINTR                     = 0x4
	EINVAL                    = 0x16
	EIO                       = 0x5
	EISCONN                   = 0x38
	EISDIR                    = 0x15
	ELAST                     = 0x5c
	ELOOP                     = 0x3e
	EMFILE                    = 0x18
	EMLINK                    = 0x1f
	EMSGSIZE                  = 0x28
	EMULTIHOP                 = 0x5a
	ENAMETOOLONG              = 0x3f
	ENEEDAUTH                 = 0x51
	ENETDOWN                  = 0x32
	ENETRESET                 = 0x34
	ENETUNREACH               = 0x33
	ENFILE                    = 0x17
	ENOATTR                   = 0x57
	ENOBUFS                   = 0x37
	ENODEV                    = 0x13
	ENOENT                    = 0x2
	ENOEXEC                   = 0x8
	ENOLCK                    = 0x4d
	ENOLINK                   = 0x5b
	ENOMEM                    = 0xc
	ENOMSG                    = 0x53
	ENOPROTOOPT               = 0x2a
	ENOSPC                    = 0x1c
	ENOSYS                    = 0x4e
	ENOTBLK                   = 0xf
	ENOTCONN                  = 0x39
	ENOTDIR                   = 0x14
	ENOTEMPTY                 = 0x42
	ENOTSOCK                  = 0x26
	ENOTSUP                   = 0x2d
	ENOTTY                    = 0x19
	ENXIO                     = 0x6
	EOPNOTSUPP                = 0x2d
	EOVERFLOW                 = 0x54
	EPERM                     = 0x1
	EPFNOSUPPORT              = 0x2e
	EPIPE                     = 0x20
	EPROCLIM                  = 0x43
	EPROCUNAVAIL              = 0x4c
	EPROGMISMATCH             = 0x4b
	EPROGUNAVAIL              = 0x4a
	EPROTO                    = 0x5c
	EPROTONOSUPPORT           = 0x2b
	EPROTOTYPE                = 0x29
	ERANGE                    = 0x22
	EREMOTE                   = 0x47
	EROFS                     = 0x1e
	ERPCMISMATCH              = 0x49
	ESHUTDOWN                 = 0x3a
	ESOCKTNOSUPPORT           = 0x2c
	ESPIPE                    = 0x1d
	ESRCH                     = 0x3
	ESTALE                    = 0x46
	ETIMEDOUT                 = 0x3c
	ETOOMANYREFS              = 0x3b
	ETXTBSY                   = 0x1a
	EUSERS                    = 0x44
	EVFILT_AIO                = -0x3
	EVFILT_FS                 = -0x9
	EVFILT_LIO                = -0xa
	EVFILT_NETDEV             = -0x8
	EVFILT_PROC               = -0x5
	EVFILT_READ               = -0x1
	EVFILT_SIGNAL             = -0x6
	EVFILT_SYSCOUNT           = 0xa
	EVFILT_TIMER              = -0x7
	EVFILT_VNODE              = -0x4
	EVFILT_WRITE              = -0x2
	EV_ADD                    = 0x1
	EV_CLEAR                  = 0x20
	EV_DELETE                 = 0x2
	EV_DISABLE                = 0x8
	EV_ENABLE                 = 0x4
	EV_EOF                    = 0x8000
	EV_ERROR                  = 0x4000
	EV_FLAG1                  = 0x2000
	EV_ONESHOT                = 0x10
	EV_SYSFLAGS               = 0xf000
	EWOULDBLOCK               = 0x23
	EXDEV                     = 0x12
	FD_CLOEXEC                = 0x1
	FD_SETSIZE                = 0x400
	F_CANCEL                  = 0x5
	F_DUP2FD                  = 0xa
	F_DUPFD                   = 0
	F_GETFD                   = 0x1
	F_GETFL                   = 0x3
	F_GETLK                   = 0xb
	F_GETOWN                  = 0x5
	F_OGETLK                  = 0x7
	F_OSETLK                  = 0x8
	F_OSETLKW                 = 0x9
	F_RDLCK                   = 0x1
	F_SETFD                   = 0x2
	F_SETFL                   = 0x4
	F_SETLK                   = 0xc
	F_SETLKW                  = 0xd
	F_SETLK_REMOTE            = 0xe
	F_SETOWN                  = 0x6
	F_UNLCK                   = 0x2
	F_UNLCKSYS                = 0x4
	F_WRLCK                   = 0x3
	IPPROTO_3PC               = 0x22
	IPPROTO_ADFS              = 0x44
	IPPROTO_AH                = 0x33
	IPPROTO_AHIP              = 0x3d
	IPPROTO_APES              = 0x63
	IPPROTO_ARGUS             = 0xd
	IPPROTO_AX25              = 0x5d
	IPPROTO_BHA               = 0x31
	IPPROTO_BLT               = 0x1e
	IPPROTO_BRSATMON          = 0x4c
	IPPROTO_CARP              = 0x70
	IPPROTO_CFTP              = 0x3e
	IPPROTO_CHAOS             = 0x10
	IPPROTO_CMTP              = 0x26
	IPPROTO_CPHB              = 0x49
	IPPROTO_CPNX              = 0x48
	IPPROTO_DDP               = 0x25
	IPPROTO_DGP               = 0x56
	IPPROTO_DIVERT            = 0x102
	IPPROTO_DONE              = 0x101
	IPPROTO_DSTOPTS           = 0x3c
	IPPROTO_EGP               = 0x8
	IPPROTO_EMCON             = 0xe
	IPPROTO_ENCAP             = 0x62
	IPPROTO_EON               = 0x50
	IPPROTO_ESP               = 0x32
	IPPROTO_ETHERIP           = 0x61
	IPPROTO_FRAGMENT          = 0x2c
	IPPROTO_GGP               = 0x3
	IPPROTO_GMTP              = 0x64
	IPPROTO_GRE               = 0x2f
	IPPROTO_HELLO             = 0x3f
	IPPROTO_HMP               = 0x14
	IPPROTO_HOPOPTS           = 0
	IPPROTO_ICMP              = 0x1
	IPPROTO_ICMPV6            = 0x3a
	IPPROTO_IDP               = 0x16
	IPPROTO_IDPR              = 0x23
	IPPROTO_IDRP              = 0x2d
	IPPROTO_IGMP              = 0x2
	IPPROTO_IGP               = 0x55
	IPPROTO_IGRP              = 0x58
	IPPROTO_IL                = 0x28
	IPPROTO_INLSP             = 0x34
	IPPROTO_INP               = 0x20
	IPPROTO_IP                = 0
	IPPROTO_IPCOMP            = 0x6c
	IPPROTO_IPCV              = 0x47
	IPPROTO_IPEIP             = 0x5e
	IPPROTO_IPIP              = 0x4
	IPPROTO_IPPC              = 0x43
	IPPROTO_IPV4              = 0x4
	IPPROTO_IPV6              = 0x29
	IPPROTO_IRTP              = 0x1c
	IPPROTO_KRYPTOLAN         = 0x41
	IPPROTO_LARP              = 0x5b
	IPPROTO_LEAF1             = 0x19
	IPPROTO_LEAF2             = 0x1a
	IPPROTO_MAX               = 0x100
	IPPROTO_MAXID             = 0x34
	IPPROTO_MEAS              = 0x13
	IPPROTO_MHRP              = 0x30
	IPPROTO_MICP              = 0x5f
	IPPROTO_MOBILE            = 0x37
	IPPROTO_MTP               = 0x5c
	IPPROTO_MUX               = 0x12
	IPPROTO_ND                = 0x4d
	IPPROTO_NHRP              = 0x36
	IPPROTO_NONE              = 0x3b
	IPPROTO_NSP               = 0x1f
	IPPROTO_NVPII             = 0xb
	IPPROTO_OLD_DIVERT        = 0xfe
	IPPROTO_OSPFIGP           = 0x59
	IPPROTO_PFSYNC            = 0xf0
	IPPROTO_PGM               = 0x71
	IPPROTO_PIGP              = 0x9
	IPPROTO_PIM               = 0x67
	IPPROTO_PRM               = 0x15
	IPPROTO_PUP               = 0xc
	IPPROTO_PVP               = 0x4b
	IPPROTO_RAW               = 0xff
	IPPROTO_RCCMON            = 0xa
	IPPROTO_RDP               = 0x1b
	IPPROTO_ROUTING           = 0x2b
	IPPROTO_RSVP              = 0x2e
	IPPROTO_RVD               = 0x42
	IPPROTO_SATEXPAK          = 0x40
	IPPROTO_SATMON            = 0x45
	IPPROTO_SCCSP             = 0x60
	IPPROTO_SCTP              = 0x84
	IPPROTO_SDRP              = 0x2a
	IPPROTO_SEP               = 0x21
	IPPROTO_SKIP              = 0x39
	IPPROTO_SPACER            = 0x7fff
	IPPROTO_SRPC              = 0x5a
	IPPROTO_ST                = 0x7
	IPPROTO_SVMTP             = 0x52
	IPPROTO_SWIPE             = 0x35
	IPPROTO_TCF               = 0x57
	IPPROTO_TCP               = 0x6
	IPPROTO_TLSP              = 0x38
	IPPROTO_TP                = 0x1d
	IPPROTO_TPXX              = 0x27
	IPPROTO_TRUNK1            = 0x17
	IPPROTO_TRUNK2            = 0x18
	IPPROTO_TTP               = 0x54
	IPPROTO_UDP               = 0x11
	IPPROTO_VINES             = 0x53
	IPPROTO_VISA              = 0x46
	IPPROTO_VMTP              = 0x51
	IPPROTO_WBEXPAK           = 0x4f
	IPPROTO_WBMON             = 0x4e
	IPPROTO_WSN               = 0x4a
	IPPROTO_XNET              = 0xf
	IPPROTO_XTP               = 0x24
	IP_ADD_MEMBERSHIP         = 0xc
	IP_ADD_SOURCE_MEMBERSHIP  = 0x46
	IP_BINDANY                = 0x18
	IP_BLOCK_SOURCE           = 0x48
	IP_DEFAULT_MULTICAST_LOOP = 0x1
	IP_DEFAULT_MULTICAST_TTL  = 0x1
	IP_DONTFRAG               = 0x43
	IP_DROP_MEMBERSHIP        = 0xd
	IP_DROP_SOURCE_MEMBERSHIP = 0x47
	IP_DUMMYNET_CONFIGURE     = 0x3c
	IP_DUMMYNET_DEL           = 0x3d
	IP_DUMMYNET_FLUSH         = 0x3e
	IP_DUMMYNET_GET           = 0x40
	IP_FAITH                  = 0x16
	IP_FW_ADD                 = 0x32
	IP_FW_DEL                 = 0x33
	IP_FW_FLUSH               = 0x34
	IP_FW_GET                 = 0x36
	IP_FW_NAT_CFG             = 0x38
	IP_FW_NAT_DEL             = 0x39
	IP_FW_NAT_GET_CONFIG      = 0x3a
	IP_FW_NAT_GET_LOG         = 0x3b
	IP_FW_RESETLOG            = 0x37
	IP_FW_TABLE_ADD           = 0x28
	IP_FW_TABLE_DEL           = 0x29
	IP_FW_TABLE_FLUSH         = 0x2a
	IP_FW_TABLE_GETSIZE       = 0x2b
	IP_FW_TABLE_LIST          = 0x2c
	IP_FW_ZERO                = 0x35
	IP_HDRINCL                = 0x2
	IP_IPSEC_POLICY           = 0x15
	IP_MAX_GROUP_SRC_FILTER   = 0x200
	IP_MAX_MEMBERSHIPS        = 0xfff
	IP_MAX_SOCK_MUTE_FILTER   = 0x80
	IP_MAX_SOCK_SRC_FILTER    = 0x80
	IP_MAX_SOURCE_FILTER      = 0x400
	IP_MINTTL                 = 0x42
	IP_MIN_MEMBERSHIPS        = 0x1f
	IP_MSFILTER               = 0x4a
	IP_MULTICAST_IF           = 0x9
	IP_MULTICAST_LOOP         = 0xb
	IP_MULTICAST_TTL          = 0xa
	IP_MULTICAST_VIF          = 0xe
	IP_ONESBCAST              = 0x17
	IP_OPTIONS                = 0x1
	IP_PORTRANGE              = 0x13
	IP_PORTRANGE_DEFAULT      = 0
	IP_PORTRANGE_HIGH         = 0x1
	IP_PORTRANGE_LOW          = 0x2
	IP_RECVDSTADDR            = 0x7
	IP_RECVIF                 = 0x14
	IP_RECVOPTS               = 0x5
	IP_RECVRETOPTS            = 0x6
	IP_RECVTTL                = 0x41
	IP_RETOPTS                = 0x8
	IP_RSVP_OFF               = 0x10
	IP_RSVP_ON                = 0xf
	IP_RSVP_VIF_OFF           = 0x12
	IP_RSVP_VIF_ON            = 0x11
	IP_SENDSRCADDR            = 0x7
	IP_TOS                    = 0x3
	IP_TTL                    = 0x4
	IP_UNBLOCK_SOURCE         = 0x49
	O_ACCMODE                 = 0x3
	O_APPEND                  = 0x8
	O_ASYNC                   = 0x40
	O_CREAT                   = 0x200
	O_DIRECT                  = 0x10000
	O_DIRECTORY               = 0x20000
	O_EXCL                    = 0x800
	O_EXEC                    = 0x40000
	O_EXLOCK                  = 0x20
	O_FSYNC                   = 0x80
	O_NDELAY                  = 0x4
	O_NOCTTY                  = 0x8000
	O_NOFOLLOW                = 0x100
	O_NONBLOCK                = 0x4
	O_RDONLY                  = 0
	O_RDWR                    = 0x2
	O_SHLOCK                  = 0x10
	O_SYNC                    = 0x80
	O_TRUNC                   = 0x400
	O_TTY_INIT                = 0x80000
	O_WRONLY                  = 0x1
	SHUT_RD                   = 0
	SHUT_RDWR                 = 0x2
	SHUT_WR                   = 0x1
	SIGABRT                   = 0x6
	SIGALRM                   = 0xe
	SIGBUS                    = 0xa
	SIGCHLD                   = 0x14
	SIGCONT                   = 0x13
	SIGEMT                    = 0x7
	SIGFPE                    = 0x8
	SIGHUP                    = 0x1
	SIGILL                    = 0x4
	SIGINFO                   = 0x1d
	SIGINT                    = 0x2
	SIGIO                     = 0x17
	SIGIOT                    = 0x6
	SIGKILL                   = 0x9
	SIGLWP                    = 0x20
	SIGPIPE                   = 0xd
	SIGPROF                   = 0x1b
	SIGQUIT                   = 0x3
	SIGSEGV                   = 0xb
	SIGSTOP                   = 0x11
	SIGSYS                    = 0xc
	SIGTERM                   = 0xf
	SIGTHR                    = 0x20
	SIGTRAP                   = 0x5
	SIGTSTP                   = 0x12
	SIGTTIN                   = 0x15
	SIGTTOU                   = 0x16
	SIGURG                    = 0x10
	SIGUSR1                   = 0x1e
	SIGUSR2                   = 0x1f
	SIGVTALRM                 = 0x1a
	SIGWINCH                  = 0x1c
	SIGXCPU                   = 0x18
	SIGXFSZ                   = 0x19
	SOCK_DGRAM                = 0x2
	SOCK_MAXADDRLEN           = 0xff
	SOCK_RAW                  = 0x3
	SOCK_RDM                  = 0x4
	SOCK_SEQPACKET            = 0x5
	SOCK_STREAM               = 0x1
	SOL_SOCKET                = 0xffff
	SOMAXCONN                 = 0x80
	SO_ACCEPTCONN             = 0x2
	SO_ACCEPTFILTER           = 0x1000
	SO_BINTIME                = 0x2000
	SO_BROADCAST              = 0x20
	SO_DEBUG                  = 0x1
	SO_DONTROUTE              = 0x10
	SO_ERROR                  = 0x1007
	SO_KEEPALIVE              = 0x8
	SO_LABEL                  = 0x1009
	SO_LINGER                 = 0x80
	SO_LISTENINCQLEN          = 0x1013
	SO_LISTENQLEN             = 0x1012
	SO_LISTENQLIMIT           = 0x1011
	SO_NOSIGPIPE              = 0x800
	SO_NO_DDP                 = 0x8000
	SO_NO_OFFLOAD             = 0x4000
	SO_OOBINLINE              = 0x100
	SO_PEERLABEL              = 0x1010
	SO_RCVBUF                 = 0x1002
	SO_RCVLOWAT               = 0x1004
	SO_RCVTIMEO               = 0x1006
	SO_REUSEADDR              = 0x4
	SO_REUSEPORT              = 0x200
	SO_SETFIB                 = 0x1014
	SO_SNDBUF                 = 0x1001
	SO_SNDLOWAT               = 0x1003
	SO_SNDTIMEO               = 0x1005
	SO_TIMESTAMP              = 0x400
	SO_TYPE                   = 0x1008
	SO_USELOOPBACK            = 0x40
	TCP_CA_NAME_MAX           = 0x10
	TCP_CONGESTION            = 0x40
	TCP_INFO                  = 0x20
	TCP_MAXBURST              = 0x4
	TCP_MAXHLEN               = 0x3c
	TCP_MAXOLEN               = 0x28
	TCP_MAXSEG                = 0x2
	TCP_MAXWIN                = 0xffff
	TCP_MAX_SACK              = 0x4
	TCP_MAX_WINSHIFT          = 0xe
	TCP_MD5SIG                = 0x10
	TCP_MINMSS                = 0xd8
	TCP_MSS                   = 0x200
	TCP_NODELAY               = 0x1
	TCP_NOOPT                 = 0x8
	TCP_NOPUSH                = 0x4
	WCONTINUED                = 0x4
	WCOREFLAG                 = 0x80
	WLINUXCLONE               = 0x80000000
	WNOHANG                   = 0x1
	WNOWAIT                   = 0x8
	WSTOPPED                  = 0x2
	WUNTRACED                 = 0x2
)

// Types


// Error table
var errors = [...]string{
	7:  "argument list too long",
	13: "permission denied",
	48: "address already in use",
	49: "can't assign requested address",
	47: "address family not supported by protocol family",
	35: "resource temporarily unavailable",
	37: "operation already in progress",
	80: "authentication error",
	9:  "bad file descriptor",
	89: "bad message",
	72: "RPC struct is bad",
	16: "device busy",
	85: "operation canceled",
	10: "no child processes",
	53: "software caused connection abort",
	61: "connection refused",
	54: "connection reset by peer",
	11: "resource deadlock avoided",
	39: "destination address required",
	33: "numerical argument out of domain",
	88: "programming error",
	69: "disc quota exceeded",
	17: "file exists",
	14: "bad address",
	27: "file too large",
	79: "inappropriate file type or format",
	64: "host is down",
	65: "no route to host",
	82: "identifier removed",
	86: "illegal byte sequence",
	36: "operation now in progress",
	4:  "interrupted system call",
	22: "invalid argument",
	5:  "input/output error",
	56: "socket is already connected",
	21: "is a directory",
	92: "protocol error",
	62: "too many levels of symbolic links",
	24: "too many open files",
	31: "too many links",
	40: "message too long",
	90: "multihop attempted",
	63: "file name too long",
	81: "need authenticator",
	50: "network is down",
	52: "network dropped connection on reset",
	51: "network is unreachable",
	23: "too many open files in system",
	87: "attribute not found",
	55: "no buffer space available",
	19: "operation not supported by device",
	2:  "no such file or directory",
	8:  "exec format error",
	77: "no locks available",
	91: "link has been severed",
	12: "cannot allocate memory",
	83: "no message of desired type",
	42: "protocol not available",
	28: "no space left on device",
	78: "function not implemented",
	15: "block device required",
	57: "socket is not connected",
	20: "not a directory",
	66: "directory not empty",
	38: "socket operation on non-socket",
	45: "operation not supported",
	25: "inappropriate ioctl for device",
	6:  "device not configured",
	84: "value too large to be stored in data type",
	1:  "operation not permitted",
	46: "protocol family not supported",
	32: "broken pipe",
	67: "too many processes",
	76: "bad procedure for program",
	75: "program version wrong",
	74: "RPC prog. not avail",
	43: "protocol not supported",
	41: "protocol wrong type for socket",
	34: "result too large",
	71: "too many levels of remote in path",
	30: "read-only file system",
	73: "RPC version wrong",
	58: "can't send after socket shutdown",
	44: "socket type not supported",
	29: "illegal seek",
	3:  "no such process",
	70: "stale NFS file handle",
	60: "operation timed out",
	59: "too many references: can't splice",
	26: "text file busy",
	68: "too many users",
	18: "cross-device link",
}
