// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPushTimesUsageResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *PushTimesUsageResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *PushTimesUsageResponse
	GetStatusCode() *int32
	SetBody(v *PushTimesUsageResponseBody) *PushTimesUsageResponse
	GetBody() *PushTimesUsageResponseBody
}

type PushTimesUsageResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PushTimesUsageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PushTimesUsageResponse) String() string {
	return dara.Prettify(s)
}

func (s PushTimesUsageResponse) GoString() string {
	return s.String()
}

func (s *PushTimesUsageResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *PushTimesUsageResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *PushTimesUsageResponse) GetBody() *PushTimesUsageResponseBody {
	return s.Body
}

func (s *PushTimesUsageResponse) SetHeaders(v map[string]*string) *PushTimesUsageResponse {
	s.Headers = v
	return s
}

func (s *PushTimesUsageResponse) SetStatusCode(v int32) *PushTimesUsageResponse {
	s.StatusCode = &v
	return s
}

func (s *PushTimesUsageResponse) SetBody(v *PushTimesUsageResponseBody) *PushTimesUsageResponse {
	s.Body = v
	return s
}

func (s *PushTimesUsageResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
