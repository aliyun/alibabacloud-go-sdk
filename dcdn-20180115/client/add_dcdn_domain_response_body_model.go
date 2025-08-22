// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddDcdnDomainResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *AddDcdnDomainResponseBody
	GetRequestId() *string
}

type AddDcdnDomainResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 15C66C7B-671A-4297-9187-2C4477247A74
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AddDcdnDomainResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AddDcdnDomainResponseBody) GoString() string {
	return s.String()
}

func (s *AddDcdnDomainResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AddDcdnDomainResponseBody) SetRequestId(v string) *AddDcdnDomainResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddDcdnDomainResponseBody) Validate() error {
	return dara.Validate(s)
}
