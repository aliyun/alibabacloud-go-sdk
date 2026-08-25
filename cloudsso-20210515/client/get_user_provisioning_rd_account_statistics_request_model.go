// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetUserProvisioningRdAccountStatisticsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDirectoryId(v string) *GetUserProvisioningRdAccountStatisticsRequest
	GetDirectoryId() *string
	SetRdMemberId(v string) *GetUserProvisioningRdAccountStatisticsRequest
	GetRdMemberId() *string
}

type GetUserProvisioningRdAccountStatisticsRequest struct {
	// The ID of the resource directory.
	//
	// example:
	//
	// d-003qew84****
	DirectoryId *string `json:"DirectoryId,omitempty" xml:"DirectoryId,omitempty"`
	// The ID of the member in the resource directory.
	//
	// example:
	//
	// 1743382******
	RdMemberId *string `json:"RdMemberId,omitempty" xml:"RdMemberId,omitempty"`
}

func (s GetUserProvisioningRdAccountStatisticsRequest) String() string {
	return dara.Prettify(s)
}

func (s GetUserProvisioningRdAccountStatisticsRequest) GoString() string {
	return s.String()
}

func (s *GetUserProvisioningRdAccountStatisticsRequest) GetDirectoryId() *string {
	return s.DirectoryId
}

func (s *GetUserProvisioningRdAccountStatisticsRequest) GetRdMemberId() *string {
	return s.RdMemberId
}

func (s *GetUserProvisioningRdAccountStatisticsRequest) SetDirectoryId(v string) *GetUserProvisioningRdAccountStatisticsRequest {
	s.DirectoryId = &v
	return s
}

func (s *GetUserProvisioningRdAccountStatisticsRequest) SetRdMemberId(v string) *GetUserProvisioningRdAccountStatisticsRequest {
	s.RdMemberId = &v
	return s
}

func (s *GetUserProvisioningRdAccountStatisticsRequest) Validate() error {
	return dara.Validate(s)
}
