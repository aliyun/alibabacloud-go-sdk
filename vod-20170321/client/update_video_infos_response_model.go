// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateVideoInfosResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateVideoInfosResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateVideoInfosResponse
	GetStatusCode() *int32
	SetBody(v *UpdateVideoInfosResponseBody) *UpdateVideoInfosResponse
	GetBody() *UpdateVideoInfosResponseBody
}

type UpdateVideoInfosResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateVideoInfosResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateVideoInfosResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateVideoInfosResponse) GoString() string {
	return s.String()
}

func (s *UpdateVideoInfosResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateVideoInfosResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateVideoInfosResponse) GetBody() *UpdateVideoInfosResponseBody {
	return s.Body
}

func (s *UpdateVideoInfosResponse) SetHeaders(v map[string]*string) *UpdateVideoInfosResponse {
	s.Headers = v
	return s
}

func (s *UpdateVideoInfosResponse) SetStatusCode(v int32) *UpdateVideoInfosResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateVideoInfosResponse) SetBody(v *UpdateVideoInfosResponseBody) *UpdateVideoInfosResponse {
	s.Body = v
	return s
}

func (s *UpdateVideoInfosResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
