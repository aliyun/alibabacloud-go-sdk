// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveVMFromServiceMeshResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *RemoveVMFromServiceMeshResponseBody
	GetRequestId() *string
}

type RemoveVMFromServiceMeshResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 4b2c0fe0-6705-4614-8521-6b9d289163c8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RemoveVMFromServiceMeshResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RemoveVMFromServiceMeshResponseBody) GoString() string {
	return s.String()
}

func (s *RemoveVMFromServiceMeshResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RemoveVMFromServiceMeshResponseBody) SetRequestId(v string) *RemoveVMFromServiceMeshResponseBody {
	s.RequestId = &v
	return s
}

func (s *RemoveVMFromServiceMeshResponseBody) Validate() error {
	return dara.Validate(s)
}
