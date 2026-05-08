// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListVoiceModelsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListVoiceModelsResponseBody
	GetCode() *string
	SetErrorCode(v string) *ListVoiceModelsResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *ListVoiceModelsResponseBody
	GetErrorMessage() *string
	SetList(v []*VoiceModelResponse) *ListVoiceModelsResponseBody
	GetList() []*VoiceModelResponse
	SetRequestId(v string) *ListVoiceModelsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListVoiceModelsResponseBody
	GetSuccess() *bool
	SetTotal(v int32) *ListVoiceModelsResponseBody
	GetTotal() *int32
}

type ListVoiceModelsResponseBody struct {
	// example:
	//
	// 200
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// example:
	//
	// 040002
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// example:
	//
	// Failed to proxy flink ui request, message: An error occurred: Invalid UUID string: jobsn
	ErrorMessage *string               `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	List         []*VoiceModelResponse `json:"list,omitempty" xml:"list,omitempty" type:"Repeated"`
	// Id of the request
	//
	// example:
	//
	// 14878724-A835-578D-9DD5-4779ADCE9221
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// True
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
	// example:
	//
	// 10
	Total *int32 `json:"total,omitempty" xml:"total,omitempty"`
}

func (s ListVoiceModelsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListVoiceModelsResponseBody) GoString() string {
	return s.String()
}

func (s *ListVoiceModelsResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListVoiceModelsResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *ListVoiceModelsResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *ListVoiceModelsResponseBody) GetList() []*VoiceModelResponse {
	return s.List
}

func (s *ListVoiceModelsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListVoiceModelsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListVoiceModelsResponseBody) GetTotal() *int32 {
	return s.Total
}

func (s *ListVoiceModelsResponseBody) SetCode(v string) *ListVoiceModelsResponseBody {
	s.Code = &v
	return s
}

func (s *ListVoiceModelsResponseBody) SetErrorCode(v string) *ListVoiceModelsResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListVoiceModelsResponseBody) SetErrorMessage(v string) *ListVoiceModelsResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *ListVoiceModelsResponseBody) SetList(v []*VoiceModelResponse) *ListVoiceModelsResponseBody {
	s.List = v
	return s
}

func (s *ListVoiceModelsResponseBody) SetRequestId(v string) *ListVoiceModelsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListVoiceModelsResponseBody) SetSuccess(v bool) *ListVoiceModelsResponseBody {
	s.Success = &v
	return s
}

func (s *ListVoiceModelsResponseBody) SetTotal(v int32) *ListVoiceModelsResponseBody {
	s.Total = &v
	return s
}

func (s *ListVoiceModelsResponseBody) Validate() error {
	if s.List != nil {
		for _, item := range s.List {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
