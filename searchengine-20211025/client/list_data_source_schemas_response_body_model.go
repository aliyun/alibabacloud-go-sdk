// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDataSourceSchemasResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ListDataSourceSchemasResponseBody
	GetRequestId() *string
	SetResult(v []*ListDataSourceSchemasResponseBodyResult) *ListDataSourceSchemasResponseBody
	GetResult() []*ListDataSourceSchemasResponseBodyResult
}

type ListDataSourceSchemasResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 022F36C7-9FB4-5D67-BEBC-3D14B0984463
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The results returned.
	Result []*ListDataSourceSchemasResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
}

func (s ListDataSourceSchemasResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListDataSourceSchemasResponseBody) GoString() string {
	return s.String()
}

func (s *ListDataSourceSchemasResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListDataSourceSchemasResponseBody) GetResult() []*ListDataSourceSchemasResponseBodyResult {
	return s.Result
}

func (s *ListDataSourceSchemasResponseBody) SetRequestId(v string) *ListDataSourceSchemasResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListDataSourceSchemasResponseBody) SetResult(v []*ListDataSourceSchemasResponseBodyResult) *ListDataSourceSchemasResponseBody {
	s.Result = v
	return s
}

func (s *ListDataSourceSchemasResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListDataSourceSchemasResponseBodyResult struct {
	// Indicates whether the field has the index attribute. Valid values: **true*	- and **false**.
	//
	// example:
	//
	// false
	AddIndex *bool `json:"addIndex,omitempty" xml:"addIndex,omitempty"`
	// Indicates whether the field is an attribute field. Valid values: **true*	- and **false**.
	//
	// example:
	//
	// false
	Attribute *bool `json:"attribute,omitempty" xml:"attribute,omitempty"`
	// Indicates whether the field is a custom field. Valid values: **true*	- and **false**.
	//
	// example:
	//
	// false
	Custom *bool `json:"custom,omitempty" xml:"custom,omitempty"`
	// The field name.
	//
	// example:
	//
	// test
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The primary key field.
	PrimaryKey *ListDataSourceSchemasResponseBodyResultPrimaryKey `json:"primaryKey,omitempty" xml:"primaryKey,omitempty" type:"Struct"`
	// Indicates whether the field can be displayed. Valid values: **true*	- and **false**.
	//
	// example:
	//
	// false
	Summary *bool `json:"summary,omitempty" xml:"summary,omitempty"`
	// The field type.
	//
	// example:
	//
	// STRING
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s ListDataSourceSchemasResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s ListDataSourceSchemasResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListDataSourceSchemasResponseBodyResult) GetAddIndex() *bool {
	return s.AddIndex
}

func (s *ListDataSourceSchemasResponseBodyResult) GetAttribute() *bool {
	return s.Attribute
}

func (s *ListDataSourceSchemasResponseBodyResult) GetCustom() *bool {
	return s.Custom
}

func (s *ListDataSourceSchemasResponseBodyResult) GetName() *string {
	return s.Name
}

func (s *ListDataSourceSchemasResponseBodyResult) GetPrimaryKey() *ListDataSourceSchemasResponseBodyResultPrimaryKey {
	return s.PrimaryKey
}

func (s *ListDataSourceSchemasResponseBodyResult) GetSummary() *bool {
	return s.Summary
}

func (s *ListDataSourceSchemasResponseBodyResult) GetType() *string {
	return s.Type
}

func (s *ListDataSourceSchemasResponseBodyResult) SetAddIndex(v bool) *ListDataSourceSchemasResponseBodyResult {
	s.AddIndex = &v
	return s
}

func (s *ListDataSourceSchemasResponseBodyResult) SetAttribute(v bool) *ListDataSourceSchemasResponseBodyResult {
	s.Attribute = &v
	return s
}

func (s *ListDataSourceSchemasResponseBodyResult) SetCustom(v bool) *ListDataSourceSchemasResponseBodyResult {
	s.Custom = &v
	return s
}

func (s *ListDataSourceSchemasResponseBodyResult) SetName(v string) *ListDataSourceSchemasResponseBodyResult {
	s.Name = &v
	return s
}

func (s *ListDataSourceSchemasResponseBodyResult) SetPrimaryKey(v *ListDataSourceSchemasResponseBodyResultPrimaryKey) *ListDataSourceSchemasResponseBodyResult {
	s.PrimaryKey = v
	return s
}

func (s *ListDataSourceSchemasResponseBodyResult) SetSummary(v bool) *ListDataSourceSchemasResponseBodyResult {
	s.Summary = &v
	return s
}

func (s *ListDataSourceSchemasResponseBodyResult) SetType(v string) *ListDataSourceSchemasResponseBodyResult {
	s.Type = &v
	return s
}

func (s *ListDataSourceSchemasResponseBodyResult) Validate() error {
	return dara.Validate(s)
}

type ListDataSourceSchemasResponseBodyResultPrimaryKey struct {
	// Indicates whether the field has the primary key attribute. Valid values: **true*	- and **false**.
	//
	// example:
	//
	// false
	HasPrimaryKeyAttribute *bool `json:"hasPrimaryKeyAttribute,omitempty" xml:"hasPrimaryKeyAttribute,omitempty"`
	// Indicates whether the field is the primary key. Valid values: **true*	- and **false**.
	//
	// example:
	//
	// false
	IsPrimaryKey *bool `json:"isPrimaryKey,omitempty" xml:"isPrimaryKey,omitempty"`
	// Indicates whether the field can be sorted. Valid values: **true*	- and **false**.
	//
	// example:
	//
	// false
	IsPrimaryKeySorted *bool `json:"isPrimaryKeySorted,omitempty" xml:"isPrimaryKeySorted,omitempty"`
}

func (s ListDataSourceSchemasResponseBodyResultPrimaryKey) String() string {
	return dara.Prettify(s)
}

func (s ListDataSourceSchemasResponseBodyResultPrimaryKey) GoString() string {
	return s.String()
}

func (s *ListDataSourceSchemasResponseBodyResultPrimaryKey) GetHasPrimaryKeyAttribute() *bool {
	return s.HasPrimaryKeyAttribute
}

func (s *ListDataSourceSchemasResponseBodyResultPrimaryKey) GetIsPrimaryKey() *bool {
	return s.IsPrimaryKey
}

func (s *ListDataSourceSchemasResponseBodyResultPrimaryKey) GetIsPrimaryKeySorted() *bool {
	return s.IsPrimaryKeySorted
}

func (s *ListDataSourceSchemasResponseBodyResultPrimaryKey) SetHasPrimaryKeyAttribute(v bool) *ListDataSourceSchemasResponseBodyResultPrimaryKey {
	s.HasPrimaryKeyAttribute = &v
	return s
}

func (s *ListDataSourceSchemasResponseBodyResultPrimaryKey) SetIsPrimaryKey(v bool) *ListDataSourceSchemasResponseBodyResultPrimaryKey {
	s.IsPrimaryKey = &v
	return s
}

func (s *ListDataSourceSchemasResponseBodyResultPrimaryKey) SetIsPrimaryKeySorted(v bool) *ListDataSourceSchemasResponseBodyResultPrimaryKey {
	s.IsPrimaryKeySorted = &v
	return s
}

func (s *ListDataSourceSchemasResponseBodyResultPrimaryKey) Validate() error {
	return dara.Validate(s)
}
