// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOpenArmsServiceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerId(v int64) *OpenArmsServiceRequest
	GetOwnerId() *int64
	SetType(v string) *OpenArmsServiceRequest
	GetType() *string
}

type OpenArmsServiceRequest struct {
	OwnerId *int64 `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// This parameter is required.
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s OpenArmsServiceRequest) String() string {
	return dara.Prettify(s)
}

func (s OpenArmsServiceRequest) GoString() string {
	return s.String()
}

func (s *OpenArmsServiceRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *OpenArmsServiceRequest) GetType() *string {
	return s.Type
}

func (s *OpenArmsServiceRequest) SetOwnerId(v int64) *OpenArmsServiceRequest {
	s.OwnerId = &v
	return s
}

func (s *OpenArmsServiceRequest) SetType(v string) *OpenArmsServiceRequest {
	s.Type = &v
	return s
}

func (s *OpenArmsServiceRequest) Validate() error {
	return dara.Validate(s)
}
