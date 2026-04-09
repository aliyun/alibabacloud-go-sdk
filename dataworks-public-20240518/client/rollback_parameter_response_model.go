// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRollbackParameterResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RollbackParameterResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RollbackParameterResponse
	GetStatusCode() *int32
	SetBody(v *RollbackParameterResponseBody) *RollbackParameterResponse
	GetBody() *RollbackParameterResponseBody
}

type RollbackParameterResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RollbackParameterResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RollbackParameterResponse) String() string {
	return dara.Prettify(s)
}

func (s RollbackParameterResponse) GoString() string {
	return s.String()
}

func (s *RollbackParameterResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RollbackParameterResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RollbackParameterResponse) GetBody() *RollbackParameterResponseBody {
	return s.Body
}

func (s *RollbackParameterResponse) SetHeaders(v map[string]*string) *RollbackParameterResponse {
	s.Headers = v
	return s
}

func (s *RollbackParameterResponse) SetStatusCode(v int32) *RollbackParameterResponse {
	s.StatusCode = &v
	return s
}

func (s *RollbackParameterResponse) SetBody(v *RollbackParameterResponseBody) *RollbackParameterResponse {
	s.Body = v
	return s
}

func (s *RollbackParameterResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
