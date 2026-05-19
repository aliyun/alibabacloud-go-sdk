// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetFileProtectClientEventRequest interface {
	dara.Model
	String() string
	GoString() string
	SetId(v int64) *GetFileProtectClientEventRequest
	GetId() *int64
}

type GetFileProtectClientEventRequest struct {
	// example:
	//
	// 131231
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s GetFileProtectClientEventRequest) String() string {
	return dara.Prettify(s)
}

func (s GetFileProtectClientEventRequest) GoString() string {
	return s.String()
}

func (s *GetFileProtectClientEventRequest) GetId() *int64 {
	return s.Id
}

func (s *GetFileProtectClientEventRequest) SetId(v int64) *GetFileProtectClientEventRequest {
	s.Id = &v
	return s
}

func (s *GetFileProtectClientEventRequest) Validate() error {
	return dara.Validate(s)
}
