// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteNetworkAclEntryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteNetworkAclEntryResponseBody
	GetRequestId() *string
}

type DeleteNetworkAclEntryResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// CEF72CEB-54B6-4AE8-B225-F876FF7BA984
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteNetworkAclEntryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteNetworkAclEntryResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteNetworkAclEntryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteNetworkAclEntryResponseBody) SetRequestId(v string) *DeleteNetworkAclEntryResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteNetworkAclEntryResponseBody) Validate() error {
	return dara.Validate(s)
}
