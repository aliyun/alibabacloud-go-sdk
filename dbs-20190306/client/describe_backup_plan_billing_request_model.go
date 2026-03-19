// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeBackupPlanBillingRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBackupPlanId(v string) *DescribeBackupPlanBillingRequest
	GetBackupPlanId() *string
	SetClientToken(v string) *DescribeBackupPlanBillingRequest
	GetClientToken() *string
	SetOwnerId(v string) *DescribeBackupPlanBillingRequest
	GetOwnerId() *string
	SetShowStorageType(v bool) *DescribeBackupPlanBillingRequest
	GetShowStorageType() *bool
}

type DescribeBackupPlanBillingRequest struct {
	// The ID of the backup plan. Call the [DescribeBackupPlanList](https://help.aliyun.com/document_detail/2869825.html) operation to obtain the value of this parameter.
	//
	// This parameter is required.
	//
	// example:
	//
	// dbstooi01ex****
	BackupPlanId *string `json:"BackupPlanId,omitempty" xml:"BackupPlanId,omitempty"`
	// A client token to ensure the idempotence of the request. The client generates this value. It must be unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length.
	//
	// example:
	//
	// ETnLKlblzczshOTUbOCzxxxxxxxxxx
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	OwnerId     *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// Specifies whether to display the storage class.
	//
	// example:
	//
	// true
	ShowStorageType *bool `json:"ShowStorageType,omitempty" xml:"ShowStorageType,omitempty"`
}

func (s DescribeBackupPlanBillingRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeBackupPlanBillingRequest) GoString() string {
	return s.String()
}

func (s *DescribeBackupPlanBillingRequest) GetBackupPlanId() *string {
	return s.BackupPlanId
}

func (s *DescribeBackupPlanBillingRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *DescribeBackupPlanBillingRequest) GetOwnerId() *string {
	return s.OwnerId
}

func (s *DescribeBackupPlanBillingRequest) GetShowStorageType() *bool {
	return s.ShowStorageType
}

func (s *DescribeBackupPlanBillingRequest) SetBackupPlanId(v string) *DescribeBackupPlanBillingRequest {
	s.BackupPlanId = &v
	return s
}

func (s *DescribeBackupPlanBillingRequest) SetClientToken(v string) *DescribeBackupPlanBillingRequest {
	s.ClientToken = &v
	return s
}

func (s *DescribeBackupPlanBillingRequest) SetOwnerId(v string) *DescribeBackupPlanBillingRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeBackupPlanBillingRequest) SetShowStorageType(v bool) *DescribeBackupPlanBillingRequest {
	s.ShowStorageType = &v
	return s
}

func (s *DescribeBackupPlanBillingRequest) Validate() error {
	return dara.Validate(s)
}
