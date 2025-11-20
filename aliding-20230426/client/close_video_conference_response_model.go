// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCloseVideoConferenceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CloseVideoConferenceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CloseVideoConferenceResponse
	GetStatusCode() *int32
	SetBody(v *CloseVideoConferenceResponseBody) *CloseVideoConferenceResponse
	GetBody() *CloseVideoConferenceResponseBody
}

type CloseVideoConferenceResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CloseVideoConferenceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CloseVideoConferenceResponse) String() string {
	return dara.Prettify(s)
}

func (s CloseVideoConferenceResponse) GoString() string {
	return s.String()
}

func (s *CloseVideoConferenceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CloseVideoConferenceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CloseVideoConferenceResponse) GetBody() *CloseVideoConferenceResponseBody {
	return s.Body
}

func (s *CloseVideoConferenceResponse) SetHeaders(v map[string]*string) *CloseVideoConferenceResponse {
	s.Headers = v
	return s
}

func (s *CloseVideoConferenceResponse) SetStatusCode(v int32) *CloseVideoConferenceResponse {
	s.StatusCode = &v
	return s
}

func (s *CloseVideoConferenceResponse) SetBody(v *CloseVideoConferenceResponseBody) *CloseVideoConferenceResponse {
	s.Body = v
	return s
}

func (s *CloseVideoConferenceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
