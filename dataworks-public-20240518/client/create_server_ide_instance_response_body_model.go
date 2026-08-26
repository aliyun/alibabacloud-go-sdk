// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateServerIdeInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *CreateServerIdeInstanceResponseBody
	GetInstanceId() *string
	SetRequestId(v string) *CreateServerIdeInstanceResponseBody
	GetRequestId() *string
}

type CreateServerIdeInstanceResponseBody struct {
	// The ID of the personal development environment instance.
	//
	// example:
	//
	// 699573
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// E7D55162-4489-1619-AAF5-3F97D5FCA948
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateServerIdeInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateServerIdeInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *CreateServerIdeInstanceResponseBody) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CreateServerIdeInstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateServerIdeInstanceResponseBody) SetInstanceId(v string) *CreateServerIdeInstanceResponseBody {
	s.InstanceId = &v
	return s
}

func (s *CreateServerIdeInstanceResponseBody) SetRequestId(v string) *CreateServerIdeInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateServerIdeInstanceResponseBody) Validate() error {
	return dara.Validate(s)
}
