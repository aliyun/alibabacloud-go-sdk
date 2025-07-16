// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetRowsVisibilityResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetId(v string) *SetRowsVisibilityResponseBody
	GetId() *string
	SetRequestId(v string) *SetRowsVisibilityResponseBody
	GetRequestId() *string
}

type SetRowsVisibilityResponseBody struct {
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

func (s SetRowsVisibilityResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SetRowsVisibilityResponseBody) GoString() string {
	return s.String()
}

func (s *SetRowsVisibilityResponseBody) GetId() *string {
	return s.Id
}

func (s *SetRowsVisibilityResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SetRowsVisibilityResponseBody) SetId(v string) *SetRowsVisibilityResponseBody {
	s.Id = &v
	return s
}

func (s *SetRowsVisibilityResponseBody) SetRequestId(v string) *SetRowsVisibilityResponseBody {
	s.RequestId = &v
	return s
}

func (s *SetRowsVisibilityResponseBody) Validate() error {
	return dara.Validate(s)
}
