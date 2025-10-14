// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeployInstanceSDGResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *DeployInstanceSDGResponseBodyData) *DeployInstanceSDGResponseBody
	GetData() *DeployInstanceSDGResponseBodyData
	SetRequestId(v string) *DeployInstanceSDGResponseBody
	GetRequestId() *string
}

type DeployInstanceSDGResponseBody struct {
	// The returned data object.
	Data *DeployInstanceSDGResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 125B04C7-3D0D-4245-AF96-14E3758E3F06
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeployInstanceSDGResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeployInstanceSDGResponseBody) GoString() string {
	return s.String()
}

func (s *DeployInstanceSDGResponseBody) GetData() *DeployInstanceSDGResponseBodyData {
	return s.Data
}

func (s *DeployInstanceSDGResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeployInstanceSDGResponseBody) SetData(v *DeployInstanceSDGResponseBodyData) *DeployInstanceSDGResponseBody {
	s.Data = v
	return s
}

func (s *DeployInstanceSDGResponseBody) SetRequestId(v string) *DeployInstanceSDGResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeployInstanceSDGResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DeployInstanceSDGResponseBodyData struct {
	// The response message. Success is returned for a successful request.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The execution result of the synchronization request.
	Result *DeployInstanceSDGResponseBodyDataResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
	// Indicates whether all tasks are successful. Valid values:
	//
	// 	- true: All tasks are successful.
	//
	// 	- false: Failed tasks exist.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeployInstanceSDGResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DeployInstanceSDGResponseBodyData) GoString() string {
	return s.String()
}

func (s *DeployInstanceSDGResponseBodyData) GetMessage() *string {
	return s.Message
}

func (s *DeployInstanceSDGResponseBodyData) GetResult() *DeployInstanceSDGResponseBodyDataResult {
	return s.Result
}

func (s *DeployInstanceSDGResponseBodyData) GetSuccess() *bool {
	return s.Success
}

func (s *DeployInstanceSDGResponseBodyData) SetMessage(v string) *DeployInstanceSDGResponseBodyData {
	s.Message = &v
	return s
}

func (s *DeployInstanceSDGResponseBodyData) SetResult(v *DeployInstanceSDGResponseBodyDataResult) *DeployInstanceSDGResponseBodyData {
	s.Result = v
	return s
}

func (s *DeployInstanceSDGResponseBodyData) SetSuccess(v bool) *DeployInstanceSDGResponseBodyData {
	s.Success = &v
	return s
}

func (s *DeployInstanceSDGResponseBodyData) Validate() error {
	if s.Result != nil {
		if err := s.Result.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DeployInstanceSDGResponseBodyDataResult struct {
	// The number of failed tasks.
	//
	// example:
	//
	// 0
	FailedCount *int64 `json:"FailedCount,omitempty" xml:"FailedCount,omitempty"`
	// Details about the failed tasks.
	FailedItems []*DeployInstanceSDGResponseBodyDataResultFailedItems `json:"FailedItems,omitempty" xml:"FailedItems,omitempty" type:"Repeated"`
	// The number of successful tasks.
	//
	// example:
	//
	// 1
	SuccessCount *int64 `json:"SuccessCount,omitempty" xml:"SuccessCount,omitempty"`
}

func (s DeployInstanceSDGResponseBodyDataResult) String() string {
	return dara.Prettify(s)
}

func (s DeployInstanceSDGResponseBodyDataResult) GoString() string {
	return s.String()
}

func (s *DeployInstanceSDGResponseBodyDataResult) GetFailedCount() *int64 {
	return s.FailedCount
}

func (s *DeployInstanceSDGResponseBodyDataResult) GetFailedItems() []*DeployInstanceSDGResponseBodyDataResultFailedItems {
	return s.FailedItems
}

func (s *DeployInstanceSDGResponseBodyDataResult) GetSuccessCount() *int64 {
	return s.SuccessCount
}

func (s *DeployInstanceSDGResponseBodyDataResult) SetFailedCount(v int64) *DeployInstanceSDGResponseBodyDataResult {
	s.FailedCount = &v
	return s
}

func (s *DeployInstanceSDGResponseBodyDataResult) SetFailedItems(v []*DeployInstanceSDGResponseBodyDataResultFailedItems) *DeployInstanceSDGResponseBodyDataResult {
	s.FailedItems = v
	return s
}

func (s *DeployInstanceSDGResponseBodyDataResult) SetSuccessCount(v int64) *DeployInstanceSDGResponseBodyDataResult {
	s.SuccessCount = &v
	return s
}

func (s *DeployInstanceSDGResponseBodyDataResult) Validate() error {
	if s.FailedItems != nil {
		for _, item := range s.FailedItems {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DeployInstanceSDGResponseBodyDataResultFailedItems struct {
	// The error message that is returned.
	//
	// example:
	//
	// sdg not found
	ErrMessage *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	// The ID of the instance.
	//
	// example:
	//
	// aic-xxxxx-0
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s DeployInstanceSDGResponseBodyDataResultFailedItems) String() string {
	return dara.Prettify(s)
}

func (s DeployInstanceSDGResponseBodyDataResultFailedItems) GoString() string {
	return s.String()
}

func (s *DeployInstanceSDGResponseBodyDataResultFailedItems) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *DeployInstanceSDGResponseBodyDataResultFailedItems) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DeployInstanceSDGResponseBodyDataResultFailedItems) SetErrMessage(v string) *DeployInstanceSDGResponseBodyDataResultFailedItems {
	s.ErrMessage = &v
	return s
}

func (s *DeployInstanceSDGResponseBodyDataResultFailedItems) SetInstanceId(v string) *DeployInstanceSDGResponseBodyDataResultFailedItems {
	s.InstanceId = &v
	return s
}

func (s *DeployInstanceSDGResponseBodyDataResultFailedItems) Validate() error {
	return dara.Validate(s)
}
