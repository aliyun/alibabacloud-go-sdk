// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteCrossAccountResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteCrossAccountResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteCrossAccountResponse
	GetStatusCode() *int32
	SetBody(v *DeleteCrossAccountResponseBody) *DeleteCrossAccountResponse
	GetBody() *DeleteCrossAccountResponseBody
}

type DeleteCrossAccountResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteCrossAccountResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteCrossAccountResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteCrossAccountResponse) GoString() string {
	return s.String()
}

func (s *DeleteCrossAccountResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteCrossAccountResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteCrossAccountResponse) GetBody() *DeleteCrossAccountResponseBody {
	return s.Body
}

func (s *DeleteCrossAccountResponse) SetHeaders(v map[string]*string) *DeleteCrossAccountResponse {
	s.Headers = v
	return s
}

func (s *DeleteCrossAccountResponse) SetStatusCode(v int32) *DeleteCrossAccountResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteCrossAccountResponse) SetBody(v *DeleteCrossAccountResponseBody) *DeleteCrossAccountResponse {
	s.Body = v
	return s
}

func (s *DeleteCrossAccountResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
