// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetOpaStrategyDetailNewResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetOpaStrategyDetailNewResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetOpaStrategyDetailNewResponse
	GetStatusCode() *int32
	SetBody(v *GetOpaStrategyDetailNewResponseBody) *GetOpaStrategyDetailNewResponse
	GetBody() *GetOpaStrategyDetailNewResponseBody
}

type GetOpaStrategyDetailNewResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetOpaStrategyDetailNewResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetOpaStrategyDetailNewResponse) String() string {
	return dara.Prettify(s)
}

func (s GetOpaStrategyDetailNewResponse) GoString() string {
	return s.String()
}

func (s *GetOpaStrategyDetailNewResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetOpaStrategyDetailNewResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetOpaStrategyDetailNewResponse) GetBody() *GetOpaStrategyDetailNewResponseBody {
	return s.Body
}

func (s *GetOpaStrategyDetailNewResponse) SetHeaders(v map[string]*string) *GetOpaStrategyDetailNewResponse {
	s.Headers = v
	return s
}

func (s *GetOpaStrategyDetailNewResponse) SetStatusCode(v int32) *GetOpaStrategyDetailNewResponse {
	s.StatusCode = &v
	return s
}

func (s *GetOpaStrategyDetailNewResponse) SetBody(v *GetOpaStrategyDetailNewResponseBody) *GetOpaStrategyDetailNewResponse {
	s.Body = v
	return s
}

func (s *GetOpaStrategyDetailNewResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
