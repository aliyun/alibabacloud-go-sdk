// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteServiceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteServiceResponseBody
	GetRequestId() *string
}

type DeleteServiceResponseBody struct {
	// Request ID
	//
	// example:
	//
	// 51B6A3E8-EA9E-5143-BE11-8E5F83474C95
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s DeleteServiceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteServiceResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteServiceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteServiceResponseBody) SetRequestId(v string) *DeleteServiceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteServiceResponseBody) Validate() error {
	return dara.Validate(s)
}
