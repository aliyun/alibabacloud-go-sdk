// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddLiveRecordVodConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AddLiveRecordVodConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AddLiveRecordVodConfigResponse
	GetStatusCode() *int32
	SetBody(v *AddLiveRecordVodConfigResponseBody) *AddLiveRecordVodConfigResponse
	GetBody() *AddLiveRecordVodConfigResponseBody
}

type AddLiveRecordVodConfigResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddLiveRecordVodConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddLiveRecordVodConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s AddLiveRecordVodConfigResponse) GoString() string {
	return s.String()
}

func (s *AddLiveRecordVodConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AddLiveRecordVodConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AddLiveRecordVodConfigResponse) GetBody() *AddLiveRecordVodConfigResponseBody {
	return s.Body
}

func (s *AddLiveRecordVodConfigResponse) SetHeaders(v map[string]*string) *AddLiveRecordVodConfigResponse {
	s.Headers = v
	return s
}

func (s *AddLiveRecordVodConfigResponse) SetStatusCode(v int32) *AddLiveRecordVodConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *AddLiveRecordVodConfigResponse) SetBody(v *AddLiveRecordVodConfigResponseBody) *AddLiveRecordVodConfigResponse {
	s.Body = v
	return s
}

func (s *AddLiveRecordVodConfigResponse) Validate() error {
	return dara.Validate(s)
}
