// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchCreateRecordsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *BatchCreateRecordsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *BatchCreateRecordsResponse
	GetStatusCode() *int32
	SetBody(v *BatchCreateRecordsResponseBody) *BatchCreateRecordsResponse
	GetBody() *BatchCreateRecordsResponseBody
}

type BatchCreateRecordsResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *BatchCreateRecordsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s BatchCreateRecordsResponse) String() string {
	return dara.Prettify(s)
}

func (s BatchCreateRecordsResponse) GoString() string {
	return s.String()
}

func (s *BatchCreateRecordsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *BatchCreateRecordsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *BatchCreateRecordsResponse) GetBody() *BatchCreateRecordsResponseBody {
	return s.Body
}

func (s *BatchCreateRecordsResponse) SetHeaders(v map[string]*string) *BatchCreateRecordsResponse {
	s.Headers = v
	return s
}

func (s *BatchCreateRecordsResponse) SetStatusCode(v int32) *BatchCreateRecordsResponse {
	s.StatusCode = &v
	return s
}

func (s *BatchCreateRecordsResponse) SetBody(v *BatchCreateRecordsResponseBody) *BatchCreateRecordsResponse {
	s.Body = v
	return s
}

func (s *BatchCreateRecordsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
