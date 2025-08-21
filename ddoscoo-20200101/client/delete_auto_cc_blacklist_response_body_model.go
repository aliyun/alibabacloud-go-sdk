// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAutoCcBlacklistResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteAutoCcBlacklistResponseBody
	GetRequestId() *string
}

type DeleteAutoCcBlacklistResponseBody struct {
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// C33EB3D5-AF96-43CA-9C7E-37A81BC06A1E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteAutoCcBlacklistResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteAutoCcBlacklistResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteAutoCcBlacklistResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteAutoCcBlacklistResponseBody) SetRequestId(v string) *DeleteAutoCcBlacklistResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteAutoCcBlacklistResponseBody) Validate() error {
	return dara.Validate(s)
}
