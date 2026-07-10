// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteLangfuseUserResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteLangfuseUserResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteLangfuseUserResponse
	GetStatusCode() *int32
	SetBody(v *DeleteLangfuseUserResponseBody) *DeleteLangfuseUserResponse
	GetBody() *DeleteLangfuseUserResponseBody
}

type DeleteLangfuseUserResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteLangfuseUserResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteLangfuseUserResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteLangfuseUserResponse) GoString() string {
	return s.String()
}

func (s *DeleteLangfuseUserResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteLangfuseUserResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteLangfuseUserResponse) GetBody() *DeleteLangfuseUserResponseBody {
	return s.Body
}

func (s *DeleteLangfuseUserResponse) SetHeaders(v map[string]*string) *DeleteLangfuseUserResponse {
	s.Headers = v
	return s
}

func (s *DeleteLangfuseUserResponse) SetStatusCode(v int32) *DeleteLangfuseUserResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteLangfuseUserResponse) SetBody(v *DeleteLangfuseUserResponseBody) *DeleteLangfuseUserResponse {
	s.Body = v
	return s
}

func (s *DeleteLangfuseUserResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
