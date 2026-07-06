// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDataConnectorsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListDataConnectorsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListDataConnectorsResponse
	GetStatusCode() *int32
	SetBody(v *ListDataConnectorsResponseBody) *ListDataConnectorsResponse
	GetBody() *ListDataConnectorsResponseBody
}

type ListDataConnectorsResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListDataConnectorsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListDataConnectorsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListDataConnectorsResponse) GoString() string {
	return s.String()
}

func (s *ListDataConnectorsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListDataConnectorsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListDataConnectorsResponse) GetBody() *ListDataConnectorsResponseBody {
	return s.Body
}

func (s *ListDataConnectorsResponse) SetHeaders(v map[string]*string) *ListDataConnectorsResponse {
	s.Headers = v
	return s
}

func (s *ListDataConnectorsResponse) SetStatusCode(v int32) *ListDataConnectorsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListDataConnectorsResponse) SetBody(v *ListDataConnectorsResponseBody) *ListDataConnectorsResponse {
	s.Body = v
	return s
}

func (s *ListDataConnectorsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
