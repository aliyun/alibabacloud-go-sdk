// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListBroadcastAudiosByIdResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListBroadcastAudiosByIdResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListBroadcastAudiosByIdResponse
	GetStatusCode() *int32
	SetBody(v *ListBroadcastAudiosByIdResponseBody) *ListBroadcastAudiosByIdResponse
	GetBody() *ListBroadcastAudiosByIdResponseBody
}

type ListBroadcastAudiosByIdResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListBroadcastAudiosByIdResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListBroadcastAudiosByIdResponse) String() string {
	return dara.Prettify(s)
}

func (s ListBroadcastAudiosByIdResponse) GoString() string {
	return s.String()
}

func (s *ListBroadcastAudiosByIdResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListBroadcastAudiosByIdResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListBroadcastAudiosByIdResponse) GetBody() *ListBroadcastAudiosByIdResponseBody {
	return s.Body
}

func (s *ListBroadcastAudiosByIdResponse) SetHeaders(v map[string]*string) *ListBroadcastAudiosByIdResponse {
	s.Headers = v
	return s
}

func (s *ListBroadcastAudiosByIdResponse) SetStatusCode(v int32) *ListBroadcastAudiosByIdResponse {
	s.StatusCode = &v
	return s
}

func (s *ListBroadcastAudiosByIdResponse) SetBody(v *ListBroadcastAudiosByIdResponseBody) *ListBroadcastAudiosByIdResponse {
	s.Body = v
	return s
}

func (s *ListBroadcastAudiosByIdResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
