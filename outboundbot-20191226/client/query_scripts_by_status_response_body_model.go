// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryScriptsByStatusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *QueryScriptsByStatusResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *QueryScriptsByStatusResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *QueryScriptsByStatusResponseBody
	GetMessage() *string
	SetRequestId(v string) *QueryScriptsByStatusResponseBody
	GetRequestId() *string
	SetScripts(v *QueryScriptsByStatusResponseBodyScripts) *QueryScriptsByStatusResponseBody
	GetScripts() *QueryScriptsByStatusResponseBodyScripts
	SetSuccess(v bool) *QueryScriptsByStatusResponseBody
	GetSuccess() *bool
}

type QueryScriptsByStatusResponseBody struct {
	// API status code
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// HTTP status code
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// API message
	//
	// example:
	//
	// Success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Request ID
	//
	// example:
	//
	// 254EB995-DEDF-48A4-9101-9CA5B72FFBCC
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Scenario List
	Scripts *QueryScriptsByStatusResponseBodyScripts `json:"Scripts,omitempty" xml:"Scripts,omitempty" type:"Struct"`
	// Indicates whether the operation succeeded.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryScriptsByStatusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryScriptsByStatusResponseBody) GoString() string {
	return s.String()
}

func (s *QueryScriptsByStatusResponseBody) GetCode() *string {
	return s.Code
}

func (s *QueryScriptsByStatusResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *QueryScriptsByStatusResponseBody) GetMessage() *string {
	return s.Message
}

func (s *QueryScriptsByStatusResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryScriptsByStatusResponseBody) GetScripts() *QueryScriptsByStatusResponseBodyScripts {
	return s.Scripts
}

func (s *QueryScriptsByStatusResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *QueryScriptsByStatusResponseBody) SetCode(v string) *QueryScriptsByStatusResponseBody {
	s.Code = &v
	return s
}

func (s *QueryScriptsByStatusResponseBody) SetHttpStatusCode(v int32) *QueryScriptsByStatusResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *QueryScriptsByStatusResponseBody) SetMessage(v string) *QueryScriptsByStatusResponseBody {
	s.Message = &v
	return s
}

func (s *QueryScriptsByStatusResponseBody) SetRequestId(v string) *QueryScriptsByStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryScriptsByStatusResponseBody) SetScripts(v *QueryScriptsByStatusResponseBodyScripts) *QueryScriptsByStatusResponseBody {
	s.Scripts = v
	return s
}

func (s *QueryScriptsByStatusResponseBody) SetSuccess(v bool) *QueryScriptsByStatusResponseBody {
	s.Success = &v
	return s
}

func (s *QueryScriptsByStatusResponseBody) Validate() error {
	if s.Scripts != nil {
		if err := s.Scripts.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QueryScriptsByStatusResponseBodyScripts struct {
	// List of scenario information
	List []*QueryScriptsByStatusResponseBodyScriptsList `json:"List,omitempty" xml:"List,omitempty" type:"Repeated"`
	// Page number
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// Page size
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// Total count
	//
	// example:
	//
	// 15
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s QueryScriptsByStatusResponseBodyScripts) String() string {
	return dara.Prettify(s)
}

func (s QueryScriptsByStatusResponseBodyScripts) GoString() string {
	return s.String()
}

func (s *QueryScriptsByStatusResponseBodyScripts) GetList() []*QueryScriptsByStatusResponseBodyScriptsList {
	return s.List
}

func (s *QueryScriptsByStatusResponseBodyScripts) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *QueryScriptsByStatusResponseBodyScripts) GetPageSize() *int32 {
	return s.PageSize
}

func (s *QueryScriptsByStatusResponseBodyScripts) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *QueryScriptsByStatusResponseBodyScripts) SetList(v []*QueryScriptsByStatusResponseBodyScriptsList) *QueryScriptsByStatusResponseBodyScripts {
	s.List = v
	return s
}

