// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateHostShareKeyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetHostShareKeyId(v int64) *CreateHostShareKeyResponseBody
	GetHostShareKeyId() *int64
	SetRequestId(v string) *CreateHostShareKeyResponseBody
	GetRequestId() *string
}

type CreateHostShareKeyResponseBody struct {
	// The ID of the shared key.
	//
	// example:
	//
	// 10235
	HostShareKeyId *int64 `json:"HostShareKeyId,omitempty" xml:"HostShareKeyId,omitempty"`
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// EC9BF0F4-8983-491A-BC8C-1B4DD94976DE
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateHostShareKeyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateHostShareKeyResponseBody) GoString() string {
	return s.String()
}

func (s *CreateHostShareKeyResponseBody) GetHostShareKeyId() *int64 {
	return s.HostShareKeyId
}

func (s *CreateHostShareKeyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateHostShareKeyResponseBody) SetHostShareKeyId(v int64) *CreateHostShareKeyResponseBody {
	s.HostShareKeyId = &v
	return s
}

func (s *CreateHostShareKeyResponseBody) SetRequestId(v string) *CreateHostShareKeyResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateHostShareKeyResponseBody) Validate() error {
	return dara.Validate(s)
}
