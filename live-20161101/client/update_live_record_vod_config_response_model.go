// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateLiveRecordVodConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateLiveRecordVodConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateLiveRecordVodConfigResponse
	GetStatusCode() *int32
	SetBody(v *UpdateLiveRecordVodConfigResponseBody) *UpdateLiveRecordVodConfigResponse
	GetBody() *UpdateLiveRecordVodConfigResponseBody
}

type UpdateLiveRecordVodConfigResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateLiveRecordVodConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateLiveRecordVodConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateLiveRecordVodConfigResponse) GoString() string {
	return s.String()
}

func (s *UpdateLiveRecordVodConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateLiveRecordVodConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateLiveRecordVodConfigResponse) GetBody() *UpdateLiveRecordVodConfigResponseBody {
	return s.Body
}

func (s *UpdateLiveRecordVodConfigResponse) SetHeaders(v map[string]*string) *UpdateLiveRecordVodConfigResponse {
	s.Headers = v
	return s
}

func (s *UpdateLiveRecordVodConfigResponse) SetStatusCode(v int32) *UpdateLiveRecordVodConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateLiveRecordVodConfigResponse) SetBody(v *UpdateLiveRecordVodConfigResponseBody) *UpdateLiveRecordVodConfigResponse {
	s.Body = v
	return s
}

func (s *UpdateLiveRecordVodConfigResponse) Validate() error {
	return dara.Validate(s)
}
