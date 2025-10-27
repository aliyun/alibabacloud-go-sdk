// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDBResourceGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateDBResourceGroupResponseBody
	GetRequestId() *string
}

type CreateDBResourceGroupResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 3A8F6106-6AFD-5A34-9C80-8DE2C42D06E8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateDBResourceGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateDBResourceGroupResponseBody) GoString() string {
	return s.String()
}

func (s *CreateDBResourceGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateDBResourceGroupResponseBody) SetRequestId(v string) *CreateDBResourceGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateDBResourceGroupResponseBody) Validate() error {
	return dara.Validate(s)
}
