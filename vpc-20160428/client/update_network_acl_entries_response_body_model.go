// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateNetworkAclEntriesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateNetworkAclEntriesResponseBody
	GetRequestId() *string
}

type UpdateNetworkAclEntriesResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 1170A5A0-E760-4331-9133-A7D38D973215
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateNetworkAclEntriesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateNetworkAclEntriesResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateNetworkAclEntriesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateNetworkAclEntriesResponseBody) SetRequestId(v string) *UpdateNetworkAclEntriesResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateNetworkAclEntriesResponseBody) Validate() error {
	return dara.Validate(s)
}
