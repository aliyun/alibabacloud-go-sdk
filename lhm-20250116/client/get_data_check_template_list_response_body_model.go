// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDataCheckTemplateListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*GetDataCheckTemplateListResponseBodyData) *GetDataCheckTemplateListResponseBody
	GetData() []*GetDataCheckTemplateListResponseBodyData
	SetErrCode(v string) *GetDataCheckTemplateListResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *GetDataCheckTemplateListResponseBody
	GetErrMessage() *string
	SetRequestId(v string) *GetDataCheckTemplateListResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetDataCheckTemplateListResponseBody
	GetSuccess() *bool
}

type GetDataCheckTemplateListResponseBody struct {
	// The data list returned by the operation. For the structure of each element, see the child field descriptions.
	Data []*GetDataCheckTemplateListResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	// The error code. An empty string is returned if the call is successful.
	//
	// example:
	//
	// Success
	ErrCode *string `json:"errCode,omitempty" xml:"errCode,omitempty"`
	// The error message. An empty string is returned if the call is successful.
	//
	// example:
	//
	// success
	ErrMessage *string `json:"errMessage,omitempty" xml:"errMessage,omitempty"`
	// The request ID, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// 4C467B38-3910-4477-9B0B-6963D83B4E72
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Indicates whether the call is successful. Valid values:
	//
	// - true: The call is successful.
	//
	// - false: The call failed. Check errCode and errMessage for details.
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s GetDataCheckTemplateListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetDataCheckTemplateListResponseBody) GoString() string {
	return s.String()
}

func (s *GetDataCheckTemplateListResponseBody) GetData() []*GetDataCheckTemplateListResponseBodyData {
	return s.Data
}

func (s *GetDataCheckTemplateListResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *GetDataCheckTemplateListResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *GetDataCheckTemplateListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetDataCheckTemplateListResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetDataCheckTemplateListResponseBody) SetData(v []*GetDataCheckTemplateListResponseBodyData) *GetDataCheckTemplateListResponseBody {
	s.Data = v
	return s
}

func (s *GetDataCheckTemplateListResponseBody) SetErrCode(v string) *GetDataCheckTemplateListResponseBody {
	s.ErrCode = &v
	return s
}

func (s *GetDataCheckTemplateListResponseBody) SetErrMessage(v string) *GetDataCheckTemplateListResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *GetDataCheckTemplateListResponseBody) SetRequestId(v string) *GetDataCheckTemplateListResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetDataCheckTemplateListResponseBody) SetSuccess(v bool) *GetDataCheckTemplateListResponseBody {
	s.Success = &v
	return s
}

