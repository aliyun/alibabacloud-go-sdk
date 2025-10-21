// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCatalogsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetCatalogsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetCatalogsResponse
	GetStatusCode() *int32
	SetBody(v *GetCatalogsResponseBody) *GetCatalogsResponse
	GetBody() *GetCatalogsResponseBody
}

type GetCatalogsResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetCatalogsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetCatalogsResponse) String() string {
	return dara.Prettify(s)
}

func (s GetCatalogsResponse) GoString() string {
	return s.String()
}

func (s *GetCatalogsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetCatalogsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetCatalogsResponse) GetBody() *GetCatalogsResponseBody {
	return s.Body
}

func (s *GetCatalogsResponse) SetHeaders(v map[string]*string) *GetCatalogsResponse {
	s.Headers = v
	return s
}

func (s *GetCatalogsResponse) SetStatusCode(v int32) *GetCatalogsResponse {
	s.StatusCode = &v
	return s
}

func (s *GetCatalogsResponse) SetBody(v *GetCatalogsResponseBody) *GetCatalogsResponse {
	s.Body = v
	return s
}

func (s *GetCatalogsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
