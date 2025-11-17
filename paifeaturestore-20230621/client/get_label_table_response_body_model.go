// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetLabelTableResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDatasourceId(v string) *GetLabelTableResponseBody
	GetDatasourceId() *string
	SetDatasourceName(v string) *GetLabelTableResponseBody
	GetDatasourceName() *string
	SetFields(v []*GetLabelTableResponseBodyFields) *GetLabelTableResponseBody
	GetFields() []*GetLabelTableResponseBodyFields
	SetGmtCreateTime(v string) *GetLabelTableResponseBody
	GetGmtCreateTime() *string
	SetGmtModifiedTime(v string) *GetLabelTableResponseBody
	GetGmtModifiedTime() *string
	SetName(v string) *GetLabelTableResponseBody
	GetName() *string
	SetOwner(v string) *GetLabelTableResponseBody
	GetOwner() *string
	SetProjectId(v string) *GetLabelTableResponseBody
	GetProjectId() *string
	SetProjectName(v string) *GetLabelTableResponseBody
	GetProjectName() *string
	SetRelatedModelFeatures(v []*string) *GetLabelTableResponseBody
	GetRelatedModelFeatures() []*string
	SetRequestId(v string) *GetLabelTableResponseBody
	GetRequestId() *string
}

type GetLabelTableResponseBody struct {
	// example:
	//
	// 1
	DatasourceId *string `json:"DatasourceId,omitempty" xml:"DatasourceId,omitempty"`
	// example:
	//
	// datasource1
	DatasourceName *string                            `json:"DatasourceName,omitempty" xml:"DatasourceName,omitempty"`
	Fields         []*GetLabelTableResponseBodyFields `json:"Fields,omitempty" xml:"Fields,omitempty" type:"Repeated"`
	// example:
	//
	// 2021-12-15T23:24:33.132+08:00
	GmtCreateTime *string `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	// example:
	//
	// 2021-12-15T23:24:33.132+08:00
	GmtModifiedTime *string `json:"GmtModifiedTime,omitempty" xml:"GmtModifiedTime,omitempty"`
	// example:
	//
	// label_table1
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// 12321312*****
	Owner *string `json:"Owner,omitempty" xml:"Owner,omitempty"`
	// example:
	//
	// 1
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// example:
	//
	// project1
	ProjectName          *string   `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	RelatedModelFeatures []*string `json:"RelatedModelFeatures,omitempty" xml:"RelatedModelFeatures,omitempty" type:"Repeated"`
	// example:
	//
	// 728C5E01-ABF6-5AA8-B9FC-B3BA05DECC77
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetLabelTableResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetLabelTableResponseBody) GoString() string {
	return s.String()
}

func (s *GetLabelTableResponseBody) GetDatasourceId() *string {
	return s.DatasourceId
}

func (s *GetLabelTableResponseBody) GetDatasourceName() *string {
	return s.DatasourceName
}

func (s *GetLabelTableResponseBody) GetFields() []*GetLabelTableResponseBodyFields {
	return s.Fields
}

func (s *GetLabelTableResponseBody) GetGmtCreateTime() *string {
	return s.GmtCreateTime
}

func (s *GetLabelTableResponseBody) GetGmtModifiedTime() *string {
	return s.GmtModifiedTime
}

func (s *GetLabelTableResponseBody) GetName() *string {
	return s.Name
}

func (s *GetLabelTableResponseBody) GetOwner() *string {
	return s.Owner
}

func (s *GetLabelTableResponseBody) GetProjectId() *string {
	return s.ProjectId
}

func (s *GetLabelTableResponseBody) GetProjectName() *string {
	return s.ProjectName
}

func (s *GetLabelTableResponseBody) GetRelatedModelFeatures() []*string {
	return s.RelatedModelFeatures
}

func (s *GetLabelTableResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetLabelTableResponseBody) SetDatasourceId(v string) *GetLabelTableResponseBody {
	s.DatasourceId = &v
	return s
}

func (s *GetLabelTableResponseBody) SetDatasourceName(v string) *GetLabelTableResponseBody {
	s.DatasourceName = &v
	return s
}

func (s *GetLabelTableResponseBody) SetFields(v []*GetLabelTableResponseBodyFields) *GetLabelTableResponseBody {
	s.Fields = v
	return s
}

func (s *GetLabelTableResponseBody) SetGmtCreateTime(v string) *GetLabelTableResponseBody {
	s.GmtCreateTime = &v
	return s
}

func (s *GetLabelTableResponseBody) SetGmtModifiedTime(v string) *GetLabelTableResponseBody {
	s.GmtModifiedTime = &v
	return s
}

func (s *GetLabelTableResponseBody) SetName(v string) *GetLabelTableResponseBody {
	s.Name = &v
	return s
}

func (s *GetLabelTableResponseBody) SetOwner(v string) *GetLabelTableResponseBody {
	s.Owner = &v
	return s
}

func (s *GetLabelTableResponseBody) SetProjectId(v string) *GetLabelTableResponseBody {
	s.ProjectId = &v
	return s
}

func (s *GetLabelTableResponseBody) SetProjectName(v string) *GetLabelTableResponseBody {
	s.ProjectName = &v
	return s
}

func (s *GetLabelTableResponseBody) SetRelatedModelFeatures(v []*string) *GetLabelTableResponseBody {
	s.RelatedModelFeatures = v
	return s
}

func (s *GetLabelTableResponseBody) SetRequestId(v string) *GetLabelTableResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetLabelTableResponseBody) Validate() error {
	if s.Fields != nil {
		for _, item := range s.Fields {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetLabelTableResponseBodyFields struct {
	Attributes []*string `json:"Attributes,omitempty" xml:"Attributes,omitempty" type:"Repeated"`
	// example:
	//
	// field1
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// INT32
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GetLabelTableResponseBodyFields) String() string {
	return dara.Prettify(s)
}

func (s GetLabelTableResponseBodyFields) GoString() string {
	return s.String()
}

func (s *GetLabelTableResponseBodyFields) GetAttributes() []*string {
	return s.Attributes
}

func (s *GetLabelTableResponseBodyFields) GetName() *string {
	return s.Name
}

func (s *GetLabelTableResponseBodyFields) GetType() *string {
	return s.Type
}

func (s *GetLabelTableResponseBodyFields) SetAttributes(v []*string) *GetLabelTableResponseBodyFields {
	s.Attributes = v
	return s
}

func (s *GetLabelTableResponseBodyFields) SetName(v string) *GetLabelTableResponseBodyFields {
	s.Name = &v
	return s
}

func (s *GetLabelTableResponseBodyFields) SetType(v string) *GetLabelTableResponseBodyFields {
	s.Type = &v
	return s
}

func (s *GetLabelTableResponseBodyFields) Validate() error {
	return dara.Validate(s)
}
