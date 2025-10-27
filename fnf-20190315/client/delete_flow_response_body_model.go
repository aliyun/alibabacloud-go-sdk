// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteFlowResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteFlowResponseBody
	GetRequestId() *string
}

type DeleteFlowResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// testRequestId
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteFlowResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteFlowResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteFlowResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteFlowResponseBody) SetRequestId(v string) *DeleteFlowResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteFlowResponseBody) Validate() error {
	return dara.Validate(s)
}
