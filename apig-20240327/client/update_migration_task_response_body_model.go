// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateMigrationTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *UpdateMigrationTaskResponseBody
	GetCode() *string
	SetMessage(v string) *UpdateMigrationTaskResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateMigrationTaskResponseBody
	GetRequestId() *string
}

type UpdateMigrationTaskResponseBody struct {
	// example:
	//
	// 200
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// E7357D40-2418-5D4B-AC2E-D0A3C5930C7F
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s UpdateMigrationTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateMigrationTaskResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateMigrationTaskResponseBody) GetCode() *string {
	return s.Code
}

func (s *UpdateMigrationTaskResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateMigrationTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateMigrationTaskResponseBody) SetCode(v string) *UpdateMigrationTaskResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateMigrationTaskResponseBody) SetMessage(v string) *UpdateMigrationTaskResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateMigrationTaskResponseBody) SetRequestId(v string) *UpdateMigrationTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateMigrationTaskResponseBody) Validate() error {
	return dara.Validate(s)
}
