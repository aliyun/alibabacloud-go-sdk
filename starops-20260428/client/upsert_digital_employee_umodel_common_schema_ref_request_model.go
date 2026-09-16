// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpsertDigitalEmployeeUmodelCommonSchemaRefRequest interface {
	dara.Model
	String() string
	GoString() string
	SetVersion(v string) *UpsertDigitalEmployeeUmodelCommonSchemaRefRequest
	GetVersion() *string
}

type UpsertDigitalEmployeeUmodelCommonSchemaRefRequest struct {
	// The version of the public schema.
	//
	// This parameter is required.
	//
	// example:
	//
	// v1
	Version *string `json:"version,omitempty" xml:"version,omitempty"`
}

func (s UpsertDigitalEmployeeUmodelCommonSchemaRefRequest) String() string {
	return dara.Prettify(s)
}

func (s UpsertDigitalEmployeeUmodelCommonSchemaRefRequest) GoString() string {
	return s.String()
}

func (s *UpsertDigitalEmployeeUmodelCommonSchemaRefRequest) GetVersion() *string {
	return s.Version
}

func (s *UpsertDigitalEmployeeUmodelCommonSchemaRefRequest) SetVersion(v string) *UpsertDigitalEmployeeUmodelCommonSchemaRefRequest {
	s.Version = &v
	return s
}

func (s *UpsertDigitalEmployeeUmodelCommonSchemaRefRequest) Validate() error {
	return dara.Validate(s)
}
