// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPdpServiceInfo interface {
	dara.Model
	String() string
	GoString() string
	SetId(v int64) *PdpServiceInfo
	GetId() *int64
	SetName(v string) *PdpServiceInfo
	GetName() *string
}

type PdpServiceInfo struct {
	// example:
	//
	// 1
	Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
	// example:
	//
	// pdp-pbc-service
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s PdpServiceInfo) String() string {
	return dara.Prettify(s)
}

func (s PdpServiceInfo) GoString() string {
	return s.String()
}

func (s *PdpServiceInfo) GetId() *int64 {
	return s.Id
}

func (s *PdpServiceInfo) GetName() *string {
	return s.Name
}

func (s *PdpServiceInfo) SetId(v int64) *PdpServiceInfo {
	s.Id = &v
	return s
}

func (s *PdpServiceInfo) SetName(v string) *PdpServiceInfo {
	s.Name = &v
	return s
}

func (s *PdpServiceInfo) Validate() error {
	return dara.Validate(s)
}
