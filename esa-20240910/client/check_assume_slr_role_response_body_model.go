// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckAssumeSlrRoleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorMsg(v string) *CheckAssumeSlrRoleResponseBody
	GetErrorMsg() *string
	SetIsExist(v string) *CheckAssumeSlrRoleResponseBody
	GetIsExist() *string
	SetRequestId(v string) *CheckAssumeSlrRoleResponseBody
	GetRequestId() *string
}

type CheckAssumeSlrRoleResponseBody struct {
	// example:
	//
	// aliuid:xxx assumeOssRole not exist,serviceName:realtimelogpush.dcdnservices.aliyuncs.com
	ErrorMsg *string `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	// example:
	//
	// true
	IsExist *string `json:"IsExist,omitempty" xml:"IsExist,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 34DCBC8A-****-****-****-6DAA11D7DDBD
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CheckAssumeSlrRoleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CheckAssumeSlrRoleResponseBody) GoString() string {
	return s.String()
}

func (s *CheckAssumeSlrRoleResponseBody) GetErrorMsg() *string {
	return s.ErrorMsg
}

func (s *CheckAssumeSlrRoleResponseBody) GetIsExist() *string {
	return s.IsExist
}

func (s *CheckAssumeSlrRoleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CheckAssumeSlrRoleResponseBody) SetErrorMsg(v string) *CheckAssumeSlrRoleResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *CheckAssumeSlrRoleResponseBody) SetIsExist(v string) *CheckAssumeSlrRoleResponseBody {
	s.IsExist = &v
	return s
}

func (s *CheckAssumeSlrRoleResponseBody) SetRequestId(v string) *CheckAssumeSlrRoleResponseBody {
	s.RequestId = &v
	return s
}

func (s *CheckAssumeSlrRoleResponseBody) Validate() error {
	return dara.Validate(s)
}
