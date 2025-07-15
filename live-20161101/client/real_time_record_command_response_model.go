// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRealTimeRecordCommandResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RealTimeRecordCommandResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RealTimeRecordCommandResponse
	GetStatusCode() *int32
	SetBody(v *RealTimeRecordCommandResponseBody) *RealTimeRecordCommandResponse
	GetBody() *RealTimeRecordCommandResponseBody
}

type RealTimeRecordCommandResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RealTimeRecordCommandResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RealTimeRecordCommandResponse) String() string {
	return dara.Prettify(s)
}

func (s RealTimeRecordCommandResponse) GoString() string {
	return s.String()
}

func (s *RealTimeRecordCommandResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RealTimeRecordCommandResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RealTimeRecordCommandResponse) GetBody() *RealTimeRecordCommandResponseBody {
	return s.Body
}

func (s *RealTimeRecordCommandResponse) SetHeaders(v map[string]*string) *RealTimeRecordCommandResponse {
	s.Headers = v
	return s
}

func (s *RealTimeRecordCommandResponse) SetStatusCode(v int32) *RealTimeRecordCommandResponse {
	s.StatusCode = &v
	return s
}

func (s *RealTimeRecordCommandResponse) SetBody(v *RealTimeRecordCommandResponseBody) *RealTimeRecordCommandResponse {
	s.Body = v
	return s
}

func (s *RealTimeRecordCommandResponse) Validate() error {
	return dara.Validate(s)
}
