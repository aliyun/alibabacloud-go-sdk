// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisableInstanceAccessControlResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DisableInstanceAccessControlResponseBody
	GetRequestId() *string
}

type DisableInstanceAccessControlResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// CEF72CEB-54B6-4AE8-B225-F876FF7BZ016
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DisableInstanceAccessControlResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DisableInstanceAccessControlResponseBody) GoString() string {
	return s.String()
}

func (s *DisableInstanceAccessControlResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DisableInstanceAccessControlResponseBody) SetRequestId(v string) *DisableInstanceAccessControlResponseBody {
	s.RequestId = &v
	return s
}

func (s *DisableInstanceAccessControlResponseBody) Validate() error {
	return dara.Validate(s)
}
