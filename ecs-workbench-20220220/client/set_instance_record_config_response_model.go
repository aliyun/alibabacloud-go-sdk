// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetInstanceRecordConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SetInstanceRecordConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SetInstanceRecordConfigResponse
	GetStatusCode() *int32
	SetBody(v *SetInstanceRecordConfigResponseBody) *SetInstanceRecordConfigResponse
	GetBody() *SetInstanceRecordConfigResponseBody
}

type SetInstanceRecordConfigResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SetInstanceRecordConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SetInstanceRecordConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s SetInstanceRecordConfigResponse) GoString() string {
	return s.String()
}

func (s *SetInstanceRecordConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SetInstanceRecordConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SetInstanceRecordConfigResponse) GetBody() *SetInstanceRecordConfigResponseBody {
	return s.Body
}

func (s *SetInstanceRecordConfigResponse) SetHeaders(v map[string]*string) *SetInstanceRecordConfigResponse {
	s.Headers = v
	return s
}

func (s *SetInstanceRecordConfigResponse) SetStatusCode(v int32) *SetInstanceRecordConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *SetInstanceRecordConfigResponse) SetBody(v *SetInstanceRecordConfigResponseBody) *SetInstanceRecordConfigResponse {
	s.Body = v
	return s
}

func (s *SetInstanceRecordConfigResponse) Validate() error {
	return dara.Validate(s)
}
