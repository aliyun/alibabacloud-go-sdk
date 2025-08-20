// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCancelRepoBuildRecordResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CancelRepoBuildRecordResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CancelRepoBuildRecordResponse
	GetStatusCode() *int32
	SetBody(v *CancelRepoBuildRecordResponseBody) *CancelRepoBuildRecordResponse
	GetBody() *CancelRepoBuildRecordResponseBody
}

type CancelRepoBuildRecordResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CancelRepoBuildRecordResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CancelRepoBuildRecordResponse) String() string {
	return dara.Prettify(s)
}

func (s CancelRepoBuildRecordResponse) GoString() string {
	return s.String()
}

func (s *CancelRepoBuildRecordResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CancelRepoBuildRecordResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CancelRepoBuildRecordResponse) GetBody() *CancelRepoBuildRecordResponseBody {
	return s.Body
}

func (s *CancelRepoBuildRecordResponse) SetHeaders(v map[string]*string) *CancelRepoBuildRecordResponse {
	s.Headers = v
	return s
}

func (s *CancelRepoBuildRecordResponse) SetStatusCode(v int32) *CancelRepoBuildRecordResponse {
	s.StatusCode = &v
	return s
}

func (s *CancelRepoBuildRecordResponse) SetBody(v *CancelRepoBuildRecordResponseBody) *CancelRepoBuildRecordResponse {
	s.Body = v
	return s
}

func (s *CancelRepoBuildRecordResponse) Validate() error {
	return dara.Validate(s)
}
