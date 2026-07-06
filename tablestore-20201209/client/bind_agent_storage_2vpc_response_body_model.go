// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBindAgentStorage2VpcResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDomain(v string) *BindAgentStorage2VpcResponseBody
	GetDomain() *string
	SetEndpoint(v string) *BindAgentStorage2VpcResponseBody
	GetEndpoint() *string
	SetRequestId(v string) *BindAgentStorage2VpcResponseBody
	GetRequestId() *string
}

type BindAgentStorage2VpcResponseBody struct {
	// The domain name.
	//
	// example:
	//
	// remua-agent-test.cn-beijing.vpc.ots.aliyuncs.com
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// The endpoint of the instance.
	//
	// example:
	//
	// 172.**.***.34
	Endpoint *string `json:"Endpoint,omitempty" xml:"Endpoint,omitempty"`
	// The request ID, which can be used to troubleshoot and locate issues.
	//
	// example:
	//
	// 39871ED2-62C0-578F-A32E-B88072D5582F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s BindAgentStorage2VpcResponseBody) String() string {
	return dara.Prettify(s)
}

func (s BindAgentStorage2VpcResponseBody) GoString() string {
	return s.String()
}

func (s *BindAgentStorage2VpcResponseBody) GetDomain() *string {
	return s.Domain
}

func (s *BindAgentStorage2VpcResponseBody) GetEndpoint() *string {
	return s.Endpoint
}

func (s *BindAgentStorage2VpcResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *BindAgentStorage2VpcResponseBody) SetDomain(v string) *BindAgentStorage2VpcResponseBody {
	s.Domain = &v
	return s
}

func (s *BindAgentStorage2VpcResponseBody) SetEndpoint(v string) *BindAgentStorage2VpcResponseBody {
	s.Endpoint = &v
	return s
}

func (s *BindAgentStorage2VpcResponseBody) SetRequestId(v string) *BindAgentStorage2VpcResponseBody {
	s.RequestId = &v
	return s
}

func (s *BindAgentStorage2VpcResponseBody) Validate() error {
	return dara.Validate(s)
}
