// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetUserAccessTokenForPartnerResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *GetUserAccessTokenForPartnerResponseBody
	GetErrorCode() *string
	SetModule(v *GetUserAccessTokenForPartnerResponseBodyModule) *GetUserAccessTokenForPartnerResponseBody
	GetModule() *GetUserAccessTokenForPartnerResponseBodyModule
	SetRequestId(v string) *GetUserAccessTokenForPartnerResponseBody
	GetRequestId() *string
	SetSuccess(v string) *GetUserAccessTokenForPartnerResponseBody
	GetSuccess() *string
}

type GetUserAccessTokenForPartnerResponseBody struct {
	// example:
	//
	// 0
	ErrorCode *string                                         `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	Module    *GetUserAccessTokenForPartnerResponseBodyModule `json:"Module,omitempty" xml:"Module,omitempty" type:"Struct"`
	// Id of the request
	//
	// example:
	//
	// 6C6B99AC-39EC-5350-874C-204128C905E6
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetUserAccessTokenForPartnerResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetUserAccessTokenForPartnerResponseBody) GoString() string {
	return s.String()
}

func (s *GetUserAccessTokenForPartnerResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *GetUserAccessTokenForPartnerResponseBody) GetModule() *GetUserAccessTokenForPartnerResponseBodyModule {
	return s.Module
}

func (s *GetUserAccessTokenForPartnerResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetUserAccessTokenForPartnerResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *GetUserAccessTokenForPartnerResponseBody) SetErrorCode(v string) *GetUserAccessTokenForPartnerResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetUserAccessTokenForPartnerResponseBody) SetModule(v *GetUserAccessTokenForPartnerResponseBodyModule) *GetUserAccessTokenForPartnerResponseBody {
	s.Module = v
	return s
}

func (s *GetUserAccessTokenForPartnerResponseBody) SetRequestId(v string) *GetUserAccessTokenForPartnerResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetUserAccessTokenForPartnerResponseBody) SetSuccess(v string) *GetUserAccessTokenForPartnerResponseBody {
	s.Success = &v
	return s
}

func (s *GetUserAccessTokenForPartnerResponseBody) Validate() error {
	if s.Module != nil {
		if err := s.Module.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetUserAccessTokenForPartnerResponseBodyModule struct {
	TokenValue *string `json:"TokenValue,omitempty" xml:"TokenValue,omitempty"`
}

func (s GetUserAccessTokenForPartnerResponseBodyModule) String() string {
	return dara.Prettify(s)
}

func (s GetUserAccessTokenForPartnerResponseBodyModule) GoString() string {
	return s.String()
}

func (s *GetUserAccessTokenForPartnerResponseBodyModule) GetTokenValue() *string {
	return s.TokenValue
}

func (s *GetUserAccessTokenForPartnerResponseBodyModule) SetTokenValue(v string) *GetUserAccessTokenForPartnerResponseBodyModule {
	s.TokenValue = &v
	return s
}

func (s *GetUserAccessTokenForPartnerResponseBodyModule) Validate() error {
	return dara.Validate(s)
}
