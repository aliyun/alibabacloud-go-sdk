// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListApiPluginsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetApiPlugins(v []*ListApiPluginsResponseBodyApiPlugins) *ListApiPluginsResponseBody
	GetApiPlugins() []*ListApiPluginsResponseBodyApiPlugins
	SetCode(v string) *ListApiPluginsResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *ListApiPluginsResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *ListApiPluginsResponseBody
	GetMessage() *string
	SetPageNumber(v int32) *ListApiPluginsResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *ListApiPluginsResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *ListApiPluginsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListApiPluginsResponseBody
	GetSuccess() *bool
	SetTotalCount(v int64) *ListApiPluginsResponseBody
	GetTotalCount() *int64
}

type ListApiPluginsResponseBody struct {
	ApiPlugins []*ListApiPluginsResponseBodyApiPlugins `json:"ApiPlugins,omitempty" xml:"ApiPlugins,omitempty" type:"Repeated"`
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// Success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 254EB995-DEDF-48A4-9101-9CA5B72FFBCC
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// example:
	//
	// 100
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListApiPluginsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListApiPluginsResponseBody) GoString() string {
	return s.String()
}

func (s *ListApiPluginsResponseBody) GetApiPlugins() []*ListApiPluginsResponseBodyApiPlugins {
	return s.ApiPlugins
}

func (s *ListApiPluginsResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListApiPluginsResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListApiPluginsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListApiPluginsResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListApiPluginsResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListApiPluginsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListApiPluginsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListApiPluginsResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListApiPluginsResponseBody) SetApiPlugins(v []*ListApiPluginsResponseBodyApiPlugins) *ListApiPluginsResponseBody {
	s.ApiPlugins = v
	return s
}

func (s *ListApiPluginsResponseBody) SetCode(v string) *ListApiPluginsResponseBody {
	s.Code = &v
	return s
}

func (s *ListApiPluginsResponseBody) SetHttpStatusCode(v int32) *ListApiPluginsResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListApiPluginsResponseBody) SetMessage(v string) *ListApiPluginsResponseBody {
	s.Message = &v
	return s
}

func (s *ListApiPluginsResponseBody) SetPageNumber(v int32) *ListApiPluginsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListApiPluginsResponseBody) SetPageSize(v int32) *ListApiPluginsResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListApiPluginsResponseBody) SetRequestId(v string) *ListApiPluginsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListApiPluginsResponseBody) SetSuccess(v bool) *ListApiPluginsResponseBody {
	s.Success = &v
	return s
}

func (s *ListApiPluginsResponseBody) SetTotalCount(v int64) *ListApiPluginsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListApiPluginsResponseBody) Validate() error {
	if s.ApiPlugins != nil {
		for _, item := range s.ApiPlugins {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListApiPluginsResponseBodyApiPlugins struct {
	Description       *string `json:"Description,omitempty" xml:"Description,omitempty"`
	DraftedConfigJson *string `json:"DraftedConfigJson,omitempty" xml:"DraftedConfigJson,omitempty"`
	// example:
	//
	// 1666230851000
	GmtCreate *int64 `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// example:
	//
	// 1641891940000
	GmtModified *int64 `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// example:
	//
	// 7c0e5b5e-a839-4999-8301-2c7d07a1f16f
	InstanceId          *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Name                *string `json:"Name,omitempty" xml:"Name,omitempty"`
	PublishedConfigJson *string `json:"PublishedConfigJson,omitempty" xml:"PublishedConfigJson,omitempty"`
	// example:
	//
	// Drafted
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// Function
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// example:
	//
	// e1a3c448-20cf-4586-8aa2-4cdca75f7c20
	Uuid *string `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
}

func (s ListApiPluginsResponseBodyApiPlugins) String() string {
	return dara.Prettify(s)
}

func (s ListApiPluginsResponseBodyApiPlugins) GoString() string {
	return s.String()
}

func (s *ListApiPluginsResponseBodyApiPlugins) GetDescription() *string {
	return s.Description
}

func (s *ListApiPluginsResponseBodyApiPlugins) GetDraftedConfigJson() *string {
	return s.DraftedConfigJson
}

func (s *ListApiPluginsResponseBodyApiPlugins) GetGmtCreate() *int64 {
	return s.GmtCreate
}

func (s *ListApiPluginsResponseBodyApiPlugins) GetGmtModified() *int64 {
	return s.GmtModified
}

func (s *ListApiPluginsResponseBodyApiPlugins) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListApiPluginsResponseBodyApiPlugins) GetName() *string {
	return s.Name
}

func (s *ListApiPluginsResponseBodyApiPlugins) GetPublishedConfigJson() *string {
	return s.PublishedConfigJson
}

func (s *ListApiPluginsResponseBodyApiPlugins) GetStatus() *string {
	return s.Status
}

func (s *ListApiPluginsResponseBodyApiPlugins) GetType() *string {
	return s.Type
}

func (s *ListApiPluginsResponseBodyApiPlugins) GetUuid() *string {
	return s.Uuid
}

func (s *ListApiPluginsResponseBodyApiPlugins) SetDescription(v string) *ListApiPluginsResponseBodyApiPlugins {
	s.Description = &v
	return s
}

func (s *ListApiPluginsResponseBodyApiPlugins) SetDraftedConfigJson(v string) *ListApiPluginsResponseBodyApiPlugins {
	s.DraftedConfigJson = &v
	return s
}

func (s *ListApiPluginsResponseBodyApiPlugins) SetGmtCreate(v int64) *ListApiPluginsResponseBodyApiPlugins {
	s.GmtCreate = &v
	return s
}

func (s *ListApiPluginsResponseBodyApiPlugins) SetGmtModified(v int64) *ListApiPluginsResponseBodyApiPlugins {
	s.GmtModified = &v
	return s
}

func (s *ListApiPluginsResponseBodyApiPlugins) SetInstanceId(v string) *ListApiPluginsResponseBodyApiPlugins {
	s.InstanceId = &v
	return s
}

func (s *ListApiPluginsResponseBodyApiPlugins) SetName(v string) *ListApiPluginsResponseBodyApiPlugins {
	s.Name = &v
	return s
}

func (s *ListApiPluginsResponseBodyApiPlugins) SetPublishedConfigJson(v string) *ListApiPluginsResponseBodyApiPlugins {
	s.PublishedConfigJson = &v
	return s
}

func (s *ListApiPluginsResponseBodyApiPlugins) SetStatus(v string) *ListApiPluginsResponseBodyApiPlugins {
	s.Status = &v
	return s
}

func (s *ListApiPluginsResponseBodyApiPlugins) SetType(v string) *ListApiPluginsResponseBodyApiPlugins {
	s.Type = &v
	return s
}

func (s *ListApiPluginsResponseBodyApiPlugins) SetUuid(v string) *ListApiPluginsResponseBodyApiPlugins {
	s.Uuid = &v
	return s
}

func (s *ListApiPluginsResponseBodyApiPlugins) Validate() error {
	return dara.Validate(s)
}
