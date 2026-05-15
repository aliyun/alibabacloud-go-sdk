// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSuspendHotlineServiceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SuspendHotlineServiceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SuspendHotlineServiceResponse
	GetStatusCode() *int32
	SetBody(v *SuspendHotlineServiceResponseBody) *SuspendHotlineServiceResponse
	GetBody() *SuspendHotlineServiceResponseBody
}

type SuspendHotlineServiceResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SuspendHotlineServiceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SuspendHotlineServiceResponse) String() string {
	return dara.Prettify(s)
}

func (s SuspendHotlineServiceResponse) GoString() string {
	return s.String()
}

func (s *SuspendHotlineServiceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SuspendHotlineServiceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SuspendHotlineServiceResponse) GetBody() *SuspendHotlineServiceResponseBody {
	return s.Body
}

func (s *SuspendHotlineServiceResponse) SetHeaders(v map[string]*string) *SuspendHotlineServiceResponse {
	s.Headers = v
	return s
}

func (s *SuspendHotlineServiceResponse) SetStatusCode(v int32) *SuspendHotlineServiceResponse {
	s.StatusCode = &v
	return s
}

func (s *SuspendHotlineServiceResponse) SetBody(v *SuspendHotlineServiceResponseBody) *SuspendHotlineServiceResponse {
	s.Body = v
	return s
}

func (s *SuspendHotlineServiceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
