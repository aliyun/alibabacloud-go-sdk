// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteFaceEntityResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteFaceEntityResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteFaceEntityResponse
	GetStatusCode() *int32
	SetBody(v *DeleteFaceEntityResponseBody) *DeleteFaceEntityResponse
	GetBody() *DeleteFaceEntityResponseBody
}

type DeleteFaceEntityResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteFaceEntityResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteFaceEntityResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteFaceEntityResponse) GoString() string {
	return s.String()
}

func (s *DeleteFaceEntityResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteFaceEntityResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteFaceEntityResponse) GetBody() *DeleteFaceEntityResponseBody {
	return s.Body
}

func (s *DeleteFaceEntityResponse) SetHeaders(v map[string]*string) *DeleteFaceEntityResponse {
	s.Headers = v
	return s
}

func (s *DeleteFaceEntityResponse) SetStatusCode(v int32) *DeleteFaceEntityResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteFaceEntityResponse) SetBody(v *DeleteFaceEntityResponseBody) *DeleteFaceEntityResponse {
	s.Body = v
	return s
}

func (s *DeleteFaceEntityResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
