// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateMediaLiveInputSecurityGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateMediaLiveInputSecurityGroupResponseBody
	GetRequestId() *string
}

type UpdateMediaLiveInputSecurityGroupResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// ******11-DB8D-4A9A-875B-275798******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateMediaLiveInputSecurityGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateMediaLiveInputSecurityGroupResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateMediaLiveInputSecurityGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateMediaLiveInputSecurityGroupResponseBody) SetRequestId(v string) *UpdateMediaLiveInputSecurityGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateMediaLiveInputSecurityGroupResponseBody) Validate() error {
	return dara.Validate(s)
}
