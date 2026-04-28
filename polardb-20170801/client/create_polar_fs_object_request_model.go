// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePolarFsObjectRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPath(v string) *CreatePolarFsObjectRequest
	GetPath() *string
	SetPolarFsInstanceId(v string) *CreatePolarFsObjectRequest
	GetPolarFsInstanceId() *string
}

type CreatePolarFsObjectRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// /test
	Path *string `json:"Path,omitempty" xml:"Path,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// pfs-2ze0i7*****
	PolarFsInstanceId *string `json:"PolarFsInstanceId,omitempty" xml:"PolarFsInstanceId,omitempty"`
}

func (s CreatePolarFsObjectRequest) String() string {
	return dara.Prettify(s)
}

func (s CreatePolarFsObjectRequest) GoString() string {
	return s.String()
}

func (s *CreatePolarFsObjectRequest) GetPath() *string {
	return s.Path
}

func (s *CreatePolarFsObjectRequest) GetPolarFsInstanceId() *string {
	return s.PolarFsInstanceId
}

func (s *CreatePolarFsObjectRequest) SetPath(v string) *CreatePolarFsObjectRequest {
	s.Path = &v
	return s
}

func (s *CreatePolarFsObjectRequest) SetPolarFsInstanceId(v string) *CreatePolarFsObjectRequest {
	s.PolarFsInstanceId = &v
	return s
}

func (s *CreatePolarFsObjectRequest) Validate() error {
	return dara.Validate(s)
}
