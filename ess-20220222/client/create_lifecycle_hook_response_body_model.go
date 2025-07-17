// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateLifecycleHookResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetLifecycleHookId(v string) *CreateLifecycleHookResponseBody
	GetLifecycleHookId() *string
	SetRequestId(v string) *CreateLifecycleHookResponseBody
	GetRequestId() *string
}

type CreateLifecycleHookResponseBody struct {
	// The ID of the lifecycle hook.
	//
	// example:
	//
	// ash-bp1at9ufhmcf9cmy****
	LifecycleHookId *string `json:"LifecycleHookId,omitempty" xml:"LifecycleHookId,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3C83E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateLifecycleHookResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateLifecycleHookResponseBody) GoString() string {
	return s.String()
}

func (s *CreateLifecycleHookResponseBody) GetLifecycleHookId() *string {
	return s.LifecycleHookId
}

func (s *CreateLifecycleHookResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateLifecycleHookResponseBody) SetLifecycleHookId(v string) *CreateLifecycleHookResponseBody {
	s.LifecycleHookId = &v
	return s
}

func (s *CreateLifecycleHookResponseBody) SetRequestId(v string) *CreateLifecycleHookResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateLifecycleHookResponseBody) Validate() error {
	return dara.Validate(s)
}
