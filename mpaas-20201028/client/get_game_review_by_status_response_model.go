// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetGameReviewByStatusResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetGameReviewByStatusResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetGameReviewByStatusResponse
	GetStatusCode() *int32
	SetBody(v *GetGameReviewByStatusResponseBody) *GetGameReviewByStatusResponse
	GetBody() *GetGameReviewByStatusResponseBody
}

type GetGameReviewByStatusResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetGameReviewByStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetGameReviewByStatusResponse) String() string {
	return dara.Prettify(s)
}

func (s GetGameReviewByStatusResponse) GoString() string {
	return s.String()
}

func (s *GetGameReviewByStatusResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetGameReviewByStatusResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetGameReviewByStatusResponse) GetBody() *GetGameReviewByStatusResponseBody {
	return s.Body
}

func (s *GetGameReviewByStatusResponse) SetHeaders(v map[string]*string) *GetGameReviewByStatusResponse {
	s.Headers = v
	return s
}

func (s *GetGameReviewByStatusResponse) SetStatusCode(v int32) *GetGameReviewByStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *GetGameReviewByStatusResponse) SetBody(v *GetGameReviewByStatusResponseBody) *GetGameReviewByStatusResponse {
	s.Body = v
	return s
}

func (s *GetGameReviewByStatusResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
