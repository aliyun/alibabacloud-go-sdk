// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteClientKeyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteClientKeyResponseBody
	GetRequestId() *string
}

type DeleteClientKeyResponseBody struct {
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// 2312e45f-b2fa-4c34-ad94-3eca50932916
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteClientKeyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteClientKeyResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteClientKeyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteClientKeyResponseBody) SetRequestId(v string) *DeleteClientKeyResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteClientKeyResponseBody) Validate() error {
	return dara.Validate(s)
}
