// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddLiveAppRecordConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AddLiveAppRecordConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AddLiveAppRecordConfigResponse
	GetStatusCode() *int32
	SetBody(v *AddLiveAppRecordConfigResponseBody) *AddLiveAppRecordConfigResponse
	GetBody() *AddLiveAppRecordConfigResponseBody
}

type AddLiveAppRecordConfigResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddLiveAppRecordConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddLiveAppRecordConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s AddLiveAppRecordConfigResponse) GoString() string {
	return s.String()
}

func (s *AddLiveAppRecordConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AddLiveAppRecordConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AddLiveAppRecordConfigResponse) GetBody() *AddLiveAppRecordConfigResponseBody {
	return s.Body
}

func (s *AddLiveAppRecordConfigResponse) SetHeaders(v map[string]*string) *AddLiveAppRecordConfigResponse {
	s.Headers = v
	return s
}

func (s *AddLiveAppRecordConfigResponse) SetStatusCode(v int32) *AddLiveAppRecordConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *AddLiveAppRecordConfigResponse) SetBody(v *AddLiveAppRecordConfigResponseBody) *AddLiveAppRecordConfigResponse {
	s.Body = v
	return s
}

func (s *AddLiveAppRecordConfigResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
