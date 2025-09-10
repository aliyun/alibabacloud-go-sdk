// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetQuotaApplicationApprovalRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApplicationId(v string) *GetQuotaApplicationApprovalRequest
	GetApplicationId() *string
}

type GetQuotaApplicationApprovalRequest struct {
	// The quota application ID.
	//
	// For more information about how to obtain the ID of a quota application, see [ListQuotaApplications](https://help.aliyun.com/document_detail/440568.html).
	//
	// example:
	//
	// d314d6ae-867d-484c-9009-3d42****
	ApplicationId *string `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
}

func (s GetQuotaApplicationApprovalRequest) String() string {
	return dara.Prettify(s)
}

func (s GetQuotaApplicationApprovalRequest) GoString() string {
	return s.String()
}

func (s *GetQuotaApplicationApprovalRequest) GetApplicationId() *string {
	return s.ApplicationId
}

func (s *GetQuotaApplicationApprovalRequest) SetApplicationId(v string) *GetQuotaApplicationApprovalRequest {
	s.ApplicationId = &v
	return s
}

func (s *GetQuotaApplicationApprovalRequest) Validate() error {
	return dara.Validate(s)
}
