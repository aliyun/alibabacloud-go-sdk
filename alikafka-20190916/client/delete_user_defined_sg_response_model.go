// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteUserDefinedSgResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteUserDefinedSgResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteUserDefinedSgResponse
	GetStatusCode() *int32
	SetBody(v *DeleteUserDefinedSgResponseBody) *DeleteUserDefinedSgResponse
	GetBody() *DeleteUserDefinedSgResponseBody
}

type DeleteUserDefinedSgResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteUserDefinedSgResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteUserDefinedSgResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteUserDefinedSgResponse) GoString() string {
	return s.String()
}

func (s *DeleteUserDefinedSgResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteUserDefinedSgResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteUserDefinedSgResponse) GetBody() *DeleteUserDefinedSgResponseBody {
	return s.Body
}

func (s *DeleteUserDefinedSgResponse) SetHeaders(v map[string]*string) *DeleteUserDefinedSgResponse {
	s.Headers = v
	return s
}

func (s *DeleteUserDefinedSgResponse) SetStatusCode(v int32) *DeleteUserDefinedSgResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteUserDefinedSgResponse) SetBody(v *DeleteUserDefinedSgResponseBody) *DeleteUserDefinedSgResponse {
	s.Body = v
	return s
}

func (s *DeleteUserDefinedSgResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
