// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateCasterResourceGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateCasterResourceGroupResponseBody
	GetRequestId() *string
}

type UpdateCasterResourceGroupResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 0D776BD1-****-59D0-9A1B-272832D999F6
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateCasterResourceGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateCasterResourceGroupResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateCasterResourceGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateCasterResourceGroupResponseBody) SetRequestId(v string) *UpdateCasterResourceGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateCasterResourceGroupResponseBody) Validate() error {
	return dara.Validate(s)
}
