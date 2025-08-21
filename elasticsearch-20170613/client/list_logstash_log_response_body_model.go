// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListLogstashLogResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ListLogstashLogResponseBody
	GetRequestId() *string
	SetResult(v []*ListLogstashLogResponseBodyResult) *ListLogstashLogResponseBody
	GetResult() []*ListLogstashLogResponseBodyResult
}

type ListLogstashLogResponseBody struct {
	// The details of the log.
	//
	// example:
	//
	// 7F40EAA1-6F1D-4DD9-8DB8-C5F00C4E****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The timestamp of log generation. Unit: ms.
	Result []*ListLogstashLogResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
}

func (s ListLogstashLogResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListLogstashLogResponseBody) GoString() string {
	return s.String()
}

func (s *ListLogstashLogResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListLogstashLogResponseBody) GetResult() []*ListLogstashLogResponseBodyResult {
	return s.Result
}

func (s *ListLogstashLogResponseBody) SetRequestId(v string) *ListLogstashLogResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListLogstashLogResponseBody) SetResult(v []*ListLogstashLogResponseBodyResult) *ListLogstashLogResponseBody {
	s.Result = v
	return s
}

func (s *ListLogstashLogResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListLogstashLogResponseBodyResult struct {
	// The IP address of the node that generates the log.
	Content *string `json:"content,omitempty" xml:"content,omitempty"`
	// example:
	//
	// 192.168.xx.xx
	Host *string `json:"host,omitempty" xml:"host,omitempty"`
	// example:
	//
	// ls-cn-v0h1kzca****
	InstanceId *string `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
	// The ID of the instance.
	//
	// example:
	//
	// info
	Level *string `json:"level,omitempty" xml:"level,omitempty"`
	// example:
	//
	// 1531985112420
	Timestamp *int64 `json:"timestamp,omitempty" xml:"timestamp,omitempty"`
}

func (s ListLogstashLogResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s ListLogstashLogResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListLogstashLogResponseBodyResult) GetContent() *string {
	return s.Content
}

func (s *ListLogstashLogResponseBodyResult) GetHost() *string {
	return s.Host
}

func (s *ListLogstashLogResponseBodyResult) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListLogstashLogResponseBodyResult) GetLevel() *string {
	return s.Level
}

func (s *ListLogstashLogResponseBodyResult) GetTimestamp() *int64 {
	return s.Timestamp
}

func (s *ListLogstashLogResponseBodyResult) SetContent(v string) *ListLogstashLogResponseBodyResult {
	s.Content = &v
	return s
}

func (s *ListLogstashLogResponseBodyResult) SetHost(v string) *ListLogstashLogResponseBodyResult {
	s.Host = &v
	return s
}

func (s *ListLogstashLogResponseBodyResult) SetInstanceId(v string) *ListLogstashLogResponseBodyResult {
	s.InstanceId = &v
	return s
}

func (s *ListLogstashLogResponseBodyResult) SetLevel(v string) *ListLogstashLogResponseBodyResult {
	s.Level = &v
	return s
}

func (s *ListLogstashLogResponseBodyResult) SetTimestamp(v int64) *ListLogstashLogResponseBodyResult {
	s.Timestamp = &v
	return s
}

func (s *ListLogstashLogResponseBodyResult) Validate() error {
	return dara.Validate(s)
}
