// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePolarFsObjectResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPath(v string) *CreatePolarFsObjectResponseBody
	GetPath() *string
	SetPolarFsInstanceId(v string) *CreatePolarFsObjectResponseBody
	GetPolarFsInstanceId() *string
	SetRequestId(v string) *CreatePolarFsObjectResponseBody
	GetRequestId() *string
}

type CreatePolarFsObjectResponseBody struct {
	// example:
	//
	// /test
	Path *string `json:"Path,omitempty" xml:"Path,omitempty"`
	// example:
	//
	// pfs-test****
	PolarFsInstanceId *string `json:"PolarFsInstanceId,omitempty" xml:"PolarFsInstanceId,omitempty"`
	// Id of the request
	//
	// example:
	//
	// C890995A-CF06-4F4D-8DB8-DD26C2******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreatePolarFsObjectResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreatePolarFsObjectResponseBody) GoString() string {
	return s.String()
}

func (s *CreatePolarFsObjectResponseBody) GetPath() *string {
	return s.Path
}

func (s *CreatePolarFsObjectResponseBody) GetPolarFsInstanceId() *string {
	return s.PolarFsInstanceId
}

func (s *CreatePolarFsObjectResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreatePolarFsObjectResponseBody) SetPath(v string) *CreatePolarFsObjectResponseBody {
	s.Path = &v
	return s
}

func (s *CreatePolarFsObjectResponseBody) SetPolarFsInstanceId(v string) *CreatePolarFsObjectResponseBody {
	s.PolarFsInstanceId = &v
	return s
}

func (s *CreatePolarFsObjectResponseBody) SetRequestId(v string) *CreatePolarFsObjectResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreatePolarFsObjectResponseBody) Validate() error {
	return dara.Validate(s)
}
