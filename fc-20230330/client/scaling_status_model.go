// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iScalingStatus interface {
	dara.Model
	String() string
	GoString() string
	SetCurrentError(v string) *ScalingStatus
	GetCurrentError() *string
	SetResourceCount(v int64) *ScalingStatus
	GetResourceCount() *int64
}

type ScalingStatus struct {
	CurrentError  *string `json:"currentError,omitempty" xml:"currentError,omitempty"`
	ResourceCount *int64  `json:"resourceCount,omitempty" xml:"resourceCount,omitempty"`
}

func (s ScalingStatus) String() string {
	return dara.Prettify(s)
}

func (s ScalingStatus) GoString() string {
	return s.String()
}

func (s *ScalingStatus) GetCurrentError() *string {
	return s.CurrentError
}

func (s *ScalingStatus) GetResourceCount() *int64 {
	return s.ResourceCount
}

func (s *ScalingStatus) SetCurrentError(v string) *ScalingStatus {
	s.CurrentError = &v
	return s
}

func (s *ScalingStatus) SetResourceCount(v int64) *ScalingStatus {
	s.ResourceCount = &v
	return s
}

func (s *ScalingStatus) Validate() error {
	return dara.Validate(s)
}
