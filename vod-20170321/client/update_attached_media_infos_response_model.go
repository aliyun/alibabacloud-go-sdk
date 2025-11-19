// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAttachedMediaInfosResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateAttachedMediaInfosResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateAttachedMediaInfosResponse
	GetStatusCode() *int32
	SetBody(v *UpdateAttachedMediaInfosResponseBody) *UpdateAttachedMediaInfosResponse
	GetBody() *UpdateAttachedMediaInfosResponseBody
}

type UpdateAttachedMediaInfosResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateAttachedMediaInfosResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateAttachedMediaInfosResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateAttachedMediaInfosResponse) GoString() string {
	return s.String()
}

func (s *UpdateAttachedMediaInfosResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateAttachedMediaInfosResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateAttachedMediaInfosResponse) GetBody() *UpdateAttachedMediaInfosResponseBody {
	return s.Body
}

func (s *UpdateAttachedMediaInfosResponse) SetHeaders(v map[string]*string) *UpdateAttachedMediaInfosResponse {
	s.Headers = v
	return s
}

func (s *UpdateAttachedMediaInfosResponse) SetStatusCode(v int32) *UpdateAttachedMediaInfosResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateAttachedMediaInfosResponse) SetBody(v *UpdateAttachedMediaInfosResponseBody) *UpdateAttachedMediaInfosResponse {
	s.Body = v
	return s
}

func (s *UpdateAttachedMediaInfosResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
