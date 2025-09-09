// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateHostGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetHostGroupId(v string) *CreateHostGroupResponseBody
	GetHostGroupId() *string
	SetRequestId(v string) *CreateHostGroupResponseBody
	GetRequestId() *string
}

type CreateHostGroupResponseBody struct {
	// The asset group ID.
	//
	// example:
	//
	// 1
	HostGroupId *string `json:"HostGroupId,omitempty" xml:"HostGroupId,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// EC9BF0F4-8983-491A-BC8C-1B4DD94976DE
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateHostGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateHostGroupResponseBody) GoString() string {
	return s.String()
}

func (s *CreateHostGroupResponseBody) GetHostGroupId() *string {
	return s.HostGroupId
}

func (s *CreateHostGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateHostGroupResponseBody) SetHostGroupId(v string) *CreateHostGroupResponseBody {
	s.HostGroupId = &v
	return s
}

func (s *CreateHostGroupResponseBody) SetRequestId(v string) *CreateHostGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateHostGroupResponseBody) Validate() error {
	return dara.Validate(s)
}
