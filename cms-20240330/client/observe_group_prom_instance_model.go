// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iObserveGroupPromInstance interface {
	dara.Model
	String() string
	GoString() string
	SetId(v string) *ObserveGroupPromInstance
	GetId() *string
	SetKind(v string) *ObserveGroupPromInstance
	GetKind() *string
	SetRegion(v string) *ObserveGroupPromInstance
	GetRegion() *string
	SetTime(v string) *ObserveGroupPromInstance
	GetTime() *string
}

type ObserveGroupPromInstance struct {
	// The ID of the Managed Service for Prometheus instance, such as rw-xxxxxxxxxx.
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// The source of the instance. Valid values:
	//
	// - system: The system automatically identifies the instance based on the workspace or UModel.
	//
	// - custom: The user manually selects the instance in the console.
	Kind *string `json:"kind,omitempty" xml:"kind,omitempty"`
	// The region where the Managed Service for Prometheus instance resides. If this parameter is left empty, the backend automatically populates the region based on the workspace to which the application group belongs.
	Region *string `json:"region,omitempty" xml:"region,omitempty"`
	// The time when the record was written or selected. Format: yyyy-MM-dd HH:mm:ss.
	Time *string `json:"time,omitempty" xml:"time,omitempty"`
}

func (s ObserveGroupPromInstance) String() string {
	return dara.Prettify(s)
}

func (s ObserveGroupPromInstance) GoString() string {
	return s.String()
}

func (s *ObserveGroupPromInstance) GetId() *string {
	return s.Id
}

func (s *ObserveGroupPromInstance) GetKind() *string {
	return s.Kind
}

func (s *ObserveGroupPromInstance) GetRegion() *string {
	return s.Region
}

func (s *ObserveGroupPromInstance) GetTime() *string {
	return s.Time
}

func (s *ObserveGroupPromInstance) SetId(v string) *ObserveGroupPromInstance {
	s.Id = &v
	return s
}

func (s *ObserveGroupPromInstance) SetKind(v string) *ObserveGroupPromInstance {
	s.Kind = &v
	return s
}

func (s *ObserveGroupPromInstance) SetRegion(v string) *ObserveGroupPromInstance {
	s.Region = &v
	return s
}

func (s *ObserveGroupPromInstance) SetTime(v string) *ObserveGroupPromInstance {
	s.Time = &v
	return s
}

func (s *ObserveGroupPromInstance) Validate() error {
	return dara.Validate(s)
}
