// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDsgPlatformQueryProjectsAndSchemaFromMetaRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEngineName(v string) *DsgPlatformQueryProjectsAndSchemaFromMetaRequest
	GetEngineName() *string
}

type DsgPlatformQueryProjectsAndSchemaFromMetaRequest struct {
	// The type of the compute engine. Valid values:
	//
	// 	- ODPS.ODPS
	//
	// 	- EMR
	//
	// 	- HOLO.POSTGRES
	//
	// This parameter is required.
	//
	// example:
	//
	// EMR
	EngineName *string `json:"EngineName,omitempty" xml:"EngineName,omitempty"`
}

func (s DsgPlatformQueryProjectsAndSchemaFromMetaRequest) String() string {
	return dara.Prettify(s)
}

func (s DsgPlatformQueryProjectsAndSchemaFromMetaRequest) GoString() string {
	return s.String()
}

func (s *DsgPlatformQueryProjectsAndSchemaFromMetaRequest) GetEngineName() *string {
	return s.EngineName
}

func (s *DsgPlatformQueryProjectsAndSchemaFromMetaRequest) SetEngineName(v string) *DsgPlatformQueryProjectsAndSchemaFromMetaRequest {
	s.EngineName = &v
	return s
}

func (s *DsgPlatformQueryProjectsAndSchemaFromMetaRequest) Validate() error {
	return dara.Validate(s)
}
