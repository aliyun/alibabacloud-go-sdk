// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePolarFsObjectsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPath(v string) *DescribePolarFsObjectsRequest
	GetPath() *string
	SetPolarFsInstanceId(v string) *DescribePolarFsObjectsRequest
	GetPolarFsInstanceId() *string
}

type DescribePolarFsObjectsRequest struct {
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
	// pfs-2ze0i74ka607*****
	PolarFsInstanceId *string `json:"PolarFsInstanceId,omitempty" xml:"PolarFsInstanceId,omitempty"`
}

func (s DescribePolarFsObjectsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribePolarFsObjectsRequest) GoString() string {
	return s.String()
}

func (s *DescribePolarFsObjectsRequest) GetPath() *string {
	return s.Path
}

func (s *DescribePolarFsObjectsRequest) GetPolarFsInstanceId() *string {
	return s.PolarFsInstanceId
}

func (s *DescribePolarFsObjectsRequest) SetPath(v string) *DescribePolarFsObjectsRequest {
	s.Path = &v
	return s
}

func (s *DescribePolarFsObjectsRequest) SetPolarFsInstanceId(v string) *DescribePolarFsObjectsRequest {
	s.PolarFsInstanceId = &v
	return s
}

func (s *DescribePolarFsObjectsRequest) Validate() error {
	return dara.Validate(s)
}
