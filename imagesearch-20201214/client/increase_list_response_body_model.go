// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iIncreaseListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *IncreaseListResponseBodyData) *IncreaseListResponseBody
	GetData() *IncreaseListResponseBodyData
	SetRequestId(v string) *IncreaseListResponseBody
	GetRequestId() *string
}

type IncreaseListResponseBody struct {
	// The information about the batch task.
	Data *IncreaseListResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// B3137727-7D6E-488C-BA21-0E034C38A879
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s IncreaseListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s IncreaseListResponseBody) GoString() string {
	return s.String()
}

func (s *IncreaseListResponseBody) GetData() *IncreaseListResponseBodyData {
	return s.Data
}

func (s *IncreaseListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *IncreaseListResponseBody) SetData(v *IncreaseListResponseBodyData) *IncreaseListResponseBody {
	s.Data = v
	return s
}

func (s *IncreaseListResponseBody) SetRequestId(v string) *IncreaseListResponseBody {
	s.RequestId = &v
	return s
}

func (s *IncreaseListResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type IncreaseListResponseBodyData struct {
	// A list of batch tasks.
	Increments *IncreaseListResponseBodyDataIncrements `json:"Increments,omitempty" xml:"Increments,omitempty" type:"Struct"`
	// The number of the page to return.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The total number of tasks.
	//
	// example:
	//
	// 15
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s IncreaseListResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s IncreaseListResponseBodyData) GoString() string {
	return s.String()
}

func (s *IncreaseListResponseBodyData) GetIncrements() *IncreaseListResponseBodyDataIncrements {
	return s.Increments
}

func (s *IncreaseListResponseBodyData) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *IncreaseListResponseBodyData) GetPageSize() *int32 {
	return s.PageSize
}

func (s *IncreaseListResponseBodyData) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *IncreaseListResponseBodyData) SetIncrements(v *IncreaseListResponseBodyDataIncrements) *IncreaseListResponseBodyData {
	s.Increments = v
	return s
}

func (s *IncreaseListResponseBodyData) SetPageNumber(v int32) *IncreaseListResponseBodyData {
	s.PageNumber = &v
	return s
}

func (s *IncreaseListResponseBodyData) SetPageSize(v int32) *IncreaseListResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *IncreaseListResponseBodyData) SetTotalCount(v int64) *IncreaseListResponseBodyData {
	s.TotalCount = &v
	return s
}

