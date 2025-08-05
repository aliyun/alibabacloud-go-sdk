// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePackageResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v string) *CreatePackageResponseBody
	GetData() *string
	SetRequestId(v string) *CreatePackageResponseBody
	GetRequestId() *string
}

type CreatePackageResponseBody struct {
	// The returned data.
	//
	// example:
	//
	// success
	Data *string `json:"data,omitempty" xml:"data,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 0bc3b4ab16684833172127321e2c25
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s CreatePackageResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreatePackageResponseBody) GoString() string {
	return s.String()
}

func (s *CreatePackageResponseBody) GetData() *string {
	return s.Data
}

func (s *CreatePackageResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreatePackageResponseBody) SetData(v string) *CreatePackageResponseBody {
	s.Data = &v
	return s
}

func (s *CreatePackageResponseBody) SetRequestId(v string) *CreatePackageResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreatePackageResponseBody) Validate() error {
	return dara.Validate(s)
}
