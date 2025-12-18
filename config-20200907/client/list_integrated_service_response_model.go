// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListIntegratedServiceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListIntegratedServiceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListIntegratedServiceResponse
	GetStatusCode() *int32
	SetBody(v *ListIntegratedServiceResponseBody) *ListIntegratedServiceResponse
	GetBody() *ListIntegratedServiceResponseBody
}

type ListIntegratedServiceResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListIntegratedServiceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListIntegratedServiceResponse) String() string {
	return dara.Prettify(s)
}

func (s ListIntegratedServiceResponse) GoString() string {
	return s.String()
}

func (s *ListIntegratedServiceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListIntegratedServiceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListIntegratedServiceResponse) GetBody() *ListIntegratedServiceResponseBody {
	return s.Body
}

func (s *ListIntegratedServiceResponse) SetHeaders(v map[string]*string) *ListIntegratedServiceResponse {
	s.Headers = v
	return s
}

func (s *ListIntegratedServiceResponse) SetStatusCode(v int32) *ListIntegratedServiceResponse {
	s.StatusCode = &v
	return s
}

func (s *ListIntegratedServiceResponse) SetBody(v *ListIntegratedServiceResponseBody) *ListIntegratedServiceResponse {
	s.Body = v
	return s
}

func (s *ListIntegratedServiceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
