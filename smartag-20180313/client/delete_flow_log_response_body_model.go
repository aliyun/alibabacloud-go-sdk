// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteFlowLogResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteFlowLogResponseBody
	GetRequestId() *string
}

type DeleteFlowLogResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// B05AF87C-0FE5-42D7-BEA3-665D90FAA71D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteFlowLogResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteFlowLogResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteFlowLogResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteFlowLogResponseBody) SetRequestId(v string) *DeleteFlowLogResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteFlowLogResponseBody) Validate() error {
	return dara.Validate(s)
}
