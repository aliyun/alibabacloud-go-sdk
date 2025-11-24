// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddVMIntoServiceMeshResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *AddVMIntoServiceMeshResponseBody
	GetRequestId() *string
}

type AddVMIntoServiceMeshResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 4b2c0fe0-6705-4614-8521-6b9d289163c8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AddVMIntoServiceMeshResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AddVMIntoServiceMeshResponseBody) GoString() string {
	return s.String()
}

func (s *AddVMIntoServiceMeshResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AddVMIntoServiceMeshResponseBody) SetRequestId(v string) *AddVMIntoServiceMeshResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddVMIntoServiceMeshResponseBody) Validate() error {
	return dara.Validate(s)
}
