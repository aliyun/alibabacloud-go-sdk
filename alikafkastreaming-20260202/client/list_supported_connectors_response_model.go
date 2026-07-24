// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSupportedConnectorsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListSupportedConnectorsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListSupportedConnectorsResponse
	GetStatusCode() *int32
	SetBody(v *ListSupportedConnectorsResponseBody) *ListSupportedConnectorsResponse
	GetBody() *ListSupportedConnectorsResponseBody
}

type ListSupportedConnectorsResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListSupportedConnectorsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListSupportedConnectorsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListSupportedConnectorsResponse) GoString() string {
	return s.String()
}

func (s *ListSupportedConnectorsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListSupportedConnectorsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListSupportedConnectorsResponse) GetBody() *ListSupportedConnectorsResponseBody {
	return s.Body
}

func (s *ListSupportedConnectorsResponse) SetHeaders(v map[string]*string) *ListSupportedConnectorsResponse {
	s.Headers = v
	return s
}

func (s *ListSupportedConnectorsResponse) SetStatusCode(v int32) *ListSupportedConnectorsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListSupportedConnectorsResponse) SetBody(v *ListSupportedConnectorsResponseBody) *ListSupportedConnectorsResponse {
	s.Body = v
	return s
}

func (s *ListSupportedConnectorsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
