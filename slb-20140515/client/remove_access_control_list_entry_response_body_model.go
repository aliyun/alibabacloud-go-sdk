// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveAccessControlListEntryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *RemoveAccessControlListEntryResponseBody
	GetRequestId() *string
}

type RemoveAccessControlListEntryResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 988CB45E-1643-48C0-87B4-928DDF77EA49
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RemoveAccessControlListEntryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RemoveAccessControlListEntryResponseBody) GoString() string {
	return s.String()
}

func (s *RemoveAccessControlListEntryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RemoveAccessControlListEntryResponseBody) SetRequestId(v string) *RemoveAccessControlListEntryResponseBody {
	s.RequestId = &v
	return s
}

func (s *RemoveAccessControlListEntryResponseBody) Validate() error {
	return dara.Validate(s)
}
