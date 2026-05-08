// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAnchorResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateAnchorResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateAnchorResponse
	GetStatusCode() *int32
	SetBody(v *CreateAnchorResponseBody) *CreateAnchorResponse
	GetBody() *CreateAnchorResponseBody
}

type CreateAnchorResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateAnchorResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateAnchorResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateAnchorResponse) GoString() string {
	return s.String()
}

func (s *CreateAnchorResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateAnchorResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateAnchorResponse) GetBody() *CreateAnchorResponseBody {
	return s.Body
}

func (s *CreateAnchorResponse) SetHeaders(v map[string]*string) *CreateAnchorResponse {
	s.Headers = v
	return s
}

func (s *CreateAnchorResponse) SetStatusCode(v int32) *CreateAnchorResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateAnchorResponse) SetBody(v *CreateAnchorResponseBody) *CreateAnchorResponse {
	s.Body = v
	return s
}

func (s *CreateAnchorResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
