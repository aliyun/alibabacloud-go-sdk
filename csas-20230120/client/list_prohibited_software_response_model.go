// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListProhibitedSoftwareResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListProhibitedSoftwareResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListProhibitedSoftwareResponse
	GetStatusCode() *int32
	SetBody(v *ListProhibitedSoftwareResponseBody) *ListProhibitedSoftwareResponse
	GetBody() *ListProhibitedSoftwareResponseBody
}

type ListProhibitedSoftwareResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListProhibitedSoftwareResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListProhibitedSoftwareResponse) String() string {
	return dara.Prettify(s)
}

func (s ListProhibitedSoftwareResponse) GoString() string {
	return s.String()
}

func (s *ListProhibitedSoftwareResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListProhibitedSoftwareResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListProhibitedSoftwareResponse) GetBody() *ListProhibitedSoftwareResponseBody {
	return s.Body
}

func (s *ListProhibitedSoftwareResponse) SetHeaders(v map[string]*string) *ListProhibitedSoftwareResponse {
	s.Headers = v
	return s
}

func (s *ListProhibitedSoftwareResponse) SetStatusCode(v int32) *ListProhibitedSoftwareResponse {
	s.StatusCode = &v
	return s
}

func (s *ListProhibitedSoftwareResponse) SetBody(v *ListProhibitedSoftwareResponseBody) *ListProhibitedSoftwareResponse {
	s.Body = v
	return s
}

func (s *ListProhibitedSoftwareResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
