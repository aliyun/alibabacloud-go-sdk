// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemindQuotaApplicationApprovalRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApplicationId(v string) *RemindQuotaApplicationApprovalRequest
	GetApplicationId() *string
}

type RemindQuotaApplicationApprovalRequest struct {
	// The quota application ID.
	//
	// For more information about how to obtain the ID of a quota application, see [ListQuotaApplications](https://help.aliyun.com/document_detail/440568.html).
	//
	// example:
	//
	// 219f1ff6-6205-495f-bda7-90d2fe945e****
	ApplicationId *string `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
}

func (s RemindQuotaApplicationApprovalRequest) String() string {
	return dara.Prettify(s)
}

func (s RemindQuotaApplicationApprovalRequest) GoString() string {
	return s.String()
}

func (s *RemindQuotaApplicationApprovalRequest) GetApplicationId() *string {
	return s.ApplicationId
}

func (s *RemindQuotaApplicationApprovalRequest) SetApplicationId(v string) *RemindQuotaApplicationApprovalRequest {
	s.ApplicationId = &v
	return s
}

func (s *RemindQuotaApplicationApprovalRequest) Validate() error {
	return dara.Validate(s)
}
