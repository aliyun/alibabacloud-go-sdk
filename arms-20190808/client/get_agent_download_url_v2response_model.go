// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAgentDownloadUrlV2Response interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetAgentDownloadUrlV2Response
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetAgentDownloadUrlV2Response
	GetStatusCode() *int32
	SetBody(v *GetAgentDownloadUrlV2ResponseBody) *GetAgentDownloadUrlV2Response
	GetBody() *GetAgentDownloadUrlV2ResponseBody
}

type GetAgentDownloadUrlV2Response struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAgentDownloadUrlV2ResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAgentDownloadUrlV2Response) String() string {
	return dara.Prettify(s)
}

func (s GetAgentDownloadUrlV2Response) GoString() string {
	return s.String()
}

func (s *GetAgentDownloadUrlV2Response) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetAgentDownloadUrlV2Response) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetAgentDownloadUrlV2Response) GetBody() *GetAgentDownloadUrlV2ResponseBody {
	return s.Body
}

func (s *GetAgentDownloadUrlV2Response) SetHeaders(v map[string]*string) *GetAgentDownloadUrlV2Response {
	s.Headers = v
	return s
}

func (s *GetAgentDownloadUrlV2Response) SetStatusCode(v int32) *GetAgentDownloadUrlV2Response {
	s.StatusCode = &v
	return s
}

func (s *GetAgentDownloadUrlV2Response) SetBody(v *GetAgentDownloadUrlV2ResponseBody) *GetAgentDownloadUrlV2Response {
	s.Body = v
	return s
}

func (s *GetAgentDownloadUrlV2Response) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
