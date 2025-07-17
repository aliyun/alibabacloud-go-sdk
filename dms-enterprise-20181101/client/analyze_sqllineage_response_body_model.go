// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAnalyzeSQLLineageResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *AnalyzeSQLLineageResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *AnalyzeSQLLineageResponseBody
	GetErrorMessage() *string
	SetLineageResult(v *AnalyzeSQLLineageResponseBodyLineageResult) *AnalyzeSQLLineageResponseBody
	GetLineageResult() *AnalyzeSQLLineageResponseBodyLineageResult
	SetRequestId(v string) *AnalyzeSQLLineageResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *AnalyzeSQLLineageResponseBody
	GetSuccess() *bool
}

type AnalyzeSQLLineageResponseBody struct {
	// The error code returned.
	//
	// example:
	//
	// UnknownError
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message returned if the request failed.
	//
	// example:
	//
	// UnknownError
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// Returned data set of SQL lineage.
	LineageResult *AnalyzeSQLLineageResponseBodyLineageResult `json:"LineageResult,omitempty" xml:"LineageResult,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// B4B07137-F6AE-4756-8474-7F92BB6C4E04
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s AnalyzeSQLLineageResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AnalyzeSQLLineageResponseBody) GoString() string {
	return s.String()
}

func (s *AnalyzeSQLLineageResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *AnalyzeSQLLineageResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *AnalyzeSQLLineageResponseBody) GetLineageResult() *AnalyzeSQLLineageResponseBodyLineageResult {
	return s.LineageResult
}

func (s *AnalyzeSQLLineageResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AnalyzeSQLLineageResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *AnalyzeSQLLineageResponseBody) SetErrorCode(v string) *AnalyzeSQLLineageResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *AnalyzeSQLLineageResponseBody) SetErrorMessage(v string) *AnalyzeSQLLineageResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *AnalyzeSQLLineageResponseBody) SetLineageResult(v *AnalyzeSQLLineageResponseBodyLineageResult) *AnalyzeSQLLineageResponseBody {
	s.LineageResult = v
	return s
}

func (s *AnalyzeSQLLineageResponseBody) SetRequestId(v string) *AnalyzeSQLLineageResponseBody {
	s.RequestId = &v
	return s
}

func (s *AnalyzeSQLLineageResponseBody) SetSuccess(v bool) *AnalyzeSQLLineageResponseBody {
	s.Success = &v
	return s
}

func (s *AnalyzeSQLLineageResponseBody) Validate() error {
	return dara.Validate(s)
}

type AnalyzeSQLLineageResponseBodyLineageResult struct {
	// The details about the lineage.
	Lineages []*AnalyzeSQLLineageResponseBodyLineageResultLineages `json:"Lineages,omitempty" xml:"Lineages,omitempty" type:"Repeated"`
	// The table and field metadata information.
	ObjectMetadata []*AnalyzeSQLLineageResponseBodyLineageResultObjectMetadata `json:"ObjectMetadata,omitempty" xml:"ObjectMetadata,omitempty" type:"Repeated"`
}

func (s AnalyzeSQLLineageResponseBodyLineageResult) String() string {
	return dara.Prettify(s)
}

func (s AnalyzeSQLLineageResponseBodyLineageResult) GoString() string {
	return s.String()
}

func (s *AnalyzeSQLLineageResponseBodyLineageResult) GetLineages() []*AnalyzeSQLLineageResponseBodyLineageResultLineages {
	return s.Lineages
}

func (s *AnalyzeSQLLineageResponseBodyLineageResult) GetObjectMetadata() []*AnalyzeSQLLineageResponseBodyLineageResultObjectMetadata {
	return s.ObjectMetadata
}

func (s *AnalyzeSQLLineageResponseBodyLineageResult) SetLineages(v []*AnalyzeSQLLineageResponseBodyLineageResultLineages) *AnalyzeSQLLineageResponseBodyLineageResult {
	s.Lineages = v
	return s
}

func (s *AnalyzeSQLLineageResponseBodyLineageResult) SetObjectMetadata(v []*AnalyzeSQLLineageResponseBodyLineageResultObjectMetadata) *AnalyzeSQLLineageResponseBodyLineageResult {
	s.ObjectMetadata = v
	return s
}

func (s *AnalyzeSQLLineageResponseBodyLineageResult) Validate() error {
	return dara.Validate(s)
}

type AnalyzeSQLLineageResponseBodyLineageResultLineages struct {
	// The target.
	//
	// example:
	//
	// dmstest.a.id
	Dst *string `json:"Dst,omitempty" xml:"Dst,omitempty"`
	// The type of the lineage. Valid values:
	//
	// 	- **FIELD_DEPEND_FIELD**: Fields depend on fields.
	//
	// 	- **TABLE_DEPEND_TABLE**: Tables depend on tables.
	//
	// 	- **FIELD_INFLU_TABLE**: Fields influence tables.
	//
	// 	- **FIELD_INFLU_FIELD**: Fields influence fields.
	//
	// 	- **FIELD_INFLU_TABLE**: Tables influence fields.
	//
	// 	- **FIELD_JOIN_FIELD**: Fields are associated with fields.
	//
	// example:
	//
	// FIELD_DEPEND_FIELD
	LineageType *string `json:"LineageType,omitempty" xml:"LineageType,omitempty"`
	// The operation type of the SQL statement in which the data lineage is generated. For example, if the operation type is SELECT, the data lineage is generated from a SELECT statement.
	//
	// >  This field is an extended field which has no practical use.
	//
	// example:
	//
	// SELECT
	OperType *string `json:"OperType,omitempty" xml:"OperType,omitempty"`
	// The handling details. This parameter is returned only when LineageType is FIELD_DEPEND_FIELD.
	ProcessDetail *AnalyzeSQLLineageResponseBodyLineageResultLineagesProcessDetail `json:"ProcessDetail,omitempty" xml:"ProcessDetail,omitempty" type:"Struct"`
	// The source.
	//
	// example:
	//
	// dmstest.b.id
	Src *string `json:"Src,omitempty" xml:"Src,omitempty"`
}

func (s AnalyzeSQLLineageResponseBodyLineageResultLineages) String() string {
	return dara.Prettify(s)
}

func (s AnalyzeSQLLineageResponseBodyLineageResultLineages) GoString() string {
	return s.String()
}

func (s *AnalyzeSQLLineageResponseBodyLineageResultLineages) GetDst() *string {
	return s.Dst
}

func (s *AnalyzeSQLLineageResponseBodyLineageResultLineages) GetLineageType() *string {
	return s.LineageType
}

func (s *AnalyzeSQLLineageResponseBodyLineageResultLineages) GetOperType() *string {
	return s.OperType
}

func (s *AnalyzeSQLLineageResponseBodyLineageResultLineages) GetProcessDetail() *AnalyzeSQLLineageResponseBodyLineageResultLineagesProcessDetail {
	return s.ProcessDetail
}

func (s *AnalyzeSQLLineageResponseBodyLineageResultLineages) GetSrc() *string {
	return s.Src
}

func (s *AnalyzeSQLLineageResponseBodyLineageResultLineages) SetDst(v string) *AnalyzeSQLLineageResponseBodyLineageResultLineages {
	s.Dst = &v
	return s
}

func (s *AnalyzeSQLLineageResponseBodyLineageResultLineages) SetLineageType(v string) *AnalyzeSQLLineageResponseBodyLineageResultLineages {
	s.LineageType = &v
	return s
}

func (s *AnalyzeSQLLineageResponseBodyLineageResultLineages) SetOperType(v string) *AnalyzeSQLLineageResponseBodyLineageResultLineages {
	s.OperType = &v
	return s
}

func (s *AnalyzeSQLLineageResponseBodyLineageResultLineages) SetProcessDetail(v *AnalyzeSQLLineageResponseBodyLineageResultLineagesProcessDetail) *AnalyzeSQLLineageResponseBodyLineageResultLineages {
	s.ProcessDetail = v
	return s
}

func (s *AnalyzeSQLLineageResponseBodyLineageResultLineages) SetSrc(v string) *AnalyzeSQLLineageResponseBodyLineageResultLineages {
	s.Src = &v
	return s
}

func (s *AnalyzeSQLLineageResponseBodyLineageResultLineages) Validate() error {
	return dara.Validate(s)
}

type AnalyzeSQLLineageResponseBodyLineageResultLineagesProcessDetail struct {
	// The calculating method. Valid values:
	//
	// 	- **DIRECT**: No function or expression is used.
	//
	// 	- **EXPR**: A function or expression is used.
	//
	// example:
	//
	// DIRECT
	CalWay *string `json:"CalWay,omitempty" xml:"CalWay,omitempty"`
	// The SQL code snippet for field processing.
	//
	// example:
	//
	// dmstest.b.id
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
}

func (s AnalyzeSQLLineageResponseBodyLineageResultLineagesProcessDetail) String() string {
	return dara.Prettify(s)
}

func (s AnalyzeSQLLineageResponseBodyLineageResultLineagesProcessDetail) GoString() string {
	return s.String()
}

func (s *AnalyzeSQLLineageResponseBodyLineageResultLineagesProcessDetail) GetCalWay() *string {
	return s.CalWay
}

func (s *AnalyzeSQLLineageResponseBodyLineageResultLineagesProcessDetail) GetCode() *string {
	return s.Code
}

func (s *AnalyzeSQLLineageResponseBodyLineageResultLineagesProcessDetail) SetCalWay(v string) *AnalyzeSQLLineageResponseBodyLineageResultLineagesProcessDetail {
	s.CalWay = &v
	return s
}

func (s *AnalyzeSQLLineageResponseBodyLineageResultLineagesProcessDetail) SetCode(v string) *AnalyzeSQLLineageResponseBodyLineageResultLineagesProcessDetail {
	s.Code = &v
	return s
}

func (s *AnalyzeSQLLineageResponseBodyLineageResultLineagesProcessDetail) Validate() error {
	return dara.Validate(s)
}

type AnalyzeSQLLineageResponseBodyLineageResultObjectMetadata struct {
	// The fields in the metatable.
	Fields []*AnalyzeSQLLineageResponseBodyLineageResultObjectMetadataFields `json:"Fields,omitempty" xml:"Fields,omitempty" type:"Repeated"`
	// The object name.
	//
	// example:
	//
	// a
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The source of metadata. Valid values:
	//
	// 	- **DDL**: The metadata comes from parsed SQL statements or definition of databases and tables collected by DMS.
	//
	// 	- **LINEAGE**: The metadata comes from lineage analysis results.
	//
	// example:
	//
	// DDL
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
	// The object type. Valid values:
	//
	// 	- **TABLE**
	//
	// 	- **VIEW**
	//
	// 	- **TMP_TABLE**
	//
	// example:
	//
	// TABLE
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s AnalyzeSQLLineageResponseBodyLineageResultObjectMetadata) String() string {
	return dara.Prettify(s)
}

func (s AnalyzeSQLLineageResponseBodyLineageResultObjectMetadata) GoString() string {
	return s.String()
}

func (s *AnalyzeSQLLineageResponseBodyLineageResultObjectMetadata) GetFields() []*AnalyzeSQLLineageResponseBodyLineageResultObjectMetadataFields {
	return s.Fields
}

func (s *AnalyzeSQLLineageResponseBodyLineageResultObjectMetadata) GetName() *string {
	return s.Name
}

func (s *AnalyzeSQLLineageResponseBodyLineageResultObjectMetadata) GetSource() *string {
	return s.Source
}

func (s *AnalyzeSQLLineageResponseBodyLineageResultObjectMetadata) GetType() *string {
	return s.Type
}

func (s *AnalyzeSQLLineageResponseBodyLineageResultObjectMetadata) SetFields(v []*AnalyzeSQLLineageResponseBodyLineageResultObjectMetadataFields) *AnalyzeSQLLineageResponseBodyLineageResultObjectMetadata {
	s.Fields = v
	return s
}

func (s *AnalyzeSQLLineageResponseBodyLineageResultObjectMetadata) SetName(v string) *AnalyzeSQLLineageResponseBodyLineageResultObjectMetadata {
	s.Name = &v
	return s
}

func (s *AnalyzeSQLLineageResponseBodyLineageResultObjectMetadata) SetSource(v string) *AnalyzeSQLLineageResponseBodyLineageResultObjectMetadata {
	s.Source = &v
	return s
}

func (s *AnalyzeSQLLineageResponseBodyLineageResultObjectMetadata) SetType(v string) *AnalyzeSQLLineageResponseBodyLineageResultObjectMetadata {
	s.Type = &v
	return s
}

func (s *AnalyzeSQLLineageResponseBodyLineageResultObjectMetadata) Validate() error {
	return dara.Validate(s)
}

type AnalyzeSQLLineageResponseBodyLineageResultObjectMetadataFields struct {
	// The name of the field.
	//
	// example:
	//
	// dmstest.a.id
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s AnalyzeSQLLineageResponseBodyLineageResultObjectMetadataFields) String() string {
	return dara.Prettify(s)
}

func (s AnalyzeSQLLineageResponseBodyLineageResultObjectMetadataFields) GoString() string {
	return s.String()
}

func (s *AnalyzeSQLLineageResponseBodyLineageResultObjectMetadataFields) GetName() *string {
	return s.Name
}

func (s *AnalyzeSQLLineageResponseBodyLineageResultObjectMetadataFields) SetName(v string) *AnalyzeSQLLineageResponseBodyLineageResultObjectMetadataFields {
	s.Name = &v
	return s
}

func (s *AnalyzeSQLLineageResponseBodyLineageResultObjectMetadataFields) Validate() error {
	return dara.Validate(s)
}
