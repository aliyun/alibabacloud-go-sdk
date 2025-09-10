// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateQuotaApplicationsForTemplateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAliyunUids(v []*string) *CreateQuotaApplicationsForTemplateResponseBody
	GetAliyunUids() []*string
	SetBatchQuotaApplicationId(v string) *CreateQuotaApplicationsForTemplateResponseBody
	GetBatchQuotaApplicationId() *string
	SetFailResults(v []*CreateQuotaApplicationsForTemplateResponseBodyFailResults) *CreateQuotaApplicationsForTemplateResponseBody
	GetFailResults() []*CreateQuotaApplicationsForTemplateResponseBodyFailResults
	SetRequestId(v string) *CreateQuotaApplicationsForTemplateResponseBody
	GetRequestId() *string
}

type CreateQuotaApplicationsForTemplateResponseBody struct {
	// The Alibaba Cloud accounts for which the quotas are applied.
	AliyunUids []*string `json:"AliyunUids,omitempty" xml:"AliyunUids,omitempty" type:"Repeated"`
	// The ID of the quota application batch.
	//
	// example:
	//
	// d314d6ae-867d-484c-9009-3d421a80****
	BatchQuotaApplicationId *string `json:"BatchQuotaApplicationId,omitempty" xml:"BatchQuotaApplicationId,omitempty"`
	// The Alibaba Cloud accounts of the members in a resource directory whose quota increase request is rejected, and the reason for the rejection.
	FailResults []*CreateQuotaApplicationsForTemplateResponseBodyFailResults `json:"FailResults,omitempty" xml:"FailResults,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 8FF8CAF0-29D9-4F11-B6A4-FD2CBCA016D3
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateQuotaApplicationsForTemplateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateQuotaApplicationsForTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *CreateQuotaApplicationsForTemplateResponseBody) GetAliyunUids() []*string {
	return s.AliyunUids
}

func (s *CreateQuotaApplicationsForTemplateResponseBody) GetBatchQuotaApplicationId() *string {
	return s.BatchQuotaApplicationId
}

func (s *CreateQuotaApplicationsForTemplateResponseBody) GetFailResults() []*CreateQuotaApplicationsForTemplateResponseBodyFailResults {
	return s.FailResults
}

func (s *CreateQuotaApplicationsForTemplateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateQuotaApplicationsForTemplateResponseBody) SetAliyunUids(v []*string) *CreateQuotaApplicationsForTemplateResponseBody {
	s.AliyunUids = v
	return s
}

func (s *CreateQuotaApplicationsForTemplateResponseBody) SetBatchQuotaApplicationId(v string) *CreateQuotaApplicationsForTemplateResponseBody {
	s.BatchQuotaApplicationId = &v
	return s
}

func (s *CreateQuotaApplicationsForTemplateResponseBody) SetFailResults(v []*CreateQuotaApplicationsForTemplateResponseBodyFailResults) *CreateQuotaApplicationsForTemplateResponseBody {
	s.FailResults = v
	return s
}

func (s *CreateQuotaApplicationsForTemplateResponseBody) SetRequestId(v string) *CreateQuotaApplicationsForTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateQuotaApplicationsForTemplateResponseBody) Validate() error {
	return dara.Validate(s)
}

type CreateQuotaApplicationsForTemplateResponseBodyFailResults struct {
	// The Alibaba Cloud account of the members in a resource directory whose quota increase request is rejected.
	//
	// example:
	//
	// 135048337611****
	AliyunUid *string `json:"AliyunUid,omitempty" xml:"AliyunUid,omitempty"`
	// The reason for the rejection.
	//
	// example:
	//
	// The quota adjustment application is being processed. Please try again later.
	Reason *string `json:"Reason,omitempty" xml:"Reason,omitempty"`
}

func (s CreateQuotaApplicationsForTemplateResponseBodyFailResults) String() string {
	return dara.Prettify(s)
}

func (s CreateQuotaApplicationsForTemplateResponseBodyFailResults) GoString() string {
	return s.String()
}

func (s *CreateQuotaApplicationsForTemplateResponseBodyFailResults) GetAliyunUid() *string {
	return s.AliyunUid
}

func (s *CreateQuotaApplicationsForTemplateResponseBodyFailResults) GetReason() *string {
	return s.Reason
}

func (s *CreateQuotaApplicationsForTemplateResponseBodyFailResults) SetAliyunUid(v string) *CreateQuotaApplicationsForTemplateResponseBodyFailResults {
	s.AliyunUid = &v
	return s
}

func (s *CreateQuotaApplicationsForTemplateResponseBodyFailResults) SetReason(v string) *CreateQuotaApplicationsForTemplateResponseBodyFailResults {
	s.Reason = &v
	return s
}

func (s *CreateQuotaApplicationsForTemplateResponseBodyFailResults) Validate() error {
	return dara.Validate(s)
}
