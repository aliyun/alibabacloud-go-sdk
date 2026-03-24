// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCopyBroadcastSceneFromTemplateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CopyBroadcastSceneFromTemplateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CopyBroadcastSceneFromTemplateResponse
	GetStatusCode() *int32
	SetBody(v *CopyBroadcastSceneFromTemplateResponseBody) *CopyBroadcastSceneFromTemplateResponse
	GetBody() *CopyBroadcastSceneFromTemplateResponseBody
}

type CopyBroadcastSceneFromTemplateResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CopyBroadcastSceneFromTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CopyBroadcastSceneFromTemplateResponse) String() string {
	return dara.Prettify(s)
}

func (s CopyBroadcastSceneFromTemplateResponse) GoString() string {
	return s.String()
}

func (s *CopyBroadcastSceneFromTemplateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CopyBroadcastSceneFromTemplateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CopyBroadcastSceneFromTemplateResponse) GetBody() *CopyBroadcastSceneFromTemplateResponseBody {
	return s.Body
}

func (s *CopyBroadcastSceneFromTemplateResponse) SetHeaders(v map[string]*string) *CopyBroadcastSceneFromTemplateResponse {
	s.Headers = v
	return s
}

func (s *CopyBroadcastSceneFromTemplateResponse) SetStatusCode(v int32) *CopyBroadcastSceneFromTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *CopyBroadcastSceneFromTemplateResponse) SetBody(v *CopyBroadcastSceneFromTemplateResponseBody) *CopyBroadcastSceneFromTemplateResponse {
	s.Body = v
	return s
}

func (s *CopyBroadcastSceneFromTemplateResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
