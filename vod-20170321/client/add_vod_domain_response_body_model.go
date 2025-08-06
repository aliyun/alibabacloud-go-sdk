// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddVodDomainResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *AddVodDomainResponseBody
	GetRequestId() *string
}

type AddVodDomainResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 15C66C7B-671A-4297-****-2C4477247A74
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AddVodDomainResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AddVodDomainResponseBody) GoString() string {
	return s.String()
}

func (s *AddVodDomainResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AddVodDomainResponseBody) SetRequestId(v string) *AddVodDomainResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddVodDomainResponseBody) Validate() error {
	return dara.Validate(s)
}
