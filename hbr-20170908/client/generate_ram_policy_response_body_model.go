// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGenerateRamPolicyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GenerateRamPolicyResponseBody
	GetCode() *string
	SetMessage(v string) *GenerateRamPolicyResponseBody
	GetMessage() *string
	SetPolicyDocument(v string) *GenerateRamPolicyResponseBody
	GetPolicyDocument() *string
	SetRequestId(v string) *GenerateRamPolicyResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GenerateRamPolicyResponseBody
	GetSuccess() *bool
}

type GenerateRamPolicyResponseBody struct {
	// The response code. The status code 200 indicates that the request was successful.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned message. If the request was successful, "successful" is returned. If the request failed, an error message is returned.
	//
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The content of the policy.
	//
	// example:
	//
	// {     "Version": "1",     "Statement": [         {             "Effect": "Deny",             "Action": [                 "hbr:CreateRestore",                 "hbr:CreateRestoreJob",                 "hbr:CreateHanaRestore",                 "hbr:CreateUniRestorePlan",                 "hbr:CreateSqlServerRestore"             ],             "Resource": [                 "acs:hbr:*:1178******531:vault/v-000******blx06",                 "acs:hbr:*:1178******531:vault/v-000******blx06/client/*"             ]         }     ] }
	PolicyDocument *string `json:"PolicyDocument,omitempty" xml:"PolicyDocument,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3C83E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GenerateRamPolicyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GenerateRamPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *GenerateRamPolicyResponseBody) GetCode() *string {
	return s.Code
}

func (s *GenerateRamPolicyResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GenerateRamPolicyResponseBody) GetPolicyDocument() *string {
	return s.PolicyDocument
}

func (s *GenerateRamPolicyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GenerateRamPolicyResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GenerateRamPolicyResponseBody) SetCode(v string) *GenerateRamPolicyResponseBody {
	s.Code = &v
	return s
}

func (s *GenerateRamPolicyResponseBody) SetMessage(v string) *GenerateRamPolicyResponseBody {
	s.Message = &v
	return s
}

func (s *GenerateRamPolicyResponseBody) SetPolicyDocument(v string) *GenerateRamPolicyResponseBody {
	s.PolicyDocument = &v
	return s
}

func (s *GenerateRamPolicyResponseBody) SetRequestId(v string) *GenerateRamPolicyResponseBody {
	s.RequestId = &v
	return s
}

func (s *GenerateRamPolicyResponseBody) SetSuccess(v bool) *GenerateRamPolicyResponseBody {
	s.Success = &v
	return s
}

func (s *GenerateRamPolicyResponseBody) Validate() error {
	return dara.Validate(s)
}
