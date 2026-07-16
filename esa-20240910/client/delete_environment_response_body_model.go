// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteEnvironmentResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteEnvironmentResponseBody
	GetRequestId() *string
}

type DeleteEnvironmentResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// C370DAF1-C838-4288-A1A0-9A87633D248E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteEnvironmentResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteEnvironmentResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteEnvironmentResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteEnvironmentResponseBody) SetRequestId(v string) *DeleteEnvironmentResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteEnvironmentResponseBody) Validate() error {
	return dara.Validate(s)
}
