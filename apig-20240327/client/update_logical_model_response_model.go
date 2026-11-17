// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateLogicalModelResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateLogicalModelResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateLogicalModelResponse
	GetStatusCode() *int32
	SetBody(v *UpdateLogicalModelResponseBody) *UpdateLogicalModelResponse
	GetBody() *UpdateLogicalModelResponseBody
}

type UpdateLogicalModelResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateLogicalModelResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateLogicalModelResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateLogicalModelResponse) GoString() string {
	return s.String()
}

func (s *UpdateLogicalModelResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateLogicalModelResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateLogicalModelResponse) GetBody() *UpdateLogicalModelResponseBody {
	return s.Body
}

func (s *UpdateLogicalModelResponse) SetHeaders(v map[string]*string) *UpdateLogicalModelResponse {
	s.Headers = v
	return s
}

func (s *UpdateLogicalModelResponse) SetStatusCode(v int32) *UpdateLogicalModelResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateLogicalModelResponse) SetBody(v *UpdateLogicalModelResponseBody) *UpdateLogicalModelResponse {
	s.Body = v
	return s
}

func (s *UpdateLogicalModelResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
