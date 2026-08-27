// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateGroupFeishuDocResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateGroupFeishuDocResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateGroupFeishuDocResponse
	GetStatusCode() *int32
	SetBody(v *CreateGroupFeishuDocResponseBody) *CreateGroupFeishuDocResponse
	GetBody() *CreateGroupFeishuDocResponseBody
}

type CreateGroupFeishuDocResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateGroupFeishuDocResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateGroupFeishuDocResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateGroupFeishuDocResponse) GoString() string {
	return s.String()
}

func (s *CreateGroupFeishuDocResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateGroupFeishuDocResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateGroupFeishuDocResponse) GetBody() *CreateGroupFeishuDocResponseBody {
	return s.Body
}

func (s *CreateGroupFeishuDocResponse) SetHeaders(v map[string]*string) *CreateGroupFeishuDocResponse {
	s.Headers = v
	return s
}

func (s *CreateGroupFeishuDocResponse) SetStatusCode(v int32) *CreateGroupFeishuDocResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateGroupFeishuDocResponse) SetBody(v *CreateGroupFeishuDocResponseBody) *CreateGroupFeishuDocResponse {
	s.Body = v
	return s
}

func (s *CreateGroupFeishuDocResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
