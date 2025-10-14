// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetPageResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetPageResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetPageResponse
	GetStatusCode() *int32
	SetBody(v *GetPageResponseBody) *GetPageResponse
	GetBody() *GetPageResponseBody
}

type GetPageResponse struct {
	Headers    map[string]*string   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetPageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetPageResponse) String() string {
	return dara.Prettify(s)
}

func (s GetPageResponse) GoString() string {
	return s.String()
}

func (s *GetPageResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetPageResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetPageResponse) GetBody() *GetPageResponseBody {
	return s.Body
}

func (s *GetPageResponse) SetHeaders(v map[string]*string) *GetPageResponse {
	s.Headers = v
	return s
}

func (s *GetPageResponse) SetStatusCode(v int32) *GetPageResponse {
	s.StatusCode = &v
	return s
}

func (s *GetPageResponse) SetBody(v *GetPageResponseBody) *GetPageResponse {
	s.Body = v
	return s
}

func (s *GetPageResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
