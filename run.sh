#!/bin/bash

./rpcserver stop

./rpcserver install -r zk://192.168.2.242 -c t  << EOF
2
EOF

./rpcserver start