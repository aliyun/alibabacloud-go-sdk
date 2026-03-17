// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeactiveFlowLogResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeactiveFlowLogResponseBody
	GetRequestId() *string
}

type DeactiveFlowLogResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 5F2BAEAF-96D4-4BE5-A031-6CAD7B145D0D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeactiveFlowLogResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeactiveFlowLogResponseBody) GoString() string {
	return s.String()
}

func (s *DeactiveFlowLogResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeactiveFlowLogResponseBody) SetRequestId(v string) *DeactiveFlowLogResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeactiveFlowLogResponseBody) Validate() error {
	return dara.Validate(s)
}
