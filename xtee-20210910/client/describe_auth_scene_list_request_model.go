// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAuthSceneListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *DescribeAuthSceneListRequest
	GetLang() *string
	SetRegId(v string) *DescribeAuthSceneListRequest
	GetRegId() *string
}

type DescribeAuthSceneListRequest struct {
	// Sets the language type for requests and received messages, with a default value of **zh**. Values:
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
}

func (s DescribeAuthSceneListRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeAuthSceneListRequest) GoString() string {
	return s.String()
}

func (s *DescribeAuthSceneListRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeAuthSceneListRequest) GetRegId() *string {
	return s.RegId
}

func (s *DescribeAuthSceneListRequest) SetLang(v string) *DescribeAuthSceneListRequest {
	s.Lang = &v
	return s
}

func (s *DescribeAuthSceneListRequest) SetRegId(v string) *DescribeAuthSceneListRequest {
	s.RegId = &v
	return s
}

func (s *DescribeAuthSceneListRequest) Validate() error {
	return dara.Validate(s)
}
