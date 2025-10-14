// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateTransformOperationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateTransformOperationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateTransformOperationResponse
	GetStatusCode() *int32
	SetBody(v *CreateTransformOperationResponseBody) *CreateTransformOperationResponse
	GetBody() *CreateTransformOperationResponseBody
}

type CreateTransformOperationResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateTransformOperationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateTransformOperationResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateTransformOperationResponse) GoString() string {
	return s.String()
}

func (s *CreateTransformOperationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateTransformOperationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateTransformOperationResponse) GetBody() *CreateTransformOperationResponseBody {
	return s.Body
}

func (s *CreateTransformOperationResponse) SetHeaders(v map[string]*string) *CreateTransformOperationResponse {
	s.Headers = v
	return s
}

func (s *CreateTransformOperationResponse) SetStatusCode(v int32) *CreateTransformOperationResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateTransformOperationResponse) SetBody(v *CreateTransformOperationResponseBody) *CreateTransformOperationResponse {
	s.Body = v
	return s
}

func (s *CreateTransformOperationResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
