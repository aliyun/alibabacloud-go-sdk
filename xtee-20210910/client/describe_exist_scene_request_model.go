// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeExistSceneRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *DescribeExistSceneRequest
	GetLang() *string
	SetSceneName(v string) *DescribeExistSceneRequest
	GetSceneName() *string
	SetRegId(v string) *DescribeExistSceneRequest
	GetRegId() *string
}

type DescribeExistSceneRequest struct {
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
	// Scene name.
	//
	// example:
	//
	// 样本调度
	SceneName *string `json:"SceneName,omitempty" xml:"SceneName,omitempty"`
	// Region code
	//
	// example:
	//
	// cn-hangzhou
	RegId *string `json:"regId,omitempty" xml:"regId,omitempty"`
}

func (s DescribeExistSceneRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeExistSceneRequest) GoString() string {
	return s.String()
}

func (s *DescribeExistSceneRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeExistSceneRequest) GetSceneName() *string {
	return s.SceneName
}

func (s *DescribeExistSceneRequest) GetRegId() *string {
	return s.RegId
}

func (s *DescribeExistSceneRequest) SetLang(v string) *DescribeExistSceneRequest {
	s.Lang = &v
	return s
}

func (s *DescribeExistSceneRequest) SetSceneName(v string) *DescribeExistSceneRequest {
	s.SceneName = &v
	return s
}

func (s *DescribeExistSceneRequest) SetRegId(v string) *DescribeExistSceneRequest {
	s.RegId = &v
	return s
}

func (s *DescribeExistSceneRequest) Validate() error {
	return dara.Validate(s)
}
