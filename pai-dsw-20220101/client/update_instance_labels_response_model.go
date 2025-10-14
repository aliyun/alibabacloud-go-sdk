// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateInstanceLabelsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateInstanceLabelsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateInstanceLabelsResponse
	GetStatusCode() *int32
	SetBody(v *UpdateInstanceLabelsResponseBody) *UpdateInstanceLabelsResponse
	GetBody() *UpdateInstanceLabelsResponseBody
}

type UpdateInstanceLabelsResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateInstanceLabelsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateInstanceLabelsResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateInstanceLabelsResponse) GoString() string {
	return s.String()
}

func (s *UpdateInstanceLabelsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateInstanceLabelsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateInstanceLabelsResponse) GetBody() *UpdateInstanceLabelsResponseBody {
	return s.Body
}

func (s *UpdateInstanceLabelsResponse) SetHeaders(v map[string]*string) *UpdateInstanceLabelsResponse {
	s.Headers = v
	return s
}

func (s *UpdateInstanceLabelsResponse) SetStatusCode(v int32) *UpdateInstanceLabelsResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateInstanceLabelsResponse) SetBody(v *UpdateInstanceLabelsResponseBody) *UpdateInstanceLabelsResponse {
	s.Body = v
	return s
}

func (s *UpdateInstanceLabelsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
