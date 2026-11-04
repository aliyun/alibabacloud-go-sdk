// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryRecordPlansResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryRecordPlansResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryRecordPlansResponse
	GetStatusCode() *int32
	SetBody(v *QueryRecordPlansResponseBody) *QueryRecordPlansResponse
	GetBody() *QueryRecordPlansResponseBody
}

type QueryRecordPlansResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryRecordPlansResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryRecordPlansResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryRecordPlansResponse) GoString() string {
	return s.String()
}

func (s *QueryRecordPlansResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryRecordPlansResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryRecordPlansResponse) GetBody() *QueryRecordPlansResponseBody {
	return s.Body
}

func (s *QueryRecordPlansResponse) SetHeaders(v map[string]*string) *QueryRecordPlansResponse {
	s.Headers = v
	return s
}

func (s *QueryRecordPlansResponse) SetStatusCode(v int32) *QueryRecordPlansResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryRecordPlansResponse) SetBody(v *QueryRecordPlansResponseBody) *QueryRecordPlansResponse {
	s.Body = v
	return s
}

func (s *QueryRecordPlansResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
