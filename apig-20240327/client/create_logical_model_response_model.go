// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateLogicalModelResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateLogicalModelResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateLogicalModelResponse
	GetStatusCode() *int32
	SetBody(v *CreateLogicalModelResponseBody) *CreateLogicalModelResponse
	GetBody() *CreateLogicalModelResponseBody
}

type CreateLogicalModelResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateLogicalModelResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateLogicalModelResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateLogicalModelResponse) GoString() string {
	return s.String()
}

func (s *CreateLogicalModelResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateLogicalModelResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateLogicalModelResponse) GetBody() *CreateLogicalModelResponseBody {
	return s.Body
}

func (s *CreateLogicalModelResponse) SetHeaders(v map[string]*string) *CreateLogicalModelResponse {
	s.Headers = v
	return s
}

func (s *CreateLogicalModelResponse) SetStatusCode(v int32) *CreateLogicalModelResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateLogicalModelResponse) SetBody(v *CreateLogicalModelResponseBody) *CreateLogicalModelResponse {
	s.Body = v
	return s
}

func (s *CreateLogicalModelResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
