// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDLAServiceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBackupPlanId(v string) *DescribeDLAServiceRequest
	GetBackupPlanId() *string
	SetClientToken(v string) *DescribeDLAServiceRequest
	GetClientToken() *string
	SetOwnerId(v string) *DescribeDLAServiceRequest
	GetOwnerId() *string
}

type DescribeDLAServiceRequest struct {
	// The ID of the backup schedule. You can call the [DescribeBackupPlanList](https://help.aliyun.com/document_detail/2869825.html) operation to obtain the ID of the backup schedule.
	//
	// This parameter is required.
	//
	// example:
	//
	// dbstooi01ex****
	BackupPlanId *string `json:"BackupPlanId,omitempty" xml:"BackupPlanId,omitempty"`
	// The client token that is used to ensure the idempotence of the request.
	//
	// example:
	//
	// DBS
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	OwnerId     *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
}

func (s DescribeDLAServiceRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDLAServiceRequest) GoString() string {
	return s.String()
}

func (s *DescribeDLAServiceRequest) GetBackupPlanId() *string {
	return s.BackupPlanId
}

func (s *DescribeDLAServiceRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *DescribeDLAServiceRequest) GetOwnerId() *string {
	return s.OwnerId
}

func (s *DescribeDLAServiceRequest) SetBackupPlanId(v string) *DescribeDLAServiceRequest {
	s.BackupPlanId = &v
	return s
}

func (s *DescribeDLAServiceRequest) SetClientToken(v string) *DescribeDLAServiceRequest {
	s.ClientToken = &v
	return s
}

func (s *DescribeDLAServiceRequest) SetOwnerId(v string) *DescribeDLAServiceRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeDLAServiceRequest) Validate() error {
	return dara.Validate(s)
}
