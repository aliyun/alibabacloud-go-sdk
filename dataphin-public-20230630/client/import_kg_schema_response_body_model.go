// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iImportKgSchemaResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ImportKgSchemaResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *ImportKgSchemaResponseBody
	GetHttpStatusCode() *int32
	SetImportResult(v *ImportKgSchemaResponseBodyImportResult) *ImportKgSchemaResponseBody
	GetImportResult() *ImportKgSchemaResponseBodyImportResult
	SetMessage(v string) *ImportKgSchemaResponseBody
	GetMessage() *string
	SetRequestId(v string) *ImportKgSchemaResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ImportKgSchemaResponseBody
	GetSuccess() *bool
}

type ImportKgSchemaResponseBody struct {
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32                                  `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	ImportResult   *ImportKgSchemaResponseBodyImportResult `json:"ImportResult,omitempty" xml:"ImportResult,omitempty" type:"Struct"`
	// example:
	//
	// internal error
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 82E78D6B-AA8F-1FEF-8AA3-5C9DA2A79140
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ImportKgSchemaResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ImportKgSchemaResponseBody) GoString() string {
	return s.String()
}

func (s *ImportKgSchemaResponseBody) GetCode() *string {
	return s.Code
}

func (s *ImportKgSchemaResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ImportKgSchemaResponseBody) GetImportResult() *ImportKgSchemaResponseBodyImportResult {
	return s.ImportResult
}

func (s *ImportKgSchemaResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ImportKgSchemaResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ImportKgSchemaResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ImportKgSchemaResponseBody) SetCode(v string) *ImportKgSchemaResponseBody {
	s.Code = &v
	return s
}

func (s *ImportKgSchemaResponseBody) SetHttpStatusCode(v int32) *ImportKgSchemaResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ImportKgSchemaResponseBody) SetImportResult(v *ImportKgSchemaResponseBodyImportResult) *ImportKgSchemaResponseBody {
	s.ImportResult = v
	return s
}

func (s *ImportKgSchemaResponseBody) SetMessage(v string) *ImportKgSchemaResponseBody {
	s.Message = &v
	return s
}

func (s *ImportKgSchemaResponseBody) SetRequestId(v string) *ImportKgSchemaResponseBody {
	s.RequestId = &v
	return s
}

func (s *ImportKgSchemaResponseBody) SetSuccess(v bool) *ImportKgSchemaResponseBody {
	s.Success = &v
	return s
}

func (s *ImportKgSchemaResponseBody) Validate() error {
	if s.ImportResult != nil {
		if err := s.ImportResult.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ImportKgSchemaResponseBodyImportResult struct {
	// example:
	//
	// 1032591
	EntityTypeCount *int32 `json:"EntityTypeCount,omitempty" xml:"EntityTypeCount,omitempty"`
	// example:
	//
	// 3380766
	RelationTypeCount *int32 `json:"RelationTypeCount,omitempty" xml:"RelationTypeCount,omitempty"`
}

func (s ImportKgSchemaResponseBodyImportResult) String() string {
	return dara.Prettify(s)
}

func (s ImportKgSchemaResponseBodyImportResult) GoString() string {
	return s.String()
}

func (s *ImportKgSchemaResponseBodyImportResult) GetEntityTypeCount() *int32 {
	return s.EntityTypeCount
}

func (s *ImportKgSchemaResponseBodyImportResult) GetRelationTypeCount() *int32 {
	return s.RelationTypeCount
}

func (s *ImportKgSchemaResponseBodyImportResult) SetEntityTypeCount(v int32) *ImportKgSchemaResponseBodyImportResult {
	s.EntityTypeCount = &v
	return s
}

func (s *ImportKgSchemaResponseBodyImportResult) SetRelationTypeCount(v int32) *ImportKgSchemaResponseBodyImportResult {
	s.RelationTypeCount = &v
	return s
}

func (s *ImportKgSchemaResponseBodyImportResult) Validate() error {
	return dara.Validate(s)
}
