// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateGroupDirectoryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateGroupDirectoryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateGroupDirectoryResponse
	GetStatusCode() *int32
	SetBody(v *CreateGroupDirectoryResponseBody) *CreateGroupDirectoryResponse
	GetBody() *CreateGroupDirectoryResponseBody
}

type CreateGroupDirectoryResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateGroupDirectoryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateGroupDirectoryResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateGroupDirectoryResponse) GoString() string {
	return s.String()
}

func (s *CreateGroupDirectoryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateGroupDirectoryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateGroupDirectoryResponse) GetBody() *CreateGroupDirectoryResponseBody {
	return s.Body
}

func (s *CreateGroupDirectoryResponse) SetHeaders(v map[string]*string) *CreateGroupDirectoryResponse {
	s.Headers = v
	return s
}

func (s *CreateGroupDirectoryResponse) SetStatusCode(v int32) *CreateGroupDirectoryResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateGroupDirectoryResponse) SetBody(v *CreateGroupDirectoryResponseBody) *CreateGroupDirectoryResponse {
	s.Body = v
	return s
}

func (s *CreateGroupDirectoryResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
