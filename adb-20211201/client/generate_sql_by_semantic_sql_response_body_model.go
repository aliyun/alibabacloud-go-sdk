// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGenerateSqlBySemanticSqlResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *GenerateSqlBySemanticSqlResponseBodyData) *GenerateSqlBySemanticSqlResponseBody
	GetData() *GenerateSqlBySemanticSqlResponseBodyData
	SetRequestId(v string) *GenerateSqlBySemanticSqlResponseBody
	GetRequestId() *string
}

type GenerateSqlBySemanticSqlResponseBody struct {
	// The returned data.
	Data *GenerateSqlBySemanticSqlResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// Id of the request
	//
	// example:
	//
	// 1AD222E9-E606-4A42-BF6D-8A4442913CEF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GenerateSqlBySemanticSqlResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GenerateSqlBySemanticSqlResponseBody) GoString() string {
	return s.String()
}

func (s *GenerateSqlBySemanticSqlResponseBody) GetData() *GenerateSqlBySemanticSqlResponseBodyData {
	return s.Data
}

func (s *GenerateSqlBySemanticSqlResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GenerateSqlBySemanticSqlResponseBody) SetData(v *GenerateSqlBySemanticSqlResponseBodyData) *GenerateSqlBySemanticSqlResponseBody {
	s.Data = v
	return s
}

func (s *GenerateSqlBySemanticSqlResponseBody) SetRequestId(v string) *GenerateSqlBySemanticSqlResponseBody {
	s.RequestId = &v
	return s
}

func (s *GenerateSqlBySemanticSqlResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GenerateSqlBySemanticSqlResponseBodyData struct {
	// The error message returned when the task fails to be created.
	//
	// example:
	//
	// Failed to rewrite semantic SQL: Ambiguous path from \\"lineitem\\" to \\"nation\\". Multiple paths found
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The generated executable SQL statement.
	//
	// example:
	//
	// select sum(amount) from orders
	GeneratedSql *string `json:"GeneratedSql,omitempty" xml:"GeneratedSql,omitempty"`
	// Indicates whether the generation request was successful. Valid values:
	//
	// - **true**: Successful.
	//
	// - **false**: Failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GenerateSqlBySemanticSqlResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GenerateSqlBySemanticSqlResponseBodyData) GoString() string {
	return s.String()
}

func (s *GenerateSqlBySemanticSqlResponseBodyData) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *GenerateSqlBySemanticSqlResponseBodyData) GetGeneratedSql() *string {
	return s.GeneratedSql
}

func (s *GenerateSqlBySemanticSqlResponseBodyData) GetSuccess() *bool {
	return s.Success
}

func (s *GenerateSqlBySemanticSqlResponseBodyData) SetErrorMessage(v string) *GenerateSqlBySemanticSqlResponseBodyData {
	s.ErrorMessage = &v
	return s
}

func (s *GenerateSqlBySemanticSqlResponseBodyData) SetGeneratedSql(v string) *GenerateSqlBySemanticSqlResponseBodyData {
	s.GeneratedSql = &v
	return s
}

func (s *GenerateSqlBySemanticSqlResponseBodyData) SetSuccess(v bool) *GenerateSqlBySemanticSqlResponseBodyData {
	s.Success = &v
	return s
}

func (s *GenerateSqlBySemanticSqlResponseBodyData) Validate() error {
	return dara.Validate(s)
}
