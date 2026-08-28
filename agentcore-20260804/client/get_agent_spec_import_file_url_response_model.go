// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAgentSpecImportFileUrlResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetAgentSpecImportFileUrlResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetAgentSpecImportFileUrlResponse
	GetStatusCode() *int32
	SetBody(v *GetAgentSpecImportFileUrlResponseBody) *GetAgentSpecImportFileUrlResponse
	GetBody() *GetAgentSpecImportFileUrlResponseBody
}

type GetAgentSpecImportFileUrlResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAgentSpecImportFileUrlResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAgentSpecImportFileUrlResponse) String() string {
	return dara.Prettify(s)
}

func (s GetAgentSpecImportFileUrlResponse) GoString() string {
	return s.String()
}

func (s *GetAgentSpecImportFileUrlResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetAgentSpecImportFileUrlResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetAgentSpecImportFileUrlResponse) GetBody() *GetAgentSpecImportFileUrlResponseBody {
	return s.Body
}

func (s *GetAgentSpecImportFileUrlResponse) SetHeaders(v map[string]*string) *GetAgentSpecImportFileUrlResponse {
	s.Headers = v
	return s
}

func (s *GetAgentSpecImportFileUrlResponse) SetStatusCode(v int32) *GetAgentSpecImportFileUrlResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAgentSpecImportFileUrlResponse) SetBody(v *GetAgentSpecImportFileUrlResponseBody) *GetAgentSpecImportFileUrlResponse {
	s.Body = v
	return s
}

func (s *GetAgentSpecImportFileUrlResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
