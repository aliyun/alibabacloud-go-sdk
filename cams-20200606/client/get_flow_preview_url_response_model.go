// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetFlowPreviewUrlResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetFlowPreviewUrlResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetFlowPreviewUrlResponse
	GetStatusCode() *int32
	SetBody(v *GetFlowPreviewUrlResponseBody) *GetFlowPreviewUrlResponse
	GetBody() *GetFlowPreviewUrlResponseBody
}

type GetFlowPreviewUrlResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetFlowPreviewUrlResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetFlowPreviewUrlResponse) String() string {
	return dara.Prettify(s)
}

func (s GetFlowPreviewUrlResponse) GoString() string {
	return s.String()
}

func (s *GetFlowPreviewUrlResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetFlowPreviewUrlResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetFlowPreviewUrlResponse) GetBody() *GetFlowPreviewUrlResponseBody {
	return s.Body
}

func (s *GetFlowPreviewUrlResponse) SetHeaders(v map[string]*string) *GetFlowPreviewUrlResponse {
	s.Headers = v
	return s
}

func (s *GetFlowPreviewUrlResponse) SetStatusCode(v int32) *GetFlowPreviewUrlResponse {
	s.StatusCode = &v
	return s
}

func (s *GetFlowPreviewUrlResponse) SetBody(v *GetFlowPreviewUrlResponseBody) *GetFlowPreviewUrlResponse {
	s.Body = v
	return s
}

func (s *GetFlowPreviewUrlResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
