// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetWorkspaceAgenticFsMountRamAuthorizeUrlResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetWorkspaceAgenticFsMountRamAuthorizeUrlResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetWorkspaceAgenticFsMountRamAuthorizeUrlResponse
	GetStatusCode() *int32
	SetBody(v *GetWorkspaceAgenticFsMountRamAuthorizeUrlResponseBody) *GetWorkspaceAgenticFsMountRamAuthorizeUrlResponse
	GetBody() *GetWorkspaceAgenticFsMountRamAuthorizeUrlResponseBody
}

type GetWorkspaceAgenticFsMountRamAuthorizeUrlResponse struct {
	Headers    map[string]*string                                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetWorkspaceAgenticFsMountRamAuthorizeUrlResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetWorkspaceAgenticFsMountRamAuthorizeUrlResponse) String() string {
	return dara.Prettify(s)
}

func (s GetWorkspaceAgenticFsMountRamAuthorizeUrlResponse) GoString() string {
	return s.String()
}

func (s *GetWorkspaceAgenticFsMountRamAuthorizeUrlResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetWorkspaceAgenticFsMountRamAuthorizeUrlResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetWorkspaceAgenticFsMountRamAuthorizeUrlResponse) GetBody() *GetWorkspaceAgenticFsMountRamAuthorizeUrlResponseBody {
	return s.Body
}

func (s *GetWorkspaceAgenticFsMountRamAuthorizeUrlResponse) SetHeaders(v map[string]*string) *GetWorkspaceAgenticFsMountRamAuthorizeUrlResponse {
	s.Headers = v
	return s
}

func (s *GetWorkspaceAgenticFsMountRamAuthorizeUrlResponse) SetStatusCode(v int32) *GetWorkspaceAgenticFsMountRamAuthorizeUrlResponse {
	s.StatusCode = &v
	return s
}

func (s *GetWorkspaceAgenticFsMountRamAuthorizeUrlResponse) SetBody(v *GetWorkspaceAgenticFsMountRamAuthorizeUrlResponseBody) *GetWorkspaceAgenticFsMountRamAuthorizeUrlResponse {
	s.Body = v
	return s
}

func (s *GetWorkspaceAgenticFsMountRamAuthorizeUrlResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
