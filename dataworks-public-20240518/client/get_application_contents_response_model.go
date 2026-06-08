// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetApplicationContentsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetApplicationContentsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetApplicationContentsResponse
	GetStatusCode() *int32
	SetBody(v *GetApplicationContentsResponseBody) *GetApplicationContentsResponse
	GetBody() *GetApplicationContentsResponseBody
}

type GetApplicationContentsResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetApplicationContentsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetApplicationContentsResponse) String() string {
	return dara.Prettify(s)
}

func (s GetApplicationContentsResponse) GoString() string {
	return s.String()
}

func (s *GetApplicationContentsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetApplicationContentsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetApplicationContentsResponse) GetBody() *GetApplicationContentsResponseBody {
	return s.Body
}

func (s *GetApplicationContentsResponse) SetHeaders(v map[string]*string) *GetApplicationContentsResponse {
	s.Headers = v
	return s
}

func (s *GetApplicationContentsResponse) SetStatusCode(v int32) *GetApplicationContentsResponse {
	s.StatusCode = &v
	return s
}

func (s *GetApplicationContentsResponse) SetBody(v *GetApplicationContentsResponseBody) *GetApplicationContentsResponse {
	s.Body = v
	return s
}

func (s *GetApplicationContentsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
