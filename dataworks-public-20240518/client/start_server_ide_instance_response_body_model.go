// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartServerIdeInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *StartServerIdeInstanceResponseBody
	GetInstanceId() *string
	SetRequestId(v string) *StartServerIdeInstanceResponseBody
	GetRequestId() *string
}

type StartServerIdeInstanceResponseBody struct {
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

func (s StartServerIdeInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s StartServerIdeInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *StartServerIdeInstanceResponseBody) GetInstanceId() *string {
	return s.InstanceId
}

func (s *StartServerIdeInstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *StartServerIdeInstanceResponseBody) SetInstanceId(v string) *StartServerIdeInstanceResponseBody {
	s.InstanceId = &v
	return s
}

func (s *StartServerIdeInstanceResponseBody) SetRequestId(v string) *StartServerIdeInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *StartServerIdeInstanceResponseBody) Validate() error {
	return dara.Validate(s)
}
