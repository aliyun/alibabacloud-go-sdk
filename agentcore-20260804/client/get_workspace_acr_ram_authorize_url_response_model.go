// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetWorkspaceAcrRamAuthorizeUrlResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetWorkspaceAcrRamAuthorizeUrlResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetWorkspaceAcrRamAuthorizeUrlResponse
	GetStatusCode() *int32
	SetBody(v *GetWorkspaceAcrRamAuthorizeUrlResponseBody) *GetWorkspaceAcrRamAuthorizeUrlResponse
	GetBody() *GetWorkspaceAcrRamAuthorizeUrlResponseBody
}

type GetWorkspaceAcrRamAuthorizeUrlResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetWorkspaceAcrRamAuthorizeUrlResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetWorkspaceAcrRamAuthorizeUrlResponse) String() string {
	return dara.Prettify(s)
}

func (s GetWorkspaceAcrRamAuthorizeUrlResponse) GoString() string {
	return s.String()
}

func (s *GetWorkspaceAcrRamAuthorizeUrlResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetWorkspaceAcrRamAuthorizeUrlResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetWorkspaceAcrRamAuthorizeUrlResponse) GetBody() *GetWorkspaceAcrRamAuthorizeUrlResponseBody {
	return s.Body
}

func (s *GetWorkspaceAcrRamAuthorizeUrlResponse) SetHeaders(v map[string]*string) *GetWorkspaceAcrRamAuthorizeUrlResponse {
	s.Headers = v
	return s
}

func (s *GetWorkspaceAcrRamAuthorizeUrlResponse) SetStatusCode(v int32) *GetWorkspaceAcrRamAuthorizeUrlResponse {
	s.StatusCode = &v
	return s
}

func (s *GetWorkspaceAcrRamAuthorizeUrlResponse) SetBody(v *GetWorkspaceAcrRamAuthorizeUrlResponseBody) *GetWorkspaceAcrRamAuthorizeUrlResponse {
	s.Body = v
	return s
}

func (s *GetWorkspaceAcrRamAuthorizeUrlResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
