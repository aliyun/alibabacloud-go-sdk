// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAgentDownloadUrlResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetAgentDownloadUrlResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetAgentDownloadUrlResponse
	GetStatusCode() *int32
	SetBody(v *GetAgentDownloadUrlResponseBody) *GetAgentDownloadUrlResponse
	GetBody() *GetAgentDownloadUrlResponseBody
}

type GetAgentDownloadUrlResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAgentDownloadUrlResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAgentDownloadUrlResponse) String() string {
	return dara.Prettify(s)
}

func (s GetAgentDownloadUrlResponse) GoString() string {
	return s.String()
}

func (s *GetAgentDownloadUrlResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetAgentDownloadUrlResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetAgentDownloadUrlResponse) GetBody() *GetAgentDownloadUrlResponseBody {
	return s.Body
}

func (s *GetAgentDownloadUrlResponse) SetHeaders(v map[string]*string) *GetAgentDownloadUrlResponse {
	s.Headers = v
	return s
}

func (s *GetAgentDownloadUrlResponse) SetStatusCode(v int32) *GetAgentDownloadUrlResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAgentDownloadUrlResponse) SetBody(v *GetAgentDownloadUrlResponseBody) *GetAgentDownloadUrlResponse {
	s.Body = v
	return s
}

func (s *GetAgentDownloadUrlResponse) Validate() error {
	return dara.Validate(s)
}