func (s *IncreaseListResponseBodyData) Validate() error {
	if s.Increments != nil {
		if err := s.Increments.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type IncreaseListResponseBodyDataIncrements struct {
	Instance []*IncreaseListResponseBodyDataIncrementsInstance `json:"Instance,omitempty" xml:"Instance,omitempty" type:"Repeated"`
}

func (s IncreaseListResponseBodyDataIncrements) String() string {
	return dara.Prettify(s)
}

func (s IncreaseListResponseBodyDataIncrements) GoString() string {
	return s.String()
}

func (s *IncreaseListResponseBodyDataIncrements) GetInstance() []*IncreaseListResponseBodyDataIncrementsInstance {
	return s.Instance
}

func (s *IncreaseListResponseBodyDataIncrements) SetInstance(v []*IncreaseListResponseBodyDataIncrementsInstance) *IncreaseListResponseBodyDataIncrements {
	s.Instance = v
	return s
}

func (s *IncreaseListResponseBodyDataIncrements) Validate() error {
	if s.Instance != nil {
		for _, item := range s.Instance {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type IncreaseListResponseBodyDataIncrementsInstance struct {
	// The name of the Object Storage Service (OSS) bucket.
	//
	// example:
	//
	// bucketName
	BucketName *string `json:"BucketName,omitempty" xml:"BucketName,omitempty"`
	// The callback address.
	//
	// example:
	//
	// http://xxxxx
	CallbackAddress *string `json:"CallbackAddress,omitempty" xml:"CallbackAddress,omitempty"`
	// The error code returned.
	//
	// 	- A value of 0 indicates that the operation is successful.
	//
	// 	- Values other than 0 indicate errors.
	//
	// example:
	//
	// 0
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The address where you can download the result. The address is valid for 2 hours.
	//
	// example:
	//
	// https://imagesearchname.oss-cn-shanghai.aliyuncs.com/xxx
	ErrorUrl *string `json:"ErrorUrl,omitempty" xml:"ErrorUrl,omitempty"`
	// The ID of the task.
	//
	// example:
	//
	// 500
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The error message returned.
	//
	// example:
	//
	// sucess
	Msg *string `json:"Msg,omitempty" xml:"Msg,omitempty"`
	// The absolute path to the increment.meta file in the bucket. The path must start with a forward slash (/) and cannot end with a forward slash (/).
	//
	// example:
	//
	// /xx/xx
	Path *string `json:"Path,omitempty" xml:"Path,omitempty"`
	// The status of the batch task.
	//
	// 	- PROCESSING: in progress
	//
	// 	- FAIL: failed
	//
	// 	- SUCCESS: successful
	//
	// example:
	//
	// SUCCESS
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The time when the task was created. Unit: milliseconds.
	//
	// example:
	//
	// 1629095713000
	UtcCreate *string `json:"UtcCreate,omitempty" xml:"UtcCreate,omitempty"`
	// The time when the task was updated. Unit: milliseconds.
	//
	// example:
	//
	// 1629095760000
	UtcModified *int64 `json:"UtcModified,omitempty" xml:"UtcModified,omitempty"`
}

func (s IncreaseListResponseBodyDataIncrementsInstance) String() string {
	return dara.Prettify(s)
}

func (s IncreaseListResponseBodyDataIncrementsInstance) GoString() string {
	return s.String()
}

func (s *IncreaseListResponseBodyDataIncrementsInstance) GetBucketName() *string {
	return s.BucketName
}

func (s *IncreaseListResponseBodyDataIncrementsInstance) GetCallbackAddress() *string {
	return s.CallbackAddress
}

func (s *IncreaseListResponseBodyDataIncrementsInstance) GetCode() *string {
	return s.Code
}

func (s *IncreaseListResponseBodyDataIncrementsInstance) GetErrorUrl() *string {
	return s.ErrorUrl
}

func (s *IncreaseListResponseBodyDataIncrementsInstance) GetId() *int64 {
	return s.Id
}

func (s *IncreaseListResponseBodyDataIncrementsInstance) GetMsg() *string {
	return s.Msg
}

func (s *IncreaseListResponseBodyDataIncrementsInstance) GetPath() *string {
	return s.Path
}

func (s *IncreaseListResponseBodyDataIncrementsInstance) GetStatus() *string {
	return s.Status
}

func (s *IncreaseListResponseBodyDataIncrementsInstance) GetUtcCreate() *string {
	return s.UtcCreate
}

func (s *IncreaseListResponseBodyDataIncrementsInstance) GetUtcModified() *int64 {
	return s.UtcModified
}

func (s *IncreaseListResponseBodyDataIncrementsInstance) SetBucketName(v string) *IncreaseListResponseBodyDataIncrementsInstance {
	s.BucketName = &v
	return s
}

func (s *IncreaseListResponseBodyDataIncrementsInstance) SetCallbackAddress(v string) *IncreaseListResponseBodyDataIncrementsInstance {
	s.CallbackAddress = &v
	return s
}

func (s *IncreaseListResponseBodyDataIncrementsInstance) SetCode(v string) *IncreaseListResponseBodyDataIncrementsInstance {
	s.Code = &v
	return s
}

func (s *IncreaseListResponseBodyDataIncrementsInstance) SetErrorUrl(v string) *IncreaseListResponseBodyDataIncrementsInstance {
	s.ErrorUrl = &v
	return s
}

func (s *IncreaseListResponseBodyDataIncrementsInstance) SetId(v int64) *IncreaseListResponseBodyDataIncrementsInstance {
	s.Id = &v
	return s
}

func (s *IncreaseListResponseBodyDataIncrementsInstance) SetMsg(v string) *IncreaseListResponseBodyDataIncrementsInstance {
	s.Msg = &v
	return s
}

func (s *IncreaseListResponseBodyDataIncrementsInstance) SetPath(v string) *IncreaseListResponseBodyDataIncrementsInstance {
	s.Path = &v
	return s
}

func (s *IncreaseListResponseBodyDataIncrementsInstance) SetStatus(v string) *IncreaseListResponseBodyDataIncrementsInstance {
	s.Status = &v
	return s
}

func (s *IncreaseListResponseBodyDataIncrementsInstance) SetUtcCreate(v string) *IncreaseListResponseBodyDataIncrementsInstance {
	s.UtcCreate = &v
	return s
}

func (s *IncreaseListResponseBodyDataIncrementsInstance) SetUtcModified(v int64) *IncreaseListResponseBodyDataIncrementsInstance {
	s.UtcModified = &v
	return s
}

func (s *IncreaseListResponseBodyDataIncrementsInstance) Validate() error {
	return dara.Validate(s)
}
