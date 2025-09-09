// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisableUserConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DisableUserConfigResponseBody
	GetRequestId() *string
}

type DisableUserConfigResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// AC314611-D907-5EBF-B6D8-70425E5A8643
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DisableUserConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DisableUserConfigResponseBody) GoString() string {
	return s.String()
}

func (s *DisableUserConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DisableUserConfigResponseBody) SetRequestId(v string) *DisableUserConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *DisableUserConfigResponseBody) Validate() error {
	return dara.Validate(s)
}
