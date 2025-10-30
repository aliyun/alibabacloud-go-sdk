// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iHandleActiveSQLRecordResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *HandleActiveSQLRecordResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *HandleActiveSQLRecordResponse
	GetStatusCode() *int32
	SetBody(v *HandleActiveSQLRecordResponseBody) *HandleActiveSQLRecordResponse
	GetBody() *HandleActiveSQLRecordResponseBody
}

type HandleActiveSQLRecordResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *HandleActiveSQLRecordResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s HandleActiveSQLRecordResponse) String() string {
	return dara.Prettify(s)
}

func (s HandleActiveSQLRecordResponse) GoString() string {
	return s.String()
}

func (s *HandleActiveSQLRecordResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *HandleActiveSQLRecordResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *HandleActiveSQLRecordResponse) GetBody() *HandleActiveSQLRecordResponseBody {
	return s.Body
}

func (s *HandleActiveSQLRecordResponse) SetHeaders(v map[string]*string) *HandleActiveSQLRecordResponse {
	s.Headers = v
	return s
}

func (s *HandleActiveSQLRecordResponse) SetStatusCode(v int32) *HandleActiveSQLRecordResponse {
	s.StatusCode = &v
	return s
}

func (s *HandleActiveSQLRecordResponse) SetBody(v *HandleActiveSQLRecordResponseBody) *HandleActiveSQLRecordResponse {
	s.Body = v
	return s
}

func (s *HandleActiveSQLRecordResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
