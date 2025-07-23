// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iIncidentPlanFieldPath interface {
	dara.Model
	String() string
	GoString() string
	SetFieldAlias(v string) *IncidentPlanFieldPath
	GetFieldAlias() *string
	SetFieldPath(v []*string) *IncidentPlanFieldPath
	GetFieldPath() []*string
}

type IncidentPlanFieldPath struct {
	FieldAlias *string   `json:"fieldAlias,omitempty" xml:"fieldAlias,omitempty"`
	FieldPath  []*string `json:"fieldPath,omitempty" xml:"fieldPath,omitempty" type:"Repeated"`
}

func (s IncidentPlanFieldPath) String() string {
	return dara.Prettify(s)
}

func (s IncidentPlanFieldPath) GoString() string {
	return s.String()
}

func (s *IncidentPlanFieldPath) GetFieldAlias() *string {
	return s.FieldAlias
}

func (s *IncidentPlanFieldPath) GetFieldPath() []*string {
	return s.FieldPath
}

func (s *IncidentPlanFieldPath) SetFieldAlias(v string) *IncidentPlanFieldPath {
	s.FieldAlias = &v
	return s
}

func (s *IncidentPlanFieldPath) SetFieldPath(v []*string) *IncidentPlanFieldPath {
	s.FieldPath = v
	return s
}

func (s *IncidentPlanFieldPath) Validate() error {
	return dara.Validate(s)
}
