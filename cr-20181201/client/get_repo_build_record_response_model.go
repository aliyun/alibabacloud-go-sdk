// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetRepoBuildRecordResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetRepoBuildRecordResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetRepoBuildRecordResponse
	GetStatusCode() *int32
	SetBody(v *GetRepoBuildRecordResponseBody) *GetRepoBuildRecordResponse
	GetBody() *GetRepoBuildRecordResponseBody
}

type GetRepoBuildRecordResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetRepoBuildRecordResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetRepoBuildRecordResponse) String() string {
	return dara.Prettify(s)
}

func (s GetRepoBuildRecordResponse) GoString() string {
	return s.String()
}

func (s *GetRepoBuildRecordResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetRepoBuildRecordResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetRepoBuildRecordResponse) GetBody() *GetRepoBuildRecordResponseBody {
	return s.Body
}

func (s *GetRepoBuildRecordResponse) SetHeaders(v map[string]*string) *GetRepoBuildRecordResponse {
	s.Headers = v
	return s
}

func (s *GetRepoBuildRecordResponse) SetStatusCode(v int32) *GetRepoBuildRecordResponse {
	s.StatusCode = &v
	return s
}

func (s *GetRepoBuildRecordResponse) SetBody(v *GetRepoBuildRecordResponseBody) *GetRepoBuildRecordResponse {
	s.Body = v
	return s
}

func (s *GetRepoBuildRecordResponse) Validate() error {
	return dara.Validate(s)
}
