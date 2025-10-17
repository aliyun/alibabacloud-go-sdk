// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteLivyComputeTokenResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteLivyComputeTokenResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteLivyComputeTokenResponse
	GetStatusCode() *int32
	SetBody(v *DeleteLivyComputeTokenResponseBody) *DeleteLivyComputeTokenResponse
	GetBody() *DeleteLivyComputeTokenResponseBody
}

type DeleteLivyComputeTokenResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteLivyComputeTokenResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteLivyComputeTokenResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteLivyComputeTokenResponse) GoString() string {
	return s.String()
}

func (s *DeleteLivyComputeTokenResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteLivyComputeTokenResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteLivyComputeTokenResponse) GetBody() *DeleteLivyComputeTokenResponseBody {
	return s.Body
}

func (s *DeleteLivyComputeTokenResponse) SetHeaders(v map[string]*string) *DeleteLivyComputeTokenResponse {
	s.Headers = v
	return s
}

func (s *DeleteLivyComputeTokenResponse) SetStatusCode(v int32) *DeleteLivyComputeTokenResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteLivyComputeTokenResponse) SetBody(v *DeleteLivyComputeTokenResponseBody) *DeleteLivyComputeTokenResponse {
	s.Body = v
	return s
}

func (s *DeleteLivyComputeTokenResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
