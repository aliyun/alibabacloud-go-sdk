// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRollbackModelResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RollbackModelResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RollbackModelResponse
	GetStatusCode() *int32
	SetBody(v *RollbackModelResponseBody) *RollbackModelResponse
	GetBody() *RollbackModelResponseBody
}

type RollbackModelResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RollbackModelResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RollbackModelResponse) String() string {
	return dara.Prettify(s)
}

func (s RollbackModelResponse) GoString() string {
	return s.String()
}

func (s *RollbackModelResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RollbackModelResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RollbackModelResponse) GetBody() *RollbackModelResponseBody {
	return s.Body
}

func (s *RollbackModelResponse) SetHeaders(v map[string]*string) *RollbackModelResponse {
	s.Headers = v
	return s
}

func (s *RollbackModelResponse) SetStatusCode(v int32) *RollbackModelResponse {
	s.StatusCode = &v
	return s
}

func (s *RollbackModelResponse) SetBody(v *RollbackModelResponseBody) *RollbackModelResponse {
	s.Body = v
	return s
}

func (s *RollbackModelResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
