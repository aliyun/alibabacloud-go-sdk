// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCompleteLifecycleActionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CompleteLifecycleActionResponseBody
	GetRequestId() *string
}

type CompleteLifecycleActionResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CompleteLifecycleActionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CompleteLifecycleActionResponseBody) GoString() string {
	return s.String()
}

func (s *CompleteLifecycleActionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CompleteLifecycleActionResponseBody) SetRequestId(v string) *CompleteLifecycleActionResponseBody {
	s.RequestId = &v
	return s
}

func (s *CompleteLifecycleActionResponseBody) Validate() error {
	return dara.Validate(s)
}
