// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetRangeScoreResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetRangeScoreResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetRangeScoreResponse
	GetStatusCode() *int32
	SetBody(v *GetRangeScoreResponseBody) *GetRangeScoreResponse
	GetBody() *GetRangeScoreResponseBody
}

type GetRangeScoreResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetRangeScoreResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetRangeScoreResponse) String() string {
	return dara.Prettify(s)
}

func (s GetRangeScoreResponse) GoString() string {
	return s.String()
}

func (s *GetRangeScoreResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetRangeScoreResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetRangeScoreResponse) GetBody() *GetRangeScoreResponseBody {
	return s.Body
}

func (s *GetRangeScoreResponse) SetHeaders(v map[string]*string) *GetRangeScoreResponse {
	s.Headers = v
	return s
}

func (s *GetRangeScoreResponse) SetStatusCode(v int32) *GetRangeScoreResponse {
	s.StatusCode = &v
	return s
}

func (s *GetRangeScoreResponse) SetBody(v *GetRangeScoreResponseBody) *GetRangeScoreResponse {
	s.Body = v
	return s
}

func (s *GetRangeScoreResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
