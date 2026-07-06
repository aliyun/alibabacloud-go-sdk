// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUnbindAgentStorage2VpcResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UnbindAgentStorage2VpcResponseBody
	GetRequestId() *string
}

type UnbindAgentStorage2VpcResponseBody struct {
	// The request ID, which can be used to troubleshoot and locate issues.
	//
	// example:
	//
	// 18DD77BF-F967-576D-80D1-79121399AB53
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UnbindAgentStorage2VpcResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UnbindAgentStorage2VpcResponseBody) GoString() string {
	return s.String()
}

func (s *UnbindAgentStorage2VpcResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UnbindAgentStorage2VpcResponseBody) SetRequestId(v string) *UnbindAgentStorage2VpcResponseBody {
	s.RequestId = &v
	return s
}

func (s *UnbindAgentStorage2VpcResponseBody) Validate() error {
	return dara.Validate(s)
}
