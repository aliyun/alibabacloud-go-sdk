// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteUserTagMetaResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteUserTagMetaResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteUserTagMetaResponse
	GetStatusCode() *int32
	SetBody(v *DeleteUserTagMetaResponseBody) *DeleteUserTagMetaResponse
	GetBody() *DeleteUserTagMetaResponseBody
}

type DeleteUserTagMetaResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteUserTagMetaResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteUserTagMetaResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteUserTagMetaResponse) GoString() string {
	return s.String()
}

func (s *DeleteUserTagMetaResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteUserTagMetaResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteUserTagMetaResponse) GetBody() *DeleteUserTagMetaResponseBody {
	return s.Body
}

func (s *DeleteUserTagMetaResponse) SetHeaders(v map[string]*string) *DeleteUserTagMetaResponse {
	s.Headers = v
	return s
}

func (s *DeleteUserTagMetaResponse) SetStatusCode(v int32) *DeleteUserTagMetaResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteUserTagMetaResponse) SetBody(v *DeleteUserTagMetaResponseBody) *DeleteUserTagMetaResponse {
	s.Body = v
	return s
}

func (s *DeleteUserTagMetaResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
