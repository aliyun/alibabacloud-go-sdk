// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyCdnDomainSchdmByPropertyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyCdnDomainSchdmByPropertyResponseBody
	GetRequestId() *string
}

type ModifyCdnDomainSchdmByPropertyResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 15C66C7B-671A-4297-9187-2C4477247A74
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyCdnDomainSchdmByPropertyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyCdnDomainSchdmByPropertyResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyCdnDomainSchdmByPropertyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyCdnDomainSchdmByPropertyResponseBody) SetRequestId(v string) *ModifyCdnDomainSchdmByPropertyResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyCdnDomainSchdmByPropertyResponseBody) Validate() error {
	return dara.Validate(s)
}
