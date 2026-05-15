// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteOuterAccountResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteOuterAccountResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteOuterAccountResponse
	GetStatusCode() *int32
	SetBody(v *DeleteOuterAccountResponseBody) *DeleteOuterAccountResponse
	GetBody() *DeleteOuterAccountResponseBody
}

type DeleteOuterAccountResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteOuterAccountResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteOuterAccountResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteOuterAccountResponse) GoString() string {
	return s.String()
}

func (s *DeleteOuterAccountResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteOuterAccountResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteOuterAccountResponse) GetBody() *DeleteOuterAccountResponseBody {
	return s.Body
}

func (s *DeleteOuterAccountResponse) SetHeaders(v map[string]*string) *DeleteOuterAccountResponse {
	s.Headers = v
	return s
}

func (s *DeleteOuterAccountResponse) SetStatusCode(v int32) *DeleteOuterAccountResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteOuterAccountResponse) SetBody(v *DeleteOuterAccountResponseBody) *DeleteOuterAccountResponse {
	s.Body = v
	return s
}

func (s *DeleteOuterAccountResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
