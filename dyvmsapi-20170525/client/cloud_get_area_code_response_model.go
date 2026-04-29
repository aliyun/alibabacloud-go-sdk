// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCloudGetAreaCodeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CloudGetAreaCodeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CloudGetAreaCodeResponse
	GetStatusCode() *int32
	SetBody(v *CloudGetAreaCodeResponseBody) *CloudGetAreaCodeResponse
	GetBody() *CloudGetAreaCodeResponseBody
}

type CloudGetAreaCodeResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CloudGetAreaCodeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CloudGetAreaCodeResponse) String() string {
	return dara.Prettify(s)
}

func (s CloudGetAreaCodeResponse) GoString() string {
	return s.String()
}

func (s *CloudGetAreaCodeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CloudGetAreaCodeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CloudGetAreaCodeResponse) GetBody() *CloudGetAreaCodeResponseBody {
	return s.Body
}

func (s *CloudGetAreaCodeResponse) SetHeaders(v map[string]*string) *CloudGetAreaCodeResponse {
	s.Headers = v
	return s
}

func (s *CloudGetAreaCodeResponse) SetStatusCode(v int32) *CloudGetAreaCodeResponse {
	s.StatusCode = &v
	return s
}

func (s *CloudGetAreaCodeResponse) SetBody(v *CloudGetAreaCodeResponseBody) *CloudGetAreaCodeResponse {
	s.Body = v
	return s
}

func (s *CloudGetAreaCodeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
