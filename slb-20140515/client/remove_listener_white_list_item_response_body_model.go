// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveListenerWhiteListItemResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *RemoveListenerWhiteListItemResponseBody
	GetRequestId() *string
}

type RemoveListenerWhiteListItemResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// CEF72CEB-54B6-4AE8-B225-F876FF7BA984
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RemoveListenerWhiteListItemResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RemoveListenerWhiteListItemResponseBody) GoString() string {
	return s.String()
}

func (s *RemoveListenerWhiteListItemResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RemoveListenerWhiteListItemResponseBody) SetRequestId(v string) *RemoveListenerWhiteListItemResponseBody {
	s.RequestId = &v
	return s
}

func (s *RemoveListenerWhiteListItemResponseBody) Validate() error {
	return dara.Validate(s)
}
