// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopStackGroupOperationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *StopStackGroupOperationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *StopStackGroupOperationResponse
	GetStatusCode() *int32
	SetBody(v *StopStackGroupOperationResponseBody) *StopStackGroupOperationResponse
	GetBody() *StopStackGroupOperationResponseBody
}

type StopStackGroupOperationResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StopStackGroupOperationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StopStackGroupOperationResponse) String() string {
	return dara.Prettify(s)
}

func (s StopStackGroupOperationResponse) GoString() string {
	return s.String()
}

func (s *StopStackGroupOperationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *StopStackGroupOperationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *StopStackGroupOperationResponse) GetBody() *StopStackGroupOperationResponseBody {
	return s.Body
}

func (s *StopStackGroupOperationResponse) SetHeaders(v map[string]*string) *StopStackGroupOperationResponse {
	s.Headers = v
	return s
}

func (s *StopStackGroupOperationResponse) SetStatusCode(v int32) *StopStackGroupOperationResponse {
	s.StatusCode = &v
	return s
}

func (s *StopStackGroupOperationResponse) SetBody(v *StopStackGroupOperationResponseBody) *StopStackGroupOperationResponse {
	s.Body = v
	return s
}

func (s *StopStackGroupOperationResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
