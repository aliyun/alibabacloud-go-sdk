// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetFunctionTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetFunctionTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetFunctionTaskResponse
	GetStatusCode() *int32
	SetBody(v *GetFunctionTaskResponseBody) *GetFunctionTaskResponse
	GetBody() *GetFunctionTaskResponseBody
}

type GetFunctionTaskResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetFunctionTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetFunctionTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s GetFunctionTaskResponse) GoString() string {
	return s.String()
}

func (s *GetFunctionTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetFunctionTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetFunctionTaskResponse) GetBody() *GetFunctionTaskResponseBody {
	return s.Body
}

func (s *GetFunctionTaskResponse) SetHeaders(v map[string]*string) *GetFunctionTaskResponse {
	s.Headers = v
	return s
}

func (s *GetFunctionTaskResponse) SetStatusCode(v int32) *GetFunctionTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *GetFunctionTaskResponse) SetBody(v *GetFunctionTaskResponseBody) *GetFunctionTaskResponse {
	s.Body = v
	return s
}

func (s *GetFunctionTaskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
