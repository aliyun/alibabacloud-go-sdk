// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGenerateSceneKeyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetBaseUrl(v string) *GenerateSceneKeyResponseBody
	GetBaseUrl() *string
	SetExpireTime(v string) *GenerateSceneKeyResponseBody
	GetExpireTime() *string
	SetKey(v string) *GenerateSceneKeyResponseBody
	GetKey() *string
	SetRequestId(v string) *GenerateSceneKeyResponseBody
	GetRequestId() *string
	SetSceneUrl(v string) *GenerateSceneKeyResponseBody
	GetSceneUrl() *string
}

type GenerateSceneKeyResponseBody struct {
	// The Agent Gateway endpoint, which includes the cluster offset mapping port. This parameter is returned only for desktop scenarios.
	//
	// example:
	//
	// gw-agent-cn-hangzhou.wuyinggw.com:18008
	BaseUrl *string `json:"BaseUrl,omitempty" xml:"BaseUrl,omitempty"`
	// The expiration time of the key.
	//
	// example:
	//
	// 2026-07-13T12:00:00Z
	ExpireTime *string `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	// The scenario access key.
	//
	// example:
	//
	// sk-0123456789abcdef0123456789abcdef
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 1CBAFFAB-B697-4049-A9B1-67E1FC5F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The scenario access URL.
	//
	// example:
	//
	// https://gw-example.cn-hangzhou.aigclinks.com:18008
	SceneUrl *string `json:"SceneUrl,omitempty" xml:"SceneUrl,omitempty"`
}

func (s GenerateSceneKeyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GenerateSceneKeyResponseBody) GoString() string {
	return s.String()
}

func (s *GenerateSceneKeyResponseBody) GetBaseUrl() *string {
	return s.BaseUrl
}

func (s *GenerateSceneKeyResponseBody) GetExpireTime() *string {
	return s.ExpireTime
}

func (s *GenerateSceneKeyResponseBody) GetKey() *string {
	return s.Key
}

func (s *GenerateSceneKeyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GenerateSceneKeyResponseBody) GetSceneUrl() *string {
	return s.SceneUrl
}

func (s *GenerateSceneKeyResponseBody) SetBaseUrl(v string) *GenerateSceneKeyResponseBody {
	s.BaseUrl = &v
	return s
}

func (s *GenerateSceneKeyResponseBody) SetExpireTime(v string) *GenerateSceneKeyResponseBody {
	s.ExpireTime = &v
	return s
}

func (s *GenerateSceneKeyResponseBody) SetKey(v string) *GenerateSceneKeyResponseBody {
	s.Key = &v
	return s
}

func (s *GenerateSceneKeyResponseBody) SetRequestId(v string) *GenerateSceneKeyResponseBody {
	s.RequestId = &v
	return s
}

func (s *GenerateSceneKeyResponseBody) SetSceneUrl(v string) *GenerateSceneKeyResponseBody {
	s.SceneUrl = &v
	return s
}

func (s *GenerateSceneKeyResponseBody) Validate() error {
	return dara.Validate(s)
}
