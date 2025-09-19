// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetFileProtectEventCountRequest interface {
	dara.Model
	String() string
	GoString() string
	SetStatus(v int32) *GetFileProtectEventCountRequest
	GetStatus() *int32
}

type GetFileProtectEventCountRequest struct {
	// The status of the event. Valid values:
	//
	// 	- 0: unhandled
	//
	// 	- 1: handled
	//
	// example:
	//
	// 1
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s GetFileProtectEventCountRequest) String() string {
	return dara.Prettify(s)
}

func (s GetFileProtectEventCountRequest) GoString() string {
	return s.String()
}

func (s *GetFileProtectEventCountRequest) GetStatus() *int32 {
	return s.Status
}

func (s *GetFileProtectEventCountRequest) SetStatus(v int32) *GetFileProtectEventCountRequest {
	s.Status = &v
	return s
}

func (s *GetFileProtectEventCountRequest) Validate() error {
	return dara.Validate(s)
}
