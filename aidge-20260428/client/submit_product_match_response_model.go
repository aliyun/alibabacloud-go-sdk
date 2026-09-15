// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitProductMatchResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SubmitProductMatchResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SubmitProductMatchResponse
	GetStatusCode() *int32
	SetBody(v *SubmitProductMatchResponseBody) *SubmitProductMatchResponse
	GetBody() *SubmitProductMatchResponseBody
}

type SubmitProductMatchResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SubmitProductMatchResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SubmitProductMatchResponse) String() string {
	return dara.Prettify(s)
}

func (s SubmitProductMatchResponse) GoString() string {
	return s.String()
}

func (s *SubmitProductMatchResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SubmitProductMatchResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SubmitProductMatchResponse) GetBody() *SubmitProductMatchResponseBody {
	return s.Body
}

func (s *SubmitProductMatchResponse) SetHeaders(v map[string]*string) *SubmitProductMatchResponse {
	s.Headers = v
	return s
}

func (s *SubmitProductMatchResponse) SetStatusCode(v int32) *SubmitProductMatchResponse {
	s.StatusCode = &v
	return s
}

func (s *SubmitProductMatchResponse) SetBody(v *SubmitProductMatchResponseBody) *SubmitProductMatchResponse {
	s.Body = v
	return s
}

func (s *SubmitProductMatchResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
