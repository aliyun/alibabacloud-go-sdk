// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSubscriptionSharedPackagesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPageNo(v int32) *ListSubscriptionSharedPackagesRequest
	GetPageNo() *int32
	SetPageSize(v int32) *ListSubscriptionSharedPackagesRequest
	GetPageSize() *int32
	SetStatusList(v []*string) *ListSubscriptionSharedPackagesRequest
	GetStatusList() []*string
}

type ListSubscriptionSharedPackagesRequest struct {
	// The page number. Default value: 1. The value must be a positive integer.
	//
	// example:
	//
	// 1
	PageNo *int32 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// The number of entries per page. Default value: 10.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The list of statuses used for filtering.
	StatusList []*string `json:"StatusList,omitempty" xml:"StatusList,omitempty" type:"Repeated"`
}

func (s ListSubscriptionSharedPackagesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListSubscriptionSharedPackagesRequest) GoString() string {
	return s.String()
}

func (s *ListSubscriptionSharedPackagesRequest) GetPageNo() *int32 {
	return s.PageNo
}

func (s *ListSubscriptionSharedPackagesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListSubscriptionSharedPackagesRequest) GetStatusList() []*string {
	return s.StatusList
}

func (s *ListSubscriptionSharedPackagesRequest) SetPageNo(v int32) *ListSubscriptionSharedPackagesRequest {
	s.PageNo = &v
	return s
}

func (s *ListSubscriptionSharedPackagesRequest) SetPageSize(v int32) *ListSubscriptionSharedPackagesRequest {
	s.PageSize = &v
	return s
}

func (s *ListSubscriptionSharedPackagesRequest) SetStatusList(v []*string) *ListSubscriptionSharedPackagesRequest {
	s.StatusList = v
	return s
}

func (s *ListSubscriptionSharedPackagesRequest) Validate() error {
	return dara.Validate(s)
}
