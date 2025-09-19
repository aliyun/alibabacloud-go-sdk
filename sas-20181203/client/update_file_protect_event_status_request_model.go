// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateFileProtectEventStatusRequest interface {
	dara.Model
	String() string
	GoString() string
	SetId(v []*int64) *UpdateFileProtectEventStatusRequest
	GetId() []*int64
	SetStatus(v int32) *UpdateFileProtectEventStatusRequest
	GetStatus() *int32
}

type UpdateFileProtectEventStatusRequest struct {
	// The IDs of the events.
	Id []*int64 `json:"Id,omitempty" xml:"Id,omitempty" type:"Repeated"`
	// The handling status of the event. Valid values:
	//
	// 	- **0**: unhandled
	//
	// 	- **1**: handled
	//
	// 	- **2**: added to the whitelist
	//
	// example:
	//
	// 0
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s UpdateFileProtectEventStatusRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateFileProtectEventStatusRequest) GoString() string {
	return s.String()
}

func (s *UpdateFileProtectEventStatusRequest) GetId() []*int64 {
	return s.Id
}

func (s *UpdateFileProtectEventStatusRequest) GetStatus() *int32 {
	return s.Status
}

func (s *UpdateFileProtectEventStatusRequest) SetId(v []*int64) *UpdateFileProtectEventStatusRequest {
	s.Id = v
	return s
}

func (s *UpdateFileProtectEventStatusRequest) SetStatus(v int32) *UpdateFileProtectEventStatusRequest {
	s.Status = &v
	return s
}

func (s *UpdateFileProtectEventStatusRequest) Validate() error {
	return dara.Validate(s)
}
