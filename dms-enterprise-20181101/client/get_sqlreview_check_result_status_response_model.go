// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSQLReviewCheckResultStatusResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetSQLReviewCheckResultStatusResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetSQLReviewCheckResultStatusResponse
	GetStatusCode() *int32
	SetBody(v *GetSQLReviewCheckResultStatusResponseBody) *GetSQLReviewCheckResultStatusResponse
	GetBody() *GetSQLReviewCheckResultStatusResponseBody
}

type GetSQLReviewCheckResultStatusResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetSQLReviewCheckResultStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetSQLReviewCheckResultStatusResponse) String() string {
	return dara.Prettify(s)
}

func (s GetSQLReviewCheckResultStatusResponse) GoString() string {
	return s.String()
}

func (s *GetSQLReviewCheckResultStatusResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetSQLReviewCheckResultStatusResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetSQLReviewCheckResultStatusResponse) GetBody() *GetSQLReviewCheckResultStatusResponseBody {
	return s.Body
}

func (s *GetSQLReviewCheckResultStatusResponse) SetHeaders(v map[string]*string) *GetSQLReviewCheckResultStatusResponse {
	s.Headers = v
	return s
}

func (s *GetSQLReviewCheckResultStatusResponse) SetStatusCode(v int32) *GetSQLReviewCheckResultStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *GetSQLReviewCheckResultStatusResponse) SetBody(v *GetSQLReviewCheckResultStatusResponseBody) *GetSQLReviewCheckResultStatusResponse {
	s.Body = v
	return s
}

func (s *GetSQLReviewCheckResultStatusResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
