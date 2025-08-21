// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyHttp2EnableResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyHttp2EnableResponseBody
	GetRequestId() *string
}

type ModifyHttp2EnableResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// C33EB3D5-AF96-43CA-9C7E-37A81BC06A1E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyHttp2EnableResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyHttp2EnableResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyHttp2EnableResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyHttp2EnableResponseBody) SetRequestId(v string) *ModifyHttp2EnableResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyHttp2EnableResponseBody) Validate() error {
	return dara.Validate(s)
}
