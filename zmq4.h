#if ZMQ_VERSION_MAJOR != 4
#error "You need ZeroMQ version 4 to build this"
#endif
#if ZMQ_VERSION_MINOR < 1
#define ZMQ_CONNECT_RID -1
#define ZMQ_GSSAPI -1
#define ZMQ_GSSAPI_PLAINTEXT -1
#define ZMQ_GSSAPI_PRINCIPAL -1
#define ZMQ_GSSAPI_SERVER -1
#define ZMQ_GSSAPI_SERVICE_PRINCIPAL -1
#define ZMQ_HANDSHAKE_IVL -1
#define ZMQ_IPC_FILTER_GID -1
#define ZMQ_IPC_FILTER_PID -1
#define ZMQ_IPC_FILTER_UID -1
#define ZMQ_ROUTER_HANDOVER -1
#define ZMQ_SOCKS_PROXY -1
#define ZMQ_THREAD_PRIORITY -1
#define ZMQ_THREAD_SCHED_POLICY -1
#define ZMQ_TOS -1
#define ZMQ_XPUB_NODROP -1
#endif
// TODO 4.2
#if ZMQ_VERSION_MINOR < 2
#define ZMQ_SERVER -1
#define ZMQ_CLIENT -2
#define ZMQ_RADIO -3
#define ZMQ_DISH -4
#define ZMQ_GATHER -5
#define ZMQ_SCATTER -6
#define ZMQ_DGRAM -7
#endif
