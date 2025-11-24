// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetVmAppMeshInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v string) *GetVmAppMeshInfoResponseBody
	GetData() *string
	SetRequestId(v string) *GetVmAppMeshInfoResponseBody
	GetRequestId() *string
}

type GetVmAppMeshInfoResponseBody struct {
	// The response parameters.
	//
	// example:
	//
	// ...
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 9522f7c9-63a1-4603-b850-37d12a****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetVmAppMeshInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetVmAppMeshInfoResponseBody) GoString() string {
	return s.String()
}

func (s *GetVmAppMeshInfoResponseBody) GetData() *string {
	return s.Data
}

func (s *GetVmAppMeshInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetVmAppMeshInfoResponseBody) SetData(v string) *GetVmAppMeshInfoResponseBody {
	s.Data = &v
	return s
}

func (s *GetVmAppMeshInfoResponseBody) SetRequestId(v string) *GetVmAppMeshInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetVmAppMeshInfoResponseBody) Validate() error {
	return dara.Validate(s)
}
