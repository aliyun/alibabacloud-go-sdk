// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAgentStorageResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateAgentStorageResponseBody
	GetRequestId() *string
}

type UpdateAgentStorageResponseBody struct {
	// The request ID, which can be used for troubleshooting.
	//
	// example:
	//
	// 39871ED2-62C0-578F-A32E-B88072D5582F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateAgentStorageResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateAgentStorageResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateAgentStorageResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateAgentStorageResponseBody) SetRequestId(v string) *UpdateAgentStorageResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateAgentStorageResponseBody) Validate() error {
	return dara.Validate(s)
}
