// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAIStaffPreviewUrlResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetAIStaffPreviewUrlResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetAIStaffPreviewUrlResponse
	GetStatusCode() *int32
	SetBody(v *GetAIStaffPreviewUrlResponseBody) *GetAIStaffPreviewUrlResponse
	GetBody() *GetAIStaffPreviewUrlResponseBody
}

type GetAIStaffPreviewUrlResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAIStaffPreviewUrlResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAIStaffPreviewUrlResponse) String() string {
	return dara.Prettify(s)
}

func (s GetAIStaffPreviewUrlResponse) GoString() string {
	return s.String()
}

func (s *GetAIStaffPreviewUrlResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetAIStaffPreviewUrlResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetAIStaffPreviewUrlResponse) GetBody() *GetAIStaffPreviewUrlResponseBody {
	return s.Body
}

func (s *GetAIStaffPreviewUrlResponse) SetHeaders(v map[string]*string) *GetAIStaffPreviewUrlResponse {
	s.Headers = v
	return s
}

func (s *GetAIStaffPreviewUrlResponse) SetStatusCode(v int32) *GetAIStaffPreviewUrlResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAIStaffPreviewUrlResponse) SetBody(v *GetAIStaffPreviewUrlResponseBody) *GetAIStaffPreviewUrlResponse {
	s.Body = v
	return s
}

func (s *GetAIStaffPreviewUrlResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
