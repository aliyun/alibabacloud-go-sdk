// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateHostAccountResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetHostAccountId(v string) *CreateHostAccountResponseBody
	GetHostAccountId() *string
	SetRequestId(v string) *CreateHostAccountResponseBody
	GetRequestId() *string
}

type CreateHostAccountResponseBody struct {
	// The operation that you want to perform. Set the value to **CreateHostAccount**.
	//
	// example:
	//
	// 1
	HostAccountId *string `json:"HostAccountId,omitempty" xml:"HostAccountId,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// EC9BF0F4-8983-491A-BC8C-1B4DD94976DE
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateHostAccountResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateHostAccountResponseBody) GoString() string {
	return s.String()
}

func (s *CreateHostAccountResponseBody) GetHostAccountId() *string {
	return s.HostAccountId
}

func (s *CreateHostAccountResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateHostAccountResponseBody) SetHostAccountId(v string) *CreateHostAccountResponseBody {
	s.HostAccountId = &v
	return s
}

func (s *CreateHostAccountResponseBody) SetRequestId(v string) *CreateHostAccountResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateHostAccountResponseBody) Validate() error {
	return dara.Validate(s)
}
