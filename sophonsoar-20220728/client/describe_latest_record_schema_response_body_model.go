// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLatestRecordSchemaResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPlaybookNodeSchema(v *DescribeLatestRecordSchemaResponseBodyPlaybookNodeSchema) *DescribeLatestRecordSchemaResponseBody
	GetPlaybookNodeSchema() *DescribeLatestRecordSchemaResponseBodyPlaybookNodeSchema
	SetRequestId(v string) *DescribeLatestRecordSchemaResponseBody
	GetRequestId() *string
}

type DescribeLatestRecordSchemaResponseBody struct {
	// The output structure information of the playbook.
	PlaybookNodeSchema *DescribeLatestRecordSchemaResponseBodyPlaybookNodeSchema `json:"PlaybookNodeSchema,omitempty" xml:"PlaybookNodeSchema,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 10B92EE1-4597-593B-A131-7A17D25EF5C9
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeLatestRecordSchemaResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeLatestRecordSchemaResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeLatestRecordSchemaResponseBody) GetPlaybookNodeSchema() *DescribeLatestRecordSchemaResponseBodyPlaybookNodeSchema {
	return s.PlaybookNodeSchema
}

func (s *DescribeLatestRecordSchemaResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeLatestRecordSchemaResponseBody) SetPlaybookNodeSchema(v *DescribeLatestRecordSchemaResponseBodyPlaybookNodeSchema) *DescribeLatestRecordSchemaResponseBody {
	s.PlaybookNodeSchema = v
	return s
}

func (s *DescribeLatestRecordSchemaResponseBody) SetRequestId(v string) *DescribeLatestRecordSchemaResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeLatestRecordSchemaResponseBody) Validate() error {
	if s.PlaybookNodeSchema != nil {
		if err := s.PlaybookNodeSchema.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeLatestRecordSchemaResponseBodyPlaybookNodeSchema struct {
	// The structure information.
	NodeSchema []*DescribeLatestRecordSchemaResponseBodyPlaybookNodeSchemaNodeSchema `json:"NodeSchema,omitempty" xml:"NodeSchema,omitempty" type:"Repeated"`
}

func (s DescribeLatestRecordSchemaResponseBodyPlaybookNodeSchema) String() string {
	return dara.Prettify(s)
}

func (s DescribeLatestRecordSchemaResponseBodyPlaybookNodeSchema) GoString() string {
	return s.String()
}

func (s *DescribeLatestRecordSchemaResponseBodyPlaybookNodeSchema) GetNodeSchema() []*DescribeLatestRecordSchemaResponseBodyPlaybookNodeSchemaNodeSchema {
	return s.NodeSchema
}

func (s *DescribeLatestRecordSchemaResponseBodyPlaybookNodeSchema) SetNodeSchema(v []*DescribeLatestRecordSchemaResponseBodyPlaybookNodeSchemaNodeSchema) *DescribeLatestRecordSchemaResponseBodyPlaybookNodeSchema {
	s.NodeSchema = v
	return s
}

func (s *DescribeLatestRecordSchemaResponseBodyPlaybookNodeSchema) Validate() error {
	if s.NodeSchema != nil {
		for _, item := range s.NodeSchema {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeLatestRecordSchemaResponseBodyPlaybookNodeSchemaNodeSchema struct {
	// The action name of the component.
	//
	// example:
	//
	// formatedata
	ActionName *string `json:"ActionName,omitempty" xml:"ActionName,omitempty"`
	// The name of the component.
	//
	// example:
	//
	// DataFormat
	ComponentName *string `json:"ComponentName,omitempty" xml:"ComponentName,omitempty"`
	// The name of the node.
	//
	// example:
	//
	// DataFormat_1
	NodeName *string `json:"NodeName,omitempty" xml:"NodeName,omitempty"`
	// The output fields.
	OutputFields []*string `json:"OutputFields,omitempty" xml:"OutputFields,omitempty" type:"Repeated"`
}

func (s DescribeLatestRecordSchemaResponseBodyPlaybookNodeSchemaNodeSchema) String() string {
	return dara.Prettify(s)
}

func (s DescribeLatestRecordSchemaResponseBodyPlaybookNodeSchemaNodeSchema) GoString() string {
	return s.String()
}

func (s *DescribeLatestRecordSchemaResponseBodyPlaybookNodeSchemaNodeSchema) GetActionName() *string {
	return s.ActionName
}

func (s *DescribeLatestRecordSchemaResponseBodyPlaybookNodeSchemaNodeSchema) GetComponentName() *string {
	return s.ComponentName
}

func (s *DescribeLatestRecordSchemaResponseBodyPlaybookNodeSchemaNodeSchema) GetNodeName() *string {
	return s.NodeName
}

func (s *DescribeLatestRecordSchemaResponseBodyPlaybookNodeSchemaNodeSchema) GetOutputFields() []*string {
	return s.OutputFields
}

func (s *DescribeLatestRecordSchemaResponseBodyPlaybookNodeSchemaNodeSchema) SetActionName(v string) *DescribeLatestRecordSchemaResponseBodyPlaybookNodeSchemaNodeSchema {
	s.ActionName = &v
	return s
}

func (s *DescribeLatestRecordSchemaResponseBodyPlaybookNodeSchemaNodeSchema) SetComponentName(v string) *DescribeLatestRecordSchemaResponseBodyPlaybookNodeSchemaNodeSchema {
	s.ComponentName = &v
	return s
}

func (s *DescribeLatestRecordSchemaResponseBodyPlaybookNodeSchemaNodeSchema) SetNodeName(v string) *DescribeLatestRecordSchemaResponseBodyPlaybookNodeSchemaNodeSchema {
	s.NodeName = &v
	return s
}

func (s *DescribeLatestRecordSchemaResponseBodyPlaybookNodeSchemaNodeSchema) SetOutputFields(v []*string) *DescribeLatestRecordSchemaResponseBodyPlaybookNodeSchemaNodeSchema {
	s.OutputFields = v
	return s
}

func (s *DescribeLatestRecordSchemaResponseBodyPlaybookNodeSchemaNodeSchema) Validate() error {
	return dara.Validate(s)
}
