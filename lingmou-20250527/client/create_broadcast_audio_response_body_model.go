// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateBroadcastAudioResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateBroadcastAudioResponseBody
	GetCode() *string
	SetData(v *BroadcastAudio) *CreateBroadcastAudioResponseBody
	GetData() *BroadcastAudio
	SetHttpStatusCode(v int32) *CreateBroadcastAudioResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *CreateBroadcastAudioResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateBroadcastAudioResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateBroadcastAudioResponseBody
	GetSuccess() *bool
}

type CreateBroadcastAudioResponseBody struct {
	// example:
	//
	// 200
	Code *string         `json:"code,omitempty" xml:"code,omitempty"`
	Data *BroadcastAudio `json:"data,omitempty" xml:"data,omitempty"`
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
	// 90C68329-A75C-5449-A928-4D0BAD7AA0FA
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// True
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s CreateBroadcastAudioResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateBroadcastAudioResponseBody) GoString() string {
	return s.String()
}

func (s *CreateBroadcastAudioResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateBroadcastAudioResponseBody) GetData() *BroadcastAudio {
	return s.Data
}

func (s *CreateBroadcastAudioResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *CreateBroadcastAudioResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateBroadcastAudioResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateBroadcastAudioResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateBroadcastAudioResponseBody) SetCode(v string) *CreateBroadcastAudioResponseBody {
	s.Code = &v
	return s
}

func (s *CreateBroadcastAudioResponseBody) SetData(v *BroadcastAudio) *CreateBroadcastAudioResponseBody {
	s.Data = v
	return s
}

func (s *CreateBroadcastAudioResponseBody) SetHttpStatusCode(v int32) *CreateBroadcastAudioResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *CreateBroadcastAudioResponseBody) SetMessage(v string) *CreateBroadcastAudioResponseBody {
	s.Message = &v
	return s
}

func (s *CreateBroadcastAudioResponseBody) SetRequestId(v string) *CreateBroadcastAudioResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateBroadcastAudioResponseBody) SetSuccess(v bool) *CreateBroadcastAudioResponseBody {
	s.Success = &v
	return s
}

func (s *CreateBroadcastAudioResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}
