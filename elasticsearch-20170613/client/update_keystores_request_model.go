// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateKeystoresRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRemove(v []*string) *UpdateKeystoresRequest
	GetRemove() []*string
	SetUpdate(v map[string]*string) *UpdateKeystoresRequest
	GetUpdate() map[string]*string
	SetForce(v string) *UpdateKeystoresRequest
	GetForce() *string
}

type UpdateKeystoresRequest struct {
	Remove []*string          `json:"remove,omitempty" xml:"remove,omitempty" type:"Repeated"`
	Update map[string]*string `json:"update,omitempty" xml:"update,omitempty"`
	// example:
	//
	// false
	Force *string `json:"force,omitempty" xml:"force,omitempty"`
}

func (s UpdateKeystoresRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateKeystoresRequest) GoString() string {
	return s.String()
}

func (s *UpdateKeystoresRequest) GetRemove() []*string {
	return s.Remove
}

func (s *UpdateKeystoresRequest) GetUpdate() map[string]*string {
	return s.Update
}

func (s *UpdateKeystoresRequest) GetForce() *string {
	return s.Force
}

func (s *UpdateKeystoresRequest) SetRemove(v []*string) *UpdateKeystoresRequest {
	s.Remove = v
	return s
}

func (s *UpdateKeystoresRequest) SetUpdate(v map[string]*string) *UpdateKeystoresRequest {
	s.Update = v
	return s
}

func (s *UpdateKeystoresRequest) SetForce(v string) *UpdateKeystoresRequest {
	s.Force = &v
	return s
}

func (s *UpdateKeystoresRequest) Validate() error {
	return dara.Validate(s)
}
