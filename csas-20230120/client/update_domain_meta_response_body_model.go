// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateDomainMetaResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateDomainMetaResponseBody
	GetRequestId() *string
}

type UpdateDomainMetaResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// 019F68A7-D149-5BE5-9B35-5D59BAE545B3
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateDomainMetaResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateDomainMetaResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateDomainMetaResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateDomainMetaResponseBody) SetRequestId(v string) *UpdateDomainMetaResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateDomainMetaResponseBody) Validate() error {
	return dara.Validate(s)
}
