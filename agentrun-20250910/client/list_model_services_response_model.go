// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListModelServicesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListModelServicesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListModelServicesResponse
	GetStatusCode() *int32
	SetBody(v *ListModelServicesResult) *ListModelServicesResponse
	GetBody() *ListModelServicesResult
}

type ListModelServicesResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListModelServicesResult `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListModelServicesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListModelServicesResponse) GoString() string {
	return s.String()
}

func (s *ListModelServicesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListModelServicesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListModelServicesResponse) GetBody() *ListModelServicesResult {
	return s.Body
}

func (s *ListModelServicesResponse) SetHeaders(v map[string]*string) *ListModelServicesResponse {
	s.Headers = v
	return s
}

func (s *ListModelServicesResponse) SetStatusCode(v int32) *ListModelServicesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListModelServicesResponse) SetBody(v *ListModelServicesResult) *ListModelServicesResponse {
	s.Body = v
	return s
}

func (s *ListModelServicesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
