// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListESAIPInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListESAIPInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListESAIPInfoResponse
	GetStatusCode() *int32
	SetBody(v *ListESAIPInfoResponseBody) *ListESAIPInfoResponse
	GetBody() *ListESAIPInfoResponseBody
}

type ListESAIPInfoResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListESAIPInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListESAIPInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s ListESAIPInfoResponse) GoString() string {
	return s.String()
}

func (s *ListESAIPInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListESAIPInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListESAIPInfoResponse) GetBody() *ListESAIPInfoResponseBody {
	return s.Body
}

func (s *ListESAIPInfoResponse) SetHeaders(v map[string]*string) *ListESAIPInfoResponse {
	s.Headers = v
	return s
}

func (s *ListESAIPInfoResponse) SetStatusCode(v int32) *ListESAIPInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *ListESAIPInfoResponse) SetBody(v *ListESAIPInfoResponseBody) *ListESAIPInfoResponse {
	s.Body = v
	return s
}

func (s *ListESAIPInfoResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
