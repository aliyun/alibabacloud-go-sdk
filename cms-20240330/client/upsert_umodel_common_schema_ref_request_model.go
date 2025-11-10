// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpsertUmodelCommonSchemaRefRequest interface {
	dara.Model
	String() string
	GoString() string
	SetGroup(v string) *UpsertUmodelCommonSchemaRefRequest
	GetGroup() *string
	SetVersion(v string) *UpsertUmodelCommonSchemaRefRequest
	GetVersion() *string
}

type UpsertUmodelCommonSchemaRefRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// card-service-daily01
	Group *string `json:"group,omitempty" xml:"group,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 0.1.0
	Version *string `json:"version,omitempty" xml:"version,omitempty"`
}

func (s UpsertUmodelCommonSchemaRefRequest) String() string {
	return dara.Prettify(s)
}

func (s UpsertUmodelCommonSchemaRefRequest) GoString() string {
	return s.String()
}

func (s *UpsertUmodelCommonSchemaRefRequest) GetGroup() *string {
	return s.Group
}

func (s *UpsertUmodelCommonSchemaRefRequest) GetVersion() *string {
	return s.Version
}

func (s *UpsertUmodelCommonSchemaRefRequest) SetGroup(v string) *UpsertUmodelCommonSchemaRefRequest {
	s.Group = &v
	return s
}

func (s *UpsertUmodelCommonSchemaRefRequest) SetVersion(v string) *UpsertUmodelCommonSchemaRefRequest {
	s.Version = &v
	return s
}

func (s *UpsertUmodelCommonSchemaRefRequest) Validate() error {
	return dara.Validate(s)
}
