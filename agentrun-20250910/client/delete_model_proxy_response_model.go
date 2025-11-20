// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteModelProxyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteModelProxyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteModelProxyResponse
	GetStatusCode() *int32
	SetBody(v *DeleteModelProxyResult) *DeleteModelProxyResponse
	GetBody() *DeleteModelProxyResult
}

type DeleteModelProxyResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteModelProxyResult `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteModelProxyResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteModelProxyResponse) GoString() string {
	return s.String()
}

func (s *DeleteModelProxyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteModelProxyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteModelProxyResponse) GetBody() *DeleteModelProxyResult {
	return s.Body
}

func (s *DeleteModelProxyResponse) SetHeaders(v map[string]*string) *DeleteModelProxyResponse {
	s.Headers = v
	return s
}

func (s *DeleteModelProxyResponse) SetStatusCode(v int32) *DeleteModelProxyResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteModelProxyResponse) SetBody(v *DeleteModelProxyResult) *DeleteModelProxyResponse {
	s.Body = v
	return s
}

func (s *DeleteModelProxyResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
