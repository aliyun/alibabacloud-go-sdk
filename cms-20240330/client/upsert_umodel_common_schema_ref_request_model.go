// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpsertUmodelCommonSchemaRefRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDryRun(v bool) *UpsertUmodelCommonSchemaRefRequest
	GetDryRun() *bool
	SetGroup(v string) *UpsertUmodelCommonSchemaRefRequest
	GetGroup() *string
	SetVersion(v string) *UpsertUmodelCommonSchemaRefRequest
	GetVersion() *string
}

type UpsertUmodelCommonSchemaRefRequest struct {
	// Specifies whether to perform a dry run for this request. Valid values:
	//
	// - true: Sends a check request without performing actual changes.
	//
	// - false: Sends a normal request.
	//
	// Default value: false.
	DryRun *bool `json:"dryRun,omitempty" xml:"dryRun,omitempty"`
	// The public Umodel schema group.
	//
	// This parameter is required.
	//
	// example:
	//
	// card-service-daily01
	Group *string `json:"group,omitempty" xml:"group,omitempty"`
	// The version number.
	//
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

func (s *UpsertUmodelCommonSchemaRefRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *UpsertUmodelCommonSchemaRefRequest) GetGroup() *string {
	return s.Group
}

func (s *UpsertUmodelCommonSchemaRefRequest) GetVersion() *string {
	return s.Version
}

func (s *UpsertUmodelCommonSchemaRefRequest) SetDryRun(v bool) *UpsertUmodelCommonSchemaRefRequest {
	s.DryRun = &v
	return s
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
