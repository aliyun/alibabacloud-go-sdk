// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAssetOptimizeLiteResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *AssetOptimizeLiteResponseBody
	GetCode() *string
	SetData(v *AssetOptimizeLiteResponseBodyData) *AssetOptimizeLiteResponseBody
	GetData() *AssetOptimizeLiteResponseBodyData
	SetMessage(v string) *AssetOptimizeLiteResponseBody
	GetMessage() *string
	SetRequestId(v string) *AssetOptimizeLiteResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *AssetOptimizeLiteResponseBody
	GetSuccess() *bool
}

type AssetOptimizeLiteResponseBody struct {
	// Error code
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Asynchronous task submission result
	Data *AssetOptimizeLiteResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// Error message
	//
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// E1AD60F1-BAC7-546B-9533-E7AD02B16E3F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Whether the call was successful
	//
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s AssetOptimizeLiteResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AssetOptimizeLiteResponseBody) GoString() string {
	return s.String()
}

func (s *AssetOptimizeLiteResponseBody) GetCode() *string {
	return s.Code
}

func (s *AssetOptimizeLiteResponseBody) GetData() *AssetOptimizeLiteResponseBodyData {
	return s.Data
}

func (s *AssetOptimizeLiteResponseBody) GetMessage() *string {
	return s.Message
}

func (s *AssetOptimizeLiteResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AssetOptimizeLiteResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *AssetOptimizeLiteResponseBody) SetCode(v string) *AssetOptimizeLiteResponseBody {
	s.Code = &v
	return s
}

func (s *AssetOptimizeLiteResponseBody) SetData(v *AssetOptimizeLiteResponseBodyData) *AssetOptimizeLiteResponseBody {
	s.Data = v
	return s
}

func (s *AssetOptimizeLiteResponseBody) SetMessage(v string) *AssetOptimizeLiteResponseBody {
	s.Message = &v
	return s
}

func (s *AssetOptimizeLiteResponseBody) SetRequestId(v string) *AssetOptimizeLiteResponseBody {
	s.RequestId = &v
	return s
}

func (s *AssetOptimizeLiteResponseBody) SetSuccess(v bool) *AssetOptimizeLiteResponseBody {
	s.Success = &v
	return s
}

func (s *AssetOptimizeLiteResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type AssetOptimizeLiteResponseBodyData struct {
	// Asynchronous task ID, used for subsequent result queries
	//
	// example:
	//
	// task-xxxxxx
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s AssetOptimizeLiteResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s AssetOptimizeLiteResponseBodyData) GoString() string {
	return s.String()
}

func (s *AssetOptimizeLiteResponseBodyData) GetTaskId() *string {
	return s.TaskId
}

func (s *AssetOptimizeLiteResponseBodyData) SetTaskId(v string) *AssetOptimizeLiteResponseBodyData {
	s.TaskId = &v
	return s
}

func (s *AssetOptimizeLiteResponseBodyData) Validate() error {
	return dara.Validate(s)
}
