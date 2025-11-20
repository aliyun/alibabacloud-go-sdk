// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryCloudRecordVideoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryCloudRecordVideoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryCloudRecordVideoResponse
	GetStatusCode() *int32
	SetBody(v *QueryCloudRecordVideoResponseBody) *QueryCloudRecordVideoResponse
	GetBody() *QueryCloudRecordVideoResponseBody
}

type QueryCloudRecordVideoResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryCloudRecordVideoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryCloudRecordVideoResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryCloudRecordVideoResponse) GoString() string {
	return s.String()
}

func (s *QueryCloudRecordVideoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryCloudRecordVideoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryCloudRecordVideoResponse) GetBody() *QueryCloudRecordVideoResponseBody {
	return s.Body
}

func (s *QueryCloudRecordVideoResponse) SetHeaders(v map[string]*string) *QueryCloudRecordVideoResponse {
	s.Headers = v
	return s
}

func (s *QueryCloudRecordVideoResponse) SetStatusCode(v int32) *QueryCloudRecordVideoResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryCloudRecordVideoResponse) SetBody(v *QueryCloudRecordVideoResponseBody) *QueryCloudRecordVideoResponse {
	s.Body = v
	return s
}

func (s *QueryCloudRecordVideoResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
