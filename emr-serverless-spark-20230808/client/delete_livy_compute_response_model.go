// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteLivyComputeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteLivyComputeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteLivyComputeResponse
	GetStatusCode() *int32
	SetBody(v *DeleteLivyComputeResponseBody) *DeleteLivyComputeResponse
	GetBody() *DeleteLivyComputeResponseBody
}

type DeleteLivyComputeResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteLivyComputeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteLivyComputeResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteLivyComputeResponse) GoString() string {
	return s.String()
}

func (s *DeleteLivyComputeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteLivyComputeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteLivyComputeResponse) GetBody() *DeleteLivyComputeResponseBody {
	return s.Body
}

func (s *DeleteLivyComputeResponse) SetHeaders(v map[string]*string) *DeleteLivyComputeResponse {
	s.Headers = v
	return s
}

func (s *DeleteLivyComputeResponse) SetStatusCode(v int32) *DeleteLivyComputeResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteLivyComputeResponse) SetBody(v *DeleteLivyComputeResponseBody) *DeleteLivyComputeResponse {
	s.Body = v
	return s
}

func (s *DeleteLivyComputeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
