// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetEngineDefaultAuthResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *GetEngineDefaultAuthResponseBody
	GetAccessDeniedDetail() *string
	SetAuthInfos(v []*GetEngineDefaultAuthResponseBodyAuthInfos) *GetEngineDefaultAuthResponseBody
	GetAuthInfos() []*GetEngineDefaultAuthResponseBodyAuthInfos
	SetInstanceId(v string) *GetEngineDefaultAuthResponseBody
	GetInstanceId() *string
	SetRequestId(v string) *GetEngineDefaultAuthResponseBody
	GetRequestId() *string
}

type GetEngineDefaultAuthResponseBody struct {
	AccessDeniedDetail *string                                      `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	AuthInfos          []*GetEngineDefaultAuthResponseBodyAuthInfos `json:"AuthInfos,omitempty" xml:"AuthInfos,omitempty" type:"Repeated"`
	InstanceId         *string                                      `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	RequestId          *string                                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetEngineDefaultAuthResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetEngineDefaultAuthResponseBody) GoString() string {
	return s.String()
}

func (s *GetEngineDefaultAuthResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *GetEngineDefaultAuthResponseBody) GetAuthInfos() []*GetEngineDefaultAuthResponseBodyAuthInfos {
	return s.AuthInfos
}

func (s *GetEngineDefaultAuthResponseBody) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetEngineDefaultAuthResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetEngineDefaultAuthResponseBody) SetAccessDeniedDetail(v string) *GetEngineDefaultAuthResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *GetEngineDefaultAuthResponseBody) SetAuthInfos(v []*GetEngineDefaultAuthResponseBodyAuthInfos) *GetEngineDefaultAuthResponseBody {
	s.AuthInfos = v
	return s
}

func (s *GetEngineDefaultAuthResponseBody) SetInstanceId(v string) *GetEngineDefaultAuthResponseBody {
	s.InstanceId = &v
	return s
}

func (s *GetEngineDefaultAuthResponseBody) SetRequestId(v string) *GetEngineDefaultAuthResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetEngineDefaultAuthResponseBody) Validate() error {
	if s.AuthInfos != nil {
		for _, item := range s.AuthInfos {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetEngineDefaultAuthResponseBodyAuthInfos struct {
	Engine   *string `json:"Engine,omitempty" xml:"Engine,omitempty"`
	Password *string `json:"Password,omitempty" xml:"Password,omitempty"`
	Username *string `json:"Username,omitempty" xml:"Username,omitempty"`
}

func (s GetEngineDefaultAuthResponseBodyAuthInfos) String() string {
	return dara.Prettify(s)
}

func (s GetEngineDefaultAuthResponseBodyAuthInfos) GoString() string {
	return s.String()
}

func (s *GetEngineDefaultAuthResponseBodyAuthInfos) GetEngine() *string {
	return s.Engine
}

func (s *GetEngineDefaultAuthResponseBodyAuthInfos) GetPassword() *string {
	return s.Password
}

func (s *GetEngineDefaultAuthResponseBodyAuthInfos) GetUsername() *string {
	return s.Username
}

func (s *GetEngineDefaultAuthResponseBodyAuthInfos) SetEngine(v string) *GetEngineDefaultAuthResponseBodyAuthInfos {
	s.Engine = &v
	return s
}

func (s *GetEngineDefaultAuthResponseBodyAuthInfos) SetPassword(v string) *GetEngineDefaultAuthResponseBodyAuthInfos {
	s.Password = &v
	return s
}

func (s *GetEngineDefaultAuthResponseBodyAuthInfos) SetUsername(v string) *GetEngineDefaultAuthResponseBodyAuthInfos {
	s.Username = &v
	return s
}

func (s *GetEngineDefaultAuthResponseBodyAuthInfos) Validate() error {
	return dara.Validate(s)
}
