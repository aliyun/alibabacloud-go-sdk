// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetExecuteStateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetExecuteStateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetExecuteStateResponse
	GetStatusCode() *int32
	SetBody(v *GetExecuteStateResponseBody) *GetExecuteStateResponse
	GetBody() *GetExecuteStateResponseBody
}

type GetExecuteStateResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetExecuteStateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetExecuteStateResponse) String() string {
	return dara.Prettify(s)
}

func (s GetExecuteStateResponse) GoString() string {
	return s.String()
}

func (s *GetExecuteStateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetExecuteStateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetExecuteStateResponse) GetBody() *GetExecuteStateResponseBody {
	return s.Body
}

func (s *GetExecuteStateResponse) SetHeaders(v map[string]*string) *GetExecuteStateResponse {
	s.Headers = v
	return s
}

func (s *GetExecuteStateResponse) SetStatusCode(v int32) *GetExecuteStateResponse {
	s.StatusCode = &v
	return s
}

func (s *GetExecuteStateResponse) SetBody(v *GetExecuteStateResponseBody) *GetExecuteStateResponse {
	s.Body = v
	return s
}

func (s *GetExecuteStateResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
