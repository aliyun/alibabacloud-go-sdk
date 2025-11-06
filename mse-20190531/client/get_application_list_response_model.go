// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetApplicationListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetApplicationListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetApplicationListResponse
	GetStatusCode() *int32
	SetBody(v *GetApplicationListResponseBody) *GetApplicationListResponse
	GetBody() *GetApplicationListResponseBody
}

type GetApplicationListResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetApplicationListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetApplicationListResponse) String() string {
	return dara.Prettify(s)
}

func (s GetApplicationListResponse) GoString() string {
	return s.String()
}

func (s *GetApplicationListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetApplicationListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetApplicationListResponse) GetBody() *GetApplicationListResponseBody {
	return s.Body
}

func (s *GetApplicationListResponse) SetHeaders(v map[string]*string) *GetApplicationListResponse {
	s.Headers = v
	return s
}

func (s *GetApplicationListResponse) SetStatusCode(v int32) *GetApplicationListResponse {
	s.StatusCode = &v
	return s
}

func (s *GetApplicationListResponse) SetBody(v *GetApplicationListResponseBody) *GetApplicationListResponse {
	s.Body = v
	return s
}

func (s *GetApplicationListResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
