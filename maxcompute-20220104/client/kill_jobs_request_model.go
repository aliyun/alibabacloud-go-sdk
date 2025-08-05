// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iKillJobsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v string) *KillJobsRequest
	GetBody() *string
	SetRegion(v string) *KillJobsRequest
	GetRegion() *string
	SetTenantId(v string) *KillJobsRequest
	GetTenantId() *string
}

type KillJobsRequest struct {
	// The request body parameters.
	//
	// example:
	//
	// [
	//
	//       {
	//
	//             "instanceId": "",
	//
	//             "projectName": ""
	//
	//       },
	//
	//       {
	//
	//             "instanceId": "",
	//
	//             "projectName": ""
	//
	//       }
	//
	// ]
	Body *string `json:"body,omitempty" xml:"body,omitempty"`
	// The ID of the region in which the instance resides.
	//
	// example:
	//
	// cn-hangzhou
	Region *string `json:"region,omitempty" xml:"region,omitempty"`
	// The ID of the tenant.
	//
	// example:
	//
	// 478403690625249
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s KillJobsRequest) String() string {
	return dara.Prettify(s)
}

func (s KillJobsRequest) GoString() string {
	return s.String()
}

func (s *KillJobsRequest) GetBody() *string {
	return s.Body
}

func (s *KillJobsRequest) GetRegion() *string {
	return s.Region
}

func (s *KillJobsRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *KillJobsRequest) SetBody(v string) *KillJobsRequest {
	s.Body = &v
	return s
}

func (s *KillJobsRequest) SetRegion(v string) *KillJobsRequest {
	s.Region = &v
	return s
}

func (s *KillJobsRequest) SetTenantId(v string) *KillJobsRequest {
	s.TenantId = &v
	return s
}

func (s *KillJobsRequest) Validate() error {
	return dara.Validate(s)
}
