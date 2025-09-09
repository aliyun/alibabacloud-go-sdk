// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateHostResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetHostId(v string) *CreateHostResponseBody
	GetHostId() *string
	SetRequestId(v string) *CreateHostResponseBody
	GetRequestId() *string
}

type CreateHostResponseBody struct {
	// The ID of the host.
	//
	// example:
	//
	// 1
	HostId *string `json:"HostId,omitempty" xml:"HostId,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// EC9BF0F4-8983-491A-BC8C-1B4DD94976DE
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateHostResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateHostResponseBody) GoString() string {
	return s.String()
}

func (s *CreateHostResponseBody) GetHostId() *string {
	return s.HostId
}

func (s *CreateHostResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateHostResponseBody) SetHostId(v string) *CreateHostResponseBody {
	s.HostId = &v
	return s
}

func (s *CreateHostResponseBody) SetRequestId(v string) *CreateHostResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateHostResponseBody) Validate() error {
	return dara.Validate(s)
}
