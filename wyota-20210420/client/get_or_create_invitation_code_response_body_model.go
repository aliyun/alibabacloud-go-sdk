// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetOrCreateInvitationCodeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetOrCreateInvitationCodeResponseBody
	GetCode() *string
	SetData(v *GetOrCreateInvitationCodeResponseBodyData) *GetOrCreateInvitationCodeResponseBody
	GetData() *GetOrCreateInvitationCodeResponseBodyData
	SetHttpStatusCode(v int32) *GetOrCreateInvitationCodeResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GetOrCreateInvitationCodeResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetOrCreateInvitationCodeResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetOrCreateInvitationCodeResponseBody
	GetSuccess() *bool
}

type GetOrCreateInvitationCodeResponseBody struct {
	// example:
	//
	// PARAMETER_ERROR
	Code *string                                    `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *GetOrCreateInvitationCodeResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// parameter error
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// C5DCE54A-B266-522E-A6ED-468AF45F5AAA
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetOrCreateInvitationCodeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetOrCreateInvitationCodeResponseBody) GoString() string {
	return s.String()
}

func (s *GetOrCreateInvitationCodeResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetOrCreateInvitationCodeResponseBody) GetData() *GetOrCreateInvitationCodeResponseBodyData {
	return s.Data
}

func (s *GetOrCreateInvitationCodeResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetOrCreateInvitationCodeResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetOrCreateInvitationCodeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetOrCreateInvitationCodeResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetOrCreateInvitationCodeResponseBody) SetCode(v string) *GetOrCreateInvitationCodeResponseBody {
	s.Code = &v
	return s
}

func (s *GetOrCreateInvitationCodeResponseBody) SetData(v *GetOrCreateInvitationCodeResponseBodyData) *GetOrCreateInvitationCodeResponseBody {
	s.Data = v
	return s
}

func (s *GetOrCreateInvitationCodeResponseBody) SetHttpStatusCode(v int32) *GetOrCreateInvitationCodeResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetOrCreateInvitationCodeResponseBody) SetMessage(v string) *GetOrCreateInvitationCodeResponseBody {
	s.Message = &v
	return s
}

func (s *GetOrCreateInvitationCodeResponseBody) SetRequestId(v string) *GetOrCreateInvitationCodeResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetOrCreateInvitationCodeResponseBody) SetSuccess(v bool) *GetOrCreateInvitationCodeResponseBody {
	s.Success = &v
	return s
}

func (s *GetOrCreateInvitationCodeResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetOrCreateInvitationCodeResponseBodyData struct {
	// example:
	//
	// 000000
	AuthCode *string `json:"AuthCode,omitempty" xml:"AuthCode,omitempty"`
	// example:
	//
	// 1772162247
	ExpireTime *string `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	// example:
	//
	// True
	Expired *bool `json:"Expired,omitempty" xml:"Expired,omitempty"`
}

func (s GetOrCreateInvitationCodeResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetOrCreateInvitationCodeResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetOrCreateInvitationCodeResponseBodyData) GetAuthCode() *string {
	return s.AuthCode
}

func (s *GetOrCreateInvitationCodeResponseBodyData) GetExpireTime() *string {
	return s.ExpireTime
}

func (s *GetOrCreateInvitationCodeResponseBodyData) GetExpired() *bool {
	return s.Expired
}

func (s *GetOrCreateInvitationCodeResponseBodyData) SetAuthCode(v string) *GetOrCreateInvitationCodeResponseBodyData {
	s.AuthCode = &v
	return s
}

func (s *GetOrCreateInvitationCodeResponseBodyData) SetExpireTime(v string) *GetOrCreateInvitationCodeResponseBodyData {
	s.ExpireTime = &v
	return s
}

func (s *GetOrCreateInvitationCodeResponseBodyData) SetExpired(v bool) *GetOrCreateInvitationCodeResponseBodyData {
	s.Expired = &v
	return s
}

func (s *GetOrCreateInvitationCodeResponseBodyData) Validate() error {
	return dara.Validate(s)
}
