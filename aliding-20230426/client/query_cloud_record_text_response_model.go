// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryCloudRecordTextResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryCloudRecordTextResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryCloudRecordTextResponse
	GetStatusCode() *int32
	SetBody(v *QueryCloudRecordTextResponseBody) *QueryCloudRecordTextResponse
	GetBody() *QueryCloudRecordTextResponseBody
}

type QueryCloudRecordTextResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryCloudRecordTextResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryCloudRecordTextResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryCloudRecordTextResponse) GoString() string {
	return s.String()
}

func (s *QueryCloudRecordTextResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryCloudRecordTextResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryCloudRecordTextResponse) GetBody() *QueryCloudRecordTextResponseBody {
	return s.Body
}

func (s *QueryCloudRecordTextResponse) SetHeaders(v map[string]*string) *QueryCloudRecordTextResponse {
	s.Headers = v
	return s
}

func (s *QueryCloudRecordTextResponse) SetStatusCode(v int32) *QueryCloudRecordTextResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryCloudRecordTextResponse) SetBody(v *QueryCloudRecordTextResponseBody) *QueryCloudRecordTextResponse {
	s.Body = v
	return s
}

func (s *QueryCloudRecordTextResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
