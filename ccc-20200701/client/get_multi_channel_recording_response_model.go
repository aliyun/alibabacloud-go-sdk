// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMultiChannelRecordingResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetMultiChannelRecordingResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetMultiChannelRecordingResponse
	GetStatusCode() *int32
	SetBody(v *GetMultiChannelRecordingResponseBody) *GetMultiChannelRecordingResponse
	GetBody() *GetMultiChannelRecordingResponseBody
}

type GetMultiChannelRecordingResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetMultiChannelRecordingResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetMultiChannelRecordingResponse) String() string {
	return dara.Prettify(s)
}

func (s GetMultiChannelRecordingResponse) GoString() string {
	return s.String()
}

func (s *GetMultiChannelRecordingResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetMultiChannelRecordingResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetMultiChannelRecordingResponse) GetBody() *GetMultiChannelRecordingResponseBody {
	return s.Body
}

func (s *GetMultiChannelRecordingResponse) SetHeaders(v map[string]*string) *GetMultiChannelRecordingResponse {
	s.Headers = v
	return s
}

func (s *GetMultiChannelRecordingResponse) SetStatusCode(v int32) *GetMultiChannelRecordingResponse {
	s.StatusCode = &v
	return s
}

func (s *GetMultiChannelRecordingResponse) SetBody(v *GetMultiChannelRecordingResponseBody) *GetMultiChannelRecordingResponse {
	s.Body = v
	return s
}

func (s *GetMultiChannelRecordingResponse) Validate() error {
	return dara.Validate(s)
}
