// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAlertResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateAlertResponseBody
	GetRequestId() *string
}

type UpdateAlertResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 6276D891-*****-55B2-87B9-74D413F7****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateAlertResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateAlertResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateAlertResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateAlertResponseBody) SetRequestId(v string) *UpdateAlertResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateAlertResponseBody) Validate() error {
	return dara.Validate(s)
}
