// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRecordPostBackResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RecordPostBackResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RecordPostBackResponse
	GetStatusCode() *int32
	SetBody(v *RecordPostBackResponseBody) *RecordPostBackResponse
	GetBody() *RecordPostBackResponseBody
}

type RecordPostBackResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RecordPostBackResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RecordPostBackResponse) String() string {
	return dara.Prettify(s)
}

func (s RecordPostBackResponse) GoString() string {
	return s.String()
}

func (s *RecordPostBackResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RecordPostBackResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RecordPostBackResponse) GetBody() *RecordPostBackResponseBody {
	return s.Body
}

func (s *RecordPostBackResponse) SetHeaders(v map[string]*string) *RecordPostBackResponse {
	s.Headers = v
	return s
}

func (s *RecordPostBackResponse) SetStatusCode(v int32) *RecordPostBackResponse {
	s.StatusCode = &v
	return s
}

func (s *RecordPostBackResponse) SetBody(v *RecordPostBackResponseBody) *RecordPostBackResponse {
	s.Body = v
	return s
}

func (s *RecordPostBackResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
