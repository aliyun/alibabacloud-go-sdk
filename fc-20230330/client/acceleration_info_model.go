// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAccelerationInfo interface {
	dara.Model
	String() string
	GoString() string
	SetStatus(v string) *AccelerationInfo
	GetStatus() *string
}

type AccelerationInfo struct {
	// example:
	//
	// deprecated
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
}

func (s AccelerationInfo) String() string {
	return dara.Prettify(s)
}

func (s AccelerationInfo) GoString() string {
	return s.String()
}

func (s *AccelerationInfo) GetStatus() *string {
	return s.Status
}

func (s *AccelerationInfo) SetStatus(v string) *AccelerationInfo {
	s.Status = &v
	return s
}

func (s *AccelerationInfo) Validate() error {
	return dara.Validate(s)
}
