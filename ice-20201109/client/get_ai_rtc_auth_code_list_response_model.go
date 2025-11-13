// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAiRtcAuthCodeListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetAiRtcAuthCodeListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetAiRtcAuthCodeListResponse
	GetStatusCode() *int32
	SetBody(v *GetAiRtcAuthCodeListResponseBody) *GetAiRtcAuthCodeListResponse
	GetBody() *GetAiRtcAuthCodeListResponseBody
}

type GetAiRtcAuthCodeListResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAiRtcAuthCodeListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAiRtcAuthCodeListResponse) String() string {
	return dara.Prettify(s)
}

func (s GetAiRtcAuthCodeListResponse) GoString() string {
	return s.String()
}

func (s *GetAiRtcAuthCodeListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetAiRtcAuthCodeListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetAiRtcAuthCodeListResponse) GetBody() *GetAiRtcAuthCodeListResponseBody {
	return s.Body
}

func (s *GetAiRtcAuthCodeListResponse) SetHeaders(v map[string]*string) *GetAiRtcAuthCodeListResponse {
	s.Headers = v
	return s
}

func (s *GetAiRtcAuthCodeListResponse) SetStatusCode(v int32) *GetAiRtcAuthCodeListResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAiRtcAuthCodeListResponse) SetBody(v *GetAiRtcAuthCodeListResponseBody) *GetAiRtcAuthCodeListResponse {
	s.Body = v
	return s
}

func (s *GetAiRtcAuthCodeListResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
