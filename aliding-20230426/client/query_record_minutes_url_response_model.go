// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryRecordMinutesUrlResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryRecordMinutesUrlResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryRecordMinutesUrlResponse
	GetStatusCode() *int32
	SetBody(v *QueryRecordMinutesUrlResponseBody) *QueryRecordMinutesUrlResponse
	GetBody() *QueryRecordMinutesUrlResponseBody
}

type QueryRecordMinutesUrlResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryRecordMinutesUrlResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryRecordMinutesUrlResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryRecordMinutesUrlResponse) GoString() string {
	return s.String()
}

func (s *QueryRecordMinutesUrlResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryRecordMinutesUrlResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryRecordMinutesUrlResponse) GetBody() *QueryRecordMinutesUrlResponseBody {
	return s.Body
}

func (s *QueryRecordMinutesUrlResponse) SetHeaders(v map[string]*string) *QueryRecordMinutesUrlResponse {
	s.Headers = v
	return s
}

func (s *QueryRecordMinutesUrlResponse) SetStatusCode(v int32) *QueryRecordMinutesUrlResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryRecordMinutesUrlResponse) SetBody(v *QueryRecordMinutesUrlResponseBody) *QueryRecordMinutesUrlResponse {
	s.Body = v
	return s
}

func (s *QueryRecordMinutesUrlResponse) Validate() error {
	return dara.Validate(s)
}
