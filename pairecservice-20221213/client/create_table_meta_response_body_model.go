// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateTableMetaResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateTableMetaResponseBody
	GetRequestId() *string
	SetTableMetaId(v string) *CreateTableMetaResponseBody
	GetTableMetaId() *string
}

type CreateTableMetaResponseBody struct {
	// example:
	//
	// 728C5E01-ABF6-5AA8-B9FC-B3BA05DECC77
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 2
	TableMetaId *string `json:"TableMetaId,omitempty" xml:"TableMetaId,omitempty"`
}

func (s CreateTableMetaResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateTableMetaResponseBody) GoString() string {
	return s.String()
}

func (s *CreateTableMetaResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateTableMetaResponseBody) GetTableMetaId() *string {
	return s.TableMetaId
}

func (s *CreateTableMetaResponseBody) SetRequestId(v string) *CreateTableMetaResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateTableMetaResponseBody) SetTableMetaId(v string) *CreateTableMetaResponseBody {
	s.TableMetaId = &v
	return s
}

func (s *CreateTableMetaResponseBody) Validate() error {
	return dara.Validate(s)
}
