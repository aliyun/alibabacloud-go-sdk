// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateDataLakeFunctionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateDataLakeFunctionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateDataLakeFunctionResponse
	GetStatusCode() *int32
	SetBody(v *UpdateDataLakeFunctionResponseBody) *UpdateDataLakeFunctionResponse
	GetBody() *UpdateDataLakeFunctionResponseBody
}

type UpdateDataLakeFunctionResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateDataLakeFunctionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateDataLakeFunctionResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateDataLakeFunctionResponse) GoString() string {
	return s.String()
}

func (s *UpdateDataLakeFunctionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateDataLakeFunctionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateDataLakeFunctionResponse) GetBody() *UpdateDataLakeFunctionResponseBody {
	return s.Body
}

func (s *UpdateDataLakeFunctionResponse) SetHeaders(v map[string]*string) *UpdateDataLakeFunctionResponse {
	s.Headers = v
	return s
}

func (s *UpdateDataLakeFunctionResponse) SetStatusCode(v int32) *UpdateDataLakeFunctionResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateDataLakeFunctionResponse) SetBody(v *UpdateDataLakeFunctionResponseBody) *UpdateDataLakeFunctionResponse {
	s.Body = v
	return s
}

func (s *UpdateDataLakeFunctionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
