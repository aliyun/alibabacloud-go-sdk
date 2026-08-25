// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListModelConnectionsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListModelConnectionsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListModelConnectionsResponse
	GetStatusCode() *int32
	SetBody(v *ListModelConnectionsResponseBody) *ListModelConnectionsResponse
	GetBody() *ListModelConnectionsResponseBody
}

type ListModelConnectionsResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListModelConnectionsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListModelConnectionsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListModelConnectionsResponse) GoString() string {
	return s.String()
}

func (s *ListModelConnectionsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListModelConnectionsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListModelConnectionsResponse) GetBody() *ListModelConnectionsResponseBody {
	return s.Body
}

func (s *ListModelConnectionsResponse) SetHeaders(v map[string]*string) *ListModelConnectionsResponse {
	s.Headers = v
	return s
}

func (s *ListModelConnectionsResponse) SetStatusCode(v int32) *ListModelConnectionsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListModelConnectionsResponse) SetBody(v *ListModelConnectionsResponseBody) *ListModelConnectionsResponse {
	s.Body = v
	return s
}

func (s *ListModelConnectionsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
