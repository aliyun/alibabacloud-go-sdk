// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeOssPolicyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *DescribeOssPolicyRequest
	GetLang() *string
	SetRegId(v string) *DescribeOssPolicyRequest
	GetRegId() *string
	SetScene(v string) *DescribeOssPolicyRequest
	GetScene() *string
}

type DescribeOssPolicyRequest struct {
	// Sets the language type for requests and received messages, default value is **zh**. Values:
	//
	// - **zh**: Chinese
	//
	// - **en**: English
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// Region code
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegId *string `json:"regId,omitempty" xml:"regId,omitempty"`
	// Scene
	//
	// This parameter is required.
	//
	// example:
	//
	// NAME_LIST
	Scene *string `json:"scene,omitempty" xml:"scene,omitempty"`
}

func (s DescribeOssPolicyRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeOssPolicyRequest) GoString() string {
	return s.String()
}

func (s *DescribeOssPolicyRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeOssPolicyRequest) GetRegId() *string {
	return s.RegId
}

func (s *DescribeOssPolicyRequest) GetScene() *string {
	return s.Scene
}

func (s *DescribeOssPolicyRequest) SetLang(v string) *DescribeOssPolicyRequest {
	s.Lang = &v
	return s
}

func (s *DescribeOssPolicyRequest) SetRegId(v string) *DescribeOssPolicyRequest {
	s.RegId = &v
	return s
}

func (s *DescribeOssPolicyRequest) SetScene(v string) *DescribeOssPolicyRequest {
	s.Scene = &v
	return s
}

func (s *DescribeOssPolicyRequest) Validate() error {
	return dara.Validate(s)
}
