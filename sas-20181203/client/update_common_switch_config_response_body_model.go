// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateCommonSwitchConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateCommonSwitchConfigResponseBody
	GetRequestId() *string
}

type UpdateCommonSwitchConfigResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 655B538A-A31B-58F2-A3FB-2EF4390D****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateCommonSwitchConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateCommonSwitchConfigResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateCommonSwitchConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateCommonSwitchConfigResponseBody) SetRequestId(v string) *UpdateCommonSwitchConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateCommonSwitchConfigResponseBody) Validate() error {
	return dara.Validate(s)
}
