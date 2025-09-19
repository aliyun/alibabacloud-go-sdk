// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdatePublishGraySwitchResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdatePublishGraySwitchResponseBody
	GetRequestId() *string
}

type UpdatePublishGraySwitchResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 231A8A81-CBB4-5AB3-A624-98A501******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdatePublishGraySwitchResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdatePublishGraySwitchResponseBody) GoString() string {
	return s.String()
}

func (s *UpdatePublishGraySwitchResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdatePublishGraySwitchResponseBody) SetRequestId(v string) *UpdatePublishGraySwitchResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdatePublishGraySwitchResponseBody) Validate() error {
	return dara.Validate(s)
}
