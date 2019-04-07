#!/bin/bash

./rpcserver install -r zk://192.168.2.242 -c t  << EOF
2
EOF

./rpcserver run -r zk://192.168.2.242 -c t