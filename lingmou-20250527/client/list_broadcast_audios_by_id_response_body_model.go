// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListBroadcastAudiosByIdResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListBroadcastAudiosByIdResponseBody
	GetCode() *string
	SetData(v []*BroadcastAudio) *ListBroadcastAudiosByIdResponseBody
	GetData() []*BroadcastAudio
	SetHttpStatusCode(v int32) *ListBroadcastAudiosByIdResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *ListBroadcastAudiosByIdResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListBroadcastAudiosByIdResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListBroadcastAudiosByIdResponseBody
	GetSuccess() *bool
}

type ListBroadcastAudiosByIdResponseBody struct {
	// example:
	//
	// 200
	Code *string           `json:"code,omitempty" xml:"code,omitempty"`
	Data []*BroadcastAudio `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
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
	// 0EC3BA89-13F5-5766-A0BA-85096092A032
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// True
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s ListBroadcastAudiosByIdResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListBroadcastAudiosByIdResponseBody) GoString() string {
	return s.String()
}

func (s *ListBroadcastAudiosByIdResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListBroadcastAudiosByIdResponseBody) GetData() []*BroadcastAudio {
	return s.Data
}

func (s *ListBroadcastAudiosByIdResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListBroadcastAudiosByIdResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListBroadcastAudiosByIdResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListBroadcastAudiosByIdResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListBroadcastAudiosByIdResponseBody) SetCode(v string) *ListBroadcastAudiosByIdResponseBody {
	s.Code = &v
	return s
}

func (s *ListBroadcastAudiosByIdResponseBody) SetData(v []*BroadcastAudio) *ListBroadcastAudiosByIdResponseBody {
	s.Data = v
	return s
}

func (s *ListBroadcastAudiosByIdResponseBody) SetHttpStatusCode(v int32) *ListBroadcastAudiosByIdResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListBroadcastAudiosByIdResponseBody) SetMessage(v string) *ListBroadcastAudiosByIdResponseBody {
	s.Message = &v
	return s
}

func (s *ListBroadcastAudiosByIdResponseBody) SetRequestId(v string) *ListBroadcastAudiosByIdResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListBroadcastAudiosByIdResponseBody) SetSuccess(v bool) *ListBroadcastAudiosByIdResponseBody {
	s.Success = &v
	return s
}

func (s *ListBroadcastAudiosByIdResponseBody) Validate() error {
	if s.Data != nil {
		for _, item := range s.Data {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
