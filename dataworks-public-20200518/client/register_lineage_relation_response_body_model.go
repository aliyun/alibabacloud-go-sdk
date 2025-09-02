// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRegisterLineageRelationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *RegisterLineageRelationResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *RegisterLineageRelationResponseBody
	GetErrorMessage() *string
	SetHttpStatusCode(v int32) *RegisterLineageRelationResponseBody
	GetHttpStatusCode() *int32
	SetLineageRelation(v *RegisterLineageRelationResponseBodyLineageRelation) *RegisterLineageRelationResponseBody
	GetLineageRelation() *RegisterLineageRelationResponseBodyLineageRelation
	SetRequestId(v string) *RegisterLineageRelationResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *RegisterLineageRelationResponseBody
	GetSuccess() *bool
}

type RegisterLineageRelationResponseBody struct {
	// The error code.
	//
	// example:
	//
	// 1010210001
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message.
	//
	// example:
	//
	// The entity of lineage not exist, xxx
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The lineage.
	LineageRelation *RegisterLineageRelationResponseBodyLineageRelation `json:"LineageRelation,omitempty" xml:"LineageRelation,omitempty" type:"Struct"`
	// The request ID. You can locate logs and troubleshoot issues based on the ID.
	//
	// example:
	//
	// EE50E05E-028C-182B-9xxx
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s RegisterLineageRelationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RegisterLineageRelationResponseBody) GoString() string {
	return s.String()
}

func (s *RegisterLineageRelationResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *RegisterLineageRelationResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *RegisterLineageRelationResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *RegisterLineageRelationResponseBody) GetLineageRelation() *RegisterLineageRelationResponseBodyLineageRelation {
	return s.LineageRelation
}

func (s *RegisterLineageRelationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RegisterLineageRelationResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *RegisterLineageRelationResponseBody) SetErrorCode(v string) *RegisterLineageRelationResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *RegisterLineageRelationResponseBody) SetErrorMessage(v string) *RegisterLineageRelationResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *RegisterLineageRelationResponseBody) SetHttpStatusCode(v int32) *RegisterLineageRelationResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *RegisterLineageRelationResponseBody) SetLineageRelation(v *RegisterLineageRelationResponseBodyLineageRelation) *RegisterLineageRelationResponseBody {
	s.LineageRelation = v
	return s
}

func (s *RegisterLineageRelationResponseBody) SetRequestId(v string) *RegisterLineageRelationResponseBody {
	s.RequestId = &v
	return s
}

func (s *RegisterLineageRelationResponseBody) SetSuccess(v bool) *RegisterLineageRelationResponseBody {
	s.Success = &v
	return s
}

func (s *RegisterLineageRelationResponseBody) Validate() error {
	return dara.Validate(s)
}

type RegisterLineageRelationResponseBodyLineageRelation struct {
	// The unique identifier of the destination entity.
	//
	// example:
	//
	// custom-report.month_stat_user
	DestEntityQualifiedName *string `json:"DestEntityQualifiedName,omitempty" xml:"DestEntityQualifiedName,omitempty"`
	// The ID of the lineage between entities.
	//
	// example:
	//
	// dfsldfdlsfdsaaaabbbb
	RelationshipGuid *string `json:"RelationshipGuid,omitempty" xml:"RelationshipGuid,omitempty"`
	// The unique identifier of the source entity.
	//
	// example:
	//
	// maxcompute-table.project.table
	SrcEntityQualifiedName *string `json:"SrcEntityQualifiedName,omitempty" xml:"SrcEntityQualifiedName,omitempty"`
}

func (s RegisterLineageRelationResponseBodyLineageRelation) String() string {
	return dara.Prettify(s)
}

func (s RegisterLineageRelationResponseBodyLineageRelation) GoString() string {
	return s.String()
}

func (s *RegisterLineageRelationResponseBodyLineageRelation) GetDestEntityQualifiedName() *string {
	return s.DestEntityQualifiedName
}

func (s *RegisterLineageRelationResponseBodyLineageRelation) GetRelationshipGuid() *string {
	return s.RelationshipGuid
}

func (s *RegisterLineageRelationResponseBodyLineageRelation) GetSrcEntityQualifiedName() *string {
	return s.SrcEntityQualifiedName
}

func (s *RegisterLineageRelationResponseBodyLineageRelation) SetDestEntityQualifiedName(v string) *RegisterLineageRelationResponseBodyLineageRelation {
	s.DestEntityQualifiedName = &v
	return s
}

func (s *RegisterLineageRelationResponseBodyLineageRelation) SetRelationshipGuid(v string) *RegisterLineageRelationResponseBodyLineageRelation {
	s.RelationshipGuid = &v
	return s
}

func (s *RegisterLineageRelationResponseBodyLineageRelation) SetSrcEntityQualifiedName(v string) *RegisterLineageRelationResponseBodyLineageRelation {
	s.SrcEntityQualifiedName = &v
	return s
}

func (s *RegisterLineageRelationResponseBodyLineageRelation) Validate() error {
	return dara.Validate(s)
}
