// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisableSceneGroupTemplateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DisableSceneGroupTemplateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DisableSceneGroupTemplateResponse
	GetStatusCode() *int32
	SetBody(v *DisableSceneGroupTemplateResponseBody) *DisableSceneGroupTemplateResponse
	GetBody() *DisableSceneGroupTemplateResponseBody
}

type DisableSceneGroupTemplateResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DisableSceneGroupTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DisableSceneGroupTemplateResponse) String() string {
	return dara.Prettify(s)
}

func (s DisableSceneGroupTemplateResponse) GoString() string {
	return s.String()
}

func (s *DisableSceneGroupTemplateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DisableSceneGroupTemplateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DisableSceneGroupTemplateResponse) GetBody() *DisableSceneGroupTemplateResponseBody {
	return s.Body
}

func (s *DisableSceneGroupTemplateResponse) SetHeaders(v map[string]*string) *DisableSceneGroupTemplateResponse {
	s.Headers = v
	return s
}

func (s *DisableSceneGroupTemplateResponse) SetStatusCode(v int32) *DisableSceneGroupTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *DisableSceneGroupTemplateResponse) SetBody(v *DisableSceneGroupTemplateResponseBody) *DisableSceneGroupTemplateResponse {
	s.Body = v
	return s
}

func (s *DisableSceneGroupTemplateResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
