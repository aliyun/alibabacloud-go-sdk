// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateTopicConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateTopicConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateTopicConfigResponse
	GetStatusCode() *int32
	SetBody(v *UpdateTopicConfigResponseBody) *UpdateTopicConfigResponse
	GetBody() *UpdateTopicConfigResponseBody
}

type UpdateTopicConfigResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateTopicConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateTopicConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateTopicConfigResponse) GoString() string {
	return s.String()
}

func (s *UpdateTopicConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateTopicConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateTopicConfigResponse) GetBody() *UpdateTopicConfigResponseBody {
	return s.Body
}

func (s *UpdateTopicConfigResponse) SetHeaders(v map[string]*string) *UpdateTopicConfigResponse {
	s.Headers = v
	return s
}

func (s *UpdateTopicConfigResponse) SetStatusCode(v int32) *UpdateTopicConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateTopicConfigResponse) SetBody(v *UpdateTopicConfigResponseBody) *UpdateTopicConfigResponse {
	s.Body = v
	return s
}

func (s *UpdateTopicConfigResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
