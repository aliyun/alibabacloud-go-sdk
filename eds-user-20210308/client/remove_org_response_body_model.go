// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveOrgResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *RemoveOrgResponseBody
	GetRequestId() *string
}

type RemoveOrgResponseBody struct {
	// example:
	//
	// 6C352609-EE7F-5603-B5E6-57C3EDDD****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RemoveOrgResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RemoveOrgResponseBody) GoString() string {
	return s.String()
}

func (s *RemoveOrgResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RemoveOrgResponseBody) SetRequestId(v string) *RemoveOrgResponseBody {
	s.RequestId = &v
	return s
}

func (s *RemoveOrgResponseBody) Validate() error {
	return dara.Validate(s)
}
