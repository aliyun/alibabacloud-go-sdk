// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteLogicalModelResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteLogicalModelResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteLogicalModelResponse
	GetStatusCode() *int32
	SetBody(v *DeleteLogicalModelResponseBody) *DeleteLogicalModelResponse
	GetBody() *DeleteLogicalModelResponseBody
}

type DeleteLogicalModelResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteLogicalModelResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteLogicalModelResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteLogicalModelResponse) GoString() string {
	return s.String()
}

func (s *DeleteLogicalModelResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteLogicalModelResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteLogicalModelResponse) GetBody() *DeleteLogicalModelResponseBody {
	return s.Body
}

func (s *DeleteLogicalModelResponse) SetHeaders(v map[string]*string) *DeleteLogicalModelResponse {
	s.Headers = v
	return s
}

func (s *DeleteLogicalModelResponse) SetStatusCode(v int32) *DeleteLogicalModelResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteLogicalModelResponse) SetBody(v *DeleteLogicalModelResponseBody) *DeleteLogicalModelResponse {
	s.Body = v
	return s
}

func (s *DeleteLogicalModelResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
