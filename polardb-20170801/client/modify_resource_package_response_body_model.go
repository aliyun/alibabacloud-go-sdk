// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyResourcePackageResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyResourcePackageResponseBody
	GetRequestId() *string
}

type ModifyResourcePackageResponseBody struct {
	// example:
	//
	// E56531A4-E552-40BA-9C58-137B80******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyResourcePackageResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyResourcePackageResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyResourcePackageResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyResourcePackageResponseBody) SetRequestId(v string) *ModifyResourcePackageResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyResourcePackageResponseBody) Validate() error {
	return dara.Validate(s)
}
