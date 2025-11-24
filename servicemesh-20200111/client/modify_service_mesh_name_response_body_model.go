// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyServiceMeshNameResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyServiceMeshNameResponseBody
	GetRequestId() *string
}

type ModifyServiceMeshNameResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// EDDC0D86-2FC3-56FB-9213-96EB0A3523F1
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyServiceMeshNameResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyServiceMeshNameResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyServiceMeshNameResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyServiceMeshNameResponseBody) SetRequestId(v string) *ModifyServiceMeshNameResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyServiceMeshNameResponseBody) Validate() error {
	return dara.Validate(s)
}
