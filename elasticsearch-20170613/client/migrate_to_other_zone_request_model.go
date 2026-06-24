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
	// Specifies whether to perform a dry run to check whether zone node migration is feasible. Valid values:
	//
	// - true: performs only a validation check without executing the migration task.
	//
	// - false: executes the migration task after the validation check succeeds.
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
