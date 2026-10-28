// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateContactResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateContactResponseBody
	GetRequestId() *string
}

type UpdateContactResponseBody struct {
	// The unique ID of the request.
	//
	// example:
	//
	// 8FDE2569-626B-5176-9844-28877A*****
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s UpdateContactResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateContactResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateContactResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateContactResponseBody) SetRequestId(v string) *UpdateContactResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateContactResponseBody) Validate() error {
	return dara.Validate(s)
}
