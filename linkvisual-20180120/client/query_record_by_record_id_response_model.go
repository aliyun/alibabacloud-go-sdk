// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryRecordByRecordIdResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryRecordByRecordIdResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryRecordByRecordIdResponse
	GetStatusCode() *int32
	SetBody(v *QueryRecordByRecordIdResponseBody) *QueryRecordByRecordIdResponse
	GetBody() *QueryRecordByRecordIdResponseBody
}

type QueryRecordByRecordIdResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryRecordByRecordIdResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryRecordByRecordIdResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryRecordByRecordIdResponse) GoString() string {
	return s.String()
}

func (s *QueryRecordByRecordIdResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryRecordByRecordIdResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryRecordByRecordIdResponse) GetBody() *QueryRecordByRecordIdResponseBody {
	return s.Body
}

func (s *QueryRecordByRecordIdResponse) SetHeaders(v map[string]*string) *QueryRecordByRecordIdResponse {
	s.Headers = v
	return s
}

func (s *QueryRecordByRecordIdResponse) SetStatusCode(v int32) *QueryRecordByRecordIdResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryRecordByRecordIdResponse) SetBody(v *QueryRecordByRecordIdResponseBody) *QueryRecordByRecordIdResponse {
	s.Body = v
	return s
}

func (s *QueryRecordByRecordIdResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
