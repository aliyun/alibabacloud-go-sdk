// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateRenderingDataPackageResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDataPackageId(v string) *CreateRenderingDataPackageResponseBody
	GetDataPackageId() *string
	SetRequestId(v string) *CreateRenderingDataPackageResponseBody
	GetRequestId() *string
}

type CreateRenderingDataPackageResponseBody struct {
	// example:
	//
	// dp-9f8c57355d224ad7beaf95e145f22111
	DataPackageId *string `json:"DataPackageId,omitempty" xml:"DataPackageId,omitempty"`
	// example:
	//
	// BEA5625F-8FCF-48F4-851B-CA63946DA664
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateRenderingDataPackageResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateRenderingDataPackageResponseBody) GoString() string {
	return s.String()
}

func (s *CreateRenderingDataPackageResponseBody) GetDataPackageId() *string {
	return s.DataPackageId
}

func (s *CreateRenderingDataPackageResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateRenderingDataPackageResponseBody) SetDataPackageId(v string) *CreateRenderingDataPackageResponseBody {
	s.DataPackageId = &v
	return s
}

func (s *CreateRenderingDataPackageResponseBody) SetRequestId(v string) *CreateRenderingDataPackageResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateRenderingDataPackageResponseBody) Validate() error {
	return dara.Validate(s)
}
