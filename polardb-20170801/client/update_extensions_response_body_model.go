// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateExtensionsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateExtensionsResponseBody
	GetRequestId() *string
}

type UpdateExtensionsResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// A83B6055-F891-5596-96E3-52D62567DFBB
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateExtensionsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateExtensionsResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateExtensionsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateExtensionsResponseBody) SetRequestId(v string) *UpdateExtensionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateExtensionsResponseBody) Validate() error {
	return dara.Validate(s)
}
