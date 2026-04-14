// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDedicatedIpChangeWarmupTypeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetId(v string) *DedicatedIpChangeWarmupTypeRequest
	GetId() *string
	SetWarmupType(v string) *DedicatedIpChangeWarmupTypeRequest
	GetWarmupType() *string
}

type DedicatedIpChangeWarmupTypeRequest struct {
	// Dedicated IP ID
	//
	// example:
	//
	// xxx
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// Warmup method
	//
	// example:
	//
	// sysCusStream
	WarmupType *string `json:"WarmupType,omitempty" xml:"WarmupType,omitempty"`
}

func (s DedicatedIpChangeWarmupTypeRequest) String() string {
	return dara.Prettify(s)
}

func (s DedicatedIpChangeWarmupTypeRequest) GoString() string {
	return s.String()
}

func (s *DedicatedIpChangeWarmupTypeRequest) GetId() *string {
	return s.Id
}

func (s *DedicatedIpChangeWarmupTypeRequest) GetWarmupType() *string {
	return s.WarmupType
}

func (s *DedicatedIpChangeWarmupTypeRequest) SetId(v string) *DedicatedIpChangeWarmupTypeRequest {
	s.Id = &v
	return s
}

func (s *DedicatedIpChangeWarmupTypeRequest) SetWarmupType(v string) *DedicatedIpChangeWarmupTypeRequest {
	s.WarmupType = &v
	return s
}

func (s *DedicatedIpChangeWarmupTypeRequest) Validate() error {
	return dara.Validate(s)
}
