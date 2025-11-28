// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetBroadcastTemplateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetBroadcastTemplateResponseBody
	GetCode() *string
	SetData(v *BroadcastTemplate) *GetBroadcastTemplateResponseBody
	GetData() *BroadcastTemplate
	SetHttpStatusCode(v int32) *GetBroadcastTemplateResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GetBroadcastTemplateResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetBroadcastTemplateResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetBroadcastTemplateResponseBody
	GetSuccess() *bool
}

type GetBroadcastTemplateResponseBody struct {
	// example:
	//
	// 200
	Code *string            `json:"code,omitempty" xml:"code,omitempty"`
	Data *BroadcastTemplate `json:"data,omitempty" xml:"data,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	// example:
	//
	// SUCCESS
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// example:
	//
	// 7239F9E5-A4DB-55BA-B701-0CE47DBDB0A8
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// True
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s GetBroadcastTemplateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetBroadcastTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *GetBroadcastTemplateResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetBroadcastTemplateResponseBody) GetData() *BroadcastTemplate {
	return s.Data
}

func (s *GetBroadcastTemplateResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetBroadcastTemplateResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetBroadcastTemplateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetBroadcastTemplateResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetBroadcastTemplateResponseBody) SetCode(v string) *GetBroadcastTemplateResponseBody {
	s.Code = &v
	return s
}

func (s *GetBroadcastTemplateResponseBody) SetData(v *BroadcastTemplate) *GetBroadcastTemplateResponseBody {
	s.Data = v
	return s
}

func (s *GetBroadcastTemplateResponseBody) SetHttpStatusCode(v int32) *GetBroadcastTemplateResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetBroadcastTemplateResponseBody) SetMessage(v string) *GetBroadcastTemplateResponseBody {
	s.Message = &v
	return s
}

func (s *GetBroadcastTemplateResponseBody) SetRequestId(v string) *GetBroadcastTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetBroadcastTemplateResponseBody) SetSuccess(v bool) *GetBroadcastTemplateResponseBody {
	s.Success = &v
	return s
}

func (s *GetBroadcastTemplateResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}
