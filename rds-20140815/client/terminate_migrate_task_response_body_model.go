// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTerminateMigrateTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *TerminateMigrateTaskResponseBody
	GetRequestId() *string
}

type TerminateMigrateTaskResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 1AD222E9-E606-4A42-BF6D-8A4442913CEF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s TerminateMigrateTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s TerminateMigrateTaskResponseBody) GoString() string {
	return s.String()
}

func (s *TerminateMigrateTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *TerminateMigrateTaskResponseBody) SetRequestId(v string) *TerminateMigrateTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *TerminateMigrateTaskResponseBody) Validate() error {
	return dara.Validate(s)
}
