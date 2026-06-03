// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSchemeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetSchemeResponseBody
	GetCode() *string
	SetData(v *GetSchemeResponseBodyData) *GetSchemeResponseBody
	GetData() *GetSchemeResponseBodyData
	SetMessage(v string) *GetSchemeResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetSchemeResponseBody
	GetRequestId() *string
}

type GetSchemeResponseBody struct {
	Code      *string                    `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *GetSchemeResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                    `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetSchemeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetSchemeResponseBody) GoString() string {
	return s.String()
}

func (s *GetSchemeResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetSchemeResponseBody) GetData() *GetSchemeResponseBodyData {
	return s.Data
}

func (s *GetSchemeResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetSchemeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetSchemeResponseBody) SetCode(v string) *GetSchemeResponseBody {
	s.Code = &v
	return s
}

func (s *GetSchemeResponseBody) SetData(v *GetSchemeResponseBodyData) *GetSchemeResponseBody {
	s.Data = v
	return s
}

func (s *GetSchemeResponseBody) SetMessage(v string) *GetSchemeResponseBody {
	s.Message = &v
	return s
}

func (s *GetSchemeResponseBody) SetRequestId(v string) *GetSchemeResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetSchemeResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetSchemeResponseBodyData struct {
	AuditDesc        *string                                `json:"AuditDesc,omitempty" xml:"AuditDesc,omitempty"`
	AuditTime        *int64                                 `json:"AuditTime,omitempty" xml:"AuditTime,omitempty"`
	BusinessTypeList []*int32                               `json:"BusinessTypeList,omitempty" xml:"BusinessTypeList,omitempty" type:"Repeated"`
	CompanyId        *int64                                 `json:"CompanyId,omitempty" xml:"CompanyId,omitempty"`
	CycleList        []*string                              `json:"CycleList,omitempty" xml:"CycleList,omitempty" type:"Repeated"`
	Description      *string                                `json:"Description,omitempty" xml:"Description,omitempty"`
	IndustryId       *string                                `json:"IndustryId,omitempty" xml:"IndustryId,omitempty"`
	ScenesList       []*GetSchemeResponseBodyDataScenesList `json:"ScenesList,omitempty" xml:"ScenesList,omitempty" type:"Repeated"`
	SchemeId         *int64                                 `json:"SchemeId,omitempty" xml:"SchemeId,omitempty"`
	SchemeName       *string                                `json:"SchemeName,omitempty" xml:"SchemeName,omitempty"`
	SchemeType       *int32                                 `json:"SchemeType,omitempty" xml:"SchemeType,omitempty"`
	Statement        *string                                `json:"Statement,omitempty" xml:"Statement,omitempty"`
	Status           *int32                                 `json:"Status,omitempty" xml:"Status,omitempty"`
	TemplateId       *int64                                 `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
}

func (s GetSchemeResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetSchemeResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetSchemeResponseBodyData) GetAuditDesc() *string {
	return s.AuditDesc
}

func (s *GetSchemeResponseBodyData) GetAuditTime() *int64 {
	return s.AuditTime
}

func (s *GetSchemeResponseBodyData) GetBusinessTypeList() []*int32 {
	return s.BusinessTypeList
}

func (s *GetSchemeResponseBodyData) GetCompanyId() *int64 {
	return s.CompanyId
}

func (s *GetSchemeResponseBodyData) GetCycleList() []*string {
	return s.CycleList
}

func (s *GetSchemeResponseBodyData) GetDescription() *string {
	return s.Description
}

func (s *GetSchemeResponseBodyData) GetIndustryId() *string {
	return s.IndustryId
}

func (s *GetSchemeResponseBodyData) GetScenesList() []*GetSchemeResponseBodyDataScenesList {
	return s.ScenesList
}

func (s *GetSchemeResponseBodyData) GetSchemeId() *int64 {
	return s.SchemeId
}

func (s *GetSchemeResponseBodyData) GetSchemeName() *string {
	return s.SchemeName
}

func (s *GetSchemeResponseBodyData) GetSchemeType() *int32 {
	return s.SchemeType
}

func (s *GetSchemeResponseBodyData) GetStatement() *string {
	return s.Statement
}

func (s *GetSchemeResponseBodyData) GetStatus() *int32 {
	return s.Status
}

func (s *GetSchemeResponseBodyData) GetTemplateId() *int64 {
	return s.TemplateId
}

func (s *GetSchemeResponseBodyData) SetAuditDesc(v string) *GetSchemeResponseBodyData {
	s.AuditDesc = &v
	return s
}

func (s *GetSchemeResponseBodyData) SetAuditTime(v int64) *GetSchemeResponseBodyData {
	s.AuditTime = &v
	return s
}

func (s *GetSchemeResponseBodyData) SetBusinessTypeList(v []*int32) *GetSchemeResponseBodyData {
	s.BusinessTypeList = v
	return s
}

func (s *GetSchemeResponseBodyData) SetCompanyId(v int64) *GetSchemeResponseBodyData {
	s.CompanyId = &v
	return s
}

func (s *GetSchemeResponseBodyData) SetCycleList(v []*string) *GetSchemeResponseBodyData {
	s.CycleList = v
	return s
}

func (s *GetSchemeResponseBodyData) SetDescription(v string) *GetSchemeResponseBodyData {
	s.Description = &v
	return s
}

func (s *GetSchemeResponseBodyData) SetIndustryId(v string) *GetSchemeResponseBodyData {
	s.IndustryId = &v
	return s
}

func (s *GetSchemeResponseBodyData) SetScenesList(v []*GetSchemeResponseBodyDataScenesList) *GetSchemeResponseBodyData {
	s.ScenesList = v
	return s
}

func (s *GetSchemeResponseBodyData) SetSchemeId(v int64) *GetSchemeResponseBodyData {
	s.SchemeId = &v
	return s
}

func (s *GetSchemeResponseBodyData) SetSchemeName(v string) *GetSchemeResponseBodyData {
	s.SchemeName = &v
	return s
}

func (s *GetSchemeResponseBodyData) SetSchemeType(v int32) *GetSchemeResponseBodyData {
	s.SchemeType = &v
	return s
}

func (s *GetSchemeResponseBodyData) SetStatement(v string) *GetSchemeResponseBodyData {
	s.Statement = &v
	return s
}

func (s *GetSchemeResponseBodyData) SetStatus(v int32) *GetSchemeResponseBodyData {
	s.Status = &v
	return s
}

func (s *GetSchemeResponseBodyData) SetTemplateId(v int64) *GetSchemeResponseBodyData {
	s.TemplateId = &v
	return s
}

func (s *GetSchemeResponseBodyData) Validate() error {
	if s.ScenesList != nil {
		for _, item := range s.ScenesList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetSchemeResponseBodyDataScenesList struct {
	ScenesId   *int64  `json:"ScenesId,omitempty" xml:"ScenesId,omitempty"`
	ScenesName *string `json:"ScenesName,omitempty" xml:"ScenesName,omitempty"`
}

func (s GetSchemeResponseBodyDataScenesList) String() string {
	return dara.Prettify(s)
}

func (s GetSchemeResponseBodyDataScenesList) GoString() string {
	return s.String()
}

func (s *GetSchemeResponseBodyDataScenesList) GetScenesId() *int64 {
	return s.ScenesId
}

func (s *GetSchemeResponseBodyDataScenesList) GetScenesName() *string {
	return s.ScenesName
}

func (s *GetSchemeResponseBodyDataScenesList) SetScenesId(v int64) *GetSchemeResponseBodyDataScenesList {
	s.ScenesId = &v
	return s
}

func (s *GetSchemeResponseBodyDataScenesList) SetScenesName(v string) *GetSchemeResponseBodyDataScenesList {
	s.ScenesName = &v
	return s
}

func (s *GetSchemeResponseBodyDataScenesList) Validate() error {
	return dara.Validate(s)
}
