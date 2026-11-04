// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryMonthRecordResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryMonthRecordResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryMonthRecordResponse
	GetStatusCode() *int32
	SetBody(v *QueryMonthRecordResponseBody) *QueryMonthRecordResponse
	GetBody() *QueryMonthRecordResponseBody
}

type QueryMonthRecordResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryMonthRecordResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryMonthRecordResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryMonthRecordResponse) GoString() string {
	return s.String()
}

func (s *QueryMonthRecordResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryMonthRecordResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryMonthRecordResponse) GetBody() *QueryMonthRecordResponseBody {
	return s.Body
}

func (s *QueryMonthRecordResponse) SetHeaders(v map[string]*string) *QueryMonthRecordResponse {
	s.Headers = v
	return s
}

func (s *QueryMonthRecordResponse) SetStatusCode(v int32) *QueryMonthRecordResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryMonthRecordResponse) SetBody(v *QueryMonthRecordResponseBody) *QueryMonthRecordResponse {
	s.Body = v
	return s
}

func (s *QueryMonthRecordResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
