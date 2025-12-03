// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSearchCodePreviewResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetSearchCodePreviewResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetSearchCodePreviewResponse
	GetStatusCode() *int32
	SetBody(v *GetSearchCodePreviewResponseBody) *GetSearchCodePreviewResponse
	GetBody() *GetSearchCodePreviewResponseBody
}

type GetSearchCodePreviewResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetSearchCodePreviewResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetSearchCodePreviewResponse) String() string {
	return dara.Prettify(s)
}

func (s GetSearchCodePreviewResponse) GoString() string {
	return s.String()
}

func (s *GetSearchCodePreviewResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetSearchCodePreviewResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetSearchCodePreviewResponse) GetBody() *GetSearchCodePreviewResponseBody {
	return s.Body
}

func (s *GetSearchCodePreviewResponse) SetHeaders(v map[string]*string) *GetSearchCodePreviewResponse {
	s.Headers = v
	return s
}

func (s *GetSearchCodePreviewResponse) SetStatusCode(v int32) *GetSearchCodePreviewResponse {
	s.StatusCode = &v
	return s
}

func (s *GetSearchCodePreviewResponse) SetBody(v *GetSearchCodePreviewResponseBody) *GetSearchCodePreviewResponse {
	s.Body = v
	return s
}

func (s *GetSearchCodePreviewResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
