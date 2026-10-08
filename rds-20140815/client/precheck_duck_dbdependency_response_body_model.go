// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPrecheckDuckDBDependencyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetFailedCheckItems(v []*PrecheckDuckDBDependencyResponseBodyFailedCheckItems) *PrecheckDuckDBDependencyResponseBody
	GetFailedCheckItems() []*PrecheckDuckDBDependencyResponseBodyFailedCheckItems
	SetResult(v bool) *PrecheckDuckDBDependencyResponseBody
	GetResult() *bool
}

type PrecheckDuckDBDependencyResponseBody struct {
	// The items that do not meet the prerequisites for creating a DuckDB-based analytical instance.
	FailedCheckItems []*PrecheckDuckDBDependencyResponseBodyFailedCheckItems `json:"FailedCheckItems,omitempty" xml:"FailedCheckItems,omitempty" type:"Repeated"`
	// Indicates whether the prerequisite check for creating a DuckDB-based analytical instance is passed. Valid values:
	//
	// - **true**: The check is passed.
	//
	// - **false**: The check is not passed.
	//
	// example:
	//
	// false
	Result *bool `json:"Result,omitempty" xml:"Result,omitempty"`
}

func (s PrecheckDuckDBDependencyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s PrecheckDuckDBDependencyResponseBody) GoString() string {
	return s.String()
}

func (s *PrecheckDuckDBDependencyResponseBody) GetFailedCheckItems() []*PrecheckDuckDBDependencyResponseBodyFailedCheckItems {
	return s.FailedCheckItems
}

func (s *PrecheckDuckDBDependencyResponseBody) GetResult() *bool {
	return s.Result
}

func (s *PrecheckDuckDBDependencyResponseBody) SetFailedCheckItems(v []*PrecheckDuckDBDependencyResponseBodyFailedCheckItems) *PrecheckDuckDBDependencyResponseBody {
	s.FailedCheckItems = v
	return s
}

func (s *PrecheckDuckDBDependencyResponseBody) SetResult(v bool) *PrecheckDuckDBDependencyResponseBody {
	s.Result = &v
	return s
}

func (s *PrecheckDuckDBDependencyResponseBody) Validate() error {
	if s.FailedCheckItems != nil {
		for _, item := range s.FailedCheckItems {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type PrecheckDuckDBDependencyResponseBodyFailedCheckItems struct {
	// Indicates whether the item can be fixed with one click.
	//
	// - **true**: The item can be fixed with one click by calling the [ModifyDBInstanceConfig](https://help.aliyun.com/document_detail/2623684.html) operation.
	//
	// - **false**: The item cannot be fixed with one click.
	//
	//
	// 	Notice: If the major engine version of the database instance does not meet the requirements, you must perform a [manual upgrade](https://help.aliyun.com/document_detail/2623684.html).
	//
	// example:
	//
	// false
	AllowAutoModify *bool `json:"AllowAutoModify,omitempty" xml:"AllowAutoModify,omitempty"`
	// The current value of the check item.
	//
	// example:
	//
	// 15.0
	CurrentValue *string `json:"CurrentValue,omitempty" xml:"CurrentValue,omitempty"`
	// The name of the check item.
	//
	// example:
	//
	// MajorVersion
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The target value or target range of the check item.
	//
	// example:
	//
	// 17.0
	RequiredValue *string `json:"RequiredValue,omitempty" xml:"RequiredValue,omitempty"`
	// The check item type. Valid values:
	//
	// - **Parameter**: parameter.
	//
	// - **MinorVersion**: minor engine version.
	//
	// - **MajorVersion**: major engine version.
	//
	// example:
	//
	// Parameter
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s PrecheckDuckDBDependencyResponseBodyFailedCheckItems) String() string {
	return dara.Prettify(s)
}

func (s PrecheckDuckDBDependencyResponseBodyFailedCheckItems) GoString() string {
	return s.String()
}

func (s *PrecheckDuckDBDependencyResponseBodyFailedCheckItems) GetAllowAutoModify() *bool {
	return s.AllowAutoModify
}

func (s *PrecheckDuckDBDependencyResponseBodyFailedCheckItems) GetCurrentValue() *string {
	return s.CurrentValue
}

func (s *PrecheckDuckDBDependencyResponseBodyFailedCheckItems) GetName() *string {
	return s.Name
}

func (s *PrecheckDuckDBDependencyResponseBodyFailedCheckItems) GetRequiredValue() *string {
	return s.RequiredValue
}

func (s *PrecheckDuckDBDependencyResponseBodyFailedCheckItems) GetType() *string {
	return s.Type
}

func (s *PrecheckDuckDBDependencyResponseBodyFailedCheckItems) SetAllowAutoModify(v bool) *PrecheckDuckDBDependencyResponseBodyFailedCheckItems {
	s.AllowAutoModify = &v
	return s
}

func (s *PrecheckDuckDBDependencyResponseBodyFailedCheckItems) SetCurrentValue(v string) *PrecheckDuckDBDependencyResponseBodyFailedCheckItems {
	s.CurrentValue = &v
	return s
}

func (s *PrecheckDuckDBDependencyResponseBodyFailedCheckItems) SetName(v string) *PrecheckDuckDBDependencyResponseBodyFailedCheckItems {
	s.Name = &v
	return s
}

func (s *PrecheckDuckDBDependencyResponseBodyFailedCheckItems) SetRequiredValue(v string) *PrecheckDuckDBDependencyResponseBodyFailedCheckItems {
	s.RequiredValue = &v
	return s
}

func (s *PrecheckDuckDBDependencyResponseBodyFailedCheckItems) SetType(v string) *PrecheckDuckDBDependencyResponseBodyFailedCheckItems {
	s.Type = &v
	return s
}

func (s *PrecheckDuckDBDependencyResponseBodyFailedCheckItems) Validate() error {
	return dara.Validate(s)
}
