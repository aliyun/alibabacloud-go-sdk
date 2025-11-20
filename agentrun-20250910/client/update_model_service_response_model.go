// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateModelServiceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateModelServiceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateModelServiceResponse
	GetStatusCode() *int32
	SetBody(v *ModelServiceResult) *UpdateModelServiceResponse
	GetBody() *ModelServiceResult
}

type UpdateModelServiceResponse struct {
	Headers    map[string]*string  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModelServiceResult `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateModelServiceResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateModelServiceResponse) GoString() string {
	return s.String()
}

func (s *UpdateModelServiceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateModelServiceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateModelServiceResponse) GetBody() *ModelServiceResult {
	return s.Body
}

func (s *UpdateModelServiceResponse) SetHeaders(v map[string]*string) *UpdateModelServiceResponse {
	s.Headers = v
	return s
}

func (s *UpdateModelServiceResponse) SetStatusCode(v int32) *UpdateModelServiceResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateModelServiceResponse) SetBody(v *ModelServiceResult) *UpdateModelServiceResponse {
	s.Body = v
	return s
}

func (s *UpdateModelServiceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
