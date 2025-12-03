// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iConvertClusterResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ConvertClusterResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ConvertClusterResponse
	GetStatusCode() *int32
	SetBody(v *ConvertClusterResponseBody) *ConvertClusterResponse
	GetBody() *ConvertClusterResponseBody
}

type ConvertClusterResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ConvertClusterResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ConvertClusterResponse) String() string {
	return dara.Prettify(s)
}

func (s ConvertClusterResponse) GoString() string {
	return s.String()
}

func (s *ConvertClusterResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ConvertClusterResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ConvertClusterResponse) GetBody() *ConvertClusterResponseBody {
	return s.Body
}

func (s *ConvertClusterResponse) SetHeaders(v map[string]*string) *ConvertClusterResponse {
	s.Headers = v
	return s
}

func (s *ConvertClusterResponse) SetStatusCode(v int32) *ConvertClusterResponse {
	s.StatusCode = &v
	return s
}

func (s *ConvertClusterResponse) SetBody(v *ConvertClusterResponseBody) *ConvertClusterResponse {
	s.Body = v
	return s
}

func (s *ConvertClusterResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
