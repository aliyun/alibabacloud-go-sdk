// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePreCheckCreateGadOrderResultResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDynamicCode(v string) *DescribePreCheckCreateGadOrderResultResponseBody
	GetDynamicCode() *string
	SetDynamicMessage(v string) *DescribePreCheckCreateGadOrderResultResponseBody
	GetDynamicMessage() *string
	SetErrCode(v string) *DescribePreCheckCreateGadOrderResultResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *DescribePreCheckCreateGadOrderResultResponseBody
	GetErrMessage() *string
	SetHttpStatusCode(v string) *DescribePreCheckCreateGadOrderResultResponseBody
	GetHttpStatusCode() *string
	SetInstanceId(v string) *DescribePreCheckCreateGadOrderResultResponseBody
	GetInstanceId() *string
	SetPreCheckItems(v *DescribePreCheckCreateGadOrderResultResponseBodyPreCheckItems) *DescribePreCheckCreateGadOrderResultResponseBody
	GetPreCheckItems() *DescribePreCheckCreateGadOrderResultResponseBodyPreCheckItems
	SetPreCheckResult(v bool) *DescribePreCheckCreateGadOrderResultResponseBody
	GetPreCheckResult() *bool
	SetRegionId(v string) *DescribePreCheckCreateGadOrderResultResponseBody
	GetRegionId() *string
	SetRequestId(v string) *DescribePreCheckCreateGadOrderResultResponseBody
	GetRequestId() *string
	SetSuccess(v string) *DescribePreCheckCreateGadOrderResultResponseBody
	GetSuccess() *string
	SetTaskId(v string) *DescribePreCheckCreateGadOrderResultResponseBody
	GetTaskId() *string
}

type DescribePreCheckCreateGadOrderResultResponseBody struct {
	// example:
	//
	// 403
	DynamicCode *string `json:"DynamicCode,omitempty" xml:"DynamicCode,omitempty"`
	// example:
	//
	// present environment is not support,so skip.
	DynamicMessage *string `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
	// example:
	//
	// InternalError
	ErrCode *string `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	// example:
	//
	// The request processing has failed due to some unknown error.
	ErrMessage *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *string `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// gad-bp1i99e8l7913****
	InstanceId    *string                                                        `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	PreCheckItems *DescribePreCheckCreateGadOrderResultResponseBodyPreCheckItems `json:"PreCheckItems,omitempty" xml:"PreCheckItems,omitempty" type:"Struct"`
	// example:
	//
	// True
	PreCheckResult *bool `json:"PreCheckResult,omitempty" xml:"PreCheckResult,omitempty"`
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// 92E1E99D-5224-4AD3-8C94-23A3516B****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
	// example:
	//
	// 11****
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s DescribePreCheckCreateGadOrderResultResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribePreCheckCreateGadOrderResultResponseBody) GoString() string {
	return s.String()
}

func (s *DescribePreCheckCreateGadOrderResultResponseBody) GetDynamicCode() *string {
	return s.DynamicCode
}

func (s *DescribePreCheckCreateGadOrderResultResponseBody) GetDynamicMessage() *string {
	return s.DynamicMessage
}

func (s *DescribePreCheckCreateGadOrderResultResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *DescribePreCheckCreateGadOrderResultResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *DescribePreCheckCreateGadOrderResultResponseBody) GetHttpStatusCode() *string {
	return s.HttpStatusCode
}

func (s *DescribePreCheckCreateGadOrderResultResponseBody) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribePreCheckCreateGadOrderResultResponseBody) GetPreCheckItems() *DescribePreCheckCreateGadOrderResultResponseBodyPreCheckItems {
	return s.PreCheckItems
}

func (s *DescribePreCheckCreateGadOrderResultResponseBody) GetPreCheckResult() *bool {
	return s.PreCheckResult
}

func (s *DescribePreCheckCreateGadOrderResultResponseBody) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribePreCheckCreateGadOrderResultResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribePreCheckCreateGadOrderResultResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *DescribePreCheckCreateGadOrderResultResponseBody) GetTaskId() *string {
	return s.TaskId
}

func (s *DescribePreCheckCreateGadOrderResultResponseBody) SetDynamicCode(v string) *DescribePreCheckCreateGadOrderResultResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *DescribePreCheckCreateGadOrderResultResponseBody) SetDynamicMessage(v string) *DescribePreCheckCreateGadOrderResultResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *DescribePreCheckCreateGadOrderResultResponseBody) SetErrCode(v string) *DescribePreCheckCreateGadOrderResultResponseBody {
	s.ErrCode = &v
	return s
}

func (s *DescribePreCheckCreateGadOrderResultResponseBody) SetErrMessage(v string) *DescribePreCheckCreateGadOrderResultResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *DescribePreCheckCreateGadOrderResultResponseBody) SetHttpStatusCode(v string) *DescribePreCheckCreateGadOrderResultResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DescribePreCheckCreateGadOrderResultResponseBody) SetInstanceId(v string) *DescribePreCheckCreateGadOrderResultResponseBody {
	s.InstanceId = &v
	return s
}

