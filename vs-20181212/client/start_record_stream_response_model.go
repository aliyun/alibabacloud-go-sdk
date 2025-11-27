// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartRecordStreamResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *StartRecordStreamResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *StartRecordStreamResponse
	GetStatusCode() *int32
	SetBody(v *StartRecordStreamResponseBody) *StartRecordStreamResponse
	GetBody() *StartRecordStreamResponseBody
}

type StartRecordStreamResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StartRecordStreamResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StartRecordStreamResponse) String() string {
	return dara.Prettify(s)
}

func (s StartRecordStreamResponse) GoString() string {
	return s.String()
}

func (s *StartRecordStreamResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *StartRecordStreamResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *StartRecordStreamResponse) GetBody() *StartRecordStreamResponseBody {
	return s.Body
}

func (s *StartRecordStreamResponse) SetHeaders(v map[string]*string) *StartRecordStreamResponse {
	s.Headers = v
	return s
}

func (s *StartRecordStreamResponse) SetStatusCode(v int32) *StartRecordStreamResponse {
	s.StatusCode = &v
	return s
}

func (s *StartRecordStreamResponse) SetBody(v *StartRecordStreamResponseBody) *StartRecordStreamResponse {
	s.Body = v
	return s
}

func (s *StartRecordStreamResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
