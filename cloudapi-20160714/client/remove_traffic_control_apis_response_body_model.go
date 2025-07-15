// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveTrafficControlApisResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *RemoveTrafficControlApisResponseBody
	GetRequestId() *string
}

type RemoveTrafficControlApisResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// CEF72CEB-54B6-4AE8-B225-F876FF7BZ004
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RemoveTrafficControlApisResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RemoveTrafficControlApisResponseBody) GoString() string {
	return s.String()
}

func (s *RemoveTrafficControlApisResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RemoveTrafficControlApisResponseBody) SetRequestId(v string) *RemoveTrafficControlApisResponseBody {
	s.RequestId = &v
	return s
}

func (s *RemoveTrafficControlApisResponseBody) Validate() error {
	return dara.Validate(s)
}
