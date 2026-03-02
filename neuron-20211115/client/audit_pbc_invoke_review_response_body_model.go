// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAuditPbcInvokeReviewResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *AuditPbcInvokeReviewResponseBody
	GetRequestId() *string
	SetResult(v *CatalogCommonResult) *AuditPbcInvokeReviewResponseBody
	GetResult() *CatalogCommonResult
}

type AuditPbcInvokeReviewResponseBody struct {
	RequestId *string              `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Result    *CatalogCommonResult `json:"result,omitempty" xml:"result,omitempty"`
}

func (s AuditPbcInvokeReviewResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AuditPbcInvokeReviewResponseBody) GoString() string {
	return s.String()
}

func (s *AuditPbcInvokeReviewResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AuditPbcInvokeReviewResponseBody) GetResult() *CatalogCommonResult {
	return s.Result
}

func (s *AuditPbcInvokeReviewResponseBody) SetRequestId(v string) *AuditPbcInvokeReviewResponseBody {
	s.RequestId = &v
	return s
}

func (s *AuditPbcInvokeReviewResponseBody) SetResult(v *CatalogCommonResult) *AuditPbcInvokeReviewResponseBody {
	s.Result = v
	return s
}

func (s *AuditPbcInvokeReviewResponseBody) Validate() error {
	if s.Result != nil {
		if err := s.Result.Validate(); err != nil {
			return err
		}
	}
	return nil
}
