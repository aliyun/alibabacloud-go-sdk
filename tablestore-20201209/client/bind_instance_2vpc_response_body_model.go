// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBindInstance2VpcResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDomain(v string) *BindInstance2VpcResponseBody
	GetDomain() *string
	SetEndpoint(v string) *BindInstance2VpcResponseBody
	GetEndpoint() *string
	SetRequestId(v string) *BindInstance2VpcResponseBody
	GetRequestId() *string
}

type BindInstance2VpcResponseBody struct {
	// example:
	//
	// xu6666-mkdd-test.cn-hangzhou.vpc.ots.aliyuncs.com
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// example:
	//
	// 172.**.***.34
	Endpoint *string `json:"Endpoint,omitempty" xml:"Endpoint,omitempty"`
	// example:
	//
	// E734979F-5A44-5993-9CE5-C23103576923
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s BindInstance2VpcResponseBody) String() string {
	return dara.Prettify(s)
}

func (s BindInstance2VpcResponseBody) GoString() string {
	return s.String()
}

func (s *BindInstance2VpcResponseBody) GetDomain() *string {
	return s.Domain
}

func (s *BindInstance2VpcResponseBody) GetEndpoint() *string {
	return s.Endpoint
}

func (s *BindInstance2VpcResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *BindInstance2VpcResponseBody) SetDomain(v string) *BindInstance2VpcResponseBody {
	s.Domain = &v
	return s
}

func (s *BindInstance2VpcResponseBody) SetEndpoint(v string) *BindInstance2VpcResponseBody {
	s.Endpoint = &v
	return s
}

func (s *BindInstance2VpcResponseBody) SetRequestId(v string) *BindInstance2VpcResponseBody {
	s.RequestId = &v
	return s
}

func (s *BindInstance2VpcResponseBody) Validate() error {
	return dara.Validate(s)
}
