// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDingTalkUserOrgByAliyunTmpCodeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetDingTalkUserOrgByAliyunTmpCodeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetDingTalkUserOrgByAliyunTmpCodeResponse
	GetStatusCode() *int32
	SetBody(v *GetDingTalkUserOrgByAliyunTmpCodeResponseBody) *GetDingTalkUserOrgByAliyunTmpCodeResponse
	GetBody() *GetDingTalkUserOrgByAliyunTmpCodeResponseBody
}

type GetDingTalkUserOrgByAliyunTmpCodeResponse struct {
	Headers    map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetDingTalkUserOrgByAliyunTmpCodeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetDingTalkUserOrgByAliyunTmpCodeResponse) String() string {
	return dara.Prettify(s)
}

func (s GetDingTalkUserOrgByAliyunTmpCodeResponse) GoString() string {
	return s.String()
}

func (s *GetDingTalkUserOrgByAliyunTmpCodeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetDingTalkUserOrgByAliyunTmpCodeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetDingTalkUserOrgByAliyunTmpCodeResponse) GetBody() *GetDingTalkUserOrgByAliyunTmpCodeResponseBody {
	return s.Body
}

func (s *GetDingTalkUserOrgByAliyunTmpCodeResponse) SetHeaders(v map[string]*string) *GetDingTalkUserOrgByAliyunTmpCodeResponse {
	s.Headers = v
	return s
}

func (s *GetDingTalkUserOrgByAliyunTmpCodeResponse) SetStatusCode(v int32) *GetDingTalkUserOrgByAliyunTmpCodeResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDingTalkUserOrgByAliyunTmpCodeResponse) SetBody(v *GetDingTalkUserOrgByAliyunTmpCodeResponseBody) *GetDingTalkUserOrgByAliyunTmpCodeResponse {
	s.Body = v
	return s
}

func (s *GetDingTalkUserOrgByAliyunTmpCodeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
