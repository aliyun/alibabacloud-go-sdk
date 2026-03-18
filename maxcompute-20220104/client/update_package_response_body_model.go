// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdatePackageResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v string) *UpdatePackageResponseBody
	GetData() *string
	SetRequestId(v string) *UpdatePackageResponseBody
	GetRequestId() *string
}

type UpdatePackageResponseBody struct {
	Data      *string `json:"data,omitempty" xml:"data,omitempty"`
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s UpdatePackageResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdatePackageResponseBody) GoString() string {
	return s.String()
}

func (s *UpdatePackageResponseBody) GetData() *string {
	return s.Data
}

func (s *UpdatePackageResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdatePackageResponseBody) SetData(v string) *UpdatePackageResponseBody {
	s.Data = &v
	return s
}

func (s *UpdatePackageResponseBody) SetRequestId(v string) *UpdatePackageResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdatePackageResponseBody) Validate() error {
	return dara.Validate(s)
}
