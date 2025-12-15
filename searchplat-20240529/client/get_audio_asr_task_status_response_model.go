// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAudioAsrTaskStatusResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetAudioAsrTaskStatusResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetAudioAsrTaskStatusResponse
	GetStatusCode() *int32
	SetBody(v *GetAudioAsrTaskStatusResponseBody) *GetAudioAsrTaskStatusResponse
	GetBody() *GetAudioAsrTaskStatusResponseBody
}

type GetAudioAsrTaskStatusResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAudioAsrTaskStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAudioAsrTaskStatusResponse) String() string {
	return dara.Prettify(s)
}

func (s GetAudioAsrTaskStatusResponse) GoString() string {
	return s.String()
}

func (s *GetAudioAsrTaskStatusResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetAudioAsrTaskStatusResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetAudioAsrTaskStatusResponse) GetBody() *GetAudioAsrTaskStatusResponseBody {
	return s.Body
}

func (s *GetAudioAsrTaskStatusResponse) SetHeaders(v map[string]*string) *GetAudioAsrTaskStatusResponse {
	s.Headers = v
	return s
}

func (s *GetAudioAsrTaskStatusResponse) SetStatusCode(v int32) *GetAudioAsrTaskStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAudioAsrTaskStatusResponse) SetBody(v *GetAudioAsrTaskStatusResponseBody) *GetAudioAsrTaskStatusResponse {
	s.Body = v
	return s
}

func (s *GetAudioAsrTaskStatusResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
