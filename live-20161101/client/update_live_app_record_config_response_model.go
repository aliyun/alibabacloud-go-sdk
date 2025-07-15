// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateLiveAppRecordConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateLiveAppRecordConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateLiveAppRecordConfigResponse
	GetStatusCode() *int32
	SetBody(v *UpdateLiveAppRecordConfigResponseBody) *UpdateLiveAppRecordConfigResponse
	GetBody() *UpdateLiveAppRecordConfigResponseBody
}

type UpdateLiveAppRecordConfigResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateLiveAppRecordConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateLiveAppRecordConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateLiveAppRecordConfigResponse) GoString() string {
	return s.String()
}

func (s *UpdateLiveAppRecordConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateLiveAppRecordConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateLiveAppRecordConfigResponse) GetBody() *UpdateLiveAppRecordConfigResponseBody {
	return s.Body
}

func (s *UpdateLiveAppRecordConfigResponse) SetHeaders(v map[string]*string) *UpdateLiveAppRecordConfigResponse {
	s.Headers = v
	return s
}

func (s *UpdateLiveAppRecordConfigResponse) SetStatusCode(v int32) *UpdateLiveAppRecordConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateLiveAppRecordConfigResponse) SetBody(v *UpdateLiveAppRecordConfigResponseBody) *UpdateLiveAppRecordConfigResponse {
	s.Body = v
	return s
}

func (s *UpdateLiveAppRecordConfigResponse) Validate() error {
	return dara.Validate(s)
}
