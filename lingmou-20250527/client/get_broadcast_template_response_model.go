// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetBroadcastTemplateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetBroadcastTemplateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetBroadcastTemplateResponse
	GetStatusCode() *int32
	SetBody(v *GetBroadcastTemplateResponseBody) *GetBroadcastTemplateResponse
	GetBody() *GetBroadcastTemplateResponseBody
}

type GetBroadcastTemplateResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetBroadcastTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetBroadcastTemplateResponse) String() string {
	return dara.Prettify(s)
}

func (s GetBroadcastTemplateResponse) GoString() string {
	return s.String()
}

func (s *GetBroadcastTemplateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetBroadcastTemplateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetBroadcastTemplateResponse) GetBody() *GetBroadcastTemplateResponseBody {
	return s.Body
}

func (s *GetBroadcastTemplateResponse) SetHeaders(v map[string]*string) *GetBroadcastTemplateResponse {
	s.Headers = v
	return s
}

func (s *GetBroadcastTemplateResponse) SetStatusCode(v int32) *GetBroadcastTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *GetBroadcastTemplateResponse) SetBody(v *GetBroadcastTemplateResponseBody) *GetBroadcastTemplateResponse {
	s.Body = v
	return s
}

func (s *GetBroadcastTemplateResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
