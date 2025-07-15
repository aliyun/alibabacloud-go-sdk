// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyIpControlResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyIpControlResponseBody
	GetRequestId() *string
}

type ModifyIpControlResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// CEF72CEB-54B6-4AE8-B225-F876FF7BZ004
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyIpControlResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyIpControlResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyIpControlResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyIpControlResponseBody) SetRequestId(v string) *ModifyIpControlResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyIpControlResponseBody) Validate() error {
	return dara.Validate(s)
}
