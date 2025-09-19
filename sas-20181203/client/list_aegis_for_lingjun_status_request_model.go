// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAegisForLingjunStatusRequest interface {
	dara.Model
	String() string
	GoString() string
	SetUuids(v []*string) *ListAegisForLingjunStatusRequest
	GetUuids() []*string
}

type ListAegisForLingjunStatusRequest struct {
	// List of unique UUIDs for Lingjun bare metal.
	Uuids []*string `json:"Uuids,omitempty" xml:"Uuids,omitempty" type:"Repeated"`
}

func (s ListAegisForLingjunStatusRequest) String() string {
	return dara.Prettify(s)
}

func (s ListAegisForLingjunStatusRequest) GoString() string {
	return s.String()
}

func (s *ListAegisForLingjunStatusRequest) GetUuids() []*string {
	return s.Uuids
}

func (s *ListAegisForLingjunStatusRequest) SetUuids(v []*string) *ListAegisForLingjunStatusRequest {
	s.Uuids = v
	return s
}

func (s *ListAegisForLingjunStatusRequest) Validate() error {
	return dara.Validate(s)
}
