// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteVerifyResultRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDeleteAfterQuery(v string) *DeleteVerifyResultRequest
	GetDeleteAfterQuery() *string
	SetDeleteType(v string) *DeleteVerifyResultRequest
	GetDeleteType() *string
	SetTransactionId(v string) *DeleteVerifyResultRequest
	GetTransactionId() *string
}

type DeleteVerifyResultRequest struct {
	// Whether to depend on the query interface when deleting data
	//
	// example:
	//
	// Y / N
	DeleteAfterQuery *string `json:"DeleteAfterQuery,omitempty" xml:"DeleteAfterQuery,omitempty"`
	// Type of data to be deleted
	//
	// example:
	//
	// Img / Text / All
	DeleteType *string `json:"DeleteType,omitempty" xml:"DeleteType,omitempty"`
	// Unique identifier of the authentication request
	//
	// example:
	//
	// 4ab0b***cbde97
	TransactionId *string `json:"TransactionId,omitempty" xml:"TransactionId,omitempty"`
}

func (s DeleteVerifyResultRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteVerifyResultRequest) GoString() string {
	return s.String()
}

func (s *DeleteVerifyResultRequest) GetDeleteAfterQuery() *string {
	return s.DeleteAfterQuery
}

func (s *DeleteVerifyResultRequest) GetDeleteType() *string {
	return s.DeleteType
}

func (s *DeleteVerifyResultRequest) GetTransactionId() *string {
	return s.TransactionId
}

func (s *DeleteVerifyResultRequest) SetDeleteAfterQuery(v string) *DeleteVerifyResultRequest {
	s.DeleteAfterQuery = &v
	return s
}

func (s *DeleteVerifyResultRequest) SetDeleteType(v string) *DeleteVerifyResultRequest {
	s.DeleteType = &v
	return s
}

func (s *DeleteVerifyResultRequest) SetTransactionId(v string) *DeleteVerifyResultRequest {
	s.TransactionId = &v
	return s
}

func (s *DeleteVerifyResultRequest) Validate() error {
	return dara.Validate(s)
}
