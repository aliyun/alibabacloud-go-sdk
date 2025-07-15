// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveIpControlApisResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *RemoveIpControlApisResponseBody
	GetRequestId() *string
}

type RemoveIpControlApisResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// CEF72CEB-54B6-4AE8-B225-F876FF7BZ004
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RemoveIpControlApisResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RemoveIpControlApisResponseBody) GoString() string {
	return s.String()
}

func (s *RemoveIpControlApisResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RemoveIpControlApisResponseBody) SetRequestId(v string) *RemoveIpControlApisResponseBody {
	s.RequestId = &v
	return s
}

func (s *RemoveIpControlApisResponseBody) Validate() error {
	return dara.Validate(s)
}
