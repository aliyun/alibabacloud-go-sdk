// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAppCodeWorkspaceDetailResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetAppCodeWorkspaceDetailResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetAppCodeWorkspaceDetailResponse
	GetStatusCode() *int32
	SetBody(v *GetAppCodeWorkspaceDetailResponseBody) *GetAppCodeWorkspaceDetailResponse
	GetBody() *GetAppCodeWorkspaceDetailResponseBody
}

type GetAppCodeWorkspaceDetailResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAppCodeWorkspaceDetailResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAppCodeWorkspaceDetailResponse) String() string {
	return dara.Prettify(s)
}

func (s GetAppCodeWorkspaceDetailResponse) GoString() string {
	return s.String()
}

func (s *GetAppCodeWorkspaceDetailResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetAppCodeWorkspaceDetailResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetAppCodeWorkspaceDetailResponse) GetBody() *GetAppCodeWorkspaceDetailResponseBody {
	return s.Body
}

func (s *GetAppCodeWorkspaceDetailResponse) SetHeaders(v map[string]*string) *GetAppCodeWorkspaceDetailResponse {
	s.Headers = v
	return s
}

func (s *GetAppCodeWorkspaceDetailResponse) SetStatusCode(v int32) *GetAppCodeWorkspaceDetailResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAppCodeWorkspaceDetailResponse) SetBody(v *GetAppCodeWorkspaceDetailResponseBody) *GetAppCodeWorkspaceDetailResponse {
	s.Body = v
	return s
}

func (s *GetAppCodeWorkspaceDetailResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
