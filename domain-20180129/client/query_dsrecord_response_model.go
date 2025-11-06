// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryDSRecordResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryDSRecordResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryDSRecordResponse
	GetStatusCode() *int32
	SetBody(v *QueryDSRecordResponseBody) *QueryDSRecordResponse
	GetBody() *QueryDSRecordResponseBody
}

type QueryDSRecordResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryDSRecordResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryDSRecordResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryDSRecordResponse) GoString() string {
	return s.String()
}

func (s *QueryDSRecordResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryDSRecordResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryDSRecordResponse) GetBody() *QueryDSRecordResponseBody {
	return s.Body
}

func (s *QueryDSRecordResponse) SetHeaders(v map[string]*string) *QueryDSRecordResponse {
	s.Headers = v
	return s
}

func (s *QueryDSRecordResponse) SetStatusCode(v int32) *QueryDSRecordResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryDSRecordResponse) SetBody(v *QueryDSRecordResponseBody) *QueryDSRecordResponse {
	s.Body = v
	return s
}

func (s *QueryDSRecordResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
