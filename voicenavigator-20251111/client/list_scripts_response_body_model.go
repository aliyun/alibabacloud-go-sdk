// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListScriptsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListScriptsResponseBody
	GetCode() *string
	SetData(v *ListScriptsResponseBodyData) *ListScriptsResponseBody
	GetData() *ListScriptsResponseBodyData
	SetHttpStatusCode(v int32) *ListScriptsResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *ListScriptsResponseBody
	GetMessage() *string
	SetParams(v []*string) *ListScriptsResponseBody
	GetParams() []*string
	SetRequestId(v string) *ListScriptsResponseBody
	GetRequestId() *string
}

type ListScriptsResponseBody struct {
	// example:
	//
	// OK
	Code *string                      `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *ListScriptsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// Instance af81a389-91f0-4157-8d82-720edd02b66a
	//
	//  does not exist.
	Message *string   `json:"Message,omitempty" xml:"Message,omitempty"`
	Params  []*string `json:"Params,omitempty" xml:"Params,omitempty" type:"Repeated"`
	// example:
	//
	// BA092E9B-3421-5862-BC75-E646B7587531
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListScriptsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListScriptsResponseBody) GoString() string {
	return s.String()
}

func (s *ListScriptsResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListScriptsResponseBody) GetData() *ListScriptsResponseBodyData {
	return s.Data
}

func (s *ListScriptsResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListScriptsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListScriptsResponseBody) GetParams() []*string {
	return s.Params
}

func (s *ListScriptsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListScriptsResponseBody) SetCode(v string) *ListScriptsResponseBody {
	s.Code = &v
	return s
}

func (s *ListScriptsResponseBody) SetData(v *ListScriptsResponseBodyData) *ListScriptsResponseBody {
	s.Data = v
	return s
}

func (s *ListScriptsResponseBody) SetHttpStatusCode(v int32) *ListScriptsResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListScriptsResponseBody) SetMessage(v string) *ListScriptsResponseBody {
	s.Message = &v
	return s
}

func (s *ListScriptsResponseBody) SetParams(v []*string) *ListScriptsResponseBody {
	s.Params = v
	return s
}

func (s *ListScriptsResponseBody) SetRequestId(v string) *ListScriptsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListScriptsResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListScriptsResponseBodyData struct {
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32                                `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Scripts  []*ListScriptsResponseBodyDataScripts `json:"Scripts,omitempty" xml:"Scripts,omitempty" type:"Repeated"`
	// example:
	//
	// 30
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListScriptsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListScriptsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListScriptsResponseBodyData) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListScriptsResponseBodyData) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListScriptsResponseBodyData) GetScripts() []*ListScriptsResponseBodyDataScripts {
	return s.Scripts
}

func (s *ListScriptsResponseBodyData) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListScriptsResponseBodyData) SetPageNumber(v int32) *ListScriptsResponseBodyData {
	s.PageNumber = &v
	return s
}

func (s *ListScriptsResponseBodyData) SetPageSize(v int32) *ListScriptsResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListScriptsResponseBodyData) SetScripts(v []*ListScriptsResponseBodyDataScripts) *ListScriptsResponseBodyData {
	s.Scripts = v
	return s
}

func (s *ListScriptsResponseBodyData) SetTotalCount(v int32) *ListScriptsResponseBodyData {
	s.TotalCount = &v
	return s
}

