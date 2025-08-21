// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveUserFromGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *RemoveUserFromGroupResponseBody
	GetRequestId() *string
}

type RemoveUserFromGroupResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 85836703-8D4F-485F-9726-4D1C730F957E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RemoveUserFromGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RemoveUserFromGroupResponseBody) GoString() string {
	return s.String()
}

func (s *RemoveUserFromGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RemoveUserFromGroupResponseBody) SetRequestId(v string) *RemoveUserFromGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *RemoveUserFromGroupResponseBody) Validate() error {
	return dara.Validate(s)
}
