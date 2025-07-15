// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveGlobalAccelerationInstanceIpResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *RemoveGlobalAccelerationInstanceIpResponseBody
	GetRequestId() *string
}

type RemoveGlobalAccelerationInstanceIpResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 5BE01CD7-5A50-472D-AC14-CA181C5C03BE
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RemoveGlobalAccelerationInstanceIpResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RemoveGlobalAccelerationInstanceIpResponseBody) GoString() string {
	return s.String()
}

func (s *RemoveGlobalAccelerationInstanceIpResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RemoveGlobalAccelerationInstanceIpResponseBody) SetRequestId(v string) *RemoveGlobalAccelerationInstanceIpResponseBody {
	s.RequestId = &v
	return s
}

func (s *RemoveGlobalAccelerationInstanceIpResponseBody) Validate() error {
	return dara.Validate(s)
}
