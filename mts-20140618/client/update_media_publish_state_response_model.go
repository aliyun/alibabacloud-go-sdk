// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateMediaPublishStateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateMediaPublishStateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateMediaPublishStateResponse
	GetStatusCode() *int32
	SetBody(v *UpdateMediaPublishStateResponseBody) *UpdateMediaPublishStateResponse
	GetBody() *UpdateMediaPublishStateResponseBody
}

type UpdateMediaPublishStateResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateMediaPublishStateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateMediaPublishStateResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateMediaPublishStateResponse) GoString() string {
	return s.String()
}

func (s *UpdateMediaPublishStateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateMediaPublishStateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateMediaPublishStateResponse) GetBody() *UpdateMediaPublishStateResponseBody {
	return s.Body
}

func (s *UpdateMediaPublishStateResponse) SetHeaders(v map[string]*string) *UpdateMediaPublishStateResponse {
	s.Headers = v
	return s
}

func (s *UpdateMediaPublishStateResponse) SetStatusCode(v int32) *UpdateMediaPublishStateResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateMediaPublishStateResponse) SetBody(v *UpdateMediaPublishStateResponseBody) *UpdateMediaPublishStateResponse {
	s.Body = v
	return s
}

func (s *UpdateMediaPublishStateResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
