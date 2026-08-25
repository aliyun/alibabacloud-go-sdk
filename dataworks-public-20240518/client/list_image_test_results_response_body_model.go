// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListImageTestResultsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPagingInfo(v *ListImageTestResultsResponseBodyPagingInfo) *ListImageTestResultsResponseBody
	GetPagingInfo() *ListImageTestResultsResponseBodyPagingInfo
	SetRequestId(v string) *ListImageTestResultsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListImageTestResultsResponseBody
	GetSuccess() *bool
}

type ListImageTestResultsResponseBody struct {
	// The pagination information.
	PagingInfo *ListImageTestResultsResponseBodyPagingInfo `json:"PagingInfo,omitempty" xml:"PagingInfo,omitempty" type:"Struct"`
	// The request ID, which is used to locate logs and troubleshoot issues.
	//
	// example:
	//
	// 6A6CBE87-9F91-1323-B680-E7A7065XXXXX
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListImageTestResultsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListImageTestResultsResponseBody) GoString() string {
	return s.String()
}

func (s *ListImageTestResultsResponseBody) GetPagingInfo() *ListImageTestResultsResponseBodyPagingInfo {
	return s.PagingInfo
}

func (s *ListImageTestResultsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListImageTestResultsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListImageTestResultsResponseBody) SetPagingInfo(v *ListImageTestResultsResponseBodyPagingInfo) *ListImageTestResultsResponseBody {
	s.PagingInfo = v
	return s
}

func (s *ListImageTestResultsResponseBody) SetRequestId(v string) *ListImageTestResultsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListImageTestResultsResponseBody) SetSuccess(v bool) *ListImageTestResultsResponseBody {
	s.Success = &v
	return s
}

func (s *ListImageTestResultsResponseBody) Validate() error {
	if s.PagingInfo != nil {
		if err := s.PagingInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListImageTestResultsResponseBodyPagingInfo struct {
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The list of image test results.
	TestResultList []*ListImageTestResultsResponseBodyPagingInfoTestResultList `json:"TestResultList,omitempty" xml:"TestResultList,omitempty" type:"Repeated"`
	// The total number of entries.
	//
	// example:
	//
	// 10
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListImageTestResultsResponseBodyPagingInfo) String() string {
	return dara.Prettify(s)
}

func (s ListImageTestResultsResponseBodyPagingInfo) GoString() string {
	return s.String()
}

func (s *ListImageTestResultsResponseBodyPagingInfo) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListImageTestResultsResponseBodyPagingInfo) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListImageTestResultsResponseBodyPagingInfo) GetTestResultList() []*ListImageTestResultsResponseBodyPagingInfoTestResultList {
	return s.TestResultList
}

func (s *ListImageTestResultsResponseBodyPagingInfo) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListImageTestResultsResponseBodyPagingInfo) SetPageNumber(v int32) *ListImageTestResultsResponseBodyPagingInfo {
	s.PageNumber = &v
	return s
}

func (s *ListImageTestResultsResponseBodyPagingInfo) SetPageSize(v int32) *ListImageTestResultsResponseBodyPagingInfo {
	s.PageSize = &v
	return s
}

func (s *ListImageTestResultsResponseBodyPagingInfo) SetTestResultList(v []*ListImageTestResultsResponseBodyPagingInfoTestResultList) *ListImageTestResultsResponseBodyPagingInfo {
	s.TestResultList = v
	return s
}

func (s *ListImageTestResultsResponseBodyPagingInfo) SetTotalCount(v int32) *ListImageTestResultsResponseBodyPagingInfo {
	s.TotalCount = &v
	return s
}

func (s *ListImageTestResultsResponseBodyPagingInfo) Validate() error {
	if s.TestResultList != nil {
		for _, item := range s.TestResultList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListImageTestResultsResponseBodyPagingInfoTestResultList struct {
	// The image ID.
	//
	// example:
	//
	// img_123456
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	// The test result message.
	//
	// example:
	//
	// test finished
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The operation time, represented as a 64-bit timestamp.
	//
	// example:
	//
	// 1727055811000
	OperateTime *int64 `json:"OperateTime,omitempty" xml:"OperateTime,omitempty"`
	// The process ID.
	//
	// example:
	//
	// 11111111-1111-1111-1111-111111111111
	ProcessId *string `json:"ProcessId,omitempty" xml:"ProcessId,omitempty"`
	// The publish stage of the image.
	//
	// example:
	//
	// UNPUBLISHED
	PublishStage *string `json:"PublishStage,omitempty" xml:"PublishStage,omitempty"`
	// The resource group ID.
	//
	// example:
	//
	// 123456
	ResourceGroupId *int64 `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The status of the test process.
	//
	// example:
	//
	// completed
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListImageTestResultsResponseBodyPagingInfoTestResultList) String() string {
	return dara.Prettify(s)
}

func (s ListImageTestResultsResponseBodyPagingInfoTestResultList) GoString() string {
	return s.String()
}

func (s *ListImageTestResultsResponseBodyPagingInfoTestResultList) GetImageId() *string {
	return s.ImageId
}

func (s *ListImageTestResultsResponseBodyPagingInfoTestResultList) GetMessage() *string {
	return s.Message
}

func (s *ListImageTestResultsResponseBodyPagingInfoTestResultList) GetOperateTime() *int64 {
	return s.OperateTime
}

func (s *ListImageTestResultsResponseBodyPagingInfoTestResultList) GetProcessId() *string {
	return s.ProcessId
}

func (s *ListImageTestResultsResponseBodyPagingInfoTestResultList) GetPublishStage() *string {
	return s.PublishStage
}

func (s *ListImageTestResultsResponseBodyPagingInfoTestResultList) GetResourceGroupId() *int64 {
	return s.ResourceGroupId
}

func (s *ListImageTestResultsResponseBodyPagingInfoTestResultList) GetStatus() *string {
	return s.Status
}

func (s *ListImageTestResultsResponseBodyPagingInfoTestResultList) SetImageId(v string) *ListImageTestResultsResponseBodyPagingInfoTestResultList {
	s.ImageId = &v
	return s
}

func (s *ListImageTestResultsResponseBodyPagingInfoTestResultList) SetMessage(v string) *ListImageTestResultsResponseBodyPagingInfoTestResultList {
	s.Message = &v
	return s
}

func (s *ListImageTestResultsResponseBodyPagingInfoTestResultList) SetOperateTime(v int64) *ListImageTestResultsResponseBodyPagingInfoTestResultList {
	s.OperateTime = &v
	return s
}

func (s *ListImageTestResultsResponseBodyPagingInfoTestResultList) SetProcessId(v string) *ListImageTestResultsResponseBodyPagingInfoTestResultList {
	s.ProcessId = &v
	return s
}

func (s *ListImageTestResultsResponseBodyPagingInfoTestResultList) SetPublishStage(v string) *ListImageTestResultsResponseBodyPagingInfoTestResultList {
	s.PublishStage = &v
	return s
}

func (s *ListImageTestResultsResponseBodyPagingInfoTestResultList) SetResourceGroupId(v int64) *ListImageTestResultsResponseBodyPagingInfoTestResultList {
	s.ResourceGroupId = &v
	return s
}

func (s *ListImageTestResultsResponseBodyPagingInfoTestResultList) SetStatus(v string) *ListImageTestResultsResponseBodyPagingInfoTestResultList {
	s.Status = &v
	return s
}

func (s *ListImageTestResultsResponseBodyPagingInfoTestResultList) Validate() error {
	return dara.Validate(s)
}
