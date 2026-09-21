// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetFileProtectEventRequest interface {
	dara.Model
	String() string
	GoString() string
	SetId(v int64) *GetFileProtectEventRequest
	GetId() *int64
}

type GetFileProtectEventRequest struct {
	// The event ID.
	//
	// This parameter is required. If this parameter is not specified, the API returns HTTP 400 with error code -101. You can call ListFileProtectEvent to obtain valid event IDs.
	//
	// example:
	//
	// 12
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s GetFileProtectEventRequest) String() string {
	return dara.Prettify(s)
}

func (s GetFileProtectEventRequest) GoString() string {
	return s.String()
}

func (s *GetFileProtectEventRequest) GetId() *int64 {
	return s.Id
}

func (s *GetFileProtectEventRequest) SetId(v int64) *GetFileProtectEventRequest {
	s.Id = &v
	return s
}

func (s *GetFileProtectEventRequest) Validate() error {
	return dara.Validate(s)
}
