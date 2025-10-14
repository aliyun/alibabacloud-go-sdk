// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchEventMigrateInstanceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *BatchEventMigrateInstanceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *BatchEventMigrateInstanceResponse
	GetStatusCode() *int32
	SetBody(v *BatchEventMigrateInstanceResponseBody) *BatchEventMigrateInstanceResponse
	GetBody() *BatchEventMigrateInstanceResponseBody
}

type BatchEventMigrateInstanceResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *BatchEventMigrateInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s BatchEventMigrateInstanceResponse) String() string {
	return dara.Prettify(s)
}

func (s BatchEventMigrateInstanceResponse) GoString() string {
	return s.String()
}

func (s *BatchEventMigrateInstanceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *BatchEventMigrateInstanceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *BatchEventMigrateInstanceResponse) GetBody() *BatchEventMigrateInstanceResponseBody {
	return s.Body
}

func (s *BatchEventMigrateInstanceResponse) SetHeaders(v map[string]*string) *BatchEventMigrateInstanceResponse {
	s.Headers = v
	return s
}

func (s *BatchEventMigrateInstanceResponse) SetStatusCode(v int32) *BatchEventMigrateInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *BatchEventMigrateInstanceResponse) SetBody(v *BatchEventMigrateInstanceResponseBody) *BatchEventMigrateInstanceResponse {
	s.Body = v
	return s
}

func (s *BatchEventMigrateInstanceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
