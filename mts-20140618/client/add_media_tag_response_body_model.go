// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddMediaTagResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *AddMediaTagResponseBody
	GetRequestId() *string
}

type AddMediaTagResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 91B6CAB9-034C-4E4E-A40B-E7F5C81E1A2K
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AddMediaTagResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AddMediaTagResponseBody) GoString() string {
	return s.String()
}

func (s *AddMediaTagResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AddMediaTagResponseBody) SetRequestId(v string) *AddMediaTagResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddMediaTagResponseBody) Validate() error {
	return dara.Validate(s)
}
