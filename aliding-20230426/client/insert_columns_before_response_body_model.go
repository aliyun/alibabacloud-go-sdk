// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInsertColumnsBeforeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetId(v string) *InsertColumnsBeforeResponseBody
	GetId() *string
	SetRequestId(v string) *InsertColumnsBeforeResponseBody
	GetRequestId() *string
}

type InsertColumnsBeforeResponseBody struct {
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

func (s InsertColumnsBeforeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s InsertColumnsBeforeResponseBody) GoString() string {
	return s.String()
}

func (s *InsertColumnsBeforeResponseBody) GetId() *string {
	return s.Id
}

func (s *InsertColumnsBeforeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *InsertColumnsBeforeResponseBody) SetId(v string) *InsertColumnsBeforeResponseBody {
	s.Id = &v
	return s
}

func (s *InsertColumnsBeforeResponseBody) SetRequestId(v string) *InsertColumnsBeforeResponseBody {
	s.RequestId = &v
	return s
}

func (s *InsertColumnsBeforeResponseBody) Validate() error {
	return dara.Validate(s)
}
