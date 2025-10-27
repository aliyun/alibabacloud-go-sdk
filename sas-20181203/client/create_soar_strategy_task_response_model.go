// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSoarStrategyTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateSoarStrategyTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateSoarStrategyTaskResponse
	GetStatusCode() *int32
	SetBody(v *CreateSoarStrategyTaskResponseBody) *CreateSoarStrategyTaskResponse
	GetBody() *CreateSoarStrategyTaskResponseBody
}

type CreateSoarStrategyTaskResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateSoarStrategyTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateSoarStrategyTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateSoarStrategyTaskResponse) GoString() string {
	return s.String()
}

func (s *CreateSoarStrategyTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateSoarStrategyTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateSoarStrategyTaskResponse) GetBody() *CreateSoarStrategyTaskResponseBody {
	return s.Body
}

func (s *CreateSoarStrategyTaskResponse) SetHeaders(v map[string]*string) *CreateSoarStrategyTaskResponse {
	s.Headers = v
	return s
}

func (s *CreateSoarStrategyTaskResponse) SetStatusCode(v int32) *CreateSoarStrategyTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateSoarStrategyTaskResponse) SetBody(v *CreateSoarStrategyTaskResponseBody) *CreateSoarStrategyTaskResponse {
	s.Body = v
	return s
}

func (s *CreateSoarStrategyTaskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
