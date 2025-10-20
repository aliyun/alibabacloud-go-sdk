// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRollbackTableRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstant(v *FullInstant) *RollbackTableRequest
	GetInstant() *FullInstant
}

type RollbackTableRequest struct {
	Instant *FullInstant `json:"instant,omitempty" xml:"instant,omitempty"`
}

func (s RollbackTableRequest) String() string {
	return dara.Prettify(s)
}

func (s RollbackTableRequest) GoString() string {
	return s.String()
}

func (s *RollbackTableRequest) GetInstant() *FullInstant {
	return s.Instant
}

func (s *RollbackTableRequest) SetInstant(v *FullInstant) *RollbackTableRequest {
	s.Instant = v
	return s
}

func (s *RollbackTableRequest) Validate() error {
	if s.Instant != nil {
		if err := s.Instant.Validate(); err != nil {
			return err
		}
	}
	return nil
}
