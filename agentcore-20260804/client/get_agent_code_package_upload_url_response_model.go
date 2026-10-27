// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAgentCodePackageUploadUrlResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetAgentCodePackageUploadUrlResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetAgentCodePackageUploadUrlResponse
	GetStatusCode() *int32
	SetBody(v *GetAgentCodePackageUploadUrlResponseBody) *GetAgentCodePackageUploadUrlResponse
	GetBody() *GetAgentCodePackageUploadUrlResponseBody
}

type GetAgentCodePackageUploadUrlResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAgentCodePackageUploadUrlResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAgentCodePackageUploadUrlResponse) String() string {
	return dara.Prettify(s)
}

func (s GetAgentCodePackageUploadUrlResponse) GoString() string {
	return s.String()
}

func (s *GetAgentCodePackageUploadUrlResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetAgentCodePackageUploadUrlResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetAgentCodePackageUploadUrlResponse) GetBody() *GetAgentCodePackageUploadUrlResponseBody {
	return s.Body
}

func (s *GetAgentCodePackageUploadUrlResponse) SetHeaders(v map[string]*string) *GetAgentCodePackageUploadUrlResponse {
	s.Headers = v
	return s
}

func (s *GetAgentCodePackageUploadUrlResponse) SetStatusCode(v int32) *GetAgentCodePackageUploadUrlResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAgentCodePackageUploadUrlResponse) SetBody(v *GetAgentCodePackageUploadUrlResponseBody) *GetAgentCodePackageUploadUrlResponse {
	s.Body = v
	return s
}

func (s *GetAgentCodePackageUploadUrlResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
