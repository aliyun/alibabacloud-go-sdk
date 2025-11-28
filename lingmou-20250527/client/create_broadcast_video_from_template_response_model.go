// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateBroadcastVideoFromTemplateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateBroadcastVideoFromTemplateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateBroadcastVideoFromTemplateResponse
	GetStatusCode() *int32
	SetBody(v *CreateBroadcastVideoFromTemplateResponseBody) *CreateBroadcastVideoFromTemplateResponse
	GetBody() *CreateBroadcastVideoFromTemplateResponseBody
}

type CreateBroadcastVideoFromTemplateResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateBroadcastVideoFromTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateBroadcastVideoFromTemplateResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateBroadcastVideoFromTemplateResponse) GoString() string {
	return s.String()
}

func (s *CreateBroadcastVideoFromTemplateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateBroadcastVideoFromTemplateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateBroadcastVideoFromTemplateResponse) GetBody() *CreateBroadcastVideoFromTemplateResponseBody {
	return s.Body
}

func (s *CreateBroadcastVideoFromTemplateResponse) SetHeaders(v map[string]*string) *CreateBroadcastVideoFromTemplateResponse {
	s.Headers = v
	return s
}

func (s *CreateBroadcastVideoFromTemplateResponse) SetStatusCode(v int32) *CreateBroadcastVideoFromTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateBroadcastVideoFromTemplateResponse) SetBody(v *CreateBroadcastVideoFromTemplateResponseBody) *CreateBroadcastVideoFromTemplateResponse {
	s.Body = v
	return s
}

func (s *CreateBroadcastVideoFromTemplateResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
