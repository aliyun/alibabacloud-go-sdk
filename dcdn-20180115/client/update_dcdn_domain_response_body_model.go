// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateDcdnDomainResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateDcdnDomainResponseBody
	GetRequestId() *string
}

type UpdateDcdnDomainResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 0AEDAF20-4DDF-4165-8750-47FF9C1929C9
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateDcdnDomainResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateDcdnDomainResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateDcdnDomainResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateDcdnDomainResponseBody) SetRequestId(v string) *UpdateDcdnDomainResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateDcdnDomainResponseBody) Validate() error {
	return dara.Validate(s)
}
