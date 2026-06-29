// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateTaskWorkforceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetWorkforce(v []*SimpleWorkforce) *UpdateTaskWorkforceRequest
	GetWorkforce() []*SimpleWorkforce
}

type UpdateTaskWorkforceRequest struct {
	// User List.
	Workforce []*SimpleWorkforce `json:"Workforce,omitempty" xml:"Workforce,omitempty" type:"Repeated"`
}

func (s UpdateTaskWorkforceRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateTaskWorkforceRequest) GoString() string {
	return s.String()
}

func (s *UpdateTaskWorkforceRequest) GetWorkforce() []*SimpleWorkforce {
	return s.Workforce
}

func (s *UpdateTaskWorkforceRequest) SetWorkforce(v []*SimpleWorkforce) *UpdateTaskWorkforceRequest {
	s.Workforce = v
	return s
}

func (s *UpdateTaskWorkforceRequest) Validate() error {
	if s.Workforce != nil {
		for _, item := range s.Workforce {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
