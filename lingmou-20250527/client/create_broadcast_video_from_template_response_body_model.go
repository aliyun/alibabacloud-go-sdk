// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateBroadcastVideoFromTemplateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateBroadcastVideoFromTemplateResponseBody
	GetCode() *string
	SetData(v *BroadcastVideo) *CreateBroadcastVideoFromTemplateResponseBody
	GetData() *BroadcastVideo
	SetHttpStatusCode(v int32) *CreateBroadcastVideoFromTemplateResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *CreateBroadcastVideoFromTemplateResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateBroadcastVideoFromTemplateResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateBroadcastVideoFromTemplateResponseBody
	GetSuccess() *bool
}

type CreateBroadcastVideoFromTemplateResponseBody struct {
	// example:
	//
	// SUCCESS
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// example:
	//
	// {"sessionId": "20250311-41523E3C-1D27-5844-8EEF-194E4714096B", "mainAccountId": 1234567, "createdAt": 1755680457}
	Data *BroadcastVideo `json:"data,omitempty" xml:"data,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// example:
	//
	// 0EC3BA89-13F5-5766-A0BA-85096092A032
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// True
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s CreateBroadcastVideoFromTemplateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateBroadcastVideoFromTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *CreateBroadcastVideoFromTemplateResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateBroadcastVideoFromTemplateResponseBody) GetData() *BroadcastVideo {
	return s.Data
}

func (s *CreateBroadcastVideoFromTemplateResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *CreateBroadcastVideoFromTemplateResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateBroadcastVideoFromTemplateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateBroadcastVideoFromTemplateResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateBroadcastVideoFromTemplateResponseBody) SetCode(v string) *CreateBroadcastVideoFromTemplateResponseBody {
	s.Code = &v
	return s
}

func (s *CreateBroadcastVideoFromTemplateResponseBody) SetData(v *BroadcastVideo) *CreateBroadcastVideoFromTemplateResponseBody {
	s.Data = v
	return s
}

func (s *CreateBroadcastVideoFromTemplateResponseBody) SetHttpStatusCode(v int32) *CreateBroadcastVideoFromTemplateResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *CreateBroadcastVideoFromTemplateResponseBody) SetMessage(v string) *CreateBroadcastVideoFromTemplateResponseBody {
	s.Message = &v
	return s
}

func (s *CreateBroadcastVideoFromTemplateResponseBody) SetRequestId(v string) *CreateBroadcastVideoFromTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateBroadcastVideoFromTemplateResponseBody) SetSuccess(v bool) *CreateBroadcastVideoFromTemplateResponseBody {
	s.Success = &v
	return s
}

func (s *CreateBroadcastVideoFromTemplateResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}
