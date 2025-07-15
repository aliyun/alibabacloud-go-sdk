// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteIpControlResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteIpControlResponseBody
	GetRequestId() *string
}

type DeleteIpControlResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// CEF72CEB-54B6-4AE8-B225-F876FF7BZ004
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteIpControlResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteIpControlResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteIpControlResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteIpControlResponseBody) SetRequestId(v string) *DeleteIpControlResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteIpControlResponseBody) Validate() error {
	return dara.Validate(s)
}
