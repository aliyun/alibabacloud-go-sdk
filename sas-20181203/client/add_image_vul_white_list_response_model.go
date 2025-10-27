// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddImageVulWhiteListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AddImageVulWhiteListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AddImageVulWhiteListResponse
	GetStatusCode() *int32
	SetBody(v *AddImageVulWhiteListResponseBody) *AddImageVulWhiteListResponse
	GetBody() *AddImageVulWhiteListResponseBody
}

type AddImageVulWhiteListResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddImageVulWhiteListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddImageVulWhiteListResponse) String() string {
	return dara.Prettify(s)
}

func (s AddImageVulWhiteListResponse) GoString() string {
	return s.String()
}

func (s *AddImageVulWhiteListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AddImageVulWhiteListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AddImageVulWhiteListResponse) GetBody() *AddImageVulWhiteListResponseBody {
	return s.Body
}

func (s *AddImageVulWhiteListResponse) SetHeaders(v map[string]*string) *AddImageVulWhiteListResponse {
	s.Headers = v
	return s
}

func (s *AddImageVulWhiteListResponse) SetStatusCode(v int32) *AddImageVulWhiteListResponse {
	s.StatusCode = &v
	return s
}

func (s *AddImageVulWhiteListResponse) SetBody(v *AddImageVulWhiteListResponseBody) *AddImageVulWhiteListResponse {
	s.Body = v
	return s
}

func (s *AddImageVulWhiteListResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
