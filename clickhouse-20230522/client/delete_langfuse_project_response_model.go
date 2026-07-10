// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteLangfuseProjectResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteLangfuseProjectResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteLangfuseProjectResponse
	GetStatusCode() *int32
	SetBody(v *DeleteLangfuseProjectResponseBody) *DeleteLangfuseProjectResponse
	GetBody() *DeleteLangfuseProjectResponseBody
}

type DeleteLangfuseProjectResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteLangfuseProjectResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteLangfuseProjectResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteLangfuseProjectResponse) GoString() string {
	return s.String()
}

func (s *DeleteLangfuseProjectResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteLangfuseProjectResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteLangfuseProjectResponse) GetBody() *DeleteLangfuseProjectResponseBody {
	return s.Body
}

func (s *DeleteLangfuseProjectResponse) SetHeaders(v map[string]*string) *DeleteLangfuseProjectResponse {
	s.Headers = v
	return s
}

func (s *DeleteLangfuseProjectResponse) SetStatusCode(v int32) *DeleteLangfuseProjectResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteLangfuseProjectResponse) SetBody(v *DeleteLangfuseProjectResponseBody) *DeleteLangfuseProjectResponse {
	s.Body = v
	return s
}

func (s *DeleteLangfuseProjectResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
