// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryCloudRecordVideoPlayInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryCloudRecordVideoPlayInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryCloudRecordVideoPlayInfoResponse
	GetStatusCode() *int32
	SetBody(v *QueryCloudRecordVideoPlayInfoResponseBody) *QueryCloudRecordVideoPlayInfoResponse
	GetBody() *QueryCloudRecordVideoPlayInfoResponseBody
}

type QueryCloudRecordVideoPlayInfoResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryCloudRecordVideoPlayInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryCloudRecordVideoPlayInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryCloudRecordVideoPlayInfoResponse) GoString() string {
	return s.String()
}

func (s *QueryCloudRecordVideoPlayInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryCloudRecordVideoPlayInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryCloudRecordVideoPlayInfoResponse) GetBody() *QueryCloudRecordVideoPlayInfoResponseBody {
	return s.Body
}

func (s *QueryCloudRecordVideoPlayInfoResponse) SetHeaders(v map[string]*string) *QueryCloudRecordVideoPlayInfoResponse {
	s.Headers = v
	return s
}

func (s *QueryCloudRecordVideoPlayInfoResponse) SetStatusCode(v int32) *QueryCloudRecordVideoPlayInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryCloudRecordVideoPlayInfoResponse) SetBody(v *QueryCloudRecordVideoPlayInfoResponseBody) *QueryCloudRecordVideoPlayInfoResponse {
	s.Body = v
	return s
}

func (s *QueryCloudRecordVideoPlayInfoResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
