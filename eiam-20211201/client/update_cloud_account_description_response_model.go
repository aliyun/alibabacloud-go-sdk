// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateCloudAccountDescriptionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateCloudAccountDescriptionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateCloudAccountDescriptionResponse
	GetStatusCode() *int32
	SetBody(v *UpdateCloudAccountDescriptionResponseBody) *UpdateCloudAccountDescriptionResponse
	GetBody() *UpdateCloudAccountDescriptionResponseBody
}

type UpdateCloudAccountDescriptionResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateCloudAccountDescriptionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateCloudAccountDescriptionResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateCloudAccountDescriptionResponse) GoString() string {
	return s.String()
}

func (s *UpdateCloudAccountDescriptionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateCloudAccountDescriptionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateCloudAccountDescriptionResponse) GetBody() *UpdateCloudAccountDescriptionResponseBody {
	return s.Body
}

func (s *UpdateCloudAccountDescriptionResponse) SetHeaders(v map[string]*string) *UpdateCloudAccountDescriptionResponse {
	s.Headers = v
	return s
}

func (s *UpdateCloudAccountDescriptionResponse) SetStatusCode(v int32) *UpdateCloudAccountDescriptionResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateCloudAccountDescriptionResponse) SetBody(v *UpdateCloudAccountDescriptionResponseBody) *UpdateCloudAccountDescriptionResponse {
	s.Body = v
	return s
}

func (s *UpdateCloudAccountDescriptionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
