// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteServerIdeInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *DeleteServerIdeInstanceResponseBody
	GetInstanceId() *string
	SetRequestId(v string) *DeleteServerIdeInstanceResponseBody
	GetRequestId() *string
}

type DeleteServerIdeInstanceResponseBody struct {
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

func (s DeleteServerIdeInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteServerIdeInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteServerIdeInstanceResponseBody) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DeleteServerIdeInstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteServerIdeInstanceResponseBody) SetInstanceId(v string) *DeleteServerIdeInstanceResponseBody {
	s.InstanceId = &v
	return s
}

func (s *DeleteServerIdeInstanceResponseBody) SetRequestId(v string) *DeleteServerIdeInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteServerIdeInstanceResponseBody) Validate() error {
	return dara.Validate(s)
}
