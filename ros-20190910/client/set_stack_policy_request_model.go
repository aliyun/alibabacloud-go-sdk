// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetStackPolicyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRegionId(v string) *SetStackPolicyRequest
	GetRegionId() *string
	SetStackId(v string) *SetStackPolicyRequest
	GetStackId() *string
	SetStackPolicyBody(v string) *SetStackPolicyRequest
	GetStackPolicyBody() *string
	SetStackPolicyURL(v string) *SetStackPolicyRequest
	GetStackPolicyURL() *string
}

type SetStackPolicyRequest struct {
	// The region ID of the stack. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/131035.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the stack.
	//
	// This parameter is required.
	//
	// example:
	//
	// 4a6c9851-3b0f-4f5f-b4ca-a14bf691****
	StackId *string `json:"StackId,omitempty" xml:"StackId,omitempty"`
	// The structure that contains the stack policy body. The stack policy body must be 1 to 16,384 bytes in length.
	//
	// You can specify one of the StackPolicyBody and StackPolicyURL parameters, but you cannot specify both of them.
	//
	// example:
	//
	// {"Statement":[{"Effect":"Allow","Action":"Update:*","Principal":"*","Resource":"*"}]}
	StackPolicyBody *string `json:"StackPolicyBody,omitempty" xml:"StackPolicyBody,omitempty"`
	// The URL for the file that contains the stack policy. The URL must point to a template located in an HTTP or HTTPS web server or an Alibaba Cloud OSS bucket. Examples: oss://ros/template/demo and oss://ros/template/demo?RegionId=cn-hangzhou. The template can be up to 16,384 bytes in length, and the URL can be up to 1,350 bytes in length.
	//
	// >  If the region of the OSS bucket is not specified, the RegionId value is used.
	//
	// You can specify one of the StackPolicyBody and StackPolicyURL parameters, but you cannot specify both of them.
	//
	// example:
	//
	// oss://ros/stack-policy/demo
	StackPolicyURL *string `json:"StackPolicyURL,omitempty" xml:"StackPolicyURL,omitempty"`
}

func (s SetStackPolicyRequest) String() string {
	return dara.Prettify(s)
}

func (s SetStackPolicyRequest) GoString() string {
	return s.String()
}

func (s *SetStackPolicyRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *SetStackPolicyRequest) GetStackId() *string {
	return s.StackId
}

func (s *SetStackPolicyRequest) GetStackPolicyBody() *string {
	return s.StackPolicyBody
}

func (s *SetStackPolicyRequest) GetStackPolicyURL() *string {
	return s.StackPolicyURL
}

func (s *SetStackPolicyRequest) SetRegionId(v string) *SetStackPolicyRequest {
	s.RegionId = &v
	return s
}

func (s *SetStackPolicyRequest) SetStackId(v string) *SetStackPolicyRequest {
	s.StackId = &v
	return s
}

func (s *SetStackPolicyRequest) SetStackPolicyBody(v string) *SetStackPolicyRequest {
	s.StackPolicyBody = &v
	return s
}

func (s *SetStackPolicyRequest) SetStackPolicyURL(v string) *SetStackPolicyRequest {
	s.StackPolicyURL = &v
	return s
}

func (s *SetStackPolicyRequest) Validate() error {
	return dara.Validate(s)
}
