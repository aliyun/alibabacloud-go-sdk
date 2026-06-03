// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetOpenSubProductStatusResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetOpenSubProductStatusResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetOpenSubProductStatusResponse
	GetStatusCode() *int32
	SetBody(v *GetOpenSubProductStatusResponseBody) *GetOpenSubProductStatusResponse
	GetBody() *GetOpenSubProductStatusResponseBody
}

type GetOpenSubProductStatusResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetOpenSubProductStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetOpenSubProductStatusResponse) String() string {
	return dara.Prettify(s)
}

func (s GetOpenSubProductStatusResponse) GoString() string {
	return s.String()
}

func (s *GetOpenSubProductStatusResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetOpenSubProductStatusResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetOpenSubProductStatusResponse) GetBody() *GetOpenSubProductStatusResponseBody {
	return s.Body
}

func (s *GetOpenSubProductStatusResponse) SetHeaders(v map[string]*string) *GetOpenSubProductStatusResponse {
	s.Headers = v
	return s
}

func (s *GetOpenSubProductStatusResponse) SetStatusCode(v int32) *GetOpenSubProductStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *GetOpenSubProductStatusResponse) SetBody(v *GetOpenSubProductStatusResponseBody) *GetOpenSubProductStatusResponse {
	s.Body = v
	return s
}

func (s *GetOpenSubProductStatusResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
