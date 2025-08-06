// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAppInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateAppInfoResponseBody
	GetRequestId() *string
}

type UpdateAppInfoResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 25818875-5F78-4A13-DF45-D7393642****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateAppInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateAppInfoResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateAppInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateAppInfoResponseBody) SetRequestId(v string) *UpdateAppInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateAppInfoResponseBody) Validate() error {
	return dara.Validate(s)
}
