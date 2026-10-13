// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetPortalResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetPortalResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetPortalResponse
	GetStatusCode() *int32
	SetBody(v *GetPortalResponseBody) *GetPortalResponse
	GetBody() *GetPortalResponseBody
}

type GetPortalResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetPortalResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetPortalResponse) String() string {
	return dara.Prettify(s)
}

func (s GetPortalResponse) GoString() string {
	return s.String()
}

func (s *GetPortalResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetPortalResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetPortalResponse) GetBody() *GetPortalResponseBody {
	return s.Body
}

func (s *GetPortalResponse) SetHeaders(v map[string]*string) *GetPortalResponse {
	s.Headers = v
	return s
}

func (s *GetPortalResponse) SetStatusCode(v int32) *GetPortalResponse {
	s.StatusCode = &v
	return s
}

func (s *GetPortalResponse) SetBody(v *GetPortalResponseBody) *GetPortalResponse {
	s.Body = v
	return s
}

func (s *GetPortalResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
