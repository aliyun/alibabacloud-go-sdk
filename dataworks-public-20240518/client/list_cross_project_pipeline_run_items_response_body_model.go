// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCrossProjectPipelineRunItemsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *ListCrossProjectPipelineRunItemsResponseBodyData) *ListCrossProjectPipelineRunItemsResponseBody
	GetData() *ListCrossProjectPipelineRunItemsResponseBodyData
	SetRequestId(v string) *ListCrossProjectPipelineRunItemsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListCrossProjectPipelineRunItemsResponseBody
	GetSuccess() *bool
}

type ListCrossProjectPipelineRunItemsResponseBody struct {
	// The business response.
	//
	// example:
	//
	// {"RequestId":"735894D1-D5E5-50B8-8A6D-041C90A98B23","PageNumber":1,"PageSize":10,"TotalCount":1,"PipelineRunItems":[{"ObjectId":"1","ObjectType":"ODPS_SQL","ObjectName":"object-1","ObjectVersion":"7","ChangeType":"ADD","IsRoot":true,"Status":"Ready"}]}
	Data *ListCrossProjectPipelineRunItemsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The request ID, which is used to locate and troubleshoot this API call.
	//
	// example:
	//
	// 735894D1-D5E5-50B8-8A6D-041C90A98B23
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListCrossProjectPipelineRunItemsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListCrossProjectPipelineRunItemsResponseBody) GoString() string {
	return s.String()
}

func (s *ListCrossProjectPipelineRunItemsResponseBody) GetData() *ListCrossProjectPipelineRunItemsResponseBodyData {
	return s.Data
}

func (s *ListCrossProjectPipelineRunItemsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListCrossProjectPipelineRunItemsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListCrossProjectPipelineRunItemsResponseBody) SetData(v *ListCrossProjectPipelineRunItemsResponseBodyData) *ListCrossProjectPipelineRunItemsResponseBody {
	s.Data = v
	return s
}

func (s *ListCrossProjectPipelineRunItemsResponseBody) SetRequestId(v string) *ListCrossProjectPipelineRunItemsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListCrossProjectPipelineRunItemsResponseBody) SetSuccess(v bool) *ListCrossProjectPipelineRunItemsResponseBody {
	s.Success = &v
	return s
}

