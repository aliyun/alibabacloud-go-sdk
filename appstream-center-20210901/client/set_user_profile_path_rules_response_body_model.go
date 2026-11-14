// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetUserProfilePathRulesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *SetUserProfilePathRulesResponseBody
	GetCode() *string
	SetMessage(v string) *SetUserProfilePathRulesResponseBody
	GetMessage() *string
	SetRequestId(v string) *SetUserProfilePathRulesResponseBody
	GetRequestId() *string
}

type SetUserProfilePathRulesResponseBody struct {
	// The error code returned when the request fails.
	//
	// example:
	//
	// InvalidParameter.UserProfileRuleType
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The error message returned when the request fails.
	//
	// example:
	//
	// The parameter UserProfileRuleType is invalid.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The unique identifier of the request, which can be used for troubleshooting.
	//
	// example:
	//
	// 01A1CEEB-CDFF-18F9-B2BF-9997600D****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SetUserProfilePathRulesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SetUserProfilePathRulesResponseBody) GoString() string {
	return s.String()
}

func (s *SetUserProfilePathRulesResponseBody) GetCode() *string {
	return s.Code
}

func (s *SetUserProfilePathRulesResponseBody) GetMessage() *string {
	return s.Message
}

func (s *SetUserProfilePathRulesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SetUserProfilePathRulesResponseBody) SetCode(v string) *SetUserProfilePathRulesResponseBody {
	s.Code = &v
	return s
}

func (s *SetUserProfilePathRulesResponseBody) SetMessage(v string) *SetUserProfilePathRulesResponseBody {
	s.Message = &v
	return s
}

func (s *SetUserProfilePathRulesResponseBody) SetRequestId(v string) *SetUserProfilePathRulesResponseBody {
	s.RequestId = &v
	return s
}

func (s *SetUserProfilePathRulesResponseBody) Validate() error {
	return dara.Validate(s)
}
