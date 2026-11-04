// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryEventRecordPlansResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryEventRecordPlansResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryEventRecordPlansResponse
	GetStatusCode() *int32
	SetBody(v *QueryEventRecordPlansResponseBody) *QueryEventRecordPlansResponse
	GetBody() *QueryEventRecordPlansResponseBody
}

type QueryEventRecordPlansResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryEventRecordPlansResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryEventRecordPlansResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryEventRecordPlansResponse) GoString() string {
	return s.String()
}

func (s *QueryEventRecordPlansResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryEventRecordPlansResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryEventRecordPlansResponse) GetBody() *QueryEventRecordPlansResponseBody {
	return s.Body
}

func (s *QueryEventRecordPlansResponse) SetHeaders(v map[string]*string) *QueryEventRecordPlansResponse {
	s.Headers = v
	return s
}

func (s *QueryEventRecordPlansResponse) SetStatusCode(v int32) *QueryEventRecordPlansResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryEventRecordPlansResponse) SetBody(v *QueryEventRecordPlansResponseBody) *QueryEventRecordPlansResponse {
	s.Body = v
	return s
}

func (s *QueryEventRecordPlansResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
