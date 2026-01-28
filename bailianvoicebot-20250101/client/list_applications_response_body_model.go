// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListApplicationsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListApplicationsResponseBody
	GetCode() *string
	SetData(v *ListApplicationsResponseBodyData) *ListApplicationsResponseBody
	GetData() *ListApplicationsResponseBodyData
	SetHttpStatusCode(v int32) *ListApplicationsResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *ListApplicationsResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListApplicationsResponseBody
	GetRequestId() *string
}

type ListApplicationsResponseBody struct {
	// example:
	//
	// OK
	Code *string                           `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *ListApplicationsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// D771A1B6-3D5F-174A-BEE1-98CE1000D337
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListApplicationsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListApplicationsResponseBody) GoString() string {
	return s.String()
}

func (s *ListApplicationsResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListApplicationsResponseBody) GetData() *ListApplicationsResponseBodyData {
	return s.Data
}

func (s *ListApplicationsResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListApplicationsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListApplicationsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListApplicationsResponseBody) SetCode(v string) *ListApplicationsResponseBody {
	s.Code = &v
	return s
}

func (s *ListApplicationsResponseBody) SetData(v *ListApplicationsResponseBodyData) *ListApplicationsResponseBody {
	s.Data = v
	return s
}

func (s *ListApplicationsResponseBody) SetHttpStatusCode(v int32) *ListApplicationsResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListApplicationsResponseBody) SetMessage(v string) *ListApplicationsResponseBody {
	s.Message = &v
	return s
}

func (s *ListApplicationsResponseBody) SetRequestId(v string) *ListApplicationsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListApplicationsResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListApplicationsResponseBodyData struct {
	Applications []*ListApplicationsResponseBodyDataApplications `json:"Applications,omitempty" xml:"Applications,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 1000
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 6
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListApplicationsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListApplicationsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListApplicationsResponseBodyData) GetApplications() []*ListApplicationsResponseBodyDataApplications {
	return s.Applications
}

func (s *ListApplicationsResponseBodyData) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListApplicationsResponseBodyData) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListApplicationsResponseBodyData) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListApplicationsResponseBodyData) SetApplications(v []*ListApplicationsResponseBodyDataApplications) *ListApplicationsResponseBodyData {
	s.Applications = v
	return s
}

func (s *ListApplicationsResponseBodyData) SetPageNumber(v int32) *ListApplicationsResponseBodyData {
	s.PageNumber = &v
	return s
}

func (s *ListApplicationsResponseBodyData) SetPageSize(v int32) *ListApplicationsResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListApplicationsResponseBodyData) SetTotalCount(v int32) *ListApplicationsResponseBodyData {
	s.TotalCount = &v
	return s
}

func (s *ListApplicationsResponseBodyData) Validate() error {
	if s.Applications != nil {
		for _, item := range s.Applications {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListApplicationsResponseBodyDataApplications struct {
	// example:
	//
	// a395011f-a247-400f-bc69-28796749fd52
	ApplicationId *string `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	// example:
	//
	// 10
	Concurrency *int32 `json:"Concurrency,omitempty" xml:"Concurrency,omitempty"`
	// example:
	//
	// 1729909690
	CreatedTime *int64  `json:"CreatedTime,omitempty" xml:"CreatedTime,omitempty"`
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// 20904943-f711-494f-9f1f-e7f340f37707
	DraftVersionId *string `json:"DraftVersionId,omitempty" xml:"DraftVersionId,omitempty"`
	Name           *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// MANAGED
	NluAccessType *string `json:"NluAccessType,omitempty" xml:"NluAccessType,omitempty"`
	// example:
	//
	// PROMPTS
	NluEngine *string `json:"NluEngine,omitempty" xml:"NluEngine,omitempty"`
	// example:
	//
	// 20904943-f711-494f-9f1f-e7f340f37707
	PublishedVersionId *string `json:"PublishedVersionId,omitempty" xml:"PublishedVersionId,omitempty"`
	// example:
	//
	// 1729909348
	UpdatedTime *int64 `json:"UpdatedTime,omitempty" xml:"UpdatedTime,omitempty"`
}

func (s ListApplicationsResponseBodyDataApplications) String() string {
	return dara.Prettify(s)
}

func (s ListApplicationsResponseBodyDataApplications) GoString() string {
	return s.String()
}

func (s *ListApplicationsResponseBodyDataApplications) GetApplicationId() *string {
	return s.ApplicationId
}

func (s *ListApplicationsResponseBodyDataApplications) GetConcurrency() *int32 {
	return s.Concurrency
}

func (s *ListApplicationsResponseBodyDataApplications) GetCreatedTime() *int64 {
	return s.CreatedTime
}

func (s *ListApplicationsResponseBodyDataApplications) GetDescription() *string {
	return s.Description
}

func (s *ListApplicationsResponseBodyDataApplications) GetDraftVersionId() *string {
	return s.DraftVersionId
}

func (s *ListApplicationsResponseBodyDataApplications) GetName() *string {
	return s.Name
}

func (s *ListApplicationsResponseBodyDataApplications) GetNluAccessType() *string {
	return s.NluAccessType
}

func (s *ListApplicationsResponseBodyDataApplications) GetNluEngine() *string {
	return s.NluEngine
}

func (s *ListApplicationsResponseBodyDataApplications) GetPublishedVersionId() *string {
	return s.PublishedVersionId
}

func (s *ListApplicationsResponseBodyDataApplications) GetUpdatedTime() *int64 {
	return s.UpdatedTime
}

func (s *ListApplicationsResponseBodyDataApplications) SetApplicationId(v string) *ListApplicationsResponseBodyDataApplications {
	s.ApplicationId = &v
	return s
}

func (s *ListApplicationsResponseBodyDataApplications) SetConcurrency(v int32) *ListApplicationsResponseBodyDataApplications {
	s.Concurrency = &v
	return s
}

func (s *ListApplicationsResponseBodyDataApplications) SetCreatedTime(v int64) *ListApplicationsResponseBodyDataApplications {
	s.CreatedTime = &v
	return s
}

func (s *ListApplicationsResponseBodyDataApplications) SetDescription(v string) *ListApplicationsResponseBodyDataApplications {
	s.Description = &v
	return s
}

func (s *ListApplicationsResponseBodyDataApplications) SetDraftVersionId(v string) *ListApplicationsResponseBodyDataApplications {
	s.DraftVersionId = &v
	return s
}

func (s *ListApplicationsResponseBodyDataApplications) SetName(v string) *ListApplicationsResponseBodyDataApplications {
	s.Name = &v
	return s
}

func (s *ListApplicationsResponseBodyDataApplications) SetNluAccessType(v string) *ListApplicationsResponseBodyDataApplications {
	s.NluAccessType = &v
	return s
}

func (s *ListApplicationsResponseBodyDataApplications) SetNluEngine(v string) *ListApplicationsResponseBodyDataApplications {
	s.NluEngine = &v
	return s
}

func (s *ListApplicationsResponseBodyDataApplications) SetPublishedVersionId(v string) *ListApplicationsResponseBodyDataApplications {
	s.PublishedVersionId = &v
	return s
}

func (s *ListApplicationsResponseBodyDataApplications) SetUpdatedTime(v int64) *ListApplicationsResponseBodyDataApplications {
	s.UpdatedTime = &v
	return s
}

func (s *ListApplicationsResponseBodyDataApplications) Validate() error {
	return dara.Validate(s)
}
