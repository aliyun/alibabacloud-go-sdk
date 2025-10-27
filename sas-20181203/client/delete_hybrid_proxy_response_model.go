// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteHybridProxyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteHybridProxyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteHybridProxyResponse
	GetStatusCode() *int32
	SetBody(v *DeleteHybridProxyResponseBody) *DeleteHybridProxyResponse
	GetBody() *DeleteHybridProxyResponseBody
}

type DeleteHybridProxyResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteHybridProxyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteHybridProxyResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteHybridProxyResponse) GoString() string {
	return s.String()
}

func (s *DeleteHybridProxyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteHybridProxyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteHybridProxyResponse) GetBody() *DeleteHybridProxyResponseBody {
	return s.Body
}

func (s *DeleteHybridProxyResponse) SetHeaders(v map[string]*string) *DeleteHybridProxyResponse {
	s.Headers = v
	return s
}

func (s *DeleteHybridProxyResponse) SetStatusCode(v int32) *DeleteHybridProxyResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteHybridProxyResponse) SetBody(v *DeleteHybridProxyResponseBody) *DeleteHybridProxyResponse {
	s.Body = v
	return s
}

func (s *DeleteHybridProxyResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
