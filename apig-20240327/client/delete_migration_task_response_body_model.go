// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteMigrationTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DeleteMigrationTaskResponseBody
	GetCode() *string
	SetMessage(v string) *DeleteMigrationTaskResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteMigrationTaskResponseBody
	GetRequestId() *string
}

type DeleteMigrationTaskResponseBody struct {
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
	// CB56C0CE-37C8-5C5A-8A07-DFBF083A40D2
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s DeleteMigrationTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteMigrationTaskResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteMigrationTaskResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeleteMigrationTaskResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteMigrationTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteMigrationTaskResponseBody) SetCode(v string) *DeleteMigrationTaskResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteMigrationTaskResponseBody) SetMessage(v string) *DeleteMigrationTaskResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteMigrationTaskResponseBody) SetRequestId(v string) *DeleteMigrationTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteMigrationTaskResponseBody) Validate() error {
	return dara.Validate(s)
}
