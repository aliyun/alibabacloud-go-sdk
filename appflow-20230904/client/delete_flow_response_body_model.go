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
	// Id of the request
	//
	// example:
	//
	// D6382BE4-9F02-5F03-B26E-BF74FC521842
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
