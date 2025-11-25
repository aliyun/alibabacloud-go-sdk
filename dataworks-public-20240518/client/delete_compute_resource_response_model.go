// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteComputeResourceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteComputeResourceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteComputeResourceResponse
	GetStatusCode() *int32
	SetBody(v *DeleteComputeResourceResponseBody) *DeleteComputeResourceResponse
	GetBody() *DeleteComputeResourceResponseBody
}

type DeleteComputeResourceResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteComputeResourceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteComputeResourceResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteComputeResourceResponse) GoString() string {
	return s.String()
}

func (s *DeleteComputeResourceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteComputeResourceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteComputeResourceResponse) GetBody() *DeleteComputeResourceResponseBody {
	return s.Body
}

func (s *DeleteComputeResourceResponse) SetHeaders(v map[string]*string) *DeleteComputeResourceResponse {
	s.Headers = v
	return s
}

func (s *DeleteComputeResourceResponse) SetStatusCode(v int32) *DeleteComputeResourceResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteComputeResourceResponse) SetBody(v *DeleteComputeResourceResponseBody) *DeleteComputeResourceResponse {
	s.Body = v
	return s
}

func (s *DeleteComputeResourceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
