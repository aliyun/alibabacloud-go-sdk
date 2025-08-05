// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSlsLogDispatchResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateSlsLogDispatchResponseBody
	GetRequestId() *string
}

type CreateSlsLogDispatchResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 850A84D6************00090125EEB1
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateSlsLogDispatchResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateSlsLogDispatchResponseBody) GoString() string {
	return s.String()
}

func (s *CreateSlsLogDispatchResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateSlsLogDispatchResponseBody) SetRequestId(v string) *CreateSlsLogDispatchResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateSlsLogDispatchResponseBody) Validate() error {
	return dara.Validate(s)
}
