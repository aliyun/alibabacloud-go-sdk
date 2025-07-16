// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchGrantPermissionsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *BatchGrantPermissionsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *BatchGrantPermissionsResponse
	GetStatusCode() *int32
	SetBody(v *BatchGrantPermissionsResponseBody) *BatchGrantPermissionsResponse
	GetBody() *BatchGrantPermissionsResponseBody
}

type BatchGrantPermissionsResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *BatchGrantPermissionsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s BatchGrantPermissionsResponse) String() string {
	return dara.Prettify(s)
}

func (s BatchGrantPermissionsResponse) GoString() string {
	return s.String()
}

func (s *BatchGrantPermissionsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *BatchGrantPermissionsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *BatchGrantPermissionsResponse) GetBody() *BatchGrantPermissionsResponseBody {
	return s.Body
}

func (s *BatchGrantPermissionsResponse) SetHeaders(v map[string]*string) *BatchGrantPermissionsResponse {
	s.Headers = v
	return s
}

func (s *BatchGrantPermissionsResponse) SetStatusCode(v int32) *BatchGrantPermissionsResponse {
	s.StatusCode = &v
	return s
}

func (s *BatchGrantPermissionsResponse) SetBody(v *BatchGrantPermissionsResponseBody) *BatchGrantPermissionsResponse {
	s.Body = v
	return s
}

func (s *BatchGrantPermissionsResponse) Validate() error {
	return dara.Validate(s)
}
