// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTrialAccessChannelResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetTrialAccessChannelResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetTrialAccessChannelResponse
	GetStatusCode() *int32
	SetBody(v *GetTrialAccessChannelResponseBody) *GetTrialAccessChannelResponse
	GetBody() *GetTrialAccessChannelResponseBody
}

type GetTrialAccessChannelResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetTrialAccessChannelResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetTrialAccessChannelResponse) String() string {
	return dara.Prettify(s)
}

func (s GetTrialAccessChannelResponse) GoString() string {
	return s.String()
}

func (s *GetTrialAccessChannelResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetTrialAccessChannelResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetTrialAccessChannelResponse) GetBody() *GetTrialAccessChannelResponseBody {
	return s.Body
}

func (s *GetTrialAccessChannelResponse) SetHeaders(v map[string]*string) *GetTrialAccessChannelResponse {
	s.Headers = v
	return s
}

func (s *GetTrialAccessChannelResponse) SetStatusCode(v int32) *GetTrialAccessChannelResponse {
	s.StatusCode = &v
	return s
}

func (s *GetTrialAccessChannelResponse) SetBody(v *GetTrialAccessChannelResponseBody) *GetTrialAccessChannelResponse {
	s.Body = v
	return s
}

func (s *GetTrialAccessChannelResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
