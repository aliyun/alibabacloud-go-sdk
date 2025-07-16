// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInsertRowsBeforeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetId(v string) *InsertRowsBeforeResponseBody
	GetId() *string
	SetRequestId(v string) *InsertRowsBeforeResponseBody
	GetRequestId() *string
}

type InsertRowsBeforeResponseBody struct {
	// example:
	//
	// stxxxx
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// requestId
	//
	// example:
	//
	// 0FAAEC9C-C6C8-5C87-AF8E-1195889BBXXX
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s InsertRowsBeforeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s InsertRowsBeforeResponseBody) GoString() string {
	return s.String()
}

func (s *InsertRowsBeforeResponseBody) GetId() *string {
	return s.Id
}

func (s *InsertRowsBeforeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *InsertRowsBeforeResponseBody) SetId(v string) *InsertRowsBeforeResponseBody {
	s.Id = &v
	return s
}

func (s *InsertRowsBeforeResponseBody) SetRequestId(v string) *InsertRowsBeforeResponseBody {
	s.RequestId = &v
	return s
}

func (s *InsertRowsBeforeResponseBody) Validate() error {
	return dara.Validate(s)
}
