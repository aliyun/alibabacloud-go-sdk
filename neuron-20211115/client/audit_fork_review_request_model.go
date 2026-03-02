// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAuditForkReviewRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v *ForkReviewAuditCmd) *AuditForkReviewRequest
	GetBody() *ForkReviewAuditCmd
}

type AuditForkReviewRequest struct {
	Body *ForkReviewAuditCmd `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AuditForkReviewRequest) String() string {
	return dara.Prettify(s)
}

func (s AuditForkReviewRequest) GoString() string {
	return s.String()
}

func (s *AuditForkReviewRequest) GetBody() *ForkReviewAuditCmd {
	return s.Body
}

func (s *AuditForkReviewRequest) SetBody(v *ForkReviewAuditCmd) *AuditForkReviewRequest {
	s.Body = v
	return s
}

func (s *AuditForkReviewRequest) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
