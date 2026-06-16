// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAssetOptimizeProResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *AssetOptimizeProResponseBody
	GetCode() *string
	SetData(v *AssetOptimizeProResponseBodyData) *AssetOptimizeProResponseBody
	GetData() *AssetOptimizeProResponseBodyData
	SetMessage(v string) *AssetOptimizeProResponseBody
	GetMessage() *string
	SetRequestId(v string) *AssetOptimizeProResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *AssetOptimizeProResponseBody
	GetSuccess() *bool
}

type AssetOptimizeProResponseBody struct {
	// The status code. The value "success" is returned for a successful call.
	//
	// example:
	//
	// success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The result data of the asynchronous task submission, which contains the asynchronous task ID.
	Data *AssetOptimizeProResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The error message. The value "Success" is returned for a successful call.
	//
	// example:
	//
	// Success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID, which uniquely identifies the API call.
	//
	// example:
	//
	// 4FCCA90A-A7A4-1D37-88C3-C17549886E70
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the call is successful. A value of true indicates success. A value of false indicates failure.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s AssetOptimizeProResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AssetOptimizeProResponseBody) GoString() string {
	return s.String()
}

func (s *AssetOptimizeProResponseBody) GetCode() *string {
	return s.Code
}

func (s *AssetOptimizeProResponseBody) GetData() *AssetOptimizeProResponseBodyData {
	return s.Data
}

func (s *AssetOptimizeProResponseBody) GetMessage() *string {
	return s.Message
}

func (s *AssetOptimizeProResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AssetOptimizeProResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *AssetOptimizeProResponseBody) SetCode(v string) *AssetOptimizeProResponseBody {
	s.Code = &v
	return s
}

func (s *AssetOptimizeProResponseBody) SetData(v *AssetOptimizeProResponseBodyData) *AssetOptimizeProResponseBody {
	s.Data = v
	return s
}

func (s *AssetOptimizeProResponseBody) SetMessage(v string) *AssetOptimizeProResponseBody {
	s.Message = &v
	return s
}

func (s *AssetOptimizeProResponseBody) SetRequestId(v string) *AssetOptimizeProResponseBody {
	s.RequestId = &v
	return s
}

func (s *AssetOptimizeProResponseBody) SetSuccess(v bool) *AssetOptimizeProResponseBody {
	s.Success = &v
	return s
}

func (s *AssetOptimizeProResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type AssetOptimizeProResponseBodyData struct {
	// The asynchronous task ID, which is used to query the processing result by calling QueryAsyncTaskResult.
	//
	// example:
	//
	// 8080345d-b28a-9e2e-9ad7-370f8e236949
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s AssetOptimizeProResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s AssetOptimizeProResponseBodyData) GoString() string {
	return s.String()
}

func (s *AssetOptimizeProResponseBodyData) GetTaskId() *string {
	return s.TaskId
}

func (s *AssetOptimizeProResponseBodyData) SetTaskId(v string) *AssetOptimizeProResponseBodyData {
	s.TaskId = &v
	return s
}

func (s *AssetOptimizeProResponseBodyData) Validate() error {
	return dara.Validate(s)
}
