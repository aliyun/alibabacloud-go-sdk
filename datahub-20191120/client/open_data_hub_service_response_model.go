// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOpenDataHubServiceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *OpenDataHubServiceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *OpenDataHubServiceResponse
	GetStatusCode() *int32
	SetBody(v *OpenDataHubServiceResponseBody) *OpenDataHubServiceResponse
	GetBody() *OpenDataHubServiceResponseBody
}

type OpenDataHubServiceResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *OpenDataHubServiceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s OpenDataHubServiceResponse) String() string {
	return dara.Prettify(s)
}

func (s OpenDataHubServiceResponse) GoString() string {
	return s.String()
}

func (s *OpenDataHubServiceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *OpenDataHubServiceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *OpenDataHubServiceResponse) GetBody() *OpenDataHubServiceResponseBody {
	return s.Body
}

func (s *OpenDataHubServiceResponse) SetHeaders(v map[string]*string) *OpenDataHubServiceResponse {
	s.Headers = v
	return s
}

func (s *OpenDataHubServiceResponse) SetStatusCode(v int32) *OpenDataHubServiceResponse {
	s.StatusCode = &v
	return s
}

func (s *OpenDataHubServiceResponse) SetBody(v *OpenDataHubServiceResponseBody) *OpenDataHubServiceResponse {
	s.Body = v
	return s
}

func (s *OpenDataHubServiceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