func (s *DescribePreCheckCreateGadOrderResultResponseBody) SetPreCheckItems(v *DescribePreCheckCreateGadOrderResultResponseBodyPreCheckItems) *DescribePreCheckCreateGadOrderResultResponseBody {
	s.PreCheckItems = v
	return s
}

func (s *DescribePreCheckCreateGadOrderResultResponseBody) SetPreCheckResult(v bool) *DescribePreCheckCreateGadOrderResultResponseBody {
	s.PreCheckResult = &v
	return s
}

func (s *DescribePreCheckCreateGadOrderResultResponseBody) SetRegionId(v string) *DescribePreCheckCreateGadOrderResultResponseBody {
	s.RegionId = &v
	return s
}

func (s *DescribePreCheckCreateGadOrderResultResponseBody) SetRequestId(v string) *DescribePreCheckCreateGadOrderResultResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribePreCheckCreateGadOrderResultResponseBody) SetSuccess(v string) *DescribePreCheckCreateGadOrderResultResponseBody {
	s.Success = &v
	return s
}

func (s *DescribePreCheckCreateGadOrderResultResponseBody) SetTaskId(v string) *DescribePreCheckCreateGadOrderResultResponseBody {
	s.TaskId = &v
	return s
}

func (s *DescribePreCheckCreateGadOrderResultResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribePreCheckCreateGadOrderResultResponseBodyPreCheckItems struct {
	PreCheckItems []*DescribePreCheckCreateGadOrderResultResponseBodyPreCheckItemsPreCheckItems `json:"PreCheckItems,omitempty" xml:"PreCheckItems,omitempty" type:"Repeated"`
}

func (s DescribePreCheckCreateGadOrderResultResponseBodyPreCheckItems) String() string {
	return dara.Prettify(s)
}

func (s DescribePreCheckCreateGadOrderResultResponseBodyPreCheckItems) GoString() string {
	return s.String()
}

func (s *DescribePreCheckCreateGadOrderResultResponseBodyPreCheckItems) GetPreCheckItems() []*DescribePreCheckCreateGadOrderResultResponseBodyPreCheckItemsPreCheckItems {
	return s.PreCheckItems
}

func (s *DescribePreCheckCreateGadOrderResultResponseBodyPreCheckItems) SetPreCheckItems(v []*DescribePreCheckCreateGadOrderResultResponseBodyPreCheckItemsPreCheckItems) *DescribePreCheckCreateGadOrderResultResponseBodyPreCheckItems {
	s.PreCheckItems = v
	return s
}

func (s *DescribePreCheckCreateGadOrderResultResponseBodyPreCheckItems) Validate() error {
	return dara.Validate(s)
}

type DescribePreCheckCreateGadOrderResultResponseBodyPreCheckItemsPreCheckItems struct {
	// example:
	//
	// CHECK_MASTER_DB_STATUS
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// test
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// Running
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribePreCheckCreateGadOrderResultResponseBodyPreCheckItemsPreCheckItems) String() string {
	return dara.Prettify(s)
}

func (s DescribePreCheckCreateGadOrderResultResponseBodyPreCheckItemsPreCheckItems) GoString() string {
	return s.String()
}

func (s *DescribePreCheckCreateGadOrderResultResponseBodyPreCheckItemsPreCheckItems) GetCode() *string {
	return s.Code
}

func (s *DescribePreCheckCreateGadOrderResultResponseBodyPreCheckItemsPreCheckItems) GetMessage() *string {
	return s.Message
}

func (s *DescribePreCheckCreateGadOrderResultResponseBodyPreCheckItemsPreCheckItems) GetStatus() *string {
	return s.Status
}

func (s *DescribePreCheckCreateGadOrderResultResponseBodyPreCheckItemsPreCheckItems) SetCode(v string) *DescribePreCheckCreateGadOrderResultResponseBodyPreCheckItemsPreCheckItems {
	s.Code = &v
	return s
}

func (s *DescribePreCheckCreateGadOrderResultResponseBodyPreCheckItemsPreCheckItems) SetMessage(v string) *DescribePreCheckCreateGadOrderResultResponseBodyPreCheckItemsPreCheckItems {
	s.Message = &v
	return s
}

func (s *DescribePreCheckCreateGadOrderResultResponseBodyPreCheckItemsPreCheckItems) SetStatus(v string) *DescribePreCheckCreateGadOrderResultResponseBodyPreCheckItemsPreCheckItems {
	s.Status = &v
	return s
}

func (s *DescribePreCheckCreateGadOrderResultResponseBodyPreCheckItemsPreCheckItems) Validate() error {
	return dara.Validate(s)
}
