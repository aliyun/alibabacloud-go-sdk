// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOperateSwitchStatusResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *OperateSwitchStatusResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *OperateSwitchStatusResponse
	GetStatusCode() *int32
	SetBody(v *OperateSwitchStatusResponseBody) *OperateSwitchStatusResponse
	GetBody() *OperateSwitchStatusResponseBody
}

type OperateSwitchStatusResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *OperateSwitchStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s OperateSwitchStatusResponse) String() string {
	return dara.Prettify(s)
}

func (s OperateSwitchStatusResponse) GoString() string {
	return s.String()
}

func (s *OperateSwitchStatusResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *OperateSwitchStatusResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *OperateSwitchStatusResponse) GetBody() *OperateSwitchStatusResponseBody {
	return s.Body
}

func (s *OperateSwitchStatusResponse) SetHeaders(v map[string]*string) *OperateSwitchStatusResponse {
	s.Headers = v
	return s
}

func (s *OperateSwitchStatusResponse) SetStatusCode(v int32) *OperateSwitchStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *OperateSwitchStatusResponse) SetBody(v *OperateSwitchStatusResponseBody) *OperateSwitchStatusResponse {
	s.Body = v
	return s
}

func (s *OperateSwitchStatusResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
