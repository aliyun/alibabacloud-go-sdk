// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteUserGroupsMappingResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteUserGroupsMappingResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteUserGroupsMappingResponse
	GetStatusCode() *int32
	SetBody(v *DeleteUserGroupsMappingResponseBody) *DeleteUserGroupsMappingResponse
	GetBody() *DeleteUserGroupsMappingResponseBody
}

type DeleteUserGroupsMappingResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteUserGroupsMappingResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteUserGroupsMappingResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteUserGroupsMappingResponse) GoString() string {
	return s.String()
}

func (s *DeleteUserGroupsMappingResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteUserGroupsMappingResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteUserGroupsMappingResponse) GetBody() *DeleteUserGroupsMappingResponseBody {
	return s.Body
}

func (s *DeleteUserGroupsMappingResponse) SetHeaders(v map[string]*string) *DeleteUserGroupsMappingResponse {
	s.Headers = v
	return s
}

func (s *DeleteUserGroupsMappingResponse) SetStatusCode(v int32) *DeleteUserGroupsMappingResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteUserGroupsMappingResponse) SetBody(v *DeleteUserGroupsMappingResponseBody) *DeleteUserGroupsMappingResponse {
	s.Body = v
	return s
}

func (s *DeleteUserGroupsMappingResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
