// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeletePolicyGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeletePolicyGroupResponseBody
	GetRequestId() *string
}

type DeletePolicyGroupResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 17C731AB-AAEE-5844-A352-D8D0352D3F0B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeletePolicyGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeletePolicyGroupResponseBody) GoString() string {
	return s.String()
}

func (s *DeletePolicyGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeletePolicyGroupResponseBody) SetRequestId(v string) *DeletePolicyGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeletePolicyGroupResponseBody) Validate() error {
	return dara.Validate(s)
}
