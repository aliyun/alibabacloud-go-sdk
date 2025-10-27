// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdatePublishCronResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdatePublishCronResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdatePublishCronResponse
	GetStatusCode() *int32
	SetBody(v *UpdatePublishCronResponseBody) *UpdatePublishCronResponse
	GetBody() *UpdatePublishCronResponseBody
}

type UpdatePublishCronResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdatePublishCronResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdatePublishCronResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdatePublishCronResponse) GoString() string {
	return s.String()
}

func (s *UpdatePublishCronResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdatePublishCronResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdatePublishCronResponse) GetBody() *UpdatePublishCronResponseBody {
	return s.Body
}

func (s *UpdatePublishCronResponse) SetHeaders(v map[string]*string) *UpdatePublishCronResponse {
	s.Headers = v
	return s
}

func (s *UpdatePublishCronResponse) SetStatusCode(v int32) *UpdatePublishCronResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdatePublishCronResponse) SetBody(v *UpdatePublishCronResponseBody) *UpdatePublishCronResponse {
	s.Body = v
	return s
}

func (s *UpdatePublishCronResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
