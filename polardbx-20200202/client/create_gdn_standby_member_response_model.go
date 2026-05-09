// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateGdnStandbyMemberResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateGdnStandbyMemberResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateGdnStandbyMemberResponse
	GetStatusCode() *int32
	SetBody(v *CreateGdnStandbyMemberResponseBody) *CreateGdnStandbyMemberResponse
	GetBody() *CreateGdnStandbyMemberResponseBody
}

type CreateGdnStandbyMemberResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateGdnStandbyMemberResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateGdnStandbyMemberResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateGdnStandbyMemberResponse) GoString() string {
	return s.String()
}

func (s *CreateGdnStandbyMemberResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateGdnStandbyMemberResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateGdnStandbyMemberResponse) GetBody() *CreateGdnStandbyMemberResponseBody {
	return s.Body
}

func (s *CreateGdnStandbyMemberResponse) SetHeaders(v map[string]*string) *CreateGdnStandbyMemberResponse {
	s.Headers = v
	return s
}

func (s *CreateGdnStandbyMemberResponse) SetStatusCode(v int32) *CreateGdnStandbyMemberResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateGdnStandbyMemberResponse) SetBody(v *CreateGdnStandbyMemberResponseBody) *CreateGdnStandbyMemberResponse {
	s.Body = v
	return s
}

func (s *CreateGdnStandbyMemberResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
