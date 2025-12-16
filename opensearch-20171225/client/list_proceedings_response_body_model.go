// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListProceedingsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ListProceedingsResponseBody
	GetRequestId() *string
}

type ListProceedingsResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// F5099063-6B86-F398-D843-905F9EFB683A
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s ListProceedingsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListProceedingsResponseBody) GoString() string {
	return s.String()
}

func (s *ListProceedingsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListProceedingsResponseBody) SetRequestId(v string) *ListProceedingsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListProceedingsResponseBody) Validate() error {
	return dara.Validate(s)
}
