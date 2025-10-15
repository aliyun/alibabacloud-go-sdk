// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddApplicationAccountToUserResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetApplicationAccountId(v string) *AddApplicationAccountToUserResponseBody
	GetApplicationAccountId() *string
	SetRequestId(v string) *AddApplicationAccountToUserResponseBody
	GetRequestId() *string
}

type AddApplicationAccountToUserResponseBody struct {
	// example:
	//
	// aac_m6z7awz5kresi2ezgajsbkxxxx
	ApplicationAccountId *string `json:"ApplicationAccountId,omitempty" xml:"ApplicationAccountId,omitempty"`
	// example:
	//
	// 0441BD79-92F3-53AA-8657-F8CE4A2B912A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AddApplicationAccountToUserResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AddApplicationAccountToUserResponseBody) GoString() string {
	return s.String()
}

func (s *AddApplicationAccountToUserResponseBody) GetApplicationAccountId() *string {
	return s.ApplicationAccountId
}

func (s *AddApplicationAccountToUserResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AddApplicationAccountToUserResponseBody) SetApplicationAccountId(v string) *AddApplicationAccountToUserResponseBody {
	s.ApplicationAccountId = &v
	return s
}

func (s *AddApplicationAccountToUserResponseBody) SetRequestId(v string) *AddApplicationAccountToUserResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddApplicationAccountToUserResponseBody) Validate() error {
	return dara.Validate(s)
}
