// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitLiveRecordJobResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SubmitLiveRecordJobResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SubmitLiveRecordJobResponse
	GetStatusCode() *int32
	SetBody(v *SubmitLiveRecordJobResponseBody) *SubmitLiveRecordJobResponse
	GetBody() *SubmitLiveRecordJobResponseBody
}

type SubmitLiveRecordJobResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SubmitLiveRecordJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SubmitLiveRecordJobResponse) String() string {
	return dara.Prettify(s)
}

func (s SubmitLiveRecordJobResponse) GoString() string {
	return s.String()
}

func (s *SubmitLiveRecordJobResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SubmitLiveRecordJobResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SubmitLiveRecordJobResponse) GetBody() *SubmitLiveRecordJobResponseBody {
	return s.Body
}

func (s *SubmitLiveRecordJobResponse) SetHeaders(v map[string]*string) *SubmitLiveRecordJobResponse {
	s.Headers = v
	return s
}

func (s *SubmitLiveRecordJobResponse) SetStatusCode(v int32) *SubmitLiveRecordJobResponse {
	s.StatusCode = &v
	return s
}

func (s *SubmitLiveRecordJobResponse) SetBody(v *SubmitLiveRecordJobResponseBody) *SubmitLiveRecordJobResponse {
	s.Body = v
	return s
}

func (s *SubmitLiveRecordJobResponse) Validate() error {
	return dara.Validate(s)
}
