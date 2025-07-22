// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopCloudRecordResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *StopCloudRecordResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *StopCloudRecordResponse
	GetStatusCode() *int32
	SetBody(v *StopCloudRecordResponseBody) *StopCloudRecordResponse
	GetBody() *StopCloudRecordResponseBody
}

type StopCloudRecordResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StopCloudRecordResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StopCloudRecordResponse) String() string {
	return dara.Prettify(s)
}

func (s StopCloudRecordResponse) GoString() string {
	return s.String()
}

func (s *StopCloudRecordResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *StopCloudRecordResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *StopCloudRecordResponse) GetBody() *StopCloudRecordResponseBody {
	return s.Body
}

func (s *StopCloudRecordResponse) SetHeaders(v map[string]*string) *StopCloudRecordResponse {
	s.Headers = v
	return s
}

func (s *StopCloudRecordResponse) SetStatusCode(v int32) *StopCloudRecordResponse {
	s.StatusCode = &v
	return s
}

func (s *StopCloudRecordResponse) SetBody(v *StopCloudRecordResponseBody) *StopCloudRecordResponse {
	s.Body = v
	return s
}

func (s *StopCloudRecordResponse) Validate() error {
	return dara.Validate(s)
}
