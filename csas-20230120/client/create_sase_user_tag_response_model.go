// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSaseUserTagResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateSaseUserTagResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateSaseUserTagResponse
	GetStatusCode() *int32
	SetBody(v *CreateSaseUserTagResponseBody) *CreateSaseUserTagResponse
	GetBody() *CreateSaseUserTagResponseBody
}

type CreateSaseUserTagResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateSaseUserTagResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateSaseUserTagResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateSaseUserTagResponse) GoString() string {
	return s.String()
}

func (s *CreateSaseUserTagResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateSaseUserTagResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateSaseUserTagResponse) GetBody() *CreateSaseUserTagResponseBody {
	return s.Body
}

func (s *CreateSaseUserTagResponse) SetHeaders(v map[string]*string) *CreateSaseUserTagResponse {
	s.Headers = v
	return s
}

func (s *CreateSaseUserTagResponse) SetStatusCode(v int32) *CreateSaseUserTagResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateSaseUserTagResponse) SetBody(v *CreateSaseUserTagResponseBody) *CreateSaseUserTagResponse {
	s.Body = v
	return s
}

func (s *CreateSaseUserTagResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
