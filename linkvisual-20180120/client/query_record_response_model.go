// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryRecordResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryRecordResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryRecordResponse
	GetStatusCode() *int32
	SetBody(v *QueryRecordResponseBody) *QueryRecordResponse
	GetBody() *QueryRecordResponseBody
}

type QueryRecordResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryRecordResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryRecordResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryRecordResponse) GoString() string {
	return s.String()
}

func (s *QueryRecordResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryRecordResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryRecordResponse) GetBody() *QueryRecordResponseBody {
	return s.Body
}

func (s *QueryRecordResponse) SetHeaders(v map[string]*string) *QueryRecordResponse {
	s.Headers = v
	return s
}

func (s *QueryRecordResponse) SetStatusCode(v int32) *QueryRecordResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryRecordResponse) SetBody(v *QueryRecordResponseBody) *QueryRecordResponse {
	s.Body = v
	return s
}

func (s *QueryRecordResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
