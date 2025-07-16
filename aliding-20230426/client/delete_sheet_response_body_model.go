// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteSheetResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteSheetResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteSheetResponseBody
	GetSuccess() *bool
}

type DeleteSheetResponseBody struct {
	// requestId
	//
	// example:
	//
	// 0FAAEC9C-C6C8-5C87-AF8E-1195889BBXXX
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s DeleteSheetResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteSheetResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteSheetResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteSheetResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteSheetResponseBody) SetRequestId(v string) *DeleteSheetResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteSheetResponseBody) SetSuccess(v bool) *DeleteSheetResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteSheetResponseBody) Validate() error {
	return dara.Validate(s)
}
