// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteHttpApiOperationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteHttpApiOperationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteHttpApiOperationResponse
	GetStatusCode() *int32
	SetBody(v *DeleteHttpApiOperationResponseBody) *DeleteHttpApiOperationResponse
	GetBody() *DeleteHttpApiOperationResponseBody
}

type DeleteHttpApiOperationResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteHttpApiOperationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteHttpApiOperationResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteHttpApiOperationResponse) GoString() string {
	return s.String()
}

func (s *DeleteHttpApiOperationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteHttpApiOperationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteHttpApiOperationResponse) GetBody() *DeleteHttpApiOperationResponseBody {
	return s.Body
}

func (s *DeleteHttpApiOperationResponse) SetHeaders(v map[string]*string) *DeleteHttpApiOperationResponse {
	s.Headers = v
	return s
}

func (s *DeleteHttpApiOperationResponse) SetStatusCode(v int32) *DeleteHttpApiOperationResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteHttpApiOperationResponse) SetBody(v *DeleteHttpApiOperationResponseBody) *DeleteHttpApiOperationResponse {
	s.Body = v
	return s
}

func (s *DeleteHttpApiOperationResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
