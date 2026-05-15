// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetIndexCurrentValueResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetIndexCurrentValueResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetIndexCurrentValueResponse
	GetStatusCode() *int32
	SetBody(v *GetIndexCurrentValueResponseBody) *GetIndexCurrentValueResponse
	GetBody() *GetIndexCurrentValueResponseBody
}

type GetIndexCurrentValueResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetIndexCurrentValueResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetIndexCurrentValueResponse) String() string {
	return dara.Prettify(s)
}

func (s GetIndexCurrentValueResponse) GoString() string {
	return s.String()
}

func (s *GetIndexCurrentValueResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetIndexCurrentValueResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetIndexCurrentValueResponse) GetBody() *GetIndexCurrentValueResponseBody {
	return s.Body
}

func (s *GetIndexCurrentValueResponse) SetHeaders(v map[string]*string) *GetIndexCurrentValueResponse {
	s.Headers = v
	return s
}

func (s *GetIndexCurrentValueResponse) SetStatusCode(v int32) *GetIndexCurrentValueResponse {
	s.StatusCode = &v
	return s
}

func (s *GetIndexCurrentValueResponse) SetBody(v *GetIndexCurrentValueResponseBody) *GetIndexCurrentValueResponse {
	s.Body = v
	return s
}

func (s *GetIndexCurrentValueResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
