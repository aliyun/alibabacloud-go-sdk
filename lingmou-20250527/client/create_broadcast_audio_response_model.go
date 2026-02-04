// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateBroadcastAudioResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateBroadcastAudioResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateBroadcastAudioResponse
	GetStatusCode() *int32
	SetBody(v *CreateBroadcastAudioResponseBody) *CreateBroadcastAudioResponse
	GetBody() *CreateBroadcastAudioResponseBody
}

type CreateBroadcastAudioResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateBroadcastAudioResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateBroadcastAudioResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateBroadcastAudioResponse) GoString() string {
	return s.String()
}

func (s *CreateBroadcastAudioResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateBroadcastAudioResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateBroadcastAudioResponse) GetBody() *CreateBroadcastAudioResponseBody {
	return s.Body
}

func (s *CreateBroadcastAudioResponse) SetHeaders(v map[string]*string) *CreateBroadcastAudioResponse {
	s.Headers = v
	return s
}

func (s *CreateBroadcastAudioResponse) SetStatusCode(v int32) *CreateBroadcastAudioResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateBroadcastAudioResponse) SetBody(v *CreateBroadcastAudioResponseBody) *CreateBroadcastAudioResponse {
	s.Body = v
	return s
}

func (s *CreateBroadcastAudioResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
