// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDebugPlaybookResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DebugPlaybookResponseBody
	GetRequestId() *string
	SetRequestUuid(v string) *DebugPlaybookResponseBody
	GetRequestUuid() *string
}

type DebugPlaybookResponseBody struct {
	// The ID of the request. Alibaba Cloud generates a unique ID for each request. Use this ID to troubleshoot and locate issues.
	//
	// example:
	//
	// 75E56B2C-C8FA-5A2F-AA08-8745E2AC33EF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The run UUID of the debugging task. Use this UUID to query information, such as the task result.
	//
	// example:
	//
	// 6d412cfa-0905-4567-8a83-xxxxxx
	RequestUuid *string `json:"RequestUuid,omitempty" xml:"RequestUuid,omitempty"`
}

func (s DebugPlaybookResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DebugPlaybookResponseBody) GoString() string {
	return s.String()
}

func (s *DebugPlaybookResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DebugPlaybookResponseBody) GetRequestUuid() *string {
	return s.RequestUuid
}

func (s *DebugPlaybookResponseBody) SetRequestId(v string) *DebugPlaybookResponseBody {
	s.RequestId = &v
	return s
}

func (s *DebugPlaybookResponseBody) SetRequestUuid(v string) *DebugPlaybookResponseBody {
	s.RequestUuid = &v
	return s
}

func (s *DebugPlaybookResponseBody) Validate() error {
	return dara.Validate(s)
}
