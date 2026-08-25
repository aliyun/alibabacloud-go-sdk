// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListUserProvisioningsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListUserProvisioningsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListUserProvisioningsResponse
	GetStatusCode() *int32
	SetBody(v *ListUserProvisioningsResponseBody) *ListUserProvisioningsResponse
	GetBody() *ListUserProvisioningsResponseBody
}

type ListUserProvisioningsResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListUserProvisioningsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListUserProvisioningsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListUserProvisioningsResponse) GoString() string {
	return s.String()
}

func (s *ListUserProvisioningsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListUserProvisioningsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListUserProvisioningsResponse) GetBody() *ListUserProvisioningsResponseBody {
	return s.Body
}

func (s *ListUserProvisioningsResponse) SetHeaders(v map[string]*string) *ListUserProvisioningsResponse {
	s.Headers = v
	return s
}

func (s *ListUserProvisioningsResponse) SetStatusCode(v int32) *ListUserProvisioningsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListUserProvisioningsResponse) SetBody(v *ListUserProvisioningsResponseBody) *ListUserProvisioningsResponse {
	s.Body = v
	return s
}

func (s *ListUserProvisioningsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
