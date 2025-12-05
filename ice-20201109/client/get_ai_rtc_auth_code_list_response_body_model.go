// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAiRtcAuthCodeListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAuthCodeList(v []*AiRtcAuthCodeDTO) *GetAiRtcAuthCodeListResponseBody
	GetAuthCodeList() []*AiRtcAuthCodeDTO
	SetCode(v string) *GetAiRtcAuthCodeListResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *GetAiRtcAuthCodeListResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GetAiRtcAuthCodeListResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetAiRtcAuthCodeListResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetAiRtcAuthCodeListResponseBody
	GetSuccess() *bool
	SetTotalCount(v int64) *GetAiRtcAuthCodeListResponseBody
	GetTotalCount() *int64
}

type GetAiRtcAuthCodeListResponseBody struct {
	// An array of AiRtcAuthCodeDTO objects, each representing an authorization code.
	AuthCodeList []*AiRtcAuthCodeDTO `json:"AuthCodeList,omitempty" xml:"AuthCodeList,omitempty" type:"Repeated"`
	// The error code.
	//
	// example:
	//
	// Success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The error message.
	//
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 7B117AF5-2A16-412C-B127-FA6175ED1***
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the call was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 10
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s GetAiRtcAuthCodeListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetAiRtcAuthCodeListResponseBody) GoString() string {
	return s.String()
}

func (s *GetAiRtcAuthCodeListResponseBody) GetAuthCodeList() []*AiRtcAuthCodeDTO {
	return s.AuthCodeList
}

func (s *GetAiRtcAuthCodeListResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetAiRtcAuthCodeListResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetAiRtcAuthCodeListResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetAiRtcAuthCodeListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetAiRtcAuthCodeListResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetAiRtcAuthCodeListResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *GetAiRtcAuthCodeListResponseBody) SetAuthCodeList(v []*AiRtcAuthCodeDTO) *GetAiRtcAuthCodeListResponseBody {
	s.AuthCodeList = v
	return s
}

func (s *GetAiRtcAuthCodeListResponseBody) SetCode(v string) *GetAiRtcAuthCodeListResponseBody {
	s.Code = &v
	return s
}

func (s *GetAiRtcAuthCodeListResponseBody) SetHttpStatusCode(v int32) *GetAiRtcAuthCodeListResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetAiRtcAuthCodeListResponseBody) SetMessage(v string) *GetAiRtcAuthCodeListResponseBody {
	s.Message = &v
	return s
}

func (s *GetAiRtcAuthCodeListResponseBody) SetRequestId(v string) *GetAiRtcAuthCodeListResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAiRtcAuthCodeListResponseBody) SetSuccess(v bool) *GetAiRtcAuthCodeListResponseBody {
	s.Success = &v
	return s
}

func (s *GetAiRtcAuthCodeListResponseBody) SetTotalCount(v int64) *GetAiRtcAuthCodeListResponseBody {
	s.TotalCount = &v
	return s
}

func (s *GetAiRtcAuthCodeListResponseBody) Validate() error {
	if s.AuthCodeList != nil {
		for _, item := range s.AuthCodeList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
