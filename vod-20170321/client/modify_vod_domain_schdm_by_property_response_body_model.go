// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyVodDomainSchdmByPropertyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyVodDomainSchdmByPropertyResponseBody
	GetRequestId() *string
}

type ModifyVodDomainSchdmByPropertyResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyVodDomainSchdmByPropertyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyVodDomainSchdmByPropertyResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyVodDomainSchdmByPropertyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyVodDomainSchdmByPropertyResponseBody) SetRequestId(v string) *ModifyVodDomainSchdmByPropertyResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyVodDomainSchdmByPropertyResponseBody) Validate() error {
	return dara.Validate(s)
}
