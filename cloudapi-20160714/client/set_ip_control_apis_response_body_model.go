// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetIpControlApisResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *SetIpControlApisResponseBody
	GetRequestId() *string
}

type SetIpControlApisResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// CEF72CEB-54B6-4AE8-B225-F876FF7BZ004
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SetIpControlApisResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SetIpControlApisResponseBody) GoString() string {
	return s.String()
}

func (s *SetIpControlApisResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SetIpControlApisResponseBody) SetRequestId(v string) *SetIpControlApisResponseBody {
	s.RequestId = &v
	return s
}

func (s *SetIpControlApisResponseBody) Validate() error {
	return dara.Validate(s)
}
