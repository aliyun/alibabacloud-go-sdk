// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateVodDomainResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateVodDomainResponseBody
	GetRequestId() *string
}

type UpdateVodDomainResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 15C66C7B-671A-4297-****-2C4477247A74
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateVodDomainResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateVodDomainResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateVodDomainResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateVodDomainResponseBody) SetRequestId(v string) *UpdateVodDomainResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateVodDomainResponseBody) Validate() error {
	return dara.Validate(s)
}
