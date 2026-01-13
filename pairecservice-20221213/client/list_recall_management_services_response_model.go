// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListRecallManagementServicesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListRecallManagementServicesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListRecallManagementServicesResponse
	GetStatusCode() *int32
	SetBody(v *ListRecallManagementServicesResponseBody) *ListRecallManagementServicesResponse
	GetBody() *ListRecallManagementServicesResponseBody
}

type ListRecallManagementServicesResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListRecallManagementServicesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListRecallManagementServicesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListRecallManagementServicesResponse) GoString() string {
	return s.String()
}

func (s *ListRecallManagementServicesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListRecallManagementServicesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListRecallManagementServicesResponse) GetBody() *ListRecallManagementServicesResponseBody {
	return s.Body
}

func (s *ListRecallManagementServicesResponse) SetHeaders(v map[string]*string) *ListRecallManagementServicesResponse {
	s.Headers = v
	return s
}

func (s *ListRecallManagementServicesResponse) SetStatusCode(v int32) *ListRecallManagementServicesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListRecallManagementServicesResponse) SetBody(v *ListRecallManagementServicesResponseBody) *ListRecallManagementServicesResponse {
	s.Body = v
	return s
}

func (s *ListRecallManagementServicesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
