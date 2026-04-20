// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCallCenterProvidersResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListCallCenterProvidersResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListCallCenterProvidersResponse
	GetStatusCode() *int32
	SetBody(v *ListCallCenterProvidersResponseBody) *ListCallCenterProvidersResponse
	GetBody() *ListCallCenterProvidersResponseBody
}

type ListCallCenterProvidersResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListCallCenterProvidersResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListCallCenterProvidersResponse) String() string {
	return dara.Prettify(s)
}

func (s ListCallCenterProvidersResponse) GoString() string {
	return s.String()
}

func (s *ListCallCenterProvidersResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListCallCenterProvidersResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListCallCenterProvidersResponse) GetBody() *ListCallCenterProvidersResponseBody {
	return s.Body
}

func (s *ListCallCenterProvidersResponse) SetHeaders(v map[string]*string) *ListCallCenterProvidersResponse {
	s.Headers = v
	return s
}

func (s *ListCallCenterProvidersResponse) SetStatusCode(v int32) *ListCallCenterProvidersResponse {
	s.StatusCode = &v
	return s
}

func (s *ListCallCenterProvidersResponse) SetBody(v *ListCallCenterProvidersResponseBody) *ListCallCenterProvidersResponse {
	s.Body = v
	return s
}

func (s *ListCallCenterProvidersResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
