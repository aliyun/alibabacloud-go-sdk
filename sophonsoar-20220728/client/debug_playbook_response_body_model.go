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
	// The request ID.
	//
	// example:
	//
	// 75E56B2C-C8FA-5A2F-AA08-8745E2AC33EF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The UUID of the debugging task. You can use the UUID to query the result and other details of the debugging task.
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
