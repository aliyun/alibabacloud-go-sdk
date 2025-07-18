// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetApprovalSchemaResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *GetApprovalSchemaResponseBody
	GetRequestId() *string
	SetSchema(v *GetApprovalSchemaResponseBodySchema) *GetApprovalSchemaResponseBody
	GetSchema() *GetApprovalSchemaResponseBodySchema
}

type GetApprovalSchemaResponseBody struct {
	// example:
	//
	// EFE7EBB2-449D-5BBB-B381-CA7839BC1649
	RequestId *string                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Schema    *GetApprovalSchemaResponseBodySchema `json:"Schema,omitempty" xml:"Schema,omitempty" type:"Struct"`
}

func (s GetApprovalSchemaResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetApprovalSchemaResponseBody) GoString() string {
	return s.String()
}

func (s *GetApprovalSchemaResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetApprovalSchemaResponseBody) GetSchema() *GetApprovalSchemaResponseBodySchema {
	return s.Schema
}

func (s *GetApprovalSchemaResponseBody) SetRequestId(v string) *GetApprovalSchemaResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetApprovalSchemaResponseBody) SetSchema(v *GetApprovalSchemaResponseBodySchema) *GetApprovalSchemaResponseBody {
	s.Schema = v
	return s
}

func (s *GetApprovalSchemaResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetApprovalSchemaResponseBodySchema struct {
	// example:
	//
	// 2022-02-14 11:57:51
	CreateTime  *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// true
	IsDefault *bool `json:"IsDefault,omitempty" xml:"IsDefault,omitempty"`
	// example:
	//
	// DlpSend
	PolicyType *string `json:"PolicyType,omitempty" xml:"PolicyType,omitempty"`
	// example:
	//
	// {"form": {"labelCol": 6,"wrapperCol": 12}}
	SchemaContent *string `json:"SchemaContent,omitempty" xml:"SchemaContent,omitempty"`
	// example:
	//
	// approval-schema-090134f1ebff****
	SchemaId   *string `json:"SchemaId,omitempty" xml:"SchemaId,omitempty"`
	SchemaName *string `json:"SchemaName,omitempty" xml:"SchemaName,omitempty"`
	// example:
	//
	// 1715680940
	SchemaVersion *int64 `json:"SchemaVersion,omitempty" xml:"SchemaVersion,omitempty"`
}

func (s GetApprovalSchemaResponseBodySchema) String() string {
	return dara.Prettify(s)
}

func (s GetApprovalSchemaResponseBodySchema) GoString() string {
	return s.String()
}

func (s *GetApprovalSchemaResponseBodySchema) GetCreateTime() *string {
	return s.CreateTime
}

func (s *GetApprovalSchemaResponseBodySchema) GetDescription() *string {
	return s.Description
}

func (s *GetApprovalSchemaResponseBodySchema) GetIsDefault() *bool {
	return s.IsDefault
}

func (s *GetApprovalSchemaResponseBodySchema) GetPolicyType() *string {
	return s.PolicyType
}

func (s *GetApprovalSchemaResponseBodySchema) GetSchemaContent() *string {
	return s.SchemaContent
}

func (s *GetApprovalSchemaResponseBodySchema) GetSchemaId() *string {
	return s.SchemaId
}

func (s *GetApprovalSchemaResponseBodySchema) GetSchemaName() *string {
	return s.SchemaName
}

func (s *GetApprovalSchemaResponseBodySchema) GetSchemaVersion() *int64 {
	return s.SchemaVersion
}

func (s *GetApprovalSchemaResponseBodySchema) SetCreateTime(v string) *GetApprovalSchemaResponseBodySchema {
	s.CreateTime = &v
	return s
}

func (s *GetApprovalSchemaResponseBodySchema) SetDescription(v string) *GetApprovalSchemaResponseBodySchema {
	s.Description = &v
	return s
}

func (s *GetApprovalSchemaResponseBodySchema) SetIsDefault(v bool) *GetApprovalSchemaResponseBodySchema {
	s.IsDefault = &v
	return s
}

func (s *GetApprovalSchemaResponseBodySchema) SetPolicyType(v string) *GetApprovalSchemaResponseBodySchema {
	s.PolicyType = &v
	return s
}

func (s *GetApprovalSchemaResponseBodySchema) SetSchemaContent(v string) *GetApprovalSchemaResponseBodySchema {
	s.SchemaContent = &v
	return s
}

func (s *GetApprovalSchemaResponseBodySchema) SetSchemaId(v string) *GetApprovalSchemaResponseBodySchema {
	s.SchemaId = &v
	return s
}

func (s *GetApprovalSchemaResponseBodySchema) SetSchemaName(v string) *GetApprovalSchemaResponseBodySchema {
	s.SchemaName = &v
	return s
}

func (s *GetApprovalSchemaResponseBodySchema) SetSchemaVersion(v int64) *GetApprovalSchemaResponseBodySchema {
	s.SchemaVersion = &v
	return s
}

func (s *GetApprovalSchemaResponseBodySchema) Validate() error {
	return dara.Validate(s)
}
