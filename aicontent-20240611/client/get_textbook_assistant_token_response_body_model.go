// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTextbookAssistantTokenResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *GetTextbookAssistantTokenResponseBodyData) *GetTextbookAssistantTokenResponseBody
	GetData() *GetTextbookAssistantTokenResponseBodyData
	SetErrCode(v string) *GetTextbookAssistantTokenResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *GetTextbookAssistantTokenResponseBody
	GetErrMessage() *string
	SetHttpStatusCode(v int32) *GetTextbookAssistantTokenResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *GetTextbookAssistantTokenResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetTextbookAssistantTokenResponseBody
	GetSuccess() *bool
}

type GetTextbookAssistantTokenResponseBody struct {
	Data *GetTextbookAssistantTokenResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// example:
	//
	// 0
	ErrCode    *string `json:"errCode,omitempty" xml:"errCode,omitempty"`
	ErrMessage *string `json:"errMessage,omitempty" xml:"errMessage,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 0A5E9849-A2F0-551D-A7D8-1A8118557BAB
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s GetTextbookAssistantTokenResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetTextbookAssistantTokenResponseBody) GoString() string {
	return s.String()
}

func (s *GetTextbookAssistantTokenResponseBody) GetData() *GetTextbookAssistantTokenResponseBodyData {
	return s.Data
}

func (s *GetTextbookAssistantTokenResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *GetTextbookAssistantTokenResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *GetTextbookAssistantTokenResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetTextbookAssistantTokenResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetTextbookAssistantTokenResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetTextbookAssistantTokenResponseBody) SetData(v *GetTextbookAssistantTokenResponseBodyData) *GetTextbookAssistantTokenResponseBody {
	s.Data = v
	return s
}

func (s *GetTextbookAssistantTokenResponseBody) SetErrCode(v string) *GetTextbookAssistantTokenResponseBody {
	s.ErrCode = &v
	return s
}

func (s *GetTextbookAssistantTokenResponseBody) SetErrMessage(v string) *GetTextbookAssistantTokenResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *GetTextbookAssistantTokenResponseBody) SetHttpStatusCode(v int32) *GetTextbookAssistantTokenResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetTextbookAssistantTokenResponseBody) SetRequestId(v string) *GetTextbookAssistantTokenResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetTextbookAssistantTokenResponseBody) SetSuccess(v bool) *GetTextbookAssistantTokenResponseBody {
	s.Success = &v
	return s
}

func (s *GetTextbookAssistantTokenResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetTextbookAssistantTokenResponseBodyData struct {
	// example:
	//
	// tc_197bf5bb81889cc79eb51ae9b8c0cea3
	AuthToken *string `json:"authToken,omitempty" xml:"authToken,omitempty"`
	// example:
	//
	// 5400
	Expire *int32 `json:"expire,omitempty" xml:"expire,omitempty"`
}

func (s GetTextbookAssistantTokenResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetTextbookAssistantTokenResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetTextbookAssistantTokenResponseBodyData) GetAuthToken() *string {
	return s.AuthToken
}

func (s *GetTextbookAssistantTokenResponseBodyData) GetExpire() *int32 {
	return s.Expire
}

func (s *GetTextbookAssistantTokenResponseBodyData) SetAuthToken(v string) *GetTextbookAssistantTokenResponseBodyData {
	s.AuthToken = &v
	return s
}

func (s *GetTextbookAssistantTokenResponseBodyData) SetExpire(v int32) *GetTextbookAssistantTokenResponseBodyData {
	s.Expire = &v
	return s
}

func (s *GetTextbookAssistantTokenResponseBodyData) Validate() error {
	return dara.Validate(s)
}
