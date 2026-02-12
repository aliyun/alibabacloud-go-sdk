// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOnsConsumerResetOffsetResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *OnsConsumerResetOffsetResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *OnsConsumerResetOffsetResponse
	GetStatusCode() *int32
	SetBody(v *OnsConsumerResetOffsetResponseBody) *OnsConsumerResetOffsetResponse
	GetBody() *OnsConsumerResetOffsetResponseBody
}

type OnsConsumerResetOffsetResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *OnsConsumerResetOffsetResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s OnsConsumerResetOffsetResponse) String() string {
	return dara.Prettify(s)
}

func (s OnsConsumerResetOffsetResponse) GoString() string {
	return s.String()
}

func (s *OnsConsumerResetOffsetResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *OnsConsumerResetOffsetResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *OnsConsumerResetOffsetResponse) GetBody() *OnsConsumerResetOffsetResponseBody {
	return s.Body
}

func (s *OnsConsumerResetOffsetResponse) SetHeaders(v map[string]*string) *OnsConsumerResetOffsetResponse {
	s.Headers = v
	return s
}

func (s *OnsConsumerResetOffsetResponse) SetStatusCode(v int32) *OnsConsumerResetOffsetResponse {
	s.StatusCode = &v
	return s
}

func (s *OnsConsumerResetOffsetResponse) SetBody(v *OnsConsumerResetOffsetResponseBody) *OnsConsumerResetOffsetResponse {
	s.Body = v
	return s
}

func (s *OnsConsumerResetOffsetResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
