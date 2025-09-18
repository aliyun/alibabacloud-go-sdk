// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPublishedMmAppResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int32) *ListPublishedMmAppResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *ListPublishedMmAppResponseBody
	GetPageSize() *int32
	SetPublishedVersionInfoList(v []*ListPublishedMmAppResponseBodyPublishedVersionInfoList) *ListPublishedMmAppResponseBody
	GetPublishedVersionInfoList() []*ListPublishedMmAppResponseBodyPublishedVersionInfoList
	SetRequestId(v string) *ListPublishedMmAppResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *ListPublishedMmAppResponseBody
	GetTotalCount() *int32
}

type ListPublishedMmAppResponseBody struct {
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize                 *int32                                                    `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PublishedVersionInfoList []*ListPublishedMmAppResponseBodyPublishedVersionInfoList `json:"PublishedVersionInfoList,omitempty" xml:"PublishedVersionInfoList,omitempty" type:"Repeated"`
	// example:
	//
	// xxxxxx
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 10
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListPublishedMmAppResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListPublishedMmAppResponseBody) GoString() string {
	return s.String()
}

func (s *ListPublishedMmAppResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListPublishedMmAppResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListPublishedMmAppResponseBody) GetPublishedVersionInfoList() []*ListPublishedMmAppResponseBodyPublishedVersionInfoList {
	return s.PublishedVersionInfoList
}

func (s *ListPublishedMmAppResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListPublishedMmAppResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListPublishedMmAppResponseBody) SetPageNumber(v int32) *ListPublishedMmAppResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListPublishedMmAppResponseBody) SetPageSize(v int32) *ListPublishedMmAppResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListPublishedMmAppResponseBody) SetPublishedVersionInfoList(v []*ListPublishedMmAppResponseBodyPublishedVersionInfoList) *ListPublishedMmAppResponseBody {
	s.PublishedVersionInfoList = v
	return s
}

func (s *ListPublishedMmAppResponseBody) SetRequestId(v string) *ListPublishedMmAppResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListPublishedMmAppResponseBody) SetTotalCount(v int32) *ListPublishedMmAppResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListPublishedMmAppResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListPublishedMmAppResponseBodyPublishedVersionInfoList struct {
	// example:
	//
	// mm_xxx
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// example:
	//
	// 多模态应用xxx
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// example:
	//
	// 234343
	CreateUserId *string `json:"CreateUserId,omitempty" xml:"CreateUserId,omitempty"`
	// example:
	//
	// ccccc
	CreateUserName *string `json:"CreateUserName,omitempty" xml:"CreateUserName,omitempty"`
	// example:
	//
	// xxxx
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// xxxx
	GmtCreate *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// example:
	//
	// xxxx
	PublishTime *string `json:"PublishTime,omitempty" xml:"PublishTime,omitempty"`
	// example:
	//
	// 1
	Version *int64 `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s ListPublishedMmAppResponseBodyPublishedVersionInfoList) String() string {
	return dara.Prettify(s)
}

func (s ListPublishedMmAppResponseBodyPublishedVersionInfoList) GoString() string {
	return s.String()
}

func (s *ListPublishedMmAppResponseBodyPublishedVersionInfoList) GetAppId() *string {
	return s.AppId
}

func (s *ListPublishedMmAppResponseBodyPublishedVersionInfoList) GetAppName() *string {
	return s.AppName
}

func (s *ListPublishedMmAppResponseBodyPublishedVersionInfoList) GetCreateUserId() *string {
	return s.CreateUserId
}

func (s *ListPublishedMmAppResponseBodyPublishedVersionInfoList) GetCreateUserName() *string {
	return s.CreateUserName
}

func (s *ListPublishedMmAppResponseBodyPublishedVersionInfoList) GetDescription() *string {
	return s.Description
}

func (s *ListPublishedMmAppResponseBodyPublishedVersionInfoList) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *ListPublishedMmAppResponseBodyPublishedVersionInfoList) GetPublishTime() *string {
	return s.PublishTime
}

func (s *ListPublishedMmAppResponseBodyPublishedVersionInfoList) GetVersion() *int64 {
	return s.Version
}

func (s *ListPublishedMmAppResponseBodyPublishedVersionInfoList) SetAppId(v string) *ListPublishedMmAppResponseBodyPublishedVersionInfoList {
	s.AppId = &v
	return s
}

func (s *ListPublishedMmAppResponseBodyPublishedVersionInfoList) SetAppName(v string) *ListPublishedMmAppResponseBodyPublishedVersionInfoList {
	s.AppName = &v
	return s
}

func (s *ListPublishedMmAppResponseBodyPublishedVersionInfoList) SetCreateUserId(v string) *ListPublishedMmAppResponseBodyPublishedVersionInfoList {
	s.CreateUserId = &v
	return s
}

func (s *ListPublishedMmAppResponseBodyPublishedVersionInfoList) SetCreateUserName(v string) *ListPublishedMmAppResponseBodyPublishedVersionInfoList {
	s.CreateUserName = &v
	return s
}

func (s *ListPublishedMmAppResponseBodyPublishedVersionInfoList) SetDescription(v string) *ListPublishedMmAppResponseBodyPublishedVersionInfoList {
	s.Description = &v
	return s
}

func (s *ListPublishedMmAppResponseBodyPublishedVersionInfoList) SetGmtCreate(v string) *ListPublishedMmAppResponseBodyPublishedVersionInfoList {
	s.GmtCreate = &v
	return s
}

func (s *ListPublishedMmAppResponseBodyPublishedVersionInfoList) SetPublishTime(v string) *ListPublishedMmAppResponseBodyPublishedVersionInfoList {
	s.PublishTime = &v
	return s
}

func (s *ListPublishedMmAppResponseBodyPublishedVersionInfoList) SetVersion(v int64) *ListPublishedMmAppResponseBodyPublishedVersionInfoList {
	s.Version = &v
	return s
}

func (s *ListPublishedMmAppResponseBodyPublishedVersionInfoList) Validate() error {
	return dara.Validate(s)
}
