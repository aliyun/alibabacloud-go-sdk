// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTableColumnsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *GetTableColumnsResponseBody
	GetRequestId() *string
	SetResult(v []*GetTableColumnsResponseBodyResult) *GetTableColumnsResponseBody
	GetResult() []*GetTableColumnsResponseBodyResult
}

type GetTableColumnsResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 1CC93E65-6734-5060-BEF7-0EB0A4862BCF
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The returned result.
	Result []*GetTableColumnsResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
}

func (s GetTableColumnsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetTableColumnsResponseBody) GoString() string {
	return s.String()
}

func (s *GetTableColumnsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetTableColumnsResponseBody) GetResult() []*GetTableColumnsResponseBodyResult {
	return s.Result
}

func (s *GetTableColumnsResponseBody) SetRequestId(v string) *GetTableColumnsResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetTableColumnsResponseBody) SetResult(v []*GetTableColumnsResponseBodyResult) *GetTableColumnsResponseBody {
	s.Result = v
	return s
}

func (s *GetTableColumnsResponseBody) Validate() error {
	if s.Result != nil {
		for _, item := range s.Result {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetTableColumnsResponseBodyResult struct {
	// The field description.
	//
	// example:
	//
	// 主键字段
	Comment *string `json:"comment,omitempty" xml:"comment,omitempty"`
	// The field name.
	//
	// example:
	//
	// id
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// Indicates whether the field is a primary key.
	//
	// example:
	//
	// true
	Primary *bool `json:"primary,omitempty" xml:"primary,omitempty"`
	// The field type.
	//
	// example:
	//
	// BIGINT
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s GetTableColumnsResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s GetTableColumnsResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GetTableColumnsResponseBodyResult) GetComment() *string {
	return s.Comment
}

func (s *GetTableColumnsResponseBodyResult) GetName() *string {
	return s.Name
}

func (s *GetTableColumnsResponseBodyResult) GetPrimary() *bool {
	return s.Primary
}

func (s *GetTableColumnsResponseBodyResult) GetType() *string {
	return s.Type
}

func (s *GetTableColumnsResponseBodyResult) SetComment(v string) *GetTableColumnsResponseBodyResult {
	s.Comment = &v
	return s
}

func (s *GetTableColumnsResponseBodyResult) SetName(v string) *GetTableColumnsResponseBodyResult {
	s.Name = &v
	return s
}

func (s *GetTableColumnsResponseBodyResult) SetPrimary(v bool) *GetTableColumnsResponseBodyResult {
	s.Primary = &v
	return s
}

func (s *GetTableColumnsResponseBodyResult) SetType(v string) *GetTableColumnsResponseBodyResult {
	s.Type = &v
	return s
}

func (s *GetTableColumnsResponseBodyResult) Validate() error {
	return dara.Validate(s)
}
