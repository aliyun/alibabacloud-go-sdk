// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteFunctionTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteFunctionTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteFunctionTaskResponse
	GetStatusCode() *int32
	SetBody(v *DeleteFunctionTaskResponseBody) *DeleteFunctionTaskResponse
	GetBody() *DeleteFunctionTaskResponseBody
}

type DeleteFunctionTaskResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteFunctionTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteFunctionTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteFunctionTaskResponse) GoString() string {
	return s.String()
}

func (s *DeleteFunctionTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteFunctionTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteFunctionTaskResponse) GetBody() *DeleteFunctionTaskResponseBody {
	return s.Body
}

func (s *DeleteFunctionTaskResponse) SetHeaders(v map[string]*string) *DeleteFunctionTaskResponse {
	s.Headers = v
	return s
}

func (s *DeleteFunctionTaskResponse) SetStatusCode(v int32) *DeleteFunctionTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteFunctionTaskResponse) SetBody(v *DeleteFunctionTaskResponseBody) *DeleteFunctionTaskResponse {
	s.Body = v
	return s
}

func (s *DeleteFunctionTaskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
