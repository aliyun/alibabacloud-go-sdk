// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyLiveDomainSchdmByPropertyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyLiveDomainSchdmByPropertyResponseBody
	GetRequestId() *string
}

type ModifyLiveDomainSchdmByPropertyResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 15C66C7B-671A-4297-9187-2C4477247A74
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyLiveDomainSchdmByPropertyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyLiveDomainSchdmByPropertyResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyLiveDomainSchdmByPropertyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyLiveDomainSchdmByPropertyResponseBody) SetRequestId(v string) *ModifyLiveDomainSchdmByPropertyResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyLiveDomainSchdmByPropertyResponseBody) Validate() error {
	return dara.Validate(s)
}
