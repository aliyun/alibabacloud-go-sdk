// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAutoDisposeRecordResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateAutoDisposeRecordResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateAutoDisposeRecordResponse
	GetStatusCode() *int32
	SetBody(v *UpdateAutoDisposeRecordResponseBody) *UpdateAutoDisposeRecordResponse
	GetBody() *UpdateAutoDisposeRecordResponseBody
}

type UpdateAutoDisposeRecordResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateAutoDisposeRecordResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateAutoDisposeRecordResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateAutoDisposeRecordResponse) GoString() string {
	return s.String()
}

func (s *UpdateAutoDisposeRecordResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateAutoDisposeRecordResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateAutoDisposeRecordResponse) GetBody() *UpdateAutoDisposeRecordResponseBody {
	return s.Body
}

func (s *UpdateAutoDisposeRecordResponse) SetHeaders(v map[string]*string) *UpdateAutoDisposeRecordResponse {
	s.Headers = v
	return s
}

func (s *UpdateAutoDisposeRecordResponse) SetStatusCode(v int32) *UpdateAutoDisposeRecordResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateAutoDisposeRecordResponse) SetBody(v *UpdateAutoDisposeRecordResponseBody) *UpdateAutoDisposeRecordResponse {
	s.Body = v
	return s
}

func (s *UpdateAutoDisposeRecordResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
