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
	// example:
	//
	// D1B18FFE-4A81-59D8-AA02-1817098977CD
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
