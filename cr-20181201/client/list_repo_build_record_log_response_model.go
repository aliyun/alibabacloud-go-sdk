// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListRepoBuildRecordLogResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListRepoBuildRecordLogResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListRepoBuildRecordLogResponse
	GetStatusCode() *int32
	SetBody(v *ListRepoBuildRecordLogResponseBody) *ListRepoBuildRecordLogResponse
	GetBody() *ListRepoBuildRecordLogResponseBody
}

type ListRepoBuildRecordLogResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListRepoBuildRecordLogResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListRepoBuildRecordLogResponse) String() string {
	return dara.Prettify(s)
}

func (s ListRepoBuildRecordLogResponse) GoString() string {
	return s.String()
}

func (s *ListRepoBuildRecordLogResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListRepoBuildRecordLogResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListRepoBuildRecordLogResponse) GetBody() *ListRepoBuildRecordLogResponseBody {
	return s.Body
}

func (s *ListRepoBuildRecordLogResponse) SetHeaders(v map[string]*string) *ListRepoBuildRecordLogResponse {
	s.Headers = v
	return s
}

func (s *ListRepoBuildRecordLogResponse) SetStatusCode(v int32) *ListRepoBuildRecordLogResponse {
	s.StatusCode = &v
	return s
}

func (s *ListRepoBuildRecordLogResponse) SetBody(v *ListRepoBuildRecordLogResponseBody) *ListRepoBuildRecordLogResponse {
	s.Body = v
	return s
}

func (s *ListRepoBuildRecordLogResponse) Validate() error {
	return dara.Validate(s)
}
