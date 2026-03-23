// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSendifyAutoLoginURLResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAutoLoginURL(v string) *GetSendifyAutoLoginURLResponseBody
	GetAutoLoginURL() *string
	SetRequestId(v string) *GetSendifyAutoLoginURLResponseBody
	GetRequestId() *string
}

type GetSendifyAutoLoginURLResponseBody struct {
	// example:
	//
	// https://dingstore.cn
	AutoLoginURL *string `json:"AutoLoginURL,omitempty" xml:"AutoLoginURL,omitempty"`
	// example:
	//
	// 123423
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetSendifyAutoLoginURLResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetSendifyAutoLoginURLResponseBody) GoString() string {
	return s.String()
}

func (s *GetSendifyAutoLoginURLResponseBody) GetAutoLoginURL() *string {
	return s.AutoLoginURL
}

func (s *GetSendifyAutoLoginURLResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetSendifyAutoLoginURLResponseBody) SetAutoLoginURL(v string) *GetSendifyAutoLoginURLResponseBody {
	s.AutoLoginURL = &v
	return s
}

func (s *GetSendifyAutoLoginURLResponseBody) SetRequestId(v string) *GetSendifyAutoLoginURLResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetSendifyAutoLoginURLResponseBody) Validate() error {
	return dara.Validate(s)
}
