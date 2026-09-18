// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateGroupAliDingDocResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateGroupAliDingDocResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateGroupAliDingDocResponse
	GetStatusCode() *int32
	SetBody(v *CreateGroupAliDingDocResponseBody) *CreateGroupAliDingDocResponse
	GetBody() *CreateGroupAliDingDocResponseBody
}

type CreateGroupAliDingDocResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateGroupAliDingDocResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateGroupAliDingDocResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateGroupAliDingDocResponse) GoString() string {
	return s.String()
}

func (s *CreateGroupAliDingDocResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateGroupAliDingDocResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateGroupAliDingDocResponse) GetBody() *CreateGroupAliDingDocResponseBody {
	return s.Body
}

func (s *CreateGroupAliDingDocResponse) SetHeaders(v map[string]*string) *CreateGroupAliDingDocResponse {
	s.Headers = v
	return s
}

func (s *CreateGroupAliDingDocResponse) SetStatusCode(v int32) *CreateGroupAliDingDocResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateGroupAliDingDocResponse) SetBody(v *CreateGroupAliDingDocResponseBody) *CreateGroupAliDingDocResponse {
	s.Body = v
	return s
}

func (s *CreateGroupAliDingDocResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