func (s *ListCrossProjectPipelineRunItemsResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListCrossProjectPipelineRunItemsResponseBodyData struct {
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
	// The list of publish items for the root objects and their child objects that are included in the cross-workspace publish pipeline.
	//
	// example:
	//
	// [{"ObjectId":"1","ObjectType":"ODPS_SQL","ObjectName":"object-1","ObjectVersion":"7","ChangeType":"ADD","IsRoot":true,"Status":"Ready"}]
	PipelineRunItems []*ListCrossProjectPipelineRunItemsResponseBodyDataPipelineRunItems `json:"PipelineRunItems,omitempty" xml:"PipelineRunItems,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 735894D1-D5E5-50B8-8A6D-041C90A98B23
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of records.
	//
	// example:
	//
	// 1
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListCrossProjectPipelineRunItemsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListCrossProjectPipelineRunItemsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListCrossProjectPipelineRunItemsResponseBodyData) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListCrossProjectPipelineRunItemsResponseBodyData) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListCrossProjectPipelineRunItemsResponseBodyData) GetPipelineRunItems() []*ListCrossProjectPipelineRunItemsResponseBodyDataPipelineRunItems {
	return s.PipelineRunItems
}

func (s *ListCrossProjectPipelineRunItemsResponseBodyData) GetRequestId() *string {
	return s.RequestId
}

func (s *ListCrossProjectPipelineRunItemsResponseBodyData) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListCrossProjectPipelineRunItemsResponseBodyData) SetPageNumber(v int32) *ListCrossProjectPipelineRunItemsResponseBodyData {
	s.PageNumber = &v
	return s
}

func (s *ListCrossProjectPipelineRunItemsResponseBodyData) SetPageSize(v int32) *ListCrossProjectPipelineRunItemsResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListCrossProjectPipelineRunItemsResponseBodyData) SetPipelineRunItems(v []*ListCrossProjectPipelineRunItemsResponseBodyDataPipelineRunItems) *ListCrossProjectPipelineRunItemsResponseBodyData {
	s.PipelineRunItems = v
	return s
}

func (s *ListCrossProjectPipelineRunItemsResponseBodyData) SetRequestId(v string) *ListCrossProjectPipelineRunItemsResponseBodyData {
	s.RequestId = &v
	return s
}

func (s *ListCrossProjectPipelineRunItemsResponseBodyData) SetTotalCount(v int32) *ListCrossProjectPipelineRunItemsResponseBodyData {
	s.TotalCount = &v
	return s
}

func (s *ListCrossProjectPipelineRunItemsResponseBodyData) Validate() error {
	if s.PipelineRunItems != nil {
		for _, item := range s.PipelineRunItems {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListCrossProjectPipelineRunItemsResponseBodyDataPipelineRunItems struct {
	// The change type.
	//
	// example:
	//
	// ADD
	ChangeType *string `json:"ChangeType,omitempty" xml:"ChangeType,omitempty"`
	// The error code.
	//
	// example:
	//
	// DeploymentItemFailed
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message.
	//
	// example:
	//
	// Deployment failed
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// Indicates whether the object is a root object.
	//
	// example:
	//
	// true
	IsRoot *bool `json:"IsRoot,omitempty" xml:"IsRoot,omitempty"`
	// The ID of the publish object.
	//
	// example:
	//
	// 1
	ObjectId *string `json:"ObjectId,omitempty" xml:"ObjectId,omitempty"`
	// The name of the publish object.
	//
	// example:
	//
	// object-1
	ObjectName *string `json:"ObjectName,omitempty" xml:"ObjectName,omitempty"`
	// The object type of the publish object.
	//
	// example:
	//
	// ODPS_SQL
	ObjectType *string `json:"ObjectType,omitempty" xml:"ObjectType,omitempty"`
	// The version of the publish object.
	//
	// example:
	//
	// 7
	ObjectVersion *string `json:"ObjectVersion,omitempty" xml:"ObjectVersion,omitempty"`
	// The ID of the parent object.
	//
	// example:
	//
	// 1
	ParentObjectId *string `json:"ParentObjectId,omitempty" xml:"ParentObjectId,omitempty"`
	// The status of the publish item.
	//
	// example:
	//
	// Ready
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListCrossProjectPipelineRunItemsResponseBodyDataPipelineRunItems) String() string {
	return dara.Prettify(s)
}

func (s ListCrossProjectPipelineRunItemsResponseBodyDataPipelineRunItems) GoString() string {
	return s.String()
}

func (s *ListCrossProjectPipelineRunItemsResponseBodyDataPipelineRunItems) GetChangeType() *string {
	return s.ChangeType
}

func (s *ListCrossProjectPipelineRunItemsResponseBodyDataPipelineRunItems) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *ListCrossProjectPipelineRunItemsResponseBodyDataPipelineRunItems) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *ListCrossProjectPipelineRunItemsResponseBodyDataPipelineRunItems) GetIsRoot() *bool {
	return s.IsRoot
}

func (s *ListCrossProjectPipelineRunItemsResponseBodyDataPipelineRunItems) GetObjectId() *string {
	return s.ObjectId
}

func (s *ListCrossProjectPipelineRunItemsResponseBodyDataPipelineRunItems) GetObjectName() *string {
	return s.ObjectName
}

func (s *ListCrossProjectPipelineRunItemsResponseBodyDataPipelineRunItems) GetObjectType() *string {
	return s.ObjectType
}

func (s *ListCrossProjectPipelineRunItemsResponseBodyDataPipelineRunItems) GetObjectVersion() *string {
	return s.ObjectVersion
}

func (s *ListCrossProjectPipelineRunItemsResponseBodyDataPipelineRunItems) GetParentObjectId() *string {
	return s.ParentObjectId
}

func (s *ListCrossProjectPipelineRunItemsResponseBodyDataPipelineRunItems) GetStatus() *string {
	return s.Status
}

func (s *ListCrossProjectPipelineRunItemsResponseBodyDataPipelineRunItems) SetChangeType(v string) *ListCrossProjectPipelineRunItemsResponseBodyDataPipelineRunItems {
	s.ChangeType = &v
	return s
}

func (s *ListCrossProjectPipelineRunItemsResponseBodyDataPipelineRunItems) SetErrorCode(v string) *ListCrossProjectPipelineRunItemsResponseBodyDataPipelineRunItems {
	s.ErrorCode = &v
	return s
}

func (s *ListCrossProjectPipelineRunItemsResponseBodyDataPipelineRunItems) SetErrorMessage(v string) *ListCrossProjectPipelineRunItemsResponseBodyDataPipelineRunItems {
	s.ErrorMessage = &v
	return s
}

func (s *ListCrossProjectPipelineRunItemsResponseBodyDataPipelineRunItems) SetIsRoot(v bool) *ListCrossProjectPipelineRunItemsResponseBodyDataPipelineRunItems {
	s.IsRoot = &v
	return s
}

func (s *ListCrossProjectPipelineRunItemsResponseBodyDataPipelineRunItems) SetObjectId(v string) *ListCrossProjectPipelineRunItemsResponseBodyDataPipelineRunItems {
	s.ObjectId = &v
	return s
}

func (s *ListCrossProjectPipelineRunItemsResponseBodyDataPipelineRunItems) SetObjectName(v string) *ListCrossProjectPipelineRunItemsResponseBodyDataPipelineRunItems {
	s.ObjectName = &v
	return s
}

func (s *ListCrossProjectPipelineRunItemsResponseBodyDataPipelineRunItems) SetObjectType(v string) *ListCrossProjectPipelineRunItemsResponseBodyDataPipelineRunItems {
	s.ObjectType = &v
	return s
}

func (s *ListCrossProjectPipelineRunItemsResponseBodyDataPipelineRunItems) SetObjectVersion(v string) *ListCrossProjectPipelineRunItemsResponseBodyDataPipelineRunItems {
	s.ObjectVersion = &v
	return s
}

func (s *ListCrossProjectPipelineRunItemsResponseBodyDataPipelineRunItems) SetParentObjectId(v string) *ListCrossProjectPipelineRunItemsResponseBodyDataPipelineRunItems {
	s.ParentObjectId = &v
	return s
}

func (s *ListCrossProjectPipelineRunItemsResponseBodyDataPipelineRunItems) SetStatus(v string) *ListCrossProjectPipelineRunItemsResponseBodyDataPipelineRunItems {
	s.Status = &v
	return s
}

func (s *ListCrossProjectPipelineRunItemsResponseBodyDataPipelineRunItems) Validate() error {
	return dara.Validate(s)
}
