// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchRevokePermissionsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *BatchRevokePermissionsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *BatchRevokePermissionsResponse
	GetStatusCode() *int32
	SetBody(v *BatchRevokePermissionsResponseBody) *BatchRevokePermissionsResponse
	GetBody() *BatchRevokePermissionsResponseBody
}

type BatchRevokePermissionsResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *BatchRevokePermissionsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s BatchRevokePermissionsResponse) String() string {
	return dara.Prettify(s)
}

func (s BatchRevokePermissionsResponse) GoString() string {
	return s.String()
}

func (s *BatchRevokePermissionsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *BatchRevokePermissionsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *BatchRevokePermissionsResponse) GetBody() *BatchRevokePermissionsResponseBody {
	return s.Body
}

func (s *BatchRevokePermissionsResponse) SetHeaders(v map[string]*string) *BatchRevokePermissionsResponse {
	s.Headers = v
	return s
}

func (s *BatchRevokePermissionsResponse) SetStatusCode(v int32) *BatchRevokePermissionsResponse {
	s.StatusCode = &v
	return s
}

func (s *BatchRevokePermissionsResponse) SetBody(v *BatchRevokePermissionsResponseBody) *BatchRevokePermissionsResponse {
	s.Body = v
	return s
}

func (s *BatchRevokePermissionsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
