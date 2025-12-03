// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateRepositoryGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateRepositoryGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateRepositoryGroupResponse
	GetStatusCode() *int32
	SetBody(v *CreateRepositoryGroupResponseBody) *CreateRepositoryGroupResponse
	GetBody() *CreateRepositoryGroupResponseBody
}

type CreateRepositoryGroupResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateRepositoryGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateRepositoryGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateRepositoryGroupResponse) GoString() string {
	return s.String()
}

func (s *CreateRepositoryGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateRepositoryGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateRepositoryGroupResponse) GetBody() *CreateRepositoryGroupResponseBody {
	return s.Body
}

func (s *CreateRepositoryGroupResponse) SetHeaders(v map[string]*string) *CreateRepositoryGroupResponse {
	s.Headers = v
	return s
}

func (s *CreateRepositoryGroupResponse) SetStatusCode(v int32) *CreateRepositoryGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateRepositoryGroupResponse) SetBody(v *CreateRepositoryGroupResponseBody) *CreateRepositoryGroupResponse {
	s.Body = v
	return s
}

func (s *CreateRepositoryGroupResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
