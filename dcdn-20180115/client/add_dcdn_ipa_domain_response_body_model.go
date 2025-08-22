// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddDcdnIpaDomainResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *AddDcdnIpaDomainResponseBody
	GetRequestId() *string
}

type AddDcdnIpaDomainResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 0AEDAF20-4DDF-4165-8750-47FF9C1929C9
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AddDcdnIpaDomainResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AddDcdnIpaDomainResponseBody) GoString() string {
	return s.String()
}

func (s *AddDcdnIpaDomainResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AddDcdnIpaDomainResponseBody) SetRequestId(v string) *AddDcdnIpaDomainResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddDcdnIpaDomainResponseBody) Validate() error {
	return dara.Validate(s)
}
