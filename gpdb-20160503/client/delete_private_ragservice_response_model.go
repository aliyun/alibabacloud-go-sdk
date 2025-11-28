// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeletePrivateRAGServiceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeletePrivateRAGServiceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeletePrivateRAGServiceResponse
	GetStatusCode() *int32
	SetBody(v *DeletePrivateRAGServiceResponseBody) *DeletePrivateRAGServiceResponse
	GetBody() *DeletePrivateRAGServiceResponseBody
}

type DeletePrivateRAGServiceResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeletePrivateRAGServiceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeletePrivateRAGServiceResponse) String() string {
	return dara.Prettify(s)
}

func (s DeletePrivateRAGServiceResponse) GoString() string {
	return s.String()
}

func (s *DeletePrivateRAGServiceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeletePrivateRAGServiceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeletePrivateRAGServiceResponse) GetBody() *DeletePrivateRAGServiceResponseBody {
	return s.Body
}

func (s *DeletePrivateRAGServiceResponse) SetHeaders(v map[string]*string) *DeletePrivateRAGServiceResponse {
	s.Headers = v
	return s
}

func (s *DeletePrivateRAGServiceResponse) SetStatusCode(v int32) *DeletePrivateRAGServiceResponse {
	s.StatusCode = &v
	return s
}

func (s *DeletePrivateRAGServiceResponse) SetBody(v *DeletePrivateRAGServiceResponseBody) *DeletePrivateRAGServiceResponse {
	s.Body = v
	return s
}

func (s *DeletePrivateRAGServiceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
