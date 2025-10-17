// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateCloudRecordResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateCloudRecordResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateCloudRecordResponse
	GetStatusCode() *int32
	SetBody(v *UpdateCloudRecordResponseBody) *UpdateCloudRecordResponse
	GetBody() *UpdateCloudRecordResponseBody
}

type UpdateCloudRecordResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateCloudRecordResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateCloudRecordResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateCloudRecordResponse) GoString() string {
	return s.String()
}

func (s *UpdateCloudRecordResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateCloudRecordResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateCloudRecordResponse) GetBody() *UpdateCloudRecordResponseBody {
	return s.Body
}

func (s *UpdateCloudRecordResponse) SetHeaders(v map[string]*string) *UpdateCloudRecordResponse {
	s.Headers = v
	return s
}

func (s *UpdateCloudRecordResponse) SetStatusCode(v int32) *UpdateCloudRecordResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateCloudRecordResponse) SetBody(v *UpdateCloudRecordResponseBody) *UpdateCloudRecordResponse {
	s.Body = v
	return s
}

func (s *UpdateCloudRecordResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
