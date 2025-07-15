// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAllCustomTemplatesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerId(v int64) *GetAllCustomTemplatesRequest
	GetOwnerId() *int64
	SetRegionId(v string) *GetAllCustomTemplatesRequest
	GetRegionId() *string
	SetUserId(v string) *GetAllCustomTemplatesRequest
	GetUserId() *string
}

type GetAllCustomTemplatesRequest struct {
	OwnerId  *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the user.
	//
	// example:
	//
	// 25346073170691****
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s GetAllCustomTemplatesRequest) String() string {
	return dara.Prettify(s)
}

func (s GetAllCustomTemplatesRequest) GoString() string {
	return s.String()
}

func (s *GetAllCustomTemplatesRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *GetAllCustomTemplatesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetAllCustomTemplatesRequest) GetUserId() *string {
	return s.UserId
}

func (s *GetAllCustomTemplatesRequest) SetOwnerId(v int64) *GetAllCustomTemplatesRequest {
	s.OwnerId = &v
	return s
}

func (s *GetAllCustomTemplatesRequest) SetRegionId(v string) *GetAllCustomTemplatesRequest {
	s.RegionId = &v
	return s
}

func (s *GetAllCustomTemplatesRequest) SetUserId(v string) *GetAllCustomTemplatesRequest {
	s.UserId = &v
	return s
}

func (s *GetAllCustomTemplatesRequest) Validate() error {
	return dara.Validate(s)
}
