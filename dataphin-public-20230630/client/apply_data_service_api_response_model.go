// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iApplyDataServiceApiResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ApplyDataServiceApiResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ApplyDataServiceApiResponse
	GetStatusCode() *int32
	SetBody(v *ApplyDataServiceApiResponseBody) *ApplyDataServiceApiResponse
	GetBody() *ApplyDataServiceApiResponseBody
}

type ApplyDataServiceApiResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ApplyDataServiceApiResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ApplyDataServiceApiResponse) String() string {
	return dara.Prettify(s)
}

func (s ApplyDataServiceApiResponse) GoString() string {
	return s.String()
}

func (s *ApplyDataServiceApiResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ApplyDataServiceApiResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ApplyDataServiceApiResponse) GetBody() *ApplyDataServiceApiResponseBody {
	return s.Body
}

func (s *ApplyDataServiceApiResponse) SetHeaders(v map[string]*string) *ApplyDataServiceApiResponse {
	s.Headers = v
	return s
}

func (s *ApplyDataServiceApiResponse) SetStatusCode(v int32) *ApplyDataServiceApiResponse {
	s.StatusCode = &v
	return s
}

func (s *ApplyDataServiceApiResponse) SetBody(v *ApplyDataServiceApiResponseBody) *ApplyDataServiceApiResponse {
	s.Body = v
	return s
}

func (s *ApplyDataServiceApiResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
