// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDataQualityTemplateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDataQualityTemplate(v *GetDataQualityTemplateResponseBodyDataQualityTemplate) *GetDataQualityTemplateResponseBody
	GetDataQualityTemplate() *GetDataQualityTemplateResponseBodyDataQualityTemplate
	SetRequestId(v string) *GetDataQualityTemplateResponseBody
	GetRequestId() *string
}

type GetDataQualityTemplateResponseBody struct {
	// The data quality rule template.
	DataQualityTemplate *GetDataQualityTemplateResponseBodyDataQualityTemplate `json:"DataQualityTemplate,omitempty" xml:"DataQualityTemplate,omitempty" type:"Struct"`
	// The API request ID, which is generated as a UUID.
	//
	// example:
	//
	// 0bc14115***159376359
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetDataQualityTemplateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetDataQualityTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *GetDataQualityTemplateResponseBody) GetDataQualityTemplate() *GetDataQualityTemplateResponseBodyDataQualityTemplate {
	return s.DataQualityTemplate
}

func (s *GetDataQualityTemplateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetDataQualityTemplateResponseBody) SetDataQualityTemplate(v *GetDataQualityTemplateResponseBodyDataQualityTemplate) *GetDataQualityTemplateResponseBody {
	s.DataQualityTemplate = v
	return s
}

func (s *GetDataQualityTemplateResponseBody) SetRequestId(v string) *GetDataQualityTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetDataQualityTemplateResponseBody) Validate() error {
	if s.DataQualityTemplate != nil {
		if err := s.DataQualityTemplate.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetDataQualityTemplateResponseBodyDataQualityTemplate struct {
	// The time when the data quality rule template was created.
	//
	// example:
	//
	// 1606724043000
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The creator of the data quality rule template.
	//
	// example:
	//
	// 238428342865
	CreateUser *string `json:"CreateUser,omitempty" xml:"CreateUser,omitempty"`
	// The ID of the data quality rule template.
	//
	// example:
	//
	// 10001
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The time when the data quality rule template was updated.
	//
	// example:
	//
	// 1606724043000
	ModifyTime *int64 `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	// The last updater of the data quality rule template.
	//
	// example:
	//
	// 238428342865
	ModifyUser *string `json:"ModifyUser,omitempty" xml:"ModifyUser,omitempty"`
	// The owner of the data quality rule template.
	//
	// example:
	//
	// 238428342865
	Owner *string `json:"Owner,omitempty" xml:"Owner,omitempty"`
	// The project ID.
	//
	// example:
	//
	// 97535
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// Specific configurations of the data quality rule template. For more information, see [Data quality Spec configuration description](~2963394~).
	//
	// example:
	//
	// {
	//
	//     "assertion": "anomaly detection fro id_not_null_cnt",
	//
	//     "id_not_null_cnt": {
	//
	//         "query": "SELECT COUNT(*) AS cnt FROM ${tableName} WHERE dt = \\"$[yyyymmdd-1]\\";"
	//
	//     },
	//
	//     "identity": "819cf1f8-29be-4f94-a9d0-c5c06c0c3d2a"
	//
	// }
	Spec *string `json:"Spec,omitempty" xml:"Spec,omitempty"`
}

func (s GetDataQualityTemplateResponseBodyDataQualityTemplate) String() string {
	return dara.Prettify(s)
}

func (s GetDataQualityTemplateResponseBodyDataQualityTemplate) GoString() string {
	return s.String()
}

func (s *GetDataQualityTemplateResponseBodyDataQualityTemplate) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *GetDataQualityTemplateResponseBodyDataQualityTemplate) GetCreateUser() *string {
	return s.CreateUser
}

func (s *GetDataQualityTemplateResponseBodyDataQualityTemplate) GetId() *string {
	return s.Id
}

func (s *GetDataQualityTemplateResponseBodyDataQualityTemplate) GetModifyTime() *int64 {
	return s.ModifyTime
}

func (s *GetDataQualityTemplateResponseBodyDataQualityTemplate) GetModifyUser() *string {
	return s.ModifyUser
}

func (s *GetDataQualityTemplateResponseBodyDataQualityTemplate) GetOwner() *string {
	return s.Owner
}

func (s *GetDataQualityTemplateResponseBodyDataQualityTemplate) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *GetDataQualityTemplateResponseBodyDataQualityTemplate) GetSpec() *string {
	return s.Spec
}

func (s *GetDataQualityTemplateResponseBodyDataQualityTemplate) SetCreateTime(v int64) *GetDataQualityTemplateResponseBodyDataQualityTemplate {
	s.CreateTime = &v
	return s
}

func (s *GetDataQualityTemplateResponseBodyDataQualityTemplate) SetCreateUser(v string) *GetDataQualityTemplateResponseBodyDataQualityTemplate {
	s.CreateUser = &v
	return s
}

func (s *GetDataQualityTemplateResponseBodyDataQualityTemplate) SetId(v string) *GetDataQualityTemplateResponseBodyDataQualityTemplate {
	s.Id = &v
	return s
}

func (s *GetDataQualityTemplateResponseBodyDataQualityTemplate) SetModifyTime(v int64) *GetDataQualityTemplateResponseBodyDataQualityTemplate {
	s.ModifyTime = &v
	return s
}

func (s *GetDataQualityTemplateResponseBodyDataQualityTemplate) SetModifyUser(v string) *GetDataQualityTemplateResponseBodyDataQualityTemplate {
	s.ModifyUser = &v
	return s
}

func (s *GetDataQualityTemplateResponseBodyDataQualityTemplate) SetOwner(v string) *GetDataQualityTemplateResponseBodyDataQualityTemplate {
	s.Owner = &v
	return s
}

func (s *GetDataQualityTemplateResponseBodyDataQualityTemplate) SetProjectId(v int64) *GetDataQualityTemplateResponseBodyDataQualityTemplate {
	s.ProjectId = &v
	return s
}

func (s *GetDataQualityTemplateResponseBodyDataQualityTemplate) SetSpec(v string) *GetDataQualityTemplateResponseBodyDataQualityTemplate {
	s.Spec = &v
	return s
}

func (s *GetDataQualityTemplateResponseBodyDataQualityTemplate) Validate() error {
	return dara.Validate(s)
}
