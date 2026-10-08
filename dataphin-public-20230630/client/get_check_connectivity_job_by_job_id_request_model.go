// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCheckConnectivityJobByJobIdRequest interface {
	dara.Model
	String() string
	GoString() string
	SetJobId(v string) *GetCheckConnectivityJobByJobIdRequest
	GetJobId() *string
	SetOpTenantId(v int64) *GetCheckConnectivityJobByJobIdRequest
	GetOpTenantId() *int64
	SetOpUserId(v string) *GetCheckConnectivityJobByJobIdRequest
	GetOpUserId() *string
}

type GetCheckConnectivityJobByJobIdRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 129837xxxx
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
	// example:
	//
	// 30001011
	OpUserId *string `json:"OpUserId,omitempty" xml:"OpUserId,omitempty"`
}

func (s GetCheckConnectivityJobByJobIdRequest) String() string {
	return dara.Prettify(s)
}

func (s GetCheckConnectivityJobByJobIdRequest) GoString() string {
	return s.String()
}

func (s *GetCheckConnectivityJobByJobIdRequest) GetJobId() *string {
	return s.JobId
}

func (s *GetCheckConnectivityJobByJobIdRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *GetCheckConnectivityJobByJobIdRequest) GetOpUserId() *string {
	return s.OpUserId
}

func (s *GetCheckConnectivityJobByJobIdRequest) SetJobId(v string) *GetCheckConnectivityJobByJobIdRequest {
	s.JobId = &v
	return s
}

func (s *GetCheckConnectivityJobByJobIdRequest) SetOpTenantId(v int64) *GetCheckConnectivityJobByJobIdRequest {
	s.OpTenantId = &v
	return s
}

func (s *GetCheckConnectivityJobByJobIdRequest) SetOpUserId(v string) *GetCheckConnectivityJobByJobIdRequest {
	s.OpUserId = &v
	return s
}

func (s *GetCheckConnectivityJobByJobIdRequest) Validate() error {
	return dara.Validate(s)
}
