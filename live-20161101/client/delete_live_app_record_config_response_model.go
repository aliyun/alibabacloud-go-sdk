// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteLiveAppRecordConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteLiveAppRecordConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteLiveAppRecordConfigResponse
	GetStatusCode() *int32
	SetBody(v *DeleteLiveAppRecordConfigResponseBody) *DeleteLiveAppRecordConfigResponse
	GetBody() *DeleteLiveAppRecordConfigResponseBody
}

type DeleteLiveAppRecordConfigResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteLiveAppRecordConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteLiveAppRecordConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteLiveAppRecordConfigResponse) GoString() string {
	return s.String()
}

func (s *DeleteLiveAppRecordConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteLiveAppRecordConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteLiveAppRecordConfigResponse) GetBody() *DeleteLiveAppRecordConfigResponseBody {
	return s.Body
}

func (s *DeleteLiveAppRecordConfigResponse) SetHeaders(v map[string]*string) *DeleteLiveAppRecordConfigResponse {
	s.Headers = v
	return s
}

func (s *DeleteLiveAppRecordConfigResponse) SetStatusCode(v int32) *DeleteLiveAppRecordConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteLiveAppRecordConfigResponse) SetBody(v *DeleteLiveAppRecordConfigResponseBody) *DeleteLiveAppRecordConfigResponse {
	s.Body = v
	return s
}

func (s *DeleteLiveAppRecordConfigResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