func (s *ListScriptsResponseBodyData) Validate() error {
	if s.Scripts != nil {
		for _, item := range s.Scripts {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListScriptsResponseBodyDataScripts struct {
	// example:
	//
	// 10
	Concurrency *int32 `json:"Concurrency,omitempty" xml:"Concurrency,omitempty"`
	// example:
	//
	// 1773228988000
	CreatedTime *int64  `json:"CreatedTime,omitempty" xml:"CreatedTime,omitempty"`
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// efbafa38-336d-4eb0-865e-c16c97a91773
	DraftVersionId *string `json:"DraftVersionId,omitempty" xml:"DraftVersionId,omitempty"`
	Name           *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// MANAGED
	NluAccessType *string `json:"NluAccessType,omitempty" xml:"NluAccessType,omitempty"`
	// example:
	//
	// BEEBOT
	NluEngine *string `json:"NluEngine,omitempty" xml:"NluEngine,omitempty"`
	// example:
	//
	// 13816111993
	Number *string `json:"Number,omitempty" xml:"Number,omitempty"`
	// example:
	//
	// efbafa38-336d-4eb0-865e-c16c97a91774
	PublishedVersionId *string `json:"PublishedVersionId,omitempty" xml:"PublishedVersionId,omitempty"`
	// example:
	//
	// efbafa38-336d-4eb0-865e-c16c97a91772
	ScriptId *string `json:"ScriptId,omitempty" xml:"ScriptId,omitempty"`
	// example:
	//
	// DRAFT
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// 1773228988000
	UpdatedTime *int64 `json:"UpdatedTime,omitempty" xml:"UpdatedTime,omitempty"`
}

func (s ListScriptsResponseBodyDataScripts) String() string {
	return dara.Prettify(s)
}

func (s ListScriptsResponseBodyDataScripts) GoString() string {
	return s.String()
}

func (s *ListScriptsResponseBodyDataScripts) GetConcurrency() *int32 {
	return s.Concurrency
}

func (s *ListScriptsResponseBodyDataScripts) GetCreatedTime() *int64 {
	return s.CreatedTime
}

func (s *ListScriptsResponseBodyDataScripts) GetDescription() *string {
	return s.Description
}

func (s *ListScriptsResponseBodyDataScripts) GetDraftVersionId() *string {
	return s.DraftVersionId
}

func (s *ListScriptsResponseBodyDataScripts) GetName() *string {
	return s.Name
}

func (s *ListScriptsResponseBodyDataScripts) GetNluAccessType() *string {
	return s.NluAccessType
}

func (s *ListScriptsResponseBodyDataScripts) GetNluEngine() *string {
	return s.NluEngine
}

func (s *ListScriptsResponseBodyDataScripts) GetNumber() *string {
	return s.Number
}

func (s *ListScriptsResponseBodyDataScripts) GetPublishedVersionId() *string {
	return s.PublishedVersionId
}

func (s *ListScriptsResponseBodyDataScripts) GetScriptId() *string {
	return s.ScriptId
}

func (s *ListScriptsResponseBodyDataScripts) GetStatus() *string {
	return s.Status
}

func (s *ListScriptsResponseBodyDataScripts) GetUpdatedTime() *int64 {
	return s.UpdatedTime
}

func (s *ListScriptsResponseBodyDataScripts) SetConcurrency(v int32) *ListScriptsResponseBodyDataScripts {
	s.Concurrency = &v
	return s
}

func (s *ListScriptsResponseBodyDataScripts) SetCreatedTime(v int64) *ListScriptsResponseBodyDataScripts {
	s.CreatedTime = &v
	return s
}

func (s *ListScriptsResponseBodyDataScripts) SetDescription(v string) *ListScriptsResponseBodyDataScripts {
	s.Description = &v
	return s
}

func (s *ListScriptsResponseBodyDataScripts) SetDraftVersionId(v string) *ListScriptsResponseBodyDataScripts {
	s.DraftVersionId = &v
	return s
}

func (s *ListScriptsResponseBodyDataScripts) SetName(v string) *ListScriptsResponseBodyDataScripts {
	s.Name = &v
	return s
}

func (s *ListScriptsResponseBodyDataScripts) SetNluAccessType(v string) *ListScriptsResponseBodyDataScripts {
	s.NluAccessType = &v
	return s
}

func (s *ListScriptsResponseBodyDataScripts) SetNluEngine(v string) *ListScriptsResponseBodyDataScripts {
	s.NluEngine = &v
	return s
}

func (s *ListScriptsResponseBodyDataScripts) SetNumber(v string) *ListScriptsResponseBodyDataScripts {
	s.Number = &v
	return s
}

func (s *ListScriptsResponseBodyDataScripts) SetPublishedVersionId(v string) *ListScriptsResponseBodyDataScripts {
	s.PublishedVersionId = &v
	return s
}

func (s *ListScriptsResponseBodyDataScripts) SetScriptId(v string) *ListScriptsResponseBodyDataScripts {
	s.ScriptId = &v
	return s
}

func (s *ListScriptsResponseBodyDataScripts) SetStatus(v string) *ListScriptsResponseBodyDataScripts {
	s.Status = &v
	return s
}

func (s *ListScriptsResponseBodyDataScripts) SetUpdatedTime(v int64) *ListScriptsResponseBodyDataScripts {
	s.UpdatedTime = &v
	return s
}

func (s *ListScriptsResponseBodyDataScripts) Validate() error {
	return dara.Validate(s)
}
