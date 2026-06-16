// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCategoryMatchResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CategoryMatchResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CategoryMatchResponse
	GetStatusCode() *int32
	SetBody(v *CategoryMatchResponseBody) *CategoryMatchResponse
	GetBody() *CategoryMatchResponseBody
}

type CategoryMatchResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CategoryMatchResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CategoryMatchResponse) String() string {
	return dara.Prettify(s)
}

func (s CategoryMatchResponse) GoString() string {
	return s.String()
}

func (s *CategoryMatchResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CategoryMatchResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CategoryMatchResponse) GetBody() *CategoryMatchResponseBody {
	return s.Body
}

func (s *CategoryMatchResponse) SetHeaders(v map[string]*string) *CategoryMatchResponse {
	s.Headers = v
	return s
}

func (s *CategoryMatchResponse) SetStatusCode(v int32) *CategoryMatchResponse {
	s.StatusCode = &v
	return s
}

func (s *CategoryMatchResponse) SetBody(v *CategoryMatchResponseBody) *CategoryMatchResponse {
	s.Body = v
	return s
}

func (s *CategoryMatchResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
