// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCallDetailRecordResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetCallDetailRecordResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetCallDetailRecordResponse
	GetStatusCode() *int32
	SetBody(v *GetCallDetailRecordResponseBody) *GetCallDetailRecordResponse
	GetBody() *GetCallDetailRecordResponseBody
}

type GetCallDetailRecordResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetCallDetailRecordResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetCallDetailRecordResponse) String() string {
	return dara.Prettify(s)
}

func (s GetCallDetailRecordResponse) GoString() string {
	return s.String()
}

func (s *GetCallDetailRecordResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetCallDetailRecordResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetCallDetailRecordResponse) GetBody() *GetCallDetailRecordResponseBody {
	return s.Body
}

func (s *GetCallDetailRecordResponse) SetHeaders(v map[string]*string) *GetCallDetailRecordResponse {
	s.Headers = v
	return s
}

func (s *GetCallDetailRecordResponse) SetStatusCode(v int32) *GetCallDetailRecordResponse {
	s.StatusCode = &v
	return s
}

func (s *GetCallDetailRecordResponse) SetBody(v *GetCallDetailRecordResponseBody) *GetCallDetailRecordResponse {
	s.Body = v
	return s
}

func (s *GetCallDetailRecordResponse) Validate() error {
	return dara.Validate(s)
}
