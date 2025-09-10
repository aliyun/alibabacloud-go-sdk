// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAddonReleaseResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateAddonReleaseResponseBody
	GetRequestId() *string
}

type UpdateAddonReleaseResponseBody struct {
	// example:
	//
	// 9F00A7AF-2728-5424-B321-79D39C00A1EC
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s UpdateAddonReleaseResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateAddonReleaseResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateAddonReleaseResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateAddonReleaseResponseBody) SetRequestId(v string) *UpdateAddonReleaseResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateAddonReleaseResponseBody) Validate() error {
	return dara.Validate(s)
}
