// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteLangfuseInstanceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteLangfuseInstanceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteLangfuseInstanceResponse
	GetStatusCode() *int32
	SetBody(v *DeleteLangfuseInstanceResponseBody) *DeleteLangfuseInstanceResponse
	GetBody() *DeleteLangfuseInstanceResponseBody
}

type DeleteLangfuseInstanceResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteLangfuseInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteLangfuseInstanceResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteLangfuseInstanceResponse) GoString() string {
	return s.String()
}

func (s *DeleteLangfuseInstanceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteLangfuseInstanceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteLangfuseInstanceResponse) GetBody() *DeleteLangfuseInstanceResponseBody {
	return s.Body
}

func (s *DeleteLangfuseInstanceResponse) SetHeaders(v map[string]*string) *DeleteLangfuseInstanceResponse {
	s.Headers = v
	return s
}

func (s *DeleteLangfuseInstanceResponse) SetStatusCode(v int32) *DeleteLangfuseInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteLangfuseInstanceResponse) SetBody(v *DeleteLangfuseInstanceResponseBody) *DeleteLangfuseInstanceResponse {
	s.Body = v
	return s
}

func (s *DeleteLangfuseInstanceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
