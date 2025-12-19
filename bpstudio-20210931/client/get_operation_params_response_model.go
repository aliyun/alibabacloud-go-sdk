// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetOperationParamsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetOperationParamsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetOperationParamsResponse
	GetStatusCode() *int32
	SetBody(v *GetOperationParamsResponseBody) *GetOperationParamsResponse
	GetBody() *GetOperationParamsResponseBody
}

type GetOperationParamsResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetOperationParamsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetOperationParamsResponse) String() string {
	return dara.Prettify(s)
}

func (s GetOperationParamsResponse) GoString() string {
	return s.String()
}

func (s *GetOperationParamsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetOperationParamsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetOperationParamsResponse) GetBody() *GetOperationParamsResponseBody {
	return s.Body
}

func (s *GetOperationParamsResponse) SetHeaders(v map[string]*string) *GetOperationParamsResponse {
	s.Headers = v
	return s
}

func (s *GetOperationParamsResponse) SetStatusCode(v int32) *GetOperationParamsResponse {
	s.StatusCode = &v
	return s
}

func (s *GetOperationParamsResponse) SetBody(v *GetOperationParamsResponseBody) *GetOperationParamsResponse {
	s.Body = v
	return s
}

func (s *GetOperationParamsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
