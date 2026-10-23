// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRebootCloudPhoneNodesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *RebootCloudPhoneNodesResponseBody
	GetRequestId() *string
}

type RebootCloudPhoneNodesResponseBody struct {
	// example:
	//
	// 4610632D-D661-5982-B3D7-5D3FD183F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RebootCloudPhoneNodesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RebootCloudPhoneNodesResponseBody) GoString() string {
	return s.String()
}

func (s *RebootCloudPhoneNodesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RebootCloudPhoneNodesResponseBody) SetRequestId(v string) *RebootCloudPhoneNodesResponseBody {
	s.RequestId = &v
	return s
}

func (s *RebootCloudPhoneNodesResponseBody) Validate() error {
	return dara.Validate(s)
}
