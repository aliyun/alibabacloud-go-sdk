// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetListRecordResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetListRecordResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetListRecordResponse
	GetStatusCode() *int32
	SetBody(v *GetListRecordResponseBody) *GetListRecordResponse
	GetBody() *GetListRecordResponseBody
}

type GetListRecordResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetListRecordResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetListRecordResponse) String() string {
	return dara.Prettify(s)
}

func (s GetListRecordResponse) GoString() string {
	return s.String()
}

func (s *GetListRecordResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetListRecordResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetListRecordResponse) GetBody() *GetListRecordResponseBody {
	return s.Body
}

func (s *GetListRecordResponse) SetHeaders(v map[string]*string) *GetListRecordResponse {
	s.Headers = v
	return s
}

func (s *GetListRecordResponse) SetStatusCode(v int32) *GetListRecordResponse {
	s.StatusCode = &v
	return s
}

func (s *GetListRecordResponse) SetBody(v *GetListRecordResponseBody) *GetListRecordResponse {
	s.Body = v
	return s
}

func (s *GetListRecordResponse) Validate() error {
	return dara.Validate(s)
}
