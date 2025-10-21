// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCustomConnectorsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListCustomConnectorsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListCustomConnectorsResponse
	GetStatusCode() *int32
	SetBody(v *ListCustomConnectorsResponseBody) *ListCustomConnectorsResponse
	GetBody() *ListCustomConnectorsResponseBody
}

type ListCustomConnectorsResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListCustomConnectorsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListCustomConnectorsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListCustomConnectorsResponse) GoString() string {
	return s.String()
}

func (s *ListCustomConnectorsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListCustomConnectorsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListCustomConnectorsResponse) GetBody() *ListCustomConnectorsResponseBody {
	return s.Body
}

func (s *ListCustomConnectorsResponse) SetHeaders(v map[string]*string) *ListCustomConnectorsResponse {
	s.Headers = v
	return s
}

func (s *ListCustomConnectorsResponse) SetStatusCode(v int32) *ListCustomConnectorsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListCustomConnectorsResponse) SetBody(v *ListCustomConnectorsResponseBody) *ListCustomConnectorsResponse {
	s.Body = v
	return s
}

func (s *ListCustomConnectorsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
