// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAccountReviewRecordResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetAccountReviewRecordResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetAccountReviewRecordResponse
	GetStatusCode() *int32
	SetBody(v *GetAccountReviewRecordResponseBody) *GetAccountReviewRecordResponse
	GetBody() *GetAccountReviewRecordResponseBody
}

type GetAccountReviewRecordResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAccountReviewRecordResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAccountReviewRecordResponse) String() string {
	return dara.Prettify(s)
}

func (s GetAccountReviewRecordResponse) GoString() string {
	return s.String()
}

func (s *GetAccountReviewRecordResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetAccountReviewRecordResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetAccountReviewRecordResponse) GetBody() *GetAccountReviewRecordResponseBody {
	return s.Body
}

func (s *GetAccountReviewRecordResponse) SetHeaders(v map[string]*string) *GetAccountReviewRecordResponse {
	s.Headers = v
	return s
}

func (s *GetAccountReviewRecordResponse) SetStatusCode(v int32) *GetAccountReviewRecordResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAccountReviewRecordResponse) SetBody(v *GetAccountReviewRecordResponseBody) *GetAccountReviewRecordResponse {
	s.Body = v
	return s
}

func (s *GetAccountReviewRecordResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
