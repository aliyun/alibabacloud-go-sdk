// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAuthScenePageListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *DescribeAuthScenePageListRequest
	GetLang() *string
	SetRegId(v string) *DescribeAuthScenePageListRequest
	GetRegId() *string
	SetSceneName(v string) *DescribeAuthScenePageListRequest
	GetSceneName() *string
}

type DescribeAuthScenePageListRequest struct {
	// The language type for the request and response messages. Default value: **zh**. Valid values:
	//
	// - **zh**: Chinese.
	//
	// - **en**: English.
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The region code.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegId *string `json:"regId,omitempty" xml:"regId,omitempty"`
	// The scenario name.
	//
	// example:
	//
	// account_abuse
	SceneName *string `json:"sceneName,omitempty" xml:"sceneName,omitempty"`
}

func (s DescribeAuthScenePageListRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeAuthScenePageListRequest) GoString() string {
	return s.String()
}

func (s *DescribeAuthScenePageListRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeAuthScenePageListRequest) GetRegId() *string {
	return s.RegId
}

func (s *DescribeAuthScenePageListRequest) GetSceneName() *string {
	return s.SceneName
}

func (s *DescribeAuthScenePageListRequest) SetLang(v string) *DescribeAuthScenePageListRequest {
	s.Lang = &v
	return s
}

func (s *DescribeAuthScenePageListRequest) SetRegId(v string) *DescribeAuthScenePageListRequest {
	s.RegId = &v
	return s
}

func (s *DescribeAuthScenePageListRequest) SetSceneName(v string) *DescribeAuthScenePageListRequest {
	s.SceneName = &v
	return s
}

func (s *DescribeAuthScenePageListRequest) Validate() error {
	return dara.Validate(s)
}
