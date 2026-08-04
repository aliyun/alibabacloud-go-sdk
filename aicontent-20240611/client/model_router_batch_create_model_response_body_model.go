// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModelRouterBatchCreateModelResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *ModelRouterBatchCreateModelResponseBodyData) *ModelRouterBatchCreateModelResponseBody
	GetData() *ModelRouterBatchCreateModelResponseBodyData
	SetErrCode(v string) *ModelRouterBatchCreateModelResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *ModelRouterBatchCreateModelResponseBody
	GetErrMessage() *string
	SetHttpStatusCode(v int32) *ModelRouterBatchCreateModelResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *ModelRouterBatchCreateModelResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ModelRouterBatchCreateModelResponseBody
	GetSuccess() *bool
}

type ModelRouterBatchCreateModelResponseBody struct {
	// The data object.
	//
	// example:
	//
	// {}
	Data *ModelRouterBatchCreateModelResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// The fault code.
	//
	// example:
	//
	// UNKNOWN_ERROR
	ErrCode *string `json:"errCode,omitempty" xml:"errCode,omitempty"`
	// The error message.
	//
	// example:
	//
	// Unknown error
	ErrMessage *string `json:"errMessage,omitempty" xml:"errMessage,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	// The request ID.
	//
	// example:
	//
	// xxxx-xxxx-xxxx-xxxxxxxx
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s ModelRouterBatchCreateModelResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModelRouterBatchCreateModelResponseBody) GoString() string {
	return s.String()
}

func (s *ModelRouterBatchCreateModelResponseBody) GetData() *ModelRouterBatchCreateModelResponseBodyData {
	return s.Data
}

func (s *ModelRouterBatchCreateModelResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *ModelRouterBatchCreateModelResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *ModelRouterBatchCreateModelResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ModelRouterBatchCreateModelResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModelRouterBatchCreateModelResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ModelRouterBatchCreateModelResponseBody) SetData(v *ModelRouterBatchCreateModelResponseBodyData) *ModelRouterBatchCreateModelResponseBody {
	s.Data = v
	return s
}

func (s *ModelRouterBatchCreateModelResponseBody) SetErrCode(v string) *ModelRouterBatchCreateModelResponseBody {
	s.ErrCode = &v
	return s
}

func (s *ModelRouterBatchCreateModelResponseBody) SetErrMessage(v string) *ModelRouterBatchCreateModelResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *ModelRouterBatchCreateModelResponseBody) SetHttpStatusCode(v int32) *ModelRouterBatchCreateModelResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ModelRouterBatchCreateModelResponseBody) SetRequestId(v string) *ModelRouterBatchCreateModelResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModelRouterBatchCreateModelResponseBody) SetSuccess(v bool) *ModelRouterBatchCreateModelResponseBody {
	s.Success = &v
	return s
}

func (s *ModelRouterBatchCreateModelResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ModelRouterBatchCreateModelResponseBodyData struct {
	// The list of models that were successfully created.
	//
	// example:
	//
	// []
	Created []*ModelDTO `json:"created,omitempty" xml:"created,omitempty" type:"Repeated"`
	// The number of models that failed or were skipped.
	//
	// example:
	//
	// 0
	FailCount *int32 `json:"failCount,omitempty" xml:"failCount,omitempty"`
	// The list of models that failed or were skipped.
	//
	// example:
	//
	// []
	Failures []*BatchModelErrorDTO `json:"failures,omitempty" xml:"failures,omitempty" type:"Repeated"`
	// The number of models that were successfully created.
	//
	// example:
	//
	// 3
	SuccessCount *int32 `json:"successCount,omitempty" xml:"successCount,omitempty"`
}

func (s ModelRouterBatchCreateModelResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ModelRouterBatchCreateModelResponseBodyData) GoString() string {
	return s.String()
}

func (s *ModelRouterBatchCreateModelResponseBodyData) GetCreated() []*ModelDTO {
	return s.Created
}

func (s *ModelRouterBatchCreateModelResponseBodyData) GetFailCount() *int32 {
	return s.FailCount
}

func (s *ModelRouterBatchCreateModelResponseBodyData) GetFailures() []*BatchModelErrorDTO {
	return s.Failures
}

func (s *ModelRouterBatchCreateModelResponseBodyData) GetSuccessCount() *int32 {
	return s.SuccessCount
}

func (s *ModelRouterBatchCreateModelResponseBodyData) SetCreated(v []*ModelDTO) *ModelRouterBatchCreateModelResponseBodyData {
	s.Created = v
	return s
}

func (s *ModelRouterBatchCreateModelResponseBodyData) SetFailCount(v int32) *ModelRouterBatchCreateModelResponseBodyData {
	s.FailCount = &v
	return s
}

func (s *ModelRouterBatchCreateModelResponseBodyData) SetFailures(v []*BatchModelErrorDTO) *ModelRouterBatchCreateModelResponseBodyData {
	s.Failures = v
	return s
}

func (s *ModelRouterBatchCreateModelResponseBodyData) SetSuccessCount(v int32) *ModelRouterBatchCreateModelResponseBodyData {
	s.SuccessCount = &v
	return s
}

func (s *ModelRouterBatchCreateModelResponseBodyData) Validate() error {
	if s.Created != nil {
		for _, item := range s.Created {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Failures != nil {
		for _, item := range s.Failures {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
