// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetRoutineAccessTokenRequest interface {
	dara.Model
	String() string
	GoString() string
	SetName(v string) *GetRoutineAccessTokenRequest
	GetName() *string
}

type GetRoutineAccessTokenRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// test-routine1
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s GetRoutineAccessTokenRequest) String() string {
	return dara.Prettify(s)
}

func (s GetRoutineAccessTokenRequest) GoString() string {
	return s.String()
}

func (s *GetRoutineAccessTokenRequest) GetName() *string {
	return s.Name
}

func (s *GetRoutineAccessTokenRequest) SetName(v string) *GetRoutineAccessTokenRequest {
	s.Name = &v
	return s
}

func (s *GetRoutineAccessTokenRequest) Validate() error {
	return dara.Validate(s)
}
