// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDataLakeFunctionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateDataLakeFunctionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateDataLakeFunctionResponse
	GetStatusCode() *int32
	SetBody(v *CreateDataLakeFunctionResponseBody) *CreateDataLakeFunctionResponse
	GetBody() *CreateDataLakeFunctionResponseBody
}

type CreateDataLakeFunctionResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateDataLakeFunctionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateDataLakeFunctionResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateDataLakeFunctionResponse) GoString() string {
	return s.String()
}

func (s *CreateDataLakeFunctionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateDataLakeFunctionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateDataLakeFunctionResponse) GetBody() *CreateDataLakeFunctionResponseBody {
	return s.Body
}

func (s *CreateDataLakeFunctionResponse) SetHeaders(v map[string]*string) *CreateDataLakeFunctionResponse {
	s.Headers = v
	return s
}

func (s *CreateDataLakeFunctionResponse) SetStatusCode(v int32) *CreateDataLakeFunctionResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateDataLakeFunctionResponse) SetBody(v *CreateDataLakeFunctionResponseBody) *CreateDataLakeFunctionResponse {
	s.Body = v
	return s
}

func (s *CreateDataLakeFunctionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
