// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQuerySmsQualificationRecordResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QuerySmsQualificationRecordResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QuerySmsQualificationRecordResponse
	GetStatusCode() *int32
	SetBody(v *QuerySmsQualificationRecordResponseBody) *QuerySmsQualificationRecordResponse
	GetBody() *QuerySmsQualificationRecordResponseBody
}

type QuerySmsQualificationRecordResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QuerySmsQualificationRecordResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QuerySmsQualificationRecordResponse) String() string {
	return dara.Prettify(s)
}

func (s QuerySmsQualificationRecordResponse) GoString() string {
	return s.String()
}

func (s *QuerySmsQualificationRecordResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QuerySmsQualificationRecordResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QuerySmsQualificationRecordResponse) GetBody() *QuerySmsQualificationRecordResponseBody {
	return s.Body
}

func (s *QuerySmsQualificationRecordResponse) SetHeaders(v map[string]*string) *QuerySmsQualificationRecordResponse {
	s.Headers = v
	return s
}

func (s *QuerySmsQualificationRecordResponse) SetStatusCode(v int32) *QuerySmsQualificationRecordResponse {
	s.StatusCode = &v
	return s
}

func (s *QuerySmsQualificationRecordResponse) SetBody(v *QuerySmsQualificationRecordResponseBody) *QuerySmsQualificationRecordResponse {
	s.Body = v
	return s
}

func (s *QuerySmsQualificationRecordResponse) Validate() error {
	return dara.Validate(s)
}
