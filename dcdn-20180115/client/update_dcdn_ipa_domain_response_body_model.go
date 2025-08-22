// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateDcdnIpaDomainResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateDcdnIpaDomainResponseBody
	GetRequestId() *string
}

type UpdateDcdnIpaDomainResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 15C66C7B-671A-4297-9187-2C4477247A74
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateDcdnIpaDomainResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateDcdnIpaDomainResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateDcdnIpaDomainResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateDcdnIpaDomainResponseBody) SetRequestId(v string) *UpdateDcdnIpaDomainResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateDcdnIpaDomainResponseBody) Validate() error {
	return dara.Validate(s)
}
