// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetApplicationInstanceListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetApplicationInstanceListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetApplicationInstanceListResponse
	GetStatusCode() *int32
	SetBody(v *GetApplicationInstanceListResponseBody) *GetApplicationInstanceListResponse
	GetBody() *GetApplicationInstanceListResponseBody
}

type GetApplicationInstanceListResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetApplicationInstanceListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetApplicationInstanceListResponse) String() string {
	return dara.Prettify(s)
}

func (s GetApplicationInstanceListResponse) GoString() string {
	return s.String()
}

func (s *GetApplicationInstanceListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetApplicationInstanceListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetApplicationInstanceListResponse) GetBody() *GetApplicationInstanceListResponseBody {
	return s.Body
}

func (s *GetApplicationInstanceListResponse) SetHeaders(v map[string]*string) *GetApplicationInstanceListResponse {
	s.Headers = v
	return s
}

func (s *GetApplicationInstanceListResponse) SetStatusCode(v int32) *GetApplicationInstanceListResponse {
	s.StatusCode = &v
	return s
}

func (s *GetApplicationInstanceListResponse) SetBody(v *GetApplicationInstanceListResponseBody) *GetApplicationInstanceListResponse {
	s.Body = v
	return s
}

func (s *GetApplicationInstanceListResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
