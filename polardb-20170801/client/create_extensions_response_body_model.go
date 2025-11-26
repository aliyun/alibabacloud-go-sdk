// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateExtensionsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateExtensionsResponseBody
	GetRequestId() *string
}

type CreateExtensionsResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateExtensionsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateExtensionsResponseBody) GoString() string {
	return s.String()
}

func (s *CreateExtensionsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateExtensionsResponseBody) SetRequestId(v string) *CreateExtensionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateExtensionsResponseBody) Validate() error {
	return dara.Validate(s)
}
