// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetIntegratedServiceStatusResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetIntegratedServiceStatusResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetIntegratedServiceStatusResponse
	GetStatusCode() *int32
	SetBody(v *GetIntegratedServiceStatusResponseBody) *GetIntegratedServiceStatusResponse
	GetBody() *GetIntegratedServiceStatusResponseBody
}

type GetIntegratedServiceStatusResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetIntegratedServiceStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetIntegratedServiceStatusResponse) String() string {
	return dara.Prettify(s)
}

func (s GetIntegratedServiceStatusResponse) GoString() string {
	return s.String()
}

func (s *GetIntegratedServiceStatusResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetIntegratedServiceStatusResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetIntegratedServiceStatusResponse) GetBody() *GetIntegratedServiceStatusResponseBody {
	return s.Body
}

func (s *GetIntegratedServiceStatusResponse) SetHeaders(v map[string]*string) *GetIntegratedServiceStatusResponse {
	s.Headers = v
	return s
}

func (s *GetIntegratedServiceStatusResponse) SetStatusCode(v int32) *GetIntegratedServiceStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *GetIntegratedServiceStatusResponse) SetBody(v *GetIntegratedServiceStatusResponseBody) *GetIntegratedServiceStatusResponse {
	s.Body = v
	return s
}

func (s *GetIntegratedServiceStatusResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
