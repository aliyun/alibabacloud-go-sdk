// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetVccRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *GetVccRequest
	GetClientToken() *string
	SetEnablePage(v bool) *GetVccRequest
	GetEnablePage() *bool
	SetPageNumber(v int32) *GetVccRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *GetVccRequest
	GetPageSize() *int32
	SetRegionId(v string) *GetVccRequest
	GetRegionId() *string
	SetVccId(v string) *GetVccRequest
	GetVccId() *string
}

type GetVccRequest struct {
	// By default, popApi is not ignored and idempotent
	//
	// example:
	//
	// c5e3130a-d02f-11ec-a7d3-0242ac110005
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Paging Parameters: The current parameters are obsolete.
	//
	// example:
	//
	// false
	EnablePage *bool `json:"EnablePage,omitempty" xml:"EnablePage,omitempty"`
	// Paging Parameters: The current parameters are obsolete.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// Paging Parameters: The current parameters are obsolete.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-wulanchabu
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the Lingjun connection instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// vcc-cn-zvp2w222001
	VccId *string `json:"VccId,omitempty" xml:"VccId,omitempty"`
}

func (s GetVccRequest) String() string {
	return dara.Prettify(s)
}

func (s GetVccRequest) GoString() string {
	return s.String()
}

func (s *GetVccRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *GetVccRequest) GetEnablePage() *bool {
	return s.EnablePage
}

func (s *GetVccRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *GetVccRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *GetVccRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetVccRequest) GetVccId() *string {
	return s.VccId
}

func (s *GetVccRequest) SetClientToken(v string) *GetVccRequest {
	s.ClientToken = &v
	return s
}

func (s *GetVccRequest) SetEnablePage(v bool) *GetVccRequest {
	s.EnablePage = &v
	return s
}

func (s *GetVccRequest) SetPageNumber(v int32) *GetVccRequest {
	s.PageNumber = &v
	return s
}

func (s *GetVccRequest) SetPageSize(v int32) *GetVccRequest {
	s.PageSize = &v
	return s
}

func (s *GetVccRequest) SetRegionId(v string) *GetVccRequest {
	s.RegionId = &v
	return s
}

func (s *GetVccRequest) SetVccId(v string) *GetVccRequest {
	s.VccId = &v
	return s
}

func (s *GetVccRequest) Validate() error {
	return dara.Validate(s)
}
