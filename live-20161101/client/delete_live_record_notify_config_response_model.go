// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteLiveRecordNotifyConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteLiveRecordNotifyConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteLiveRecordNotifyConfigResponse
	GetStatusCode() *int32
	SetBody(v *DeleteLiveRecordNotifyConfigResponseBody) *DeleteLiveRecordNotifyConfigResponse
	GetBody() *DeleteLiveRecordNotifyConfigResponseBody
}

type DeleteLiveRecordNotifyConfigResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteLiveRecordNotifyConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteLiveRecordNotifyConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteLiveRecordNotifyConfigResponse) GoString() string {
	return s.String()
}

func (s *DeleteLiveRecordNotifyConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteLiveRecordNotifyConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteLiveRecordNotifyConfigResponse) GetBody() *DeleteLiveRecordNotifyConfigResponseBody {
	return s.Body
}

func (s *DeleteLiveRecordNotifyConfigResponse) SetHeaders(v map[string]*string) *DeleteLiveRecordNotifyConfigResponse {
	s.Headers = v
	return s
}

func (s *DeleteLiveRecordNotifyConfigResponse) SetStatusCode(v int32) *DeleteLiveRecordNotifyConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteLiveRecordNotifyConfigResponse) SetBody(v *DeleteLiveRecordNotifyConfigResponseBody) *DeleteLiveRecordNotifyConfigResponse {
	s.Body = v
	return s
}

func (s *DeleteLiveRecordNotifyConfigResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
