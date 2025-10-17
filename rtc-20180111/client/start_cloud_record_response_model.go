// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartCloudRecordResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *StartCloudRecordResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *StartCloudRecordResponse
	GetStatusCode() *int32
	SetBody(v *StartCloudRecordResponseBody) *StartCloudRecordResponse
	GetBody() *StartCloudRecordResponseBody
}

type StartCloudRecordResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StartCloudRecordResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StartCloudRecordResponse) String() string {
	return dara.Prettify(s)
}

func (s StartCloudRecordResponse) GoString() string {
	return s.String()
}

func (s *StartCloudRecordResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *StartCloudRecordResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *StartCloudRecordResponse) GetBody() *StartCloudRecordResponseBody {
	return s.Body
}

func (s *StartCloudRecordResponse) SetHeaders(v map[string]*string) *StartCloudRecordResponse {
	s.Headers = v
	return s
}

func (s *StartCloudRecordResponse) SetStatusCode(v int32) *StartCloudRecordResponse {
	s.StatusCode = &v
	return s
}

func (s *StartCloudRecordResponse) SetBody(v *StartCloudRecordResponseBody) *StartCloudRecordResponse {
	s.Body = v
	return s
}

func (s *StartCloudRecordResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
