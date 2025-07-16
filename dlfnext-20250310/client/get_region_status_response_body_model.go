// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetRegionStatusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetServiceRoleExists(v bool) *GetRegionStatusResponseBody
	GetServiceRoleExists() *bool
	SetStatus(v string) *GetRegionStatusResponseBody
	GetStatus() *string
}

type GetRegionStatusResponseBody struct {
	// example:
	//
	// true
	ServiceRoleExists *bool `json:"serviceRoleExists,omitempty" xml:"serviceRoleExists,omitempty"`
	// example:
	//
	// READY
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
}

func (s GetRegionStatusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetRegionStatusResponseBody) GoString() string {
	return s.String()
}

func (s *GetRegionStatusResponseBody) GetServiceRoleExists() *bool {
	return s.ServiceRoleExists
}

func (s *GetRegionStatusResponseBody) GetStatus() *string {
	return s.Status
}

func (s *GetRegionStatusResponseBody) SetServiceRoleExists(v bool) *GetRegionStatusResponseBody {
	s.ServiceRoleExists = &v
	return s
}

func (s *GetRegionStatusResponseBody) SetStatus(v string) *GetRegionStatusResponseBody {
	s.Status = &v
	return s
}

func (s *GetRegionStatusResponseBody) Validate() error {
	return dara.Validate(s)
}
