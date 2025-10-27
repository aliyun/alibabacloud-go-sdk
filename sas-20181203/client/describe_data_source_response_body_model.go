// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDataSourceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMetaDatas(v []*DescribeDataSourceResponseBodyMetaDatas) *DescribeDataSourceResponseBody
	GetMetaDatas() []*DescribeDataSourceResponseBodyMetaDatas
	SetRequestId(v string) *DescribeDataSourceResponseBody
	GetRequestId() *string
}

type DescribeDataSourceResponseBody struct {
	// The metadata of the data sources.
	MetaDatas []*DescribeDataSourceResponseBodyMetaDatas `json:"MetaDatas,omitempty" xml:"MetaDatas,omitempty" type:"Repeated"`
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// 7532B7EE-7CE7-5F4D-BF04-B12447DDXXXX
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeDataSourceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDataSourceResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDataSourceResponseBody) GetMetaDatas() []*DescribeDataSourceResponseBodyMetaDatas {
	return s.MetaDatas
}

func (s *DescribeDataSourceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDataSourceResponseBody) SetMetaDatas(v []*DescribeDataSourceResponseBodyMetaDatas) *DescribeDataSourceResponseBody {
	s.MetaDatas = v
	return s
}

func (s *DescribeDataSourceResponseBody) SetRequestId(v string) *DescribeDataSourceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDataSourceResponseBody) Validate() error {
	if s.MetaDatas != nil {
		for _, item := range s.MetaDatas {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeDataSourceResponseBodyMetaDatas struct {
	// The ID of the data source.
	//
	// example:
	//
	// 1753
	DataSourceId *int32 `json:"DataSourceId,omitempty" xml:"DataSourceId,omitempty"`
	// The name of the data source.
	//
	// example:
	//
	// sas_analysis_pre-sas-operation-log-sas-event-suspicious
	DataSourceName *string `json:"DataSourceName,omitempty" xml:"DataSourceName,omitempty"`
	// The description of the data source.
	//
	// example:
	//
	// dingtalk_suspicious
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The metadata files.
	MetaDataFields []*DescribeDataSourceResponseBodyMetaDatasMetaDataFields `json:"MetaDataFields,omitempty" xml:"MetaDataFields,omitempty" type:"Repeated"`
}

func (s DescribeDataSourceResponseBodyMetaDatas) String() string {
	return dara.Prettify(s)
}

func (s DescribeDataSourceResponseBodyMetaDatas) GoString() string {
	return s.String()
}

func (s *DescribeDataSourceResponseBodyMetaDatas) GetDataSourceId() *int32 {
	return s.DataSourceId
}

func (s *DescribeDataSourceResponseBodyMetaDatas) GetDataSourceName() *string {
	return s.DataSourceName
}

func (s *DescribeDataSourceResponseBodyMetaDatas) GetDescription() *string {
	return s.Description
}

func (s *DescribeDataSourceResponseBodyMetaDatas) GetMetaDataFields() []*DescribeDataSourceResponseBodyMetaDatasMetaDataFields {
	return s.MetaDataFields
}

func (s *DescribeDataSourceResponseBodyMetaDatas) SetDataSourceId(v int32) *DescribeDataSourceResponseBodyMetaDatas {
	s.DataSourceId = &v
	return s
}

func (s *DescribeDataSourceResponseBodyMetaDatas) SetDataSourceName(v string) *DescribeDataSourceResponseBodyMetaDatas {
	s.DataSourceName = &v
	return s
}

func (s *DescribeDataSourceResponseBodyMetaDatas) SetDescription(v string) *DescribeDataSourceResponseBodyMetaDatas {
	s.Description = &v
	return s
}

func (s *DescribeDataSourceResponseBodyMetaDatas) SetMetaDataFields(v []*DescribeDataSourceResponseBodyMetaDatasMetaDataFields) *DescribeDataSourceResponseBodyMetaDatas {
	s.MetaDataFields = v
	return s
}

func (s *DescribeDataSourceResponseBodyMetaDatas) Validate() error {
	if s.MetaDataFields != nil {
		for _, item := range s.MetaDataFields {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeDataSourceResponseBodyMetaDatasMetaDataFields struct {
	// The key of the field.
	//
	// example:
	//
	// type
	Filed *string `json:"Filed,omitempty" xml:"Filed,omitempty"`
	// The name of the field.
	//
	// example:
	//
	// dingtalk_vul_type
	FiledName *string `json:"FiledName,omitempty" xml:"FiledName,omitempty"`
	// The operators.
	OperatorList []*DescribeDataSourceResponseBodyMetaDatasMetaDataFieldsOperatorList `json:"OperatorList,omitempty" xml:"OperatorList,omitempty" type:"Repeated"`
	// The sample field.
	//
	// example:
	//
	// all:dingtalk_all;cms:dingtalk_vul_cms;oval:dingtalk_vul_cve;sys:dingtalk_vul_sys;emg:dingtalk_vul_emg
	Sample *string `json:"Sample,omitempty" xml:"Sample,omitempty"`
	// The value type of the field.
	//
	// example:
	//
	// string
	ValueType *string `json:"ValueType,omitempty" xml:"ValueType,omitempty"`
}

func (s DescribeDataSourceResponseBodyMetaDatasMetaDataFields) String() string {
	return dara.Prettify(s)
}

func (s DescribeDataSourceResponseBodyMetaDatasMetaDataFields) GoString() string {
	return s.String()
}

func (s *DescribeDataSourceResponseBodyMetaDatasMetaDataFields) GetFiled() *string {
	return s.Filed
}

func (s *DescribeDataSourceResponseBodyMetaDatasMetaDataFields) GetFiledName() *string {
	return s.FiledName
}

func (s *DescribeDataSourceResponseBodyMetaDatasMetaDataFields) GetOperatorList() []*DescribeDataSourceResponseBodyMetaDatasMetaDataFieldsOperatorList {
	return s.OperatorList
}

func (s *DescribeDataSourceResponseBodyMetaDatasMetaDataFields) GetSample() *string {
	return s.Sample
}

func (s *DescribeDataSourceResponseBodyMetaDatasMetaDataFields) GetValueType() *string {
	return s.ValueType
}

func (s *DescribeDataSourceResponseBodyMetaDatasMetaDataFields) SetFiled(v string) *DescribeDataSourceResponseBodyMetaDatasMetaDataFields {
	s.Filed = &v
	return s
}

func (s *DescribeDataSourceResponseBodyMetaDatasMetaDataFields) SetFiledName(v string) *DescribeDataSourceResponseBodyMetaDatasMetaDataFields {
	s.FiledName = &v
	return s
}

func (s *DescribeDataSourceResponseBodyMetaDatasMetaDataFields) SetOperatorList(v []*DescribeDataSourceResponseBodyMetaDatasMetaDataFieldsOperatorList) *DescribeDataSourceResponseBodyMetaDatasMetaDataFields {
	s.OperatorList = v
	return s
}

func (s *DescribeDataSourceResponseBodyMetaDatasMetaDataFields) SetSample(v string) *DescribeDataSourceResponseBodyMetaDatasMetaDataFields {
	s.Sample = &v
	return s
}

func (s *DescribeDataSourceResponseBodyMetaDatasMetaDataFields) SetValueType(v string) *DescribeDataSourceResponseBodyMetaDatasMetaDataFields {
	s.ValueType = &v
	return s
}

func (s *DescribeDataSourceResponseBodyMetaDatasMetaDataFields) Validate() error {
	if s.OperatorList != nil {
		for _, item := range s.OperatorList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeDataSourceResponseBodyMetaDatasMetaDataFieldsOperatorList struct {
	// The description of the operator.
	//
	// example:
	//
	// dingtalk_vul
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name of the operator.
	//
	// example:
	//
	// regex
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s DescribeDataSourceResponseBodyMetaDatasMetaDataFieldsOperatorList) String() string {
	return dara.Prettify(s)
}

func (s DescribeDataSourceResponseBodyMetaDatasMetaDataFieldsOperatorList) GoString() string {
	return s.String()
}

func (s *DescribeDataSourceResponseBodyMetaDatasMetaDataFieldsOperatorList) GetDescription() *string {
	return s.Description
}

func (s *DescribeDataSourceResponseBodyMetaDatasMetaDataFieldsOperatorList) GetName() *string {
	return s.Name
}

func (s *DescribeDataSourceResponseBodyMetaDatasMetaDataFieldsOperatorList) SetDescription(v string) *DescribeDataSourceResponseBodyMetaDatasMetaDataFieldsOperatorList {
	s.Description = &v
	return s
}

func (s *DescribeDataSourceResponseBodyMetaDatasMetaDataFieldsOperatorList) SetName(v string) *DescribeDataSourceResponseBodyMetaDatasMetaDataFieldsOperatorList {
	s.Name = &v
	return s
}

func (s *DescribeDataSourceResponseBodyMetaDatasMetaDataFieldsOperatorList) Validate() error {
	return dara.Validate(s)
}
