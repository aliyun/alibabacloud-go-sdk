// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveRuntimeModelTemplateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RemoveRuntimeModelTemplateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RemoveRuntimeModelTemplateResponse
	GetStatusCode() *int32
	SetBody(v *RemoveRuntimeModelTemplateResponseBody) *RemoveRuntimeModelTemplateResponse
	GetBody() *RemoveRuntimeModelTemplateResponseBody
}

type RemoveRuntimeModelTemplateResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RemoveRuntimeModelTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RemoveRuntimeModelTemplateResponse) String() string {
	return dara.Prettify(s)
}

func (s RemoveRuntimeModelTemplateResponse) GoString() string {
	return s.String()
}

func (s *RemoveRuntimeModelTemplateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RemoveRuntimeModelTemplateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RemoveRuntimeModelTemplateResponse) GetBody() *RemoveRuntimeModelTemplateResponseBody {
	return s.Body
}

func (s *RemoveRuntimeModelTemplateResponse) SetHeaders(v map[string]*string) *RemoveRuntimeModelTemplateResponse {
	s.Headers = v
	return s
}

func (s *RemoveRuntimeModelTemplateResponse) SetStatusCode(v int32) *RemoveRuntimeModelTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *RemoveRuntimeModelTemplateResponse) SetBody(v *RemoveRuntimeModelTemplateResponseBody) *RemoveRuntimeModelTemplateResponse {
	s.Body = v
	return s
}

func (s *RemoveRuntimeModelTemplateResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
