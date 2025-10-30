// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListApprovalSchemasResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ListApprovalSchemasResponseBody
	GetRequestId() *string
	SetSchemas(v []*ListApprovalSchemasResponseBodySchemas) *ListApprovalSchemasResponseBody
	GetSchemas() []*ListApprovalSchemasResponseBodySchemas
	SetTotalNum(v string) *ListApprovalSchemasResponseBody
	GetTotalNum() *string
}

type ListApprovalSchemasResponseBody struct {
	// example:
	//
	// 102350E7-1A20-58F5-9D63-ABEA820AE6E1
	RequestId *string                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Schemas   []*ListApprovalSchemasResponseBodySchemas `json:"Schemas,omitempty" xml:"Schemas,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	TotalNum *string `json:"TotalNum,omitempty" xml:"TotalNum,omitempty"`
}

func (s ListApprovalSchemasResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListApprovalSchemasResponseBody) GoString() string {
	return s.String()
}

func (s *ListApprovalSchemasResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListApprovalSchemasResponseBody) GetSchemas() []*ListApprovalSchemasResponseBodySchemas {
	return s.Schemas
}

func (s *ListApprovalSchemasResponseBody) GetTotalNum() *string {
	return s.TotalNum
}

func (s *ListApprovalSchemasResponseBody) SetRequestId(v string) *ListApprovalSchemasResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListApprovalSchemasResponseBody) SetSchemas(v []*ListApprovalSchemasResponseBodySchemas) *ListApprovalSchemasResponseBody {
	s.Schemas = v
	return s
}

func (s *ListApprovalSchemasResponseBody) SetTotalNum(v string) *ListApprovalSchemasResponseBody {
	s.TotalNum = &v
	return s
}

func (s *ListApprovalSchemasResponseBody) Validate() error {
	if s.Schemas != nil {
		for _, item := range s.Schemas {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListApprovalSchemasResponseBodySchemas struct {
	// example:
	//
	// 2024-03-15 14:44:07
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

func (s ListApprovalSchemasResponseBodySchemas) String() string {
	return dara.Prettify(s)
}

func (s ListApprovalSchemasResponseBodySchemas) GoString() string {
	return s.String()
}

func (s *ListApprovalSchemasResponseBodySchemas) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListApprovalSchemasResponseBodySchemas) GetDescription() *string {
	return s.Description
}

func (s *ListApprovalSchemasResponseBodySchemas) GetIsDefault() *bool {
	return s.IsDefault
}

func (s *ListApprovalSchemasResponseBodySchemas) GetPolicyType() *string {
	return s.PolicyType
}

func (s *ListApprovalSchemasResponseBodySchemas) GetSchemaContent() *string {
	return s.SchemaContent
}

func (s *ListApprovalSchemasResponseBodySchemas) GetSchemaId() *string {
	return s.SchemaId
}

func (s *ListApprovalSchemasResponseBodySchemas) GetSchemaName() *string {
	return s.SchemaName
}

func (s *ListApprovalSchemasResponseBodySchemas) GetSchemaVersion() *int64 {
	return s.SchemaVersion
}

func (s *ListApprovalSchemasResponseBodySchemas) SetCreateTime(v string) *ListApprovalSchemasResponseBodySchemas {
	s.CreateTime = &v
	return s
}

func (s *ListApprovalSchemasResponseBodySchemas) SetDescription(v string) *ListApprovalSchemasResponseBodySchemas {
	s.Description = &v
	return s
}

func (s *ListApprovalSchemasResponseBodySchemas) SetIsDefault(v bool) *ListApprovalSchemasResponseBodySchemas {
	s.IsDefault = &v
	return s
}

func (s *ListApprovalSchemasResponseBodySchemas) SetPolicyType(v string) *ListApprovalSchemasResponseBodySchemas {
	s.PolicyType = &v
	return s
}

func (s *ListApprovalSchemasResponseBodySchemas) SetSchemaContent(v string) *ListApprovalSchemasResponseBodySchemas {
	s.SchemaContent = &v
	return s
}

func (s *ListApprovalSchemasResponseBodySchemas) SetSchemaId(v string) *ListApprovalSchemasResponseBodySchemas {
	s.SchemaId = &v
	return s
}

func (s *ListApprovalSchemasResponseBodySchemas) SetSchemaName(v string) *ListApprovalSchemasResponseBodySchemas {
	s.SchemaName = &v
	return s
}

func (s *ListApprovalSchemasResponseBodySchemas) SetSchemaVersion(v int64) *ListApprovalSchemasResponseBodySchemas {
	s.SchemaVersion = &v
	return s
}

func (s *ListApprovalSchemasResponseBodySchemas) Validate() error {
	return dara.Validate(s)
}
