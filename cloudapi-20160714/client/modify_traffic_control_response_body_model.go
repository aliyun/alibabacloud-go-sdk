// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyTrafficControlResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyTrafficControlResponseBody
	GetRequestId() *string
}

type ModifyTrafficControlResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// CEF72CEB-54B6-4AE8-B225-F876FF7BZ004
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyTrafficControlResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyTrafficControlResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyTrafficControlResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyTrafficControlResponseBody) SetRequestId(v string) *ModifyTrafficControlResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyTrafficControlResponseBody) Validate() error {
	return dara.Validate(s)
}
