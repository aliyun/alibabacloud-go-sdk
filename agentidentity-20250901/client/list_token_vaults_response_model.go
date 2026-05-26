// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTokenVaultsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListTokenVaultsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListTokenVaultsResponse
	GetStatusCode() *int32
	SetBody(v *ListTokenVaultsResponseBody) *ListTokenVaultsResponse
	GetBody() *ListTokenVaultsResponseBody
}

type ListTokenVaultsResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListTokenVaultsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListTokenVaultsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListTokenVaultsResponse) GoString() string {
	return s.String()
}

func (s *ListTokenVaultsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListTokenVaultsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListTokenVaultsResponse) GetBody() *ListTokenVaultsResponseBody {
	return s.Body
}

func (s *ListTokenVaultsResponse) SetHeaders(v map[string]*string) *ListTokenVaultsResponse {
	s.Headers = v
	return s
}

func (s *ListTokenVaultsResponse) SetStatusCode(v int32) *ListTokenVaultsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListTokenVaultsResponse) SetBody(v *ListTokenVaultsResponseBody) *ListTokenVaultsResponse {
	s.Body = v
	return s
}

func (s *ListTokenVaultsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
