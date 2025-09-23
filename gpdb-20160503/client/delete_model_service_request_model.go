// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteModelServiceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceId(v string) *DeleteModelServiceRequest
	GetDBInstanceId() *string
	SetModelServiceId(v string) *DeleteModelServiceRequest
	GetModelServiceId() *string
}

type DeleteModelServiceRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// gp-xxxxxxxxx
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// mx-xxxxxxxxx
	ModelServiceId *string `json:"ModelServiceId,omitempty" xml:"ModelServiceId,omitempty"`
}

func (s DeleteModelServiceRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteModelServiceRequest) GoString() string {
	return s.String()
}

func (s *DeleteModelServiceRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *DeleteModelServiceRequest) GetModelServiceId() *string {
	return s.ModelServiceId
}

func (s *DeleteModelServiceRequest) SetDBInstanceId(v string) *DeleteModelServiceRequest {
	s.DBInstanceId = &v
	return s
}

func (s *DeleteModelServiceRequest) SetModelServiceId(v string) *DeleteModelServiceRequest {
	s.ModelServiceId = &v
	return s
}

func (s *DeleteModelServiceRequest) Validate() error {
	return dara.Validate(s)
}
