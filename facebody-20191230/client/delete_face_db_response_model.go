// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteFaceDbResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteFaceDbResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteFaceDbResponse
	GetStatusCode() *int32
	SetBody(v *DeleteFaceDbResponseBody) *DeleteFaceDbResponse
	GetBody() *DeleteFaceDbResponseBody
}

type DeleteFaceDbResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteFaceDbResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteFaceDbResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteFaceDbResponse) GoString() string {
	return s.String()
}

func (s *DeleteFaceDbResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteFaceDbResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteFaceDbResponse) GetBody() *DeleteFaceDbResponseBody {
	return s.Body
}

func (s *DeleteFaceDbResponse) SetHeaders(v map[string]*string) *DeleteFaceDbResponse {
	s.Headers = v
	return s
}

func (s *DeleteFaceDbResponse) SetStatusCode(v int32) *DeleteFaceDbResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteFaceDbResponse) SetBody(v *DeleteFaceDbResponseBody) *DeleteFaceDbResponse {
	s.Body = v
	return s
}

func (s *DeleteFaceDbResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
