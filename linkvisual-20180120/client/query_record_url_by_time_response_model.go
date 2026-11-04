// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryRecordUrlByTimeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryRecordUrlByTimeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryRecordUrlByTimeResponse
	GetStatusCode() *int32
	SetBody(v *QueryRecordUrlByTimeResponseBody) *QueryRecordUrlByTimeResponse
	GetBody() *QueryRecordUrlByTimeResponseBody
}

type QueryRecordUrlByTimeResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryRecordUrlByTimeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryRecordUrlByTimeResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryRecordUrlByTimeResponse) GoString() string {
	return s.String()
}

func (s *QueryRecordUrlByTimeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryRecordUrlByTimeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryRecordUrlByTimeResponse) GetBody() *QueryRecordUrlByTimeResponseBody {
	return s.Body
}

func (s *QueryRecordUrlByTimeResponse) SetHeaders(v map[string]*string) *QueryRecordUrlByTimeResponse {
	s.Headers = v
	return s
}

func (s *QueryRecordUrlByTimeResponse) SetStatusCode(v int32) *QueryRecordUrlByTimeResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryRecordUrlByTimeResponse) SetBody(v *QueryRecordUrlByTimeResponseBody) *QueryRecordUrlByTimeResponse {
	s.Body = v
	return s
}

func (s *QueryRecordUrlByTimeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
