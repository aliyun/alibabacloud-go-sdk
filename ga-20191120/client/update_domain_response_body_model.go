// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateDomainResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateDomainResponseBody
	GetRequestId() *string
}

type UpdateDomainResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 04F0F334-1335-436C-A1D7-6C044FE73368
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateDomainResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateDomainResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateDomainResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateDomainResponseBody) SetRequestId(v string) *UpdateDomainResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateDomainResponseBody) Validate() error {
	return dara.Validate(s)
}
