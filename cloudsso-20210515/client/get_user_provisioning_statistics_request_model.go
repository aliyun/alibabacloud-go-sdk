// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetUserProvisioningStatisticsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDirectoryId(v string) *GetUserProvisioningStatisticsRequest
	GetDirectoryId() *string
	SetUserProvisioningId(v string) *GetUserProvisioningStatisticsRequest
	GetUserProvisioningId() *string
}

type GetUserProvisioningStatisticsRequest struct {
	// The ID of the resource directory.
	//
	// example:
	//
	// d-003qew84****
	DirectoryId *string `json:"DirectoryId,omitempty" xml:"DirectoryId,omitempty"`
	// The ID of the RAM user provisioning.
	//
	// example:
	//
	// up-002axzhapcbz6e63****
	UserProvisioningId *string `json:"UserProvisioningId,omitempty" xml:"UserProvisioningId,omitempty"`
}

func (s GetUserProvisioningStatisticsRequest) String() string {
	return dara.Prettify(s)
}

func (s GetUserProvisioningStatisticsRequest) GoString() string {
	return s.String()
}

func (s *GetUserProvisioningStatisticsRequest) GetDirectoryId() *string {
	return s.DirectoryId
}

func (s *GetUserProvisioningStatisticsRequest) GetUserProvisioningId() *string {
	return s.UserProvisioningId
}

func (s *GetUserProvisioningStatisticsRequest) SetDirectoryId(v string) *GetUserProvisioningStatisticsRequest {
	s.DirectoryId = &v
	return s
}

func (s *GetUserProvisioningStatisticsRequest) SetUserProvisioningId(v string) *GetUserProvisioningStatisticsRequest {
	s.UserProvisioningId = &v
	return s
}

func (s *GetUserProvisioningStatisticsRequest) Validate() error {
	return dara.Validate(s)
}
