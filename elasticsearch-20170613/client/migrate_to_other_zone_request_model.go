// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMigrateToOtherZoneRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v string) *MigrateToOtherZoneRequest
	GetBody() *string
	SetDryRun(v bool) *MigrateToOtherZoneRequest
	GetDryRun() *bool
}

type MigrateToOtherZoneRequest struct {
	Body *string `json:"body,omitempty" xml:"body,omitempty"`
	// Verify whether the zone node can be migrated. true indicates that the data is only verified and the migration task is not executed. false indicates that the migration task is executed after the verification is successful.
	//
	// This parameter is required.
	//
	// example:
	//
	// false
	DryRun *bool `json:"dryRun,omitempty" xml:"dryRun,omitempty"`
}

func (s MigrateToOtherZoneRequest) String() string {
	return dara.Prettify(s)
}

func (s MigrateToOtherZoneRequest) GoString() string {
	return s.String()
}

func (s *MigrateToOtherZoneRequest) GetBody() *string {
	return s.Body
}

func (s *MigrateToOtherZoneRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *MigrateToOtherZoneRequest) SetBody(v string) *MigrateToOtherZoneRequest {
	s.Body = &v
	return s
}

func (s *MigrateToOtherZoneRequest) SetDryRun(v bool) *MigrateToOtherZoneRequest {
	s.DryRun = &v
	return s
}

func (s *MigrateToOtherZoneRequest) Validate() error {
	return dara.Validate(s)
}
