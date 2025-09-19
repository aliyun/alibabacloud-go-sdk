// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdatePublishAutoUpgradeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdatePublishAutoUpgradeResponseBody
	GetRequestId() *string
}

type UpdatePublishAutoUpgradeResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 1EE7B150-D67E-53FD-A52D-3E8E669A****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdatePublishAutoUpgradeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdatePublishAutoUpgradeResponseBody) GoString() string {
	return s.String()
}

func (s *UpdatePublishAutoUpgradeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdatePublishAutoUpgradeResponseBody) SetRequestId(v string) *UpdatePublishAutoUpgradeResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdatePublishAutoUpgradeResponseBody) Validate() error {
	return dara.Validate(s)
}
