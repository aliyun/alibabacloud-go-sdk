// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListRepoBuildRecordResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListRepoBuildRecordResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListRepoBuildRecordResponse
	GetStatusCode() *int32
	SetBody(v *ListRepoBuildRecordResponseBody) *ListRepoBuildRecordResponse
	GetBody() *ListRepoBuildRecordResponseBody
}

type ListRepoBuildRecordResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListRepoBuildRecordResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListRepoBuildRecordResponse) String() string {
	return dara.Prettify(s)
}

func (s ListRepoBuildRecordResponse) GoString() string {
	return s.String()
}

func (s *ListRepoBuildRecordResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListRepoBuildRecordResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListRepoBuildRecordResponse) GetBody() *ListRepoBuildRecordResponseBody {
	return s.Body
}

func (s *ListRepoBuildRecordResponse) SetHeaders(v map[string]*string) *ListRepoBuildRecordResponse {
	s.Headers = v
	return s
}

func (s *ListRepoBuildRecordResponse) SetStatusCode(v int32) *ListRepoBuildRecordResponse {
	s.StatusCode = &v
	return s
}

func (s *ListRepoBuildRecordResponse) SetBody(v *ListRepoBuildRecordResponseBody) *ListRepoBuildRecordResponse {
	s.Body = v
	return s
}

func (s *ListRepoBuildRecordResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
