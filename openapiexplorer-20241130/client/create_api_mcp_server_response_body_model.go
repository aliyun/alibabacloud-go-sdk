// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateApiMcpServerResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetId(v string) *CreateApiMcpServerResponseBody
	GetId() *string
	SetRequestId(v string) *CreateApiMcpServerResponseBody
	GetRequestId() *string
	SetUrls(v *CreateApiMcpServerResponseBodyUrls) *CreateApiMcpServerResponseBody
	GetUrls() *CreateApiMcpServerResponseBodyUrls
}

type CreateApiMcpServerResponseBody struct {
	// example:
	//
	// v6ZZ7ftCzEILW***
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// example:
	//
	// 9BFC4AC1-6BE4-5405-BDEC-CA288D404812
	RequestId *string                             `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Urls      *CreateApiMcpServerResponseBodyUrls `json:"urls,omitempty" xml:"urls,omitempty" type:"Struct"`
}

func (s CreateApiMcpServerResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateApiMcpServerResponseBody) GoString() string {
	return s.String()
}

func (s *CreateApiMcpServerResponseBody) GetId() *string {
	return s.Id
}

func (s *CreateApiMcpServerResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateApiMcpServerResponseBody) GetUrls() *CreateApiMcpServerResponseBodyUrls {
	return s.Urls
}

func (s *CreateApiMcpServerResponseBody) SetId(v string) *CreateApiMcpServerResponseBody {
	s.Id = &v
	return s
}

func (s *CreateApiMcpServerResponseBody) SetRequestId(v string) *CreateApiMcpServerResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateApiMcpServerResponseBody) SetUrls(v *CreateApiMcpServerResponseBodyUrls) *CreateApiMcpServerResponseBody {
	s.Urls = v
	return s
}

func (s *CreateApiMcpServerResponseBody) Validate() error {
	if s.Urls != nil {
		if err := s.Urls.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateApiMcpServerResponseBodyUrls struct {
	// example:
	//
	// https://mcpserverinner-pre.cn-zhangjiakou.aliyuncs.com/accounts/xxxx/custom/xxx/id/xxxx/mcp
	Mcp *string `json:"mcp,omitempty" xml:"mcp,omitempty"`
	// example:
	//
	// https://mcpserverinner-pre.cn-zhangjiakou.aliyuncs.com/accounts/xxxx/custom/xxx/id/xxxx/sse
	Sse    *string `json:"sse,omitempty" xml:"sse,omitempty"`
	VpcMcp *string `json:"vpcMcp,omitempty" xml:"vpcMcp,omitempty"`
	VpcSse *string `json:"vpcSse,omitempty" xml:"vpcSse,omitempty"`
}

func (s CreateApiMcpServerResponseBodyUrls) String() string {
	return dara.Prettify(s)
}

func (s CreateApiMcpServerResponseBodyUrls) GoString() string {
	return s.String()
}

func (s *CreateApiMcpServerResponseBodyUrls) GetMcp() *string {
	return s.Mcp
}

func (s *CreateApiMcpServerResponseBodyUrls) GetSse() *string {
	return s.Sse
}

func (s *CreateApiMcpServerResponseBodyUrls) GetVpcMcp() *string {
	return s.VpcMcp
}

func (s *CreateApiMcpServerResponseBodyUrls) GetVpcSse() *string {
	return s.VpcSse
}

func (s *CreateApiMcpServerResponseBodyUrls) SetMcp(v string) *CreateApiMcpServerResponseBodyUrls {
	s.Mcp = &v
	return s
}

func (s *CreateApiMcpServerResponseBodyUrls) SetSse(v string) *CreateApiMcpServerResponseBodyUrls {
	s.Sse = &v
	return s
}

func (s *CreateApiMcpServerResponseBodyUrls) SetVpcMcp(v string) *CreateApiMcpServerResponseBodyUrls {
	s.VpcMcp = &v
	return s
}

func (s *CreateApiMcpServerResponseBodyUrls) SetVpcSse(v string) *CreateApiMcpServerResponseBodyUrls {
	s.VpcSse = &v
	return s
}

func (s *CreateApiMcpServerResponseBodyUrls) Validate() error {
	return dara.Validate(s)
}
