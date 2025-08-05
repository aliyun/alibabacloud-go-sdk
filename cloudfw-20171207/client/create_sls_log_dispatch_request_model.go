// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSlsLogDispatchRequest interface {
	dara.Model
	String() string
	GoString() string
	SetSlsRegionId(v string) *CreateSlsLogDispatchRequest
	GetSlsRegionId() *string
	SetTtl(v int64) *CreateSlsLogDispatchRequest
	GetTtl() *int64
}

type CreateSlsLogDispatchRequest struct {
	// The region ID of the Simple Log Service project.
	//
	// example:
	//
	// ap-southeast-1
	SlsRegionId *string `json:"SlsRegionId,omitempty" xml:"SlsRegionId,omitempty"`
	// The log retention period. Unit: days.
	//
	// example:
	//
	// 20
	Ttl *int64 `json:"Ttl,omitempty" xml:"Ttl,omitempty"`
}

func (s CreateSlsLogDispatchRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateSlsLogDispatchRequest) GoString() string {
	return s.String()
}

func (s *CreateSlsLogDispatchRequest) GetSlsRegionId() *string {
	return s.SlsRegionId
}

func (s *CreateSlsLogDispatchRequest) GetTtl() *int64 {
	return s.Ttl
}

func (s *CreateSlsLogDispatchRequest) SetSlsRegionId(v string) *CreateSlsLogDispatchRequest {
	s.SlsRegionId = &v
	return s
}

func (s *CreateSlsLogDispatchRequest) SetTtl(v int64) *CreateSlsLogDispatchRequest {
	s.Ttl = &v
	return s
}

func (s *CreateSlsLogDispatchRequest) Validate() error {
	return dara.Validate(s)
}
