// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDCdnDomainSchdmByPropertyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyDCdnDomainSchdmByPropertyResponseBody
	GetRequestId() *string
}

type ModifyDCdnDomainSchdmByPropertyResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 15C66C7B-671A-4297-9187-2C4477247A74
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyDCdnDomainSchdmByPropertyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyDCdnDomainSchdmByPropertyResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyDCdnDomainSchdmByPropertyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyDCdnDomainSchdmByPropertyResponseBody) SetRequestId(v string) *ModifyDCdnDomainSchdmByPropertyResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyDCdnDomainSchdmByPropertyResponseBody) Validate() error {
	return dara.Validate(s)
}
