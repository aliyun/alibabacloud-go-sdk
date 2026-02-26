// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSimpleQuery interface {
	dara.Model
	String() string
	GoString() string
	SetField(v string) *SimpleQuery
	GetField() *string
	SetOperation(v string) *SimpleQuery
	GetOperation() *string
	SetSubQueries(v []*SimpleQuery) *SimpleQuery
	GetSubQueries() []*SimpleQuery
	SetValue(v string) *SimpleQuery
	GetValue() *string
}

type SimpleQuery struct {
	// This parameter is required. The field name. For a list of the supported fields, see [Supported fields and operators](https://help.aliyun.com/document_detail/252856.html).
	//
	// example:
	//
	// Size
	Field *string `json:"Field,omitempty" xml:"Field,omitempty"`
	// This parameter is required. The operator.
	//
	// Enumerated values:
	//
	// 	- exist: exists query.
	//
	// 	- not: logical NOT.
	//
	// 	- or: logical OR.
	//
	// 	- prefix: prefix query.
	//
	// 	- and: logical AND.
	//
	// 	- It: less than.
	//
	// 	- match-phrase: string match query.
	//
	// 	- gte: greater than or equal to.
	//
	// 	- eq: equal to.
	//
	// 	- lte: less than or equal to.
	//
	// 	- gt: greater than.
	//
	// 	- nested: You can perform logical condition queries within the same object when the data type of a field is ARRAY.
	//
	// example:
	//
	// and
	Operation *string `json:"Operation,omitempty" xml:"Operation,omitempty"`
	// The subquery structure.
	//
	// You can configure Subquery conditions only if you set the Operation parameter to and, or, not, or nested.
	//
	// If you set the Operation parameter to and, or, or not, all query conditions specified by the SubQueries parameter must comply with the logical relationship of the parent query condition.
	//
	// If you set the Operation parameter to nested, the parent field of a subquery must be of the ARRAY type, such as Labels. The operator of a subquery condition must be one or more of the following operators: and, or, and not. The field of the subquery must be a sub-field of the parent field.
	//
	// For information about how to call the SimpleQuery operation, see [SimpleQuery](https://help.aliyun.com/document_detail/478175.html).
	SubQueries []*SimpleQuery `json:"SubQueries,omitempty" xml:"SubQueries,omitempty" type:"Repeated"`
	// The field value. If you set the Operation parameter to and, or, not, or nested, this parameter is invalid.
	//
	// example:
	//
	// 10
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s SimpleQuery) String() string {
	return dara.Prettify(s)
}

func (s SimpleQuery) GoString() string {
	return s.String()
}

func (s *SimpleQuery) GetField() *string {
	return s.Field
}

func (s *SimpleQuery) GetOperation() *string {
	return s.Operation
}

func (s *SimpleQuery) GetSubQueries() []*SimpleQuery {
	return s.SubQueries
}

func (s *SimpleQuery) GetValue() *string {
	return s.Value
}

func (s *SimpleQuery) SetField(v string) *SimpleQuery {
	s.Field = &v
	return s
}

func (s *SimpleQuery) SetOperation(v string) *SimpleQuery {
	s.Operation = &v
	return s
}

func (s *SimpleQuery) SetSubQueries(v []*SimpleQuery) *SimpleQuery {
	s.SubQueries = v
	return s
}

func (s *SimpleQuery) SetValue(v string) *SimpleQuery {
	s.Value = &v
	return s
}

func (s *SimpleQuery) Validate() error {
	if s.SubQueries != nil {
		for _, item := range s.SubQueries {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
