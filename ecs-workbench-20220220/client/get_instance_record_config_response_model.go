// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetInstanceRecordConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetInstanceRecordConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetInstanceRecordConfigResponse
	GetStatusCode() *int32
	SetBody(v *GetInstanceRecordConfigResponseBody) *GetInstanceRecordConfigResponse
	GetBody() *GetInstanceRecordConfigResponseBody
}

type GetInstanceRecordConfigResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetInstanceRecordConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetInstanceRecordConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s GetInstanceRecordConfigResponse) GoString() string {
	return s.String()
}

func (s *GetInstanceRecordConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetInstanceRecordConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetInstanceRecordConfigResponse) GetBody() *GetInstanceRecordConfigResponseBody {
	return s.Body
}

func (s *GetInstanceRecordConfigResponse) SetHeaders(v map[string]*string) *GetInstanceRecordConfigResponse {
	s.Headers = v
	return s
}

func (s *GetInstanceRecordConfigResponse) SetStatusCode(v int32) *GetInstanceRecordConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *GetInstanceRecordConfigResponse) SetBody(v *GetInstanceRecordConfigResponseBody) *GetInstanceRecordConfigResponse {
	s.Body = v
	return s
}

func (s *GetInstanceRecordConfigResponse) Validate() error {
	return dara.Validate(s)
}
