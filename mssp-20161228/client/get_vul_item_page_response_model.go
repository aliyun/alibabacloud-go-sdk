// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetVulItemPageResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetVulItemPageResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetVulItemPageResponse
	GetStatusCode() *int32
	SetBody(v *GetVulItemPageResponseBody) *GetVulItemPageResponse
	GetBody() *GetVulItemPageResponseBody
}

type GetVulItemPageResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetVulItemPageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetVulItemPageResponse) String() string {
	return dara.Prettify(s)
}

func (s GetVulItemPageResponse) GoString() string {
	return s.String()
}

func (s *GetVulItemPageResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetVulItemPageResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetVulItemPageResponse) GetBody() *GetVulItemPageResponseBody {
	return s.Body
}

func (s *GetVulItemPageResponse) SetHeaders(v map[string]*string) *GetVulItemPageResponse {
	s.Headers = v
	return s
}

func (s *GetVulItemPageResponse) SetStatusCode(v int32) *GetVulItemPageResponse {
	s.StatusCode = &v
	return s
}

func (s *GetVulItemPageResponse) SetBody(v *GetVulItemPageResponseBody) *GetVulItemPageResponse {
	s.Body = v
	return s
}

func (s *GetVulItemPageResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
