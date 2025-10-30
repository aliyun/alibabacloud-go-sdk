// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteModelServiceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteModelServiceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteModelServiceResponse
	GetStatusCode() *int32
	SetBody(v *DeleteModelServiceResponseBody) *DeleteModelServiceResponse
	GetBody() *DeleteModelServiceResponseBody
}

type DeleteModelServiceResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteModelServiceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteModelServiceResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteModelServiceResponse) GoString() string {
	return s.String()
}

func (s *DeleteModelServiceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteModelServiceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteModelServiceResponse) GetBody() *DeleteModelServiceResponseBody {
	return s.Body
}

func (s *DeleteModelServiceResponse) SetHeaders(v map[string]*string) *DeleteModelServiceResponse {
	s.Headers = v
	return s
}

func (s *DeleteModelServiceResponse) SetStatusCode(v int32) *DeleteModelServiceResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteModelServiceResponse) SetBody(v *DeleteModelServiceResponseBody) *DeleteModelServiceResponse {
	s.Body = v
	return s
}

func (s *DeleteModelServiceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
