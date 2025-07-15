// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeletePolicyGroupsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeletePolicyGroupsResponseBody
	GetRequestId() *string
}

type DeletePolicyGroupsResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 1CBAFFAB-B697-4049-A9B1-67E1FC5F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeletePolicyGroupsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeletePolicyGroupsResponseBody) GoString() string {
	return s.String()
}

func (s *DeletePolicyGroupsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeletePolicyGroupsResponseBody) SetRequestId(v string) *DeletePolicyGroupsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeletePolicyGroupsResponseBody) Validate() error {
	return dara.Validate(s)
}
