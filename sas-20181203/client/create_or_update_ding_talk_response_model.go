// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateOrUpdateDingTalkResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateOrUpdateDingTalkResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateOrUpdateDingTalkResponse
	GetStatusCode() *int32
	SetBody(v *CreateOrUpdateDingTalkResponseBody) *CreateOrUpdateDingTalkResponse
	GetBody() *CreateOrUpdateDingTalkResponseBody
}

type CreateOrUpdateDingTalkResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateOrUpdateDingTalkResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateOrUpdateDingTalkResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateOrUpdateDingTalkResponse) GoString() string {
	return s.String()
}

func (s *CreateOrUpdateDingTalkResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateOrUpdateDingTalkResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateOrUpdateDingTalkResponse) GetBody() *CreateOrUpdateDingTalkResponseBody {
	return s.Body
}

func (s *CreateOrUpdateDingTalkResponse) SetHeaders(v map[string]*string) *CreateOrUpdateDingTalkResponse {
	s.Headers = v
	return s
}

func (s *CreateOrUpdateDingTalkResponse) SetStatusCode(v int32) *CreateOrUpdateDingTalkResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateOrUpdateDingTalkResponse) SetBody(v *CreateOrUpdateDingTalkResponseBody) *CreateOrUpdateDingTalkResponse {
	s.Body = v
	return s
}

func (s *CreateOrUpdateDingTalkResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
