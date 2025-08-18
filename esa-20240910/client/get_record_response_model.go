// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetRecordResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetRecordResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetRecordResponse
	GetStatusCode() *int32
	SetBody(v *GetRecordResponseBody) *GetRecordResponse
	GetBody() *GetRecordResponseBody
}

type GetRecordResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetRecordResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetRecordResponse) String() string {
	return dara.Prettify(s)
}

func (s GetRecordResponse) GoString() string {
	return s.String()
}

func (s *GetRecordResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetRecordResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetRecordResponse) GetBody() *GetRecordResponseBody {
	return s.Body
}

func (s *GetRecordResponse) SetHeaders(v map[string]*string) *GetRecordResponse {
	s.Headers = v
	return s
}

func (s *GetRecordResponse) SetStatusCode(v int32) *GetRecordResponse {
	s.StatusCode = &v
	return s
}

func (s *GetRecordResponse) SetBody(v *GetRecordResponseBody) *GetRecordResponse {
	s.Body = v
	return s
}

func (s *GetRecordResponse) Validate() error {
	return dara.Validate(s)
}
