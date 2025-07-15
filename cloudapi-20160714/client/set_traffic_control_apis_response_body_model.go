// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetTrafficControlApisResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *SetTrafficControlApisResponseBody
	GetRequestId() *string
}

type SetTrafficControlApisResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// CEF72CEB-54B6-4AE8-B225-F876FF7BZ004
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SetTrafficControlApisResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SetTrafficControlApisResponseBody) GoString() string {
	return s.String()
}

func (s *SetTrafficControlApisResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SetTrafficControlApisResponseBody) SetRequestId(v string) *SetTrafficControlApisResponseBody {
	s.RequestId = &v
	return s
}

func (s *SetTrafficControlApisResponseBody) Validate() error {
	return dara.Validate(s)
}
