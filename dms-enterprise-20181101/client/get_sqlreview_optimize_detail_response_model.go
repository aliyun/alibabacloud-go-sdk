// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSQLReviewOptimizeDetailResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetSQLReviewOptimizeDetailResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetSQLReviewOptimizeDetailResponse
	GetStatusCode() *int32
	SetBody(v *GetSQLReviewOptimizeDetailResponseBody) *GetSQLReviewOptimizeDetailResponse
	GetBody() *GetSQLReviewOptimizeDetailResponseBody
}

type GetSQLReviewOptimizeDetailResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetSQLReviewOptimizeDetailResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetSQLReviewOptimizeDetailResponse) String() string {
	return dara.Prettify(s)
}

func (s GetSQLReviewOptimizeDetailResponse) GoString() string {
	return s.String()
}

func (s *GetSQLReviewOptimizeDetailResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetSQLReviewOptimizeDetailResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetSQLReviewOptimizeDetailResponse) GetBody() *GetSQLReviewOptimizeDetailResponseBody {
	return s.Body
}

func (s *GetSQLReviewOptimizeDetailResponse) SetHeaders(v map[string]*string) *GetSQLReviewOptimizeDetailResponse {
	s.Headers = v
	return s
}

func (s *GetSQLReviewOptimizeDetailResponse) SetStatusCode(v int32) *GetSQLReviewOptimizeDetailResponse {
	s.StatusCode = &v
	return s
}

func (s *GetSQLReviewOptimizeDetailResponse) SetBody(v *GetSQLReviewOptimizeDetailResponseBody) *GetSQLReviewOptimizeDetailResponse {
	s.Body = v
	return s
}

func (s *GetSQLReviewOptimizeDetailResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
