// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteInstanceResourceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteInstanceResourceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteInstanceResourceResponse
	GetStatusCode() *int32
	SetBody(v *DeleteInstanceResourceResponseBody) *DeleteInstanceResourceResponse
	GetBody() *DeleteInstanceResourceResponseBody
}

type DeleteInstanceResourceResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteInstanceResourceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteInstanceResourceResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteInstanceResourceResponse) GoString() string {
	return s.String()
}

func (s *DeleteInstanceResourceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteInstanceResourceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteInstanceResourceResponse) GetBody() *DeleteInstanceResourceResponseBody {
	return s.Body
}

func (s *DeleteInstanceResourceResponse) SetHeaders(v map[string]*string) *DeleteInstanceResourceResponse {
	s.Headers = v
	return s
}

func (s *DeleteInstanceResourceResponse) SetStatusCode(v int32) *DeleteInstanceResourceResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteInstanceResourceResponse) SetBody(v *DeleteInstanceResourceResponseBody) *DeleteInstanceResourceResponse {
	s.Body = v
	return s
}

func (s *DeleteInstanceResourceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
