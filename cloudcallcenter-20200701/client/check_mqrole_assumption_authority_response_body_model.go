// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckMQRoleAssumptionAuthorityResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CheckMQRoleAssumptionAuthorityResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *CheckMQRoleAssumptionAuthorityResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *CheckMQRoleAssumptionAuthorityResponseBody
	GetMessage() *string
	SetParams(v []*string) *CheckMQRoleAssumptionAuthorityResponseBody
	GetParams() []*string
	SetRequestId(v string) *CheckMQRoleAssumptionAuthorityResponseBody
	GetRequestId() *string
}

type CheckMQRoleAssumptionAuthorityResponseBody struct {
	Code           *string   `json:"Code,omitempty" xml:"Code,omitempty"`
	HttpStatusCode *int32    `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string   `json:"Message,omitempty" xml:"Message,omitempty"`
	Params         []*string `json:"Params,omitempty" xml:"Params,omitempty" type:"Repeated"`
	RequestId      *string   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CheckMQRoleAssumptionAuthorityResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CheckMQRoleAssumptionAuthorityResponseBody) GoString() string {
	return s.String()
}

func (s *CheckMQRoleAssumptionAuthorityResponseBody) GetCode() *string {
	return s.Code
}

func (s *CheckMQRoleAssumptionAuthorityResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *CheckMQRoleAssumptionAuthorityResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CheckMQRoleAssumptionAuthorityResponseBody) GetParams() []*string {
	return s.Params
}

func (s *CheckMQRoleAssumptionAuthorityResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CheckMQRoleAssumptionAuthorityResponseBody) SetCode(v string) *CheckMQRoleAssumptionAuthorityResponseBody {
	s.Code = &v
	return s
}

func (s *CheckMQRoleAssumptionAuthorityResponseBody) SetHttpStatusCode(v int32) *CheckMQRoleAssumptionAuthorityResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *CheckMQRoleAssumptionAuthorityResponseBody) SetMessage(v string) *CheckMQRoleAssumptionAuthorityResponseBody {
	s.Message = &v
	return s
}

func (s *CheckMQRoleAssumptionAuthorityResponseBody) SetParams(v []*string) *CheckMQRoleAssumptionAuthorityResponseBody {
	s.Params = v
	return s
}

func (s *CheckMQRoleAssumptionAuthorityResponseBody) SetRequestId(v string) *CheckMQRoleAssumptionAuthorityResponseBody {
	s.RequestId = &v
	return s
}

func (s *CheckMQRoleAssumptionAuthorityResponseBody) Validate() error {
	return dara.Validate(s)
}
