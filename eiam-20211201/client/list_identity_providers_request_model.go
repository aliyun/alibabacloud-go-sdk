// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListIdentityProvidersRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDirection(v string) *ListIdentityProvidersRequest
	GetDirection() *string
	SetInstanceId(v string) *ListIdentityProvidersRequest
	GetInstanceId() *string
	SetPageNumber(v int64) *ListIdentityProvidersRequest
	GetPageNumber() *int64
	SetPageSize(v int64) *ListIdentityProvidersRequest
	GetPageSize() *int64
}

type ListIdentityProvidersRequest struct {
	// example:
	//
	// pull
	Direction *string `json:"Direction,omitempty" xml:"Direction,omitempty"`
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The number of the page to return.
	//
	// example:
	//
	// 1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page.
	//
	// example:
	//
	// 20
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListIdentityProvidersRequest) String() string {
	return dara.Prettify(s)
}

func (s ListIdentityProvidersRequest) GoString() string {
	return s.String()
}

func (s *ListIdentityProvidersRequest) GetDirection() *string {
	return s.Direction
}

func (s *ListIdentityProvidersRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListIdentityProvidersRequest) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *ListIdentityProvidersRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListIdentityProvidersRequest) SetDirection(v string) *ListIdentityProvidersRequest {
	s.Direction = &v
	return s
}

func (s *ListIdentityProvidersRequest) SetInstanceId(v string) *ListIdentityProvidersRequest {
	s.InstanceId = &v
	return s
}

func (s *ListIdentityProvidersRequest) SetPageNumber(v int64) *ListIdentityProvidersRequest {
	s.PageNumber = &v
	return s
}

func (s *ListIdentityProvidersRequest) SetPageSize(v int64) *ListIdentityProvidersRequest {
	s.PageSize = &v
	return s
}

func (s *ListIdentityProvidersRequest) Validate() error {
	return dara.Validate(s)
}
