// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iKibanaNodeConfiguration interface {
	dara.Model
	String() string
	GoString() string
	SetAmount(v int64) *KibanaNodeConfiguration
	GetAmount() *int64
	SetDisk(v int64) *KibanaNodeConfiguration
	GetDisk() *int64
	SetSpec(v string) *KibanaNodeConfiguration
	GetSpec() *string
}

type KibanaNodeConfiguration struct {
	Amount *int64 `json:"amount,omitempty" xml:"amount,omitempty"`
	Disk   *int64 `json:"disk,omitempty" xml:"disk,omitempty"`
	// This parameter is required.
	Spec *string `json:"spec,omitempty" xml:"spec,omitempty"`
}

func (s KibanaNodeConfiguration) String() string {
	return dara.Prettify(s)
}

func (s KibanaNodeConfiguration) GoString() string {
	return s.String()
}

func (s *KibanaNodeConfiguration) GetAmount() *int64 {
	return s.Amount
}

func (s *KibanaNodeConfiguration) GetDisk() *int64 {
	return s.Disk
}

func (s *KibanaNodeConfiguration) GetSpec() *string {
	return s.Spec
}

func (s *KibanaNodeConfiguration) SetAmount(v int64) *KibanaNodeConfiguration {
	s.Amount = &v
	return s
}

func (s *KibanaNodeConfiguration) SetDisk(v int64) *KibanaNodeConfiguration {
	s.Disk = &v
	return s
}

func (s *KibanaNodeConfiguration) SetSpec(v string) *KibanaNodeConfiguration {
	s.Spec = &v
	return s
}

func (s *KibanaNodeConfiguration) Validate() error {
	return dara.Validate(s)
}
