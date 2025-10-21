// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iResourceQuota interface {
	dara.Model
	String() string
	GoString() string
	SetLimit(v *ResourceSpec) *ResourceQuota
	GetLimit() *ResourceSpec
	SetRequest(v *ResourceSpec) *ResourceQuota
	GetRequest() *ResourceSpec
	SetUsed(v *ResourceSpec) *ResourceQuota
	GetUsed() *ResourceSpec
}

type ResourceQuota struct {
	Limit   *ResourceSpec `json:"limit,omitempty" xml:"limit,omitempty"`
	Request *ResourceSpec `json:"request,omitempty" xml:"request,omitempty"`
	Used    *ResourceSpec `json:"used,omitempty" xml:"used,omitempty"`
}

func (s ResourceQuota) String() string {
	return dara.Prettify(s)
}

func (s ResourceQuota) GoString() string {
	return s.String()
}

func (s *ResourceQuota) GetLimit() *ResourceSpec {
	return s.Limit
}

func (s *ResourceQuota) GetRequest() *ResourceSpec {
	return s.Request
}

func (s *ResourceQuota) GetUsed() *ResourceSpec {
	return s.Used
}

func (s *ResourceQuota) SetLimit(v *ResourceSpec) *ResourceQuota {
	s.Limit = v
	return s
}

func (s *ResourceQuota) SetRequest(v *ResourceSpec) *ResourceQuota {
	s.Request = v
	return s
}

func (s *ResourceQuota) SetUsed(v *ResourceSpec) *ResourceQuota {
	s.Used = v
	return s
}

func (s *ResourceQuota) Validate() error {
	if s.Limit != nil {
		if err := s.Limit.Validate(); err != nil {
			return err
		}
	}
	if s.Request != nil {
		if err := s.Request.Validate(); err != nil {
			return err
		}
	}
	if s.Used != nil {
		if err := s.Used.Validate(); err != nil {
			return err
		}
	}
	return nil
}
