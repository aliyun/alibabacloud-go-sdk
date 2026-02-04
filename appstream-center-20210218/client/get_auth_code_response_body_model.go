// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAuthCodeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAuthModel(v *GetAuthCodeResponseBodyAuthModel) *GetAuthCodeResponseBody
	GetAuthModel() *GetAuthCodeResponseBodyAuthModel
	SetRequestId(v string) *GetAuthCodeResponseBody
	GetRequestId() *string
}

type GetAuthCodeResponseBody struct {
	AuthModel *GetAuthCodeResponseBodyAuthModel `json:"AuthModel,omitempty" xml:"AuthModel,omitempty" type:"Struct"`
	// example:
	//
	// 1CBAFFAB-B697-4049-A9B1-67E1FC5F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetAuthCodeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetAuthCodeResponseBody) GoString() string {
	return s.String()
}

func (s *GetAuthCodeResponseBody) GetAuthModel() *GetAuthCodeResponseBodyAuthModel {
	return s.AuthModel
}

func (s *GetAuthCodeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetAuthCodeResponseBody) SetAuthModel(v *GetAuthCodeResponseBodyAuthModel) *GetAuthCodeResponseBody {
	s.AuthModel = v
	return s
}

func (s *GetAuthCodeResponseBody) SetRequestId(v string) *GetAuthCodeResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAuthCodeResponseBody) Validate() error {
	if s.AuthModel != nil {
		if err := s.AuthModel.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetAuthCodeResponseBodyAuthModel struct {
	// example:
	//
	// acv2ZEq2TNSqOlX+DvyetHGRT08iPhbWg/os1W4HojpBxkMQZkAnbKSfz/wNvS0E149IQZ5TogvBUE8ghCSVV+QBnv48Y+sn4z9fY5ywZA1peI5s4TplQI0TADBhPZXEIzMOdmbNsDGGlGcKOAq8ixd1Epzdu9XfvRzdLOO6o3zQwTZxXsjOgGBW8JTcgCP9W9iYiIjHIdZ2VCAnmMzuyVkyED8gymEevH10QoclyH/O3s8Yq6VF34gMXihDQeiI8PyiAZrW0JWXSWpdy3n4yGtAqMAgv46FzTA+XJ6C********
	AuthCode *string `json:"AuthCode,omitempty" xml:"AuthCode,omitempty"`
	// example:
	//
	// alice
	EndUserId *string `json:"EndUserId,omitempty" xml:"EndUserId,omitempty"`
	// example:
	//
	// 2015-01-23T12:33:18Z
	ExpireTime *string `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
}

func (s GetAuthCodeResponseBodyAuthModel) String() string {
	return dara.Prettify(s)
}

func (s GetAuthCodeResponseBodyAuthModel) GoString() string {
	return s.String()
}

func (s *GetAuthCodeResponseBodyAuthModel) GetAuthCode() *string {
	return s.AuthCode
}

func (s *GetAuthCodeResponseBodyAuthModel) GetEndUserId() *string {
	return s.EndUserId
}

func (s *GetAuthCodeResponseBodyAuthModel) GetExpireTime() *string {
	return s.ExpireTime
}

func (s *GetAuthCodeResponseBodyAuthModel) SetAuthCode(v string) *GetAuthCodeResponseBodyAuthModel {
	s.AuthCode = &v
	return s
}

func (s *GetAuthCodeResponseBodyAuthModel) SetEndUserId(v string) *GetAuthCodeResponseBodyAuthModel {
	s.EndUserId = &v
	return s
}

func (s *GetAuthCodeResponseBodyAuthModel) SetExpireTime(v string) *GetAuthCodeResponseBodyAuthModel {
	s.ExpireTime = &v
	return s
}

func (s *GetAuthCodeResponseBodyAuthModel) Validate() error {
	return dara.Validate(s)
}
