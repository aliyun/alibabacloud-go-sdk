// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePolarFsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetOrderId(v string) *CreatePolarFsResponseBody
	GetOrderId() *string
	SetPolarFsInstanceId(v string) *CreatePolarFsResponseBody
	GetPolarFsInstanceId() *string
	SetPolarFsPath(v string) *CreatePolarFsResponseBody
	GetPolarFsPath() *string
	SetPolarFsStatus(v string) *CreatePolarFsResponseBody
	GetPolarFsStatus() *string
	SetRequestId(v string) *CreatePolarFsResponseBody
	GetRequestId() *string
}

type CreatePolarFsResponseBody struct {
	// The order ID.
	//
	// example:
	//
	// 2035638*******
	OrderId *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// The PolarFS instance ID.
	//
	// example:
	//
	// pfs-2ze0i74ka607wck3
	PolarFsInstanceId *string `json:"PolarFsInstanceId,omitempty" xml:"PolarFsInstanceId,omitempty"`
	// The PolarFS file system path.
	//
	// example:
	//
	// pfs-xxx*******
	PolarFsPath *string `json:"PolarFsPath,omitempty" xml:"PolarFsPath,omitempty"`
	// The PolarFS instance status.
	//
	// example:
	//
	// Creating
	PolarFsStatus *string `json:"PolarFsStatus,omitempty" xml:"PolarFsStatus,omitempty"`
	// The request ID.
	//
	// example:
	//
	// CED079B7-A408-41A1-BFF1-EC608E******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreatePolarFsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreatePolarFsResponseBody) GoString() string {
	return s.String()
}

func (s *CreatePolarFsResponseBody) GetOrderId() *string {
	return s.OrderId
}

func (s *CreatePolarFsResponseBody) GetPolarFsInstanceId() *string {
	return s.PolarFsInstanceId
}

func (s *CreatePolarFsResponseBody) GetPolarFsPath() *string {
	return s.PolarFsPath
}

func (s *CreatePolarFsResponseBody) GetPolarFsStatus() *string {
	return s.PolarFsStatus
}

func (s *CreatePolarFsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreatePolarFsResponseBody) SetOrderId(v string) *CreatePolarFsResponseBody {
	s.OrderId = &v
	return s
}

func (s *CreatePolarFsResponseBody) SetPolarFsInstanceId(v string) *CreatePolarFsResponseBody {
	s.PolarFsInstanceId = &v
	return s
}

func (s *CreatePolarFsResponseBody) SetPolarFsPath(v string) *CreatePolarFsResponseBody {
	s.PolarFsPath = &v
	return s
}

func (s *CreatePolarFsResponseBody) SetPolarFsStatus(v string) *CreatePolarFsResponseBody {
	s.PolarFsStatus = &v
	return s
}

func (s *CreatePolarFsResponseBody) SetRequestId(v string) *CreatePolarFsResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreatePolarFsResponseBody) Validate() error {
	return dara.Validate(s)
}
