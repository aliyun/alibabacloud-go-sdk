// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDBVersionDetail interface {
	dara.Model
	String() string
	GoString() string
	SetSpecs(v []*DBVersionDetailSpecs) *DBVersionDetail
	GetSpecs() []*DBVersionDetailSpecs
	SetStatus(v string) *DBVersionDetail
	GetStatus() *string
	SetVersion(v string) *DBVersionDetail
	GetVersion() *string
}

type DBVersionDetail struct {
	// The specifications supported by the version.
	Specs []*DBVersionDetailSpecs `json:"specs,omitempty" xml:"specs,omitempty" type:"Repeated"`
	// The status of the version.
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// The version number.
	Version *string `json:"version,omitempty" xml:"version,omitempty"`
}

func (s DBVersionDetail) String() string {
	return dara.Prettify(s)
}

func (s DBVersionDetail) GoString() string {
	return s.String()
}

func (s *DBVersionDetail) GetSpecs() []*DBVersionDetailSpecs {
	return s.Specs
}

func (s *DBVersionDetail) GetStatus() *string {
	return s.Status
}

func (s *DBVersionDetail) GetVersion() *string {
	return s.Version
}

func (s *DBVersionDetail) SetSpecs(v []*DBVersionDetailSpecs) *DBVersionDetail {
	s.Specs = v
	return s
}

func (s *DBVersionDetail) SetStatus(v string) *DBVersionDetail {
	s.Status = &v
	return s
}

func (s *DBVersionDetail) SetVersion(v string) *DBVersionDetail {
	s.Version = &v
	return s
}

func (s *DBVersionDetail) Validate() error {
	if s.Specs != nil {
		for _, item := range s.Specs {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DBVersionDetailSpecs struct {
	// The component specifications.
	ComponentSpecs []*DBVersionDetailSpecsComponentSpecs `json:"componentSpecs,omitempty" xml:"componentSpecs,omitempty" type:"Repeated"`
	// Indicates whether the high availability (HA) edition is supported.
	IsHA *bool `json:"isHA,omitempty" xml:"isHA,omitempty"`
	// Indicates whether the Standalone edition is supported.
	IsStandalone *bool `json:"isStandalone,omitempty" xml:"isStandalone,omitempty"`
	// Indicates whether the Standalone Pro edition is supported.
	IsStandalonePro *bool `json:"isStandalonePro,omitempty" xml:"isStandalonePro,omitempty"`
	// The zone mode of the specification.
	ZoneMode *string `json:"zoneMode,omitempty" xml:"zoneMode,omitempty"`
}

func (s DBVersionDetailSpecs) String() string {
	return dara.Prettify(s)
}

func (s DBVersionDetailSpecs) GoString() string {
	return s.String()
}

func (s *DBVersionDetailSpecs) GetComponentSpecs() []*DBVersionDetailSpecsComponentSpecs {
	return s.ComponentSpecs
}

func (s *DBVersionDetailSpecs) GetIsHA() *bool {
	return s.IsHA
}

func (s *DBVersionDetailSpecs) GetIsStandalone() *bool {
	return s.IsStandalone
}

func (s *DBVersionDetailSpecs) GetIsStandalonePro() *bool {
	return s.IsStandalonePro
}

func (s *DBVersionDetailSpecs) GetZoneMode() *string {
	return s.ZoneMode
}

func (s *DBVersionDetailSpecs) SetComponentSpecs(v []*DBVersionDetailSpecsComponentSpecs) *DBVersionDetailSpecs {
	s.ComponentSpecs = v
	return s
}

func (s *DBVersionDetailSpecs) SetIsHA(v bool) *DBVersionDetailSpecs {
	s.IsHA = &v
	return s
}

func (s *DBVersionDetailSpecs) SetIsStandalone(v bool) *DBVersionDetailSpecs {
	s.IsStandalone = &v
	return s
}

func (s *DBVersionDetailSpecs) SetIsStandalonePro(v bool) *DBVersionDetailSpecs {
	s.IsStandalonePro = &v
	return s
}

func (s *DBVersionDetailSpecs) SetZoneMode(v string) *DBVersionDetailSpecs {
	s.ZoneMode = &v
	return s
}

func (s *DBVersionDetailSpecs) Validate() error {
	if s.ComponentSpecs != nil {
		for _, item := range s.ComponentSpecs {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DBVersionDetailSpecsComponentSpecs struct {
	// The default number of replicas for the component.
	DefaultReplicas *int32 `json:"defaultReplicas,omitempty" xml:"defaultReplicas,omitempty"`
	// The maximum number of replicas for the component.
	MaxReplicas *int32 `json:"maxReplicas,omitempty" xml:"maxReplicas,omitempty"`
	// The minimum number of replicas for the component.
	MinReplicas *int32 `json:"minReplicas,omitempty" xml:"minReplicas,omitempty"`
	// The name of the component.
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The available specifications for the component.
	Specs []*string `json:"specs,omitempty" xml:"specs,omitempty" type:"Repeated"`
	// The step size for scaling the number of replicas.
	Step *int32 `json:"step,omitempty" xml:"step,omitempty"`
	// The type of the component.
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s DBVersionDetailSpecsComponentSpecs) String() string {
	return dara.Prettify(s)
}

func (s DBVersionDetailSpecsComponentSpecs) GoString() string {
	return s.String()
}

func (s *DBVersionDetailSpecsComponentSpecs) GetDefaultReplicas() *int32 {
	return s.DefaultReplicas
}

func (s *DBVersionDetailSpecsComponentSpecs) GetMaxReplicas() *int32 {
	return s.MaxReplicas
}

func (s *DBVersionDetailSpecsComponentSpecs) GetMinReplicas() *int32 {
	return s.MinReplicas
}

func (s *DBVersionDetailSpecsComponentSpecs) GetName() *string {
	return s.Name
}

func (s *DBVersionDetailSpecsComponentSpecs) GetSpecs() []*string {
	return s.Specs
}

func (s *DBVersionDetailSpecsComponentSpecs) GetStep() *int32 {
	return s.Step
}

func (s *DBVersionDetailSpecsComponentSpecs) GetType() *string {
	return s.Type
}

func (s *DBVersionDetailSpecsComponentSpecs) SetDefaultReplicas(v int32) *DBVersionDetailSpecsComponentSpecs {
	s.DefaultReplicas = &v
	return s
}

func (s *DBVersionDetailSpecsComponentSpecs) SetMaxReplicas(v int32) *DBVersionDetailSpecsComponentSpecs {
	s.MaxReplicas = &v
	return s
}

func (s *DBVersionDetailSpecsComponentSpecs) SetMinReplicas(v int32) *DBVersionDetailSpecsComponentSpecs {
	s.MinReplicas = &v
	return s
}

func (s *DBVersionDetailSpecsComponentSpecs) SetName(v string) *DBVersionDetailSpecsComponentSpecs {
	s.Name = &v
	return s
}

func (s *DBVersionDetailSpecsComponentSpecs) SetSpecs(v []*string) *DBVersionDetailSpecsComponentSpecs {
	s.Specs = v
	return s
}

func (s *DBVersionDetailSpecsComponentSpecs) SetStep(v int32) *DBVersionDetailSpecsComponentSpecs {
	s.Step = &v
	return s
}

func (s *DBVersionDetailSpecsComponentSpecs) SetType(v string) *DBVersionDetailSpecsComponentSpecs {
	s.Type = &v
	return s
}

func (s *DBVersionDetailSpecsComponentSpecs) Validate() error {
	return dara.Validate(s)
}
