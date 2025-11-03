// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOpenPhysicalConnectionServiceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *OpenPhysicalConnectionServiceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *OpenPhysicalConnectionServiceResponse
	GetStatusCode() *int32
	SetBody(v *OpenPhysicalConnectionServiceResponseBody) *OpenPhysicalConnectionServiceResponse
	GetBody() *OpenPhysicalConnectionServiceResponseBody
}

type OpenPhysicalConnectionServiceResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *OpenPhysicalConnectionServiceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s OpenPhysicalConnectionServiceResponse) String() string {
	return dara.Prettify(s)
}

func (s OpenPhysicalConnectionServiceResponse) GoString() string {
	return s.String()
}

func (s *OpenPhysicalConnectionServiceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *OpenPhysicalConnectionServiceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *OpenPhysicalConnectionServiceResponse) GetBody() *OpenPhysicalConnectionServiceResponseBody {
	return s.Body
}

func (s *OpenPhysicalConnectionServiceResponse) SetHeaders(v map[string]*string) *OpenPhysicalConnectionServiceResponse {
	s.Headers = v
	return s
}

func (s *OpenPhysicalConnectionServiceResponse) SetStatusCode(v int32) *OpenPhysicalConnectionServiceResponse {
	s.StatusCode = &v
	return s
}

func (s *OpenPhysicalConnectionServiceResponse) SetBody(v *OpenPhysicalConnectionServiceResponseBody) *OpenPhysicalConnectionServiceResponse {
	s.Body = v
	return s
}

func (s *OpenPhysicalConnectionServiceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
