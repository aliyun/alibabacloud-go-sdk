// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeOperatorListBySceneRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *DescribeOperatorListBySceneRequest
	GetLang() *string
	SetRegId(v string) *DescribeOperatorListBySceneRequest
	GetRegId() *string
	SetScene(v string) *DescribeOperatorListBySceneRequest
	GetScene() *string
}

type DescribeOperatorListBySceneRequest struct {
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
	// example:
	//
	// cn-hangzhou
	RegId *string `json:"regId,omitempty" xml:"regId,omitempty"`
	// Scene type.
	//
	// example:
	//
	// VELOCITY
	Scene *string `json:"scene,omitempty" xml:"scene,omitempty"`
}

func (s DescribeOperatorListBySceneRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeOperatorListBySceneRequest) GoString() string {
	return s.String()
}

func (s *DescribeOperatorListBySceneRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeOperatorListBySceneRequest) GetRegId() *string {
	return s.RegId
}

func (s *DescribeOperatorListBySceneRequest) GetScene() *string {
	return s.Scene
}

func (s *DescribeOperatorListBySceneRequest) SetLang(v string) *DescribeOperatorListBySceneRequest {
	s.Lang = &v
	return s
}

func (s *DescribeOperatorListBySceneRequest) SetRegId(v string) *DescribeOperatorListBySceneRequest {
	s.RegId = &v
	return s
}

func (s *DescribeOperatorListBySceneRequest) SetScene(v string) *DescribeOperatorListBySceneRequest {
	s.Scene = &v
	return s
}

func (s *DescribeOperatorListBySceneRequest) Validate() error {
	return dara.Validate(s)
}
