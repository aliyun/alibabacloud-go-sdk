// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteMediaLiveInputSecurityGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteMediaLiveInputSecurityGroupResponseBody
	GetRequestId() *string
}

type DeleteMediaLiveInputSecurityGroupResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// ******11-DB8D-4A9A-875B-275798******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteMediaLiveInputSecurityGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteMediaLiveInputSecurityGroupResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteMediaLiveInputSecurityGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteMediaLiveInputSecurityGroupResponseBody) SetRequestId(v string) *DeleteMediaLiveInputSecurityGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteMediaLiveInputSecurityGroupResponseBody) Validate() error {
	return dara.Validate(s)
}
