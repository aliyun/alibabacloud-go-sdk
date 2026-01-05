// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQuerySmsAppIcpRecordResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QuerySmsAppIcpRecordResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QuerySmsAppIcpRecordResponse
	GetStatusCode() *int32
	SetBody(v *QuerySmsAppIcpRecordResponseBody) *QuerySmsAppIcpRecordResponse
	GetBody() *QuerySmsAppIcpRecordResponseBody
}

type QuerySmsAppIcpRecordResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QuerySmsAppIcpRecordResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QuerySmsAppIcpRecordResponse) String() string {
	return dara.Prettify(s)
}

func (s QuerySmsAppIcpRecordResponse) GoString() string {
	return s.String()
}

func (s *QuerySmsAppIcpRecordResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QuerySmsAppIcpRecordResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QuerySmsAppIcpRecordResponse) GetBody() *QuerySmsAppIcpRecordResponseBody {
	return s.Body
}

func (s *QuerySmsAppIcpRecordResponse) SetHeaders(v map[string]*string) *QuerySmsAppIcpRecordResponse {
	s.Headers = v
	return s
}

func (s *QuerySmsAppIcpRecordResponse) SetStatusCode(v int32) *QuerySmsAppIcpRecordResponse {
	s.StatusCode = &v
	return s
}

func (s *QuerySmsAppIcpRecordResponse) SetBody(v *QuerySmsAppIcpRecordResponseBody) *QuerySmsAppIcpRecordResponse {
	s.Body = v
	return s
}

func (s *QuerySmsAppIcpRecordResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