func (s *QueryScriptsByStatusResponseBodyScripts) SetPageNumber(v int32) *QueryScriptsByStatusResponseBodyScripts {
	s.PageNumber = &v
	return s
}

func (s *QueryScriptsByStatusResponseBodyScripts) SetPageSize(v int32) *QueryScriptsByStatusResponseBodyScripts {
	s.PageSize = &v
	return s
}

func (s *QueryScriptsByStatusResponseBodyScripts) SetTotalCount(v int32) *QueryScriptsByStatusResponseBodyScripts {
	s.TotalCount = &v
	return s
}

func (s *QueryScriptsByStatusResponseBodyScripts) Validate() error {
	if s.List != nil {
		for _, item := range s.List {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type QueryScriptsByStatusResponseBodyScriptsList struct {
	// Version ID
	//
	// example:
	//
	// 1579055782000
	AppliedVersion *string `json:"AppliedVersion,omitempty" xml:"AppliedVersion,omitempty"`
	// Debug status
	//
	// example:
	//
	// PUBLISHED
	DebugStatus *string `json:"DebugStatus,omitempty" xml:"DebugStatus,omitempty"`
	// Debug Version
	//
	// example:
	//
	// 01bb2df2-a273-42bb-a214-d3c626b13296
	DebugVersion *string `json:"DebugVersion,omitempty" xml:"DebugVersion,omitempty"`
	// Industry
	//
	// example:
	//
	// 通用
	Industry *string `json:"Industry,omitempty" xml:"Industry,omitempty"`
	// Indicates whether the debug version is in Draft status
	//
	// example:
	//
	// false
	IsDebugDrafted *bool `json:"IsDebugDrafted,omitempty" xml:"IsDebugDrafted,omitempty"`
	// Indicates whether the version is in Draft status
	//
	// example:
	//
	// false
	IsDrafted *bool `json:"IsDrafted,omitempty" xml:"IsDrafted,omitempty"`
	// Scenario
	//
	// example:
	//
	// 电销
	Scene *string `json:"Scene,omitempty" xml:"Scene,omitempty"`
	// Script Description
	//
	// example:
	//
	// 销售话术
	ScriptDescription *string `json:"ScriptDescription,omitempty" xml:"ScriptDescription,omitempty"`
	// Scenario ID
	//
	// example:
	//
	// fa0e21e9-caab-4629-9121-1e341243d599
	ScriptId *string `json:"ScriptId,omitempty" xml:"ScriptId,omitempty"`
	// Scenario Name
	//
	// example:
	//
	// 销售话术
	ScriptName *string `json:"ScriptName,omitempty" xml:"ScriptName,omitempty"`
	// Status of the application version. Valid values:
	//
	// - **DRAFTED**: Draft
	//
	// - **INITIALIZE_IN_PROGRESS**: Initializing
	//
	// - **PUBLISHED**: Published
	//
	// - **PUBLISH_IN_PROGRESS**: Publishing
	//
	// - **ROLLBACK_IN_PROGRESS**: Rolling back
	//
	// - **EXAMINE_IN_PROGRESS**: Pending Review
	//
	// - **PUBLISHED_AND_EXAMINE_IN_PROGRESS**: Published and pending review
	//
	// - **PUBLISH_FAILED**: Publish failed
	//
	// - **ROLLBACK_FAILED**: Rollback failed
	//
	// - **IMPORT_IN_PROGRESS**: Importing
	//
	// - **IMPORT_FAILED**: Import failed
	//
	// example:
	//
	// PUBLISHED
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// Update Time
	//
	// example:
	//
	// 1579055782000
	UpdateTime *int64 `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s QueryScriptsByStatusResponseBodyScriptsList) String() string {
	return dara.Prettify(s)
}

func (s QueryScriptsByStatusResponseBodyScriptsList) GoString() string {
	return s.String()
}

func (s *QueryScriptsByStatusResponseBodyScriptsList) GetAppliedVersion() *string {
	return s.AppliedVersion
}

func (s *QueryScriptsByStatusResponseBodyScriptsList) GetDebugStatus() *string {
	return s.DebugStatus
}

func (s *QueryScriptsByStatusResponseBodyScriptsList) GetDebugVersion() *string {
	return s.DebugVersion
}

func (s *QueryScriptsByStatusResponseBodyScriptsList) GetIndustry() *string {
	return s.Industry
}

func (s *QueryScriptsByStatusResponseBodyScriptsList) GetIsDebugDrafted() *bool {
	return s.IsDebugDrafted
}

func (s *QueryScriptsByStatusResponseBodyScriptsList) GetIsDrafted() *bool {
	return s.IsDrafted
}

func (s *QueryScriptsByStatusResponseBodyScriptsList) GetScene() *string {
	return s.Scene
}

func (s *QueryScriptsByStatusResponseBodyScriptsList) GetScriptDescription() *string {
	return s.ScriptDescription
}

func (s *QueryScriptsByStatusResponseBodyScriptsList) GetScriptId() *string {
	return s.ScriptId
}

func (s *QueryScriptsByStatusResponseBodyScriptsList) GetScriptName() *string {
	return s.ScriptName
}

func (s *QueryScriptsByStatusResponseBodyScriptsList) GetStatus() *string {
	return s.Status
}

func (s *QueryScriptsByStatusResponseBodyScriptsList) GetUpdateTime() *int64 {
	return s.UpdateTime
}

func (s *QueryScriptsByStatusResponseBodyScriptsList) SetAppliedVersion(v string) *QueryScriptsByStatusResponseBodyScriptsList {
	s.AppliedVersion = &v
	return s
}

func (s *QueryScriptsByStatusResponseBodyScriptsList) SetDebugStatus(v string) *QueryScriptsByStatusResponseBodyScriptsList {
	s.DebugStatus = &v
	return s
}

func (s *QueryScriptsByStatusResponseBodyScriptsList) SetDebugVersion(v string) *QueryScriptsByStatusResponseBodyScriptsList {
	s.DebugVersion = &v
	return s
}

func (s *QueryScriptsByStatusResponseBodyScriptsList) SetIndustry(v string) *QueryScriptsByStatusResponseBodyScriptsList {
	s.Industry = &v
	return s
}

func (s *QueryScriptsByStatusResponseBodyScriptsList) SetIsDebugDrafted(v bool) *QueryScriptsByStatusResponseBodyScriptsList {
	s.IsDebugDrafted = &v
	return s
}

func (s *QueryScriptsByStatusResponseBodyScriptsList) SetIsDrafted(v bool) *QueryScriptsByStatusResponseBodyScriptsList {
	s.IsDrafted = &v
	return s
}

func (s *QueryScriptsByStatusResponseBodyScriptsList) SetScene(v string) *QueryScriptsByStatusResponseBodyScriptsList {
	s.Scene = &v
	return s
}

func (s *QueryScriptsByStatusResponseBodyScriptsList) SetScriptDescription(v string) *QueryScriptsByStatusResponseBodyScriptsList {
	s.ScriptDescription = &v
	return s
}

func (s *QueryScriptsByStatusResponseBodyScriptsList) SetScriptId(v string) *QueryScriptsByStatusResponseBodyScriptsList {
	s.ScriptId = &v
	return s
}

func (s *QueryScriptsByStatusResponseBodyScriptsList) SetScriptName(v string) *QueryScriptsByStatusResponseBodyScriptsList {
	s.ScriptName = &v
	return s
}

func (s *QueryScriptsByStatusResponseBodyScriptsList) SetStatus(v string) *QueryScriptsByStatusResponseBodyScriptsList {
	s.Status = &v
	return s
}

func (s *QueryScriptsByStatusResponseBodyScriptsList) SetUpdateTime(v int64) *QueryScriptsByStatusResponseBodyScriptsList {
	s.UpdateTime = &v
	return s
}

func (s *QueryScriptsByStatusResponseBodyScriptsList) Validate() error {
	return dara.Validate(s)
}
