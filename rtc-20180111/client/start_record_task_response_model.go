// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartRecordTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *StartRecordTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *StartRecordTaskResponse
	GetStatusCode() *int32
	SetBody(v *StartRecordTaskResponseBody) *StartRecordTaskResponse
	GetBody() *StartRecordTaskResponseBody
}

type StartRecordTaskResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StartRecordTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StartRecordTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s StartRecordTaskResponse) GoString() string {
	return s.String()
}

func (s *StartRecordTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *StartRecordTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *StartRecordTaskResponse) GetBody() *StartRecordTaskResponseBody {
	return s.Body
}

func (s *StartRecordTaskResponse) SetHeaders(v map[string]*string) *StartRecordTaskResponse {
	s.Headers = v
	return s
}

func (s *StartRecordTaskResponse) SetStatusCode(v int32) *StartRecordTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *StartRecordTaskResponse) SetBody(v *StartRecordTaskResponseBody) *StartRecordTaskResponse {
	s.Body = v
	return s
}

func (s *StartRecordTaskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
