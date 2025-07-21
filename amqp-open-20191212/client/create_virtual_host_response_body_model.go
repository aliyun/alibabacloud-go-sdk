// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateVirtualHostResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateVirtualHostResponseBody
	GetRequestId() *string
}

type CreateVirtualHostResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 628705FD-03EE-4ABE-BB21-E1672960***
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateVirtualHostResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateVirtualHostResponseBody) GoString() string {
	return s.String()
}

func (s *CreateVirtualHostResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateVirtualHostResponseBody) SetRequestId(v string) *CreateVirtualHostResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateVirtualHostResponseBody) Validate() error {
	return dara.Validate(s)
}
