// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetScoreInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetScoreInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetScoreInfoResponse
	GetStatusCode() *int32
	SetBody(v *GetScoreInfoResponseBody) *GetScoreInfoResponse
	GetBody() *GetScoreInfoResponseBody
}

type GetScoreInfoResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetScoreInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetScoreInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s GetScoreInfoResponse) GoString() string {
	return s.String()
}

func (s *GetScoreInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetScoreInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetScoreInfoResponse) GetBody() *GetScoreInfoResponseBody {
	return s.Body
}

func (s *GetScoreInfoResponse) SetHeaders(v map[string]*string) *GetScoreInfoResponse {
	s.Headers = v
	return s
}

func (s *GetScoreInfoResponse) SetStatusCode(v int32) *GetScoreInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *GetScoreInfoResponse) SetBody(v *GetScoreInfoResponseBody) *GetScoreInfoResponse {
	s.Body = v
	return s
}

func (s *GetScoreInfoResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
