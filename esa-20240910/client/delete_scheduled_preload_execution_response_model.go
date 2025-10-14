// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteScheduledPreloadExecutionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteScheduledPreloadExecutionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteScheduledPreloadExecutionResponse
	GetStatusCode() *int32
	SetBody(v *DeleteScheduledPreloadExecutionResponseBody) *DeleteScheduledPreloadExecutionResponse
	GetBody() *DeleteScheduledPreloadExecutionResponseBody
}

type DeleteScheduledPreloadExecutionResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteScheduledPreloadExecutionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteScheduledPreloadExecutionResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteScheduledPreloadExecutionResponse) GoString() string {
	return s.String()
}

func (s *DeleteScheduledPreloadExecutionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteScheduledPreloadExecutionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteScheduledPreloadExecutionResponse) GetBody() *DeleteScheduledPreloadExecutionResponseBody {
	return s.Body
}

func (s *DeleteScheduledPreloadExecutionResponse) SetHeaders(v map[string]*string) *DeleteScheduledPreloadExecutionResponse {
	s.Headers = v
	return s
}

func (s *DeleteScheduledPreloadExecutionResponse) SetStatusCode(v int32) *DeleteScheduledPreloadExecutionResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteScheduledPreloadExecutionResponse) SetBody(v *DeleteScheduledPreloadExecutionResponseBody) *DeleteScheduledPreloadExecutionResponse {
	s.Body = v
	return s
}

func (s *DeleteScheduledPreloadExecutionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
