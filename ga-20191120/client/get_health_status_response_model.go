// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetHealthStatusResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetHealthStatusResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetHealthStatusResponse
	GetStatusCode() *int32
	SetBody(v *GetHealthStatusResponseBody) *GetHealthStatusResponse
	GetBody() *GetHealthStatusResponseBody
}

type GetHealthStatusResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetHealthStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetHealthStatusResponse) String() string {
	return dara.Prettify(s)
}

func (s GetHealthStatusResponse) GoString() string {
	return s.String()
}

func (s *GetHealthStatusResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetHealthStatusResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetHealthStatusResponse) GetBody() *GetHealthStatusResponseBody {
	return s.Body
}

func (s *GetHealthStatusResponse) SetHeaders(v map[string]*string) *GetHealthStatusResponse {
	s.Headers = v
	return s
}

func (s *GetHealthStatusResponse) SetStatusCode(v int32) *GetHealthStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *GetHealthStatusResponse) SetBody(v *GetHealthStatusResponseBody) *GetHealthStatusResponse {
	s.Body = v
	return s
}

func (s *GetHealthStatusResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
