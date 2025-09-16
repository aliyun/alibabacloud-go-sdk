// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDatabaseSchemaResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *GetDatabaseSchemaResponseBody
	GetRequestId() *string
	SetResult(v []*GetDatabaseSchemaResponseBodyResult) *GetDatabaseSchemaResponseBody
	GetResult() []*GetDatabaseSchemaResponseBodyResult
}

type GetDatabaseSchemaResponseBody struct {
	// id of request
	//
	// example:
	//
	// 022F36C7-9FB4-5D67-BEBC-3D14B0984463
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// List
	Result []*GetDatabaseSchemaResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
}

func (s GetDatabaseSchemaResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetDatabaseSchemaResponseBody) GoString() string {
	return s.String()
}

func (s *GetDatabaseSchemaResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetDatabaseSchemaResponseBody) GetResult() []*GetDatabaseSchemaResponseBodyResult {
	return s.Result
}

func (s *GetDatabaseSchemaResponseBody) SetRequestId(v string) *GetDatabaseSchemaResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetDatabaseSchemaResponseBody) SetResult(v []*GetDatabaseSchemaResponseBodyResult) *GetDatabaseSchemaResponseBody {
	s.Result = v
	return s
}

func (s *GetDatabaseSchemaResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetDatabaseSchemaResponseBodyResult struct {
	// example:
	//
	// id
	FieldName *string `json:"fieldName,omitempty" xml:"fieldName,omitempty"`
	// example:
	//
	// STRING
	FieldType *string `json:"fieldType,omitempty" xml:"fieldType,omitempty"`
	// example:
	//
	// FT_UINT64
	FieldTypeDetail map[string]interface{} `json:"fieldTypeDetail,omitempty" xml:"fieldTypeDetail,omitempty"`
	// example:
	//
	// test_tusou_v2
	IndexName *string `json:"indexName,omitempty" xml:"indexName,omitempty"`
	// example:
	//
	// NUMBER
	IndexType *string `json:"indexType,omitempty" xml:"indexType,omitempty"`
}

func (s GetDatabaseSchemaResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s GetDatabaseSchemaResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GetDatabaseSchemaResponseBodyResult) GetFieldName() *string {
	return s.FieldName
}

func (s *GetDatabaseSchemaResponseBodyResult) GetFieldType() *string {
	return s.FieldType
}

func (s *GetDatabaseSchemaResponseBodyResult) GetFieldTypeDetail() map[string]interface{} {
	return s.FieldTypeDetail
}

func (s *GetDatabaseSchemaResponseBodyResult) GetIndexName() *string {
	return s.IndexName
}

func (s *GetDatabaseSchemaResponseBodyResult) GetIndexType() *string {
	return s.IndexType
}

func (s *GetDatabaseSchemaResponseBodyResult) SetFieldName(v string) *GetDatabaseSchemaResponseBodyResult {
	s.FieldName = &v
	return s
}

func (s *GetDatabaseSchemaResponseBodyResult) SetFieldType(v string) *GetDatabaseSchemaResponseBodyResult {
	s.FieldType = &v
	return s
}

func (s *GetDatabaseSchemaResponseBodyResult) SetFieldTypeDetail(v map[string]interface{}) *GetDatabaseSchemaResponseBodyResult {
	s.FieldTypeDetail = v
	return s
}

func (s *GetDatabaseSchemaResponseBodyResult) SetIndexName(v string) *GetDatabaseSchemaResponseBodyResult {
	s.IndexName = &v
	return s
}

func (s *GetDatabaseSchemaResponseBodyResult) SetIndexType(v string) *GetDatabaseSchemaResponseBodyResult {
	s.IndexType = &v
	return s
}

func (s *GetDatabaseSchemaResponseBodyResult) Validate() error {
	return dara.Validate(s)
}
