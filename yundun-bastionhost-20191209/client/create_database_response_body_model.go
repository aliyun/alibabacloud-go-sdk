// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDatabaseResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDatabaseId(v string) *CreateDatabaseResponseBody
	GetDatabaseId() *string
	SetRequestId(v string) *CreateDatabaseResponseBody
	GetRequestId() *string
}

type CreateDatabaseResponseBody struct {
	// The database ID.
	//
	// example:
	//
	// 334
	DatabaseId *string `json:"DatabaseId,omitempty" xml:"DatabaseId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 98EDD923-236C-5A88-88E7-4979A91B9325
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateDatabaseResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateDatabaseResponseBody) GoString() string {
	return s.String()
}

func (s *CreateDatabaseResponseBody) GetDatabaseId() *string {
	return s.DatabaseId
}

func (s *CreateDatabaseResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateDatabaseResponseBody) SetDatabaseId(v string) *CreateDatabaseResponseBody {
	s.DatabaseId = &v
	return s
}

func (s *CreateDatabaseResponseBody) SetRequestId(v string) *CreateDatabaseResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateDatabaseResponseBody) Validate() error {
	return dara.Validate(s)
}
