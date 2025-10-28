// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateChangeSetResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateChangeSetResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateChangeSetResponse
	GetStatusCode() *int32
	SetBody(v *CreateChangeSetResponseBody) *CreateChangeSetResponse
	GetBody() *CreateChangeSetResponseBody
}

type CreateChangeSetResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateChangeSetResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateChangeSetResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateChangeSetResponse) GoString() string {
	return s.String()
}

func (s *CreateChangeSetResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateChangeSetResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateChangeSetResponse) GetBody() *CreateChangeSetResponseBody {
	return s.Body
}

func (s *CreateChangeSetResponse) SetHeaders(v map[string]*string) *CreateChangeSetResponse {
	s.Headers = v
	return s
}

func (s *CreateChangeSetResponse) SetStatusCode(v int32) *CreateChangeSetResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateChangeSetResponse) SetBody(v *CreateChangeSetResponseBody) *CreateChangeSetResponse {
	s.Body = v
	return s
}

func (s *CreateChangeSetResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
