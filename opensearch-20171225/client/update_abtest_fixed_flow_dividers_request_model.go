// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateABTestFixedFlowDividersRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v []*string) *UpdateABTestFixedFlowDividersRequest
	GetBody() []*string
}

type UpdateABTestFixedFlowDividersRequest struct {
	// The request body.
	Body []*string `json:"body,omitempty" xml:"body,omitempty" type:"Repeated"`
}

func (s UpdateABTestFixedFlowDividersRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateABTestFixedFlowDividersRequest) GoString() string {
	return s.String()
}

func (s *UpdateABTestFixedFlowDividersRequest) GetBody() []*string {
	return s.Body
}

func (s *UpdateABTestFixedFlowDividersRequest) SetBody(v []*string) *UpdateABTestFixedFlowDividersRequest {
	s.Body = v
	return s
}

func (s *UpdateABTestFixedFlowDividersRequest) Validate() error {
	return dara.Validate(s)
}
