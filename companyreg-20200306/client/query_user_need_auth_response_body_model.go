// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryUserNeedAuthResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetNeedAuth(v bool) *QueryUserNeedAuthResponseBody
	GetNeedAuth() *bool
	SetRequestId(v string) *QueryUserNeedAuthResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *QueryUserNeedAuthResponseBody
	GetSuccess() *bool
}

type QueryUserNeedAuthResponseBody struct {
	// example:
	//
	// True
	NeedAuth *bool `json:"NeedAuth,omitempty" xml:"NeedAuth,omitempty"`
	// example:
	//
	// 2C859C36-886C-5BE7-A606-01F38A5374D6
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryUserNeedAuthResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryUserNeedAuthResponseBody) GoString() string {
	return s.String()
}

func (s *QueryUserNeedAuthResponseBody) GetNeedAuth() *bool {
	return s.NeedAuth
}

func (s *QueryUserNeedAuthResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryUserNeedAuthResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *QueryUserNeedAuthResponseBody) SetNeedAuth(v bool) *QueryUserNeedAuthResponseBody {
	s.NeedAuth = &v
	return s
}

func (s *QueryUserNeedAuthResponseBody) SetRequestId(v string) *QueryUserNeedAuthResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryUserNeedAuthResponseBody) SetSuccess(v bool) *QueryUserNeedAuthResponseBody {
	s.Success = &v
	return s
}

func (s *QueryUserNeedAuthResponseBody) Validate() error {
	return dara.Validate(s)
}
