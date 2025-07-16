// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOpenAlimtServiceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerId(v int64) *OpenAlimtServiceRequest
	GetOwnerId() *int64
	SetType(v string) *OpenAlimtServiceRequest
	GetType() *string
}

type OpenAlimtServiceRequest struct {
	OwnerId *int64 `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// example:
	//
	// id
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s OpenAlimtServiceRequest) String() string {
	return dara.Prettify(s)
}

func (s OpenAlimtServiceRequest) GoString() string {
	return s.String()
}

func (s *OpenAlimtServiceRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *OpenAlimtServiceRequest) GetType() *string {
	return s.Type
}

func (s *OpenAlimtServiceRequest) SetOwnerId(v int64) *OpenAlimtServiceRequest {
	s.OwnerId = &v
	return s
}

func (s *OpenAlimtServiceRequest) SetType(v string) *OpenAlimtServiceRequest {
	s.Type = &v
	return s
}

func (s *OpenAlimtServiceRequest) Validate() error {
	return dara.Validate(s)
}
