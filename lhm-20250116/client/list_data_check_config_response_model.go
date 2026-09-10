// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDataCheckConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListDataCheckConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListDataCheckConfigResponse
	GetStatusCode() *int32
	SetBody(v *ListDataCheckConfigResponseBody) *ListDataCheckConfigResponse
	GetBody() *ListDataCheckConfigResponseBody
}

type ListDataCheckConfigResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListDataCheckConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListDataCheckConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s ListDataCheckConfigResponse) GoString() string {
	return s.String()
}

func (s *ListDataCheckConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListDataCheckConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListDataCheckConfigResponse) GetBody() *ListDataCheckConfigResponseBody {
	return s.Body
}

func (s *ListDataCheckConfigResponse) SetHeaders(v map[string]*string) *ListDataCheckConfigResponse {
	s.Headers = v
	return s
}

func (s *ListDataCheckConfigResponse) SetStatusCode(v int32) *ListDataCheckConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *ListDataCheckConfigResponse) SetBody(v *ListDataCheckConfigResponseBody) *ListDataCheckConfigResponse {
	s.Body = v
	return s
}

func (s *ListDataCheckConfigResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
