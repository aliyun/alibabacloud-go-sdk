// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateRecordResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateRecordResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateRecordResponse
	GetStatusCode() *int32
	SetBody(v *CreateRecordResponseBody) *CreateRecordResponse
	GetBody() *CreateRecordResponseBody
}

type CreateRecordResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateRecordResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateRecordResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateRecordResponse) GoString() string {
	return s.String()
}

func (s *CreateRecordResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateRecordResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateRecordResponse) GetBody() *CreateRecordResponseBody {
	return s.Body
}

func (s *CreateRecordResponse) SetHeaders(v map[string]*string) *CreateRecordResponse {
	s.Headers = v
	return s
}

func (s *CreateRecordResponse) SetStatusCode(v int32) *CreateRecordResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateRecordResponse) SetBody(v *CreateRecordResponseBody) *CreateRecordResponse {
	s.Body = v
	return s
}

func (s *CreateRecordResponse) Validate() error {
	return dara.Validate(s)
}
