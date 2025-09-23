// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteExtensionsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteExtensionsResponseBody
	GetRequestId() *string
}

type DeleteExtensionsResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// 11EB8492-C17F-5FC2-9E27-1440A5C2F583
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteExtensionsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteExtensionsResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteExtensionsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteExtensionsResponseBody) SetRequestId(v string) *DeleteExtensionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteExtensionsResponseBody) Validate() error {
	return dara.Validate(s)
}
