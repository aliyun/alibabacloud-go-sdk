// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartInstanceRefreshResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceRefreshTaskId(v string) *StartInstanceRefreshResponseBody
	GetInstanceRefreshTaskId() *string
	SetRequestId(v string) *StartInstanceRefreshResponseBody
	GetRequestId() *string
}

type StartInstanceRefreshResponseBody struct {
	// The ID of the instance refresh task.
	//
	// example:
	//
	// ir-a12ds234fasd*****
	InstanceRefreshTaskId *string `json:"InstanceRefreshTaskId,omitempty" xml:"InstanceRefreshTaskId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s StartInstanceRefreshResponseBody) String() string {
	return dara.Prettify(s)
}

func (s StartInstanceRefreshResponseBody) GoString() string {
	return s.String()
}

func (s *StartInstanceRefreshResponseBody) GetInstanceRefreshTaskId() *string {
	return s.InstanceRefreshTaskId
}

func (s *StartInstanceRefreshResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *StartInstanceRefreshResponseBody) SetInstanceRefreshTaskId(v string) *StartInstanceRefreshResponseBody {
	s.InstanceRefreshTaskId = &v
	return s
}

func (s *StartInstanceRefreshResponseBody) SetRequestId(v string) *StartInstanceRefreshResponseBody {
	s.RequestId = &v
	return s
}

func (s *StartInstanceRefreshResponseBody) Validate() error {
	return dara.Validate(s)
}
