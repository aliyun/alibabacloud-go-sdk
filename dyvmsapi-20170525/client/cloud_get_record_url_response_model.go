// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCloudGetRecordUrlResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CloudGetRecordUrlResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CloudGetRecordUrlResponse
	GetStatusCode() *int32
	SetBody(v *CloudGetRecordUrlResponseBody) *CloudGetRecordUrlResponse
	GetBody() *CloudGetRecordUrlResponseBody
}

type CloudGetRecordUrlResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CloudGetRecordUrlResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CloudGetRecordUrlResponse) String() string {
	return dara.Prettify(s)
}

func (s CloudGetRecordUrlResponse) GoString() string {
	return s.String()
}

func (s *CloudGetRecordUrlResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CloudGetRecordUrlResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CloudGetRecordUrlResponse) GetBody() *CloudGetRecordUrlResponseBody {
	return s.Body
}

func (s *CloudGetRecordUrlResponse) SetHeaders(v map[string]*string) *CloudGetRecordUrlResponse {
	s.Headers = v
	return s
}

func (s *CloudGetRecordUrlResponse) SetStatusCode(v int32) *CloudGetRecordUrlResponse {
	s.StatusCode = &v
	return s
}

func (s *CloudGetRecordUrlResponse) SetBody(v *CloudGetRecordUrlResponseBody) *CloudGetRecordUrlResponse {
	s.Body = v
	return s
}

func (s *CloudGetRecordUrlResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
