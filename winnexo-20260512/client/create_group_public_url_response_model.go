// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateGroupPublicUrlResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateGroupPublicUrlResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateGroupPublicUrlResponse
	GetStatusCode() *int32
	SetBody(v *CreateGroupPublicUrlResponseBody) *CreateGroupPublicUrlResponse
	GetBody() *CreateGroupPublicUrlResponseBody
}

type CreateGroupPublicUrlResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateGroupPublicUrlResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateGroupPublicUrlResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateGroupPublicUrlResponse) GoString() string {
	return s.String()
}

func (s *CreateGroupPublicUrlResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateGroupPublicUrlResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateGroupPublicUrlResponse) GetBody() *CreateGroupPublicUrlResponseBody {
	return s.Body
}

func (s *CreateGroupPublicUrlResponse) SetHeaders(v map[string]*string) *CreateGroupPublicUrlResponse {
	s.Headers = v
	return s
}

func (s *CreateGroupPublicUrlResponse) SetStatusCode(v int32) *CreateGroupPublicUrlResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateGroupPublicUrlResponse) SetBody(v *CreateGroupPublicUrlResponseBody) *CreateGroupPublicUrlResponse {
	s.Body = v
	return s
}

func (s *CreateGroupPublicUrlResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
