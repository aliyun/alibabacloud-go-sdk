// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetPlatformUserInfoForPartnerResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetEncryptedOpenId(v string) *GetPlatformUserInfoForPartnerResponseBody
	GetEncryptedOpenId() *string
	SetEncryptedUnionId(v string) *GetPlatformUserInfoForPartnerResponseBody
	GetEncryptedUnionId() *string
	SetErrorMsg(v string) *GetPlatformUserInfoForPartnerResponseBody
	GetErrorMsg() *string
	SetRequestId(v string) *GetPlatformUserInfoForPartnerResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetPlatformUserInfoForPartnerResponseBody
	GetSuccess() *bool
}

type GetPlatformUserInfoForPartnerResponseBody struct {
	// example:
	//
	// djisdpfOjofjifojfajaspsdkasdada
	EncryptedOpenId *string `json:"EncryptedOpenId,omitempty" xml:"EncryptedOpenId,omitempty"`
	// example:
	//
	// djisdpfOjofjifojfajaspsdkasdada
	EncryptedUnionId *string `json:"EncryptedUnionId,omitempty" xml:"EncryptedUnionId,omitempty"`
	// example:
	//
	// 11111111111111111111111
	ErrorMsg *string `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 0320C9F4-5EDC-5355-9D7E-DF4CF6C2B3BB
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// success
	//
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetPlatformUserInfoForPartnerResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetPlatformUserInfoForPartnerResponseBody) GoString() string {
	return s.String()
}

func (s *GetPlatformUserInfoForPartnerResponseBody) GetEncryptedOpenId() *string {
	return s.EncryptedOpenId
}

func (s *GetPlatformUserInfoForPartnerResponseBody) GetEncryptedUnionId() *string {
	return s.EncryptedUnionId
}

func (s *GetPlatformUserInfoForPartnerResponseBody) GetErrorMsg() *string {
	return s.ErrorMsg
}

func (s *GetPlatformUserInfoForPartnerResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetPlatformUserInfoForPartnerResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetPlatformUserInfoForPartnerResponseBody) SetEncryptedOpenId(v string) *GetPlatformUserInfoForPartnerResponseBody {
	s.EncryptedOpenId = &v
	return s
}

func (s *GetPlatformUserInfoForPartnerResponseBody) SetEncryptedUnionId(v string) *GetPlatformUserInfoForPartnerResponseBody {
	s.EncryptedUnionId = &v
	return s
}

func (s *GetPlatformUserInfoForPartnerResponseBody) SetErrorMsg(v string) *GetPlatformUserInfoForPartnerResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *GetPlatformUserInfoForPartnerResponseBody) SetRequestId(v string) *GetPlatformUserInfoForPartnerResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetPlatformUserInfoForPartnerResponseBody) SetSuccess(v bool) *GetPlatformUserInfoForPartnerResponseBody {
	s.Success = &v
	return s
}

func (s *GetPlatformUserInfoForPartnerResponseBody) Validate() error {
	return dara.Validate(s)
}
