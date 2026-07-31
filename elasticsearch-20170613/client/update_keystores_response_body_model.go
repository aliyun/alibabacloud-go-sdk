// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateKeystoresResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateKeystoresResponseBody
	GetRequestId() *string
}

type UpdateKeystoresResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// 5FFD9ED4-C2EC-4E89-B22B-1ACB6FE1D****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateKeystoresResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateKeystoresResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateKeystoresResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateKeystoresResponseBody) SetRequestId(v string) *UpdateKeystoresResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateKeystoresResponseBody) Validate() error {
	return dara.Validate(s)
}