func (s *GetDataCheckTemplateListResponseBody) Validate() error {
	if s.Data != nil {
		for _, item := range s.Data {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetDataCheckTemplateListResponseBodyData struct {
	// The validation rule type. Valid values:
	//
	// - 0: data volume comparison.
	//
	// - 1: metric comparison.
	//
	// - 2: weak content comparison.
	//
	// - 3: custom comparison.
	//
	// - 4: full-text comparison.
	//
	// - 5: null rate comparison.
	//
	// example:
	//
	// 1
	CheckType *int32 `json:"checkType,omitempty" xml:"checkType,omitempty"`
	// The display name of the check type, used in exported reports.
	//
	// example:
	//
	// Metric Comparison
	CheckTypeExport *string `json:"checkTypeExport,omitempty" xml:"checkTypeExport,omitempty"`
	// The name of the check type.
	//
	// example:
	//
	// 1
	CheckTypeName *int32 `json:"checkTypeName,omitempty" xml:"checkTypeName,omitempty"`
	// The list of covered data source types. Multiple values are separated by commas.
	//
	// example:
	//
	// Hive,MaxCompute
	DsTypes *string `json:"dsTypes,omitempty" xml:"dsTypes,omitempty"`
	// The list of covered validation engine types, such as Tez and MapReduce. When returned as a string, multiple values are separated by commas.
	//
	// example:
	//
	// Tez,MapReduce
	EngineTypes *string `json:"engineTypes,omitempty" xml:"engineTypes,omitempty"`
	// The modification time.
	//
	// example:
	//
	// 2026-01-16 10:00:00
	GmtModified *string `json:"gmtModified,omitempty" xml:"gmtModified,omitempty"`
	// Specifies whether the template is built-in. Valid values:
	//
	// - 0: No. The template is a custom template.
	//
	// - 1: Yes. The template is a built-in template.
	//
	// example:
	//
	// 0
	IsBuiltin *int32 `json:"isBuiltin,omitempty" xml:"isBuiltin,omitempty"`
	// Indicates whether the template is referenced by a validation task. Valid values:
	//
	// - true: The template is referenced.
	//
	// - false: The template is not referenced.
	//
	// The delete operation does not verify this reference relationship. Confirm before deleting.
	IsUsedByTask *bool `json:"isUsedByTask,omitempty" xml:"isUsedByTask,omitempty"`
	// The template description.
	//
	// example:
	//
	// Description of the data volume validation template
	TemplateDesc *string `json:"templateDesc,omitempty" xml:"templateDesc,omitempty"`
	// The validation template ID (logical foreign key) that uniquely identifies a validation template.
	//
	// example:
	//
	// 1001
	TemplateId *string `json:"templateId,omitempty" xml:"templateId,omitempty"`
	// The name of the validation template.
	//
	// example:
	//
	// DataVolumeValidationTemplate
	TemplateName *string `json:"templateName,omitempty" xml:"templateName,omitempty"`
}

func (s GetDataCheckTemplateListResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetDataCheckTemplateListResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetDataCheckTemplateListResponseBodyData) GetCheckType() *int32 {
	return s.CheckType
}

func (s *GetDataCheckTemplateListResponseBodyData) GetCheckTypeExport() *string {
	return s.CheckTypeExport
}

func (s *GetDataCheckTemplateListResponseBodyData) GetCheckTypeName() *int32 {
	return s.CheckTypeName
}

func (s *GetDataCheckTemplateListResponseBodyData) GetDsTypes() *string {
	return s.DsTypes
}

func (s *GetDataCheckTemplateListResponseBodyData) GetEngineTypes() *string {
	return s.EngineTypes
}

func (s *GetDataCheckTemplateListResponseBodyData) GetGmtModified() *string {
	return s.GmtModified
}

func (s *GetDataCheckTemplateListResponseBodyData) GetIsBuiltin() *int32 {
	return s.IsBuiltin
}

func (s *GetDataCheckTemplateListResponseBodyData) GetIsUsedByTask() *bool {
	return s.IsUsedByTask
}

func (s *GetDataCheckTemplateListResponseBodyData) GetTemplateDesc() *string {
	return s.TemplateDesc
}

func (s *GetDataCheckTemplateListResponseBodyData) GetTemplateId() *string {
	return s.TemplateId
}

func (s *GetDataCheckTemplateListResponseBodyData) GetTemplateName() *string {
	return s.TemplateName
}

func (s *GetDataCheckTemplateListResponseBodyData) SetCheckType(v int32) *GetDataCheckTemplateListResponseBodyData {
	s.CheckType = &v
	return s
}

func (s *GetDataCheckTemplateListResponseBodyData) SetCheckTypeExport(v string) *GetDataCheckTemplateListResponseBodyData {
	s.CheckTypeExport = &v
	return s
}

func (s *GetDataCheckTemplateListResponseBodyData) SetCheckTypeName(v int32) *GetDataCheckTemplateListResponseBodyData {
	s.CheckTypeName = &v
	return s
}

func (s *GetDataCheckTemplateListResponseBodyData) SetDsTypes(v string) *GetDataCheckTemplateListResponseBodyData {
	s.DsTypes = &v
	return s
}

func (s *GetDataCheckTemplateListResponseBodyData) SetEngineTypes(v string) *GetDataCheckTemplateListResponseBodyData {
	s.EngineTypes = &v
	return s
}

func (s *GetDataCheckTemplateListResponseBodyData) SetGmtModified(v string) *GetDataCheckTemplateListResponseBodyData {
	s.GmtModified = &v
	return s
}

func (s *GetDataCheckTemplateListResponseBodyData) SetIsBuiltin(v int32) *GetDataCheckTemplateListResponseBodyData {
	s.IsBuiltin = &v
	return s
}

func (s *GetDataCheckTemplateListResponseBodyData) SetIsUsedByTask(v bool) *GetDataCheckTemplateListResponseBodyData {
	s.IsUsedByTask = &v
	return s
}

func (s *GetDataCheckTemplateListResponseBodyData) SetTemplateDesc(v string) *GetDataCheckTemplateListResponseBodyData {
	s.TemplateDesc = &v
	return s
}

func (s *GetDataCheckTemplateListResponseBodyData) SetTemplateId(v string) *GetDataCheckTemplateListResponseBodyData {
	s.TemplateId = &v
	return s
}

func (s *GetDataCheckTemplateListResponseBodyData) SetTemplateName(v string) *GetDataCheckTemplateListResponseBodyData {
	s.TemplateName = &v
	return s
}

func (s *GetDataCheckTemplateListResponseBodyData) Validate() error {
	return dara.Validate(s)
}
