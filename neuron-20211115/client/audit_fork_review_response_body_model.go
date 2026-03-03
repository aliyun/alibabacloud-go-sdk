// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAuditForkReviewResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *AuditForkReviewResponseBody
	GetRequestId() *string
	SetResult(v string) *AuditForkReviewResponseBody
	GetResult() *string
}

type AuditForkReviewResponseBody struct {
	// example:
	//
	// fsahkfkjsjfsdjlfalsf
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Result *string `json:"result,omitempty" xml:"result,omitempty"`
}

func (s AuditForkReviewResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AuditForkReviewResponseBody) GoString() string {
	return s.String()
}

func (s *AuditForkReviewResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AuditForkReviewResponseBody) GetResult() *string {
	return s.Result
}

func (s *AuditForkReviewResponseBody) SetRequestId(v string) *AuditForkReviewResponseBody {
	s.RequestId = &v
	return s
}

func (s *AuditForkReviewResponseBody) SetResult(v string) *AuditForkReviewResponseBody {
	s.Result = &v
	return s
}

func (s *AuditForkReviewResponseBody) Validate() error {
	return dara.Validate(s)
}
