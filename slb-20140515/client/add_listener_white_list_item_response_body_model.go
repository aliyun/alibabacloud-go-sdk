// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddListenerWhiteListItemResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *AddListenerWhiteListItemResponseBody
	GetRequestId() *string
}

type AddListenerWhiteListItemResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// CEF72CEB-54B6-4AE8-B225-F876FF7BA984
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AddListenerWhiteListItemResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AddListenerWhiteListItemResponseBody) GoString() string {
	return s.String()
}

func (s *AddListenerWhiteListItemResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AddListenerWhiteListItemResponseBody) SetRequestId(v string) *AddListenerWhiteListItemResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddListenerWhiteListItemResponseBody) Validate() error {
	return dara.Validate(s)
}
