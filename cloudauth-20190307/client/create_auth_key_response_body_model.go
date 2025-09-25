// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAuthKeyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAuthKey(v string) *CreateAuthKeyResponseBody
	GetAuthKey() *string
	SetRequestId(v string) *CreateAuthKeyResponseBody
	GetRequestId() *string
}

type CreateAuthKeyResponseBody struct {
	// The key that can be used for authorization activation. The authorization key is valid for 30 minutes and cannot be reused. It is recommended to re-obtain it before each activation.
	//
	// example:
	//
	// auth.1KQMcnLd4m37LN2D0F0WCD-1qtQI$
	AuthKey *string `json:"AuthKey,omitempty" xml:"AuthKey,omitempty"`
	// The ID of this request.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3C83E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateAuthKeyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateAuthKeyResponseBody) GoString() string {
	return s.String()
}

func (s *CreateAuthKeyResponseBody) GetAuthKey() *string {
	return s.AuthKey
}

func (s *CreateAuthKeyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateAuthKeyResponseBody) SetAuthKey(v string) *CreateAuthKeyResponseBody {
	s.AuthKey = &v
	return s
}

func (s *CreateAuthKeyResponseBody) SetRequestId(v string) *CreateAuthKeyResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateAuthKeyResponseBody) Validate() error {
	return dara.Validate(s)
}
