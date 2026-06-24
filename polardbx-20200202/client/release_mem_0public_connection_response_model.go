// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReleaseMem0PublicConnectionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ReleaseMem0PublicConnectionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ReleaseMem0PublicConnectionResponse
	GetStatusCode() *int32
	SetBody(v *ReleaseMem0PublicConnectionResponseBody) *ReleaseMem0PublicConnectionResponse
	GetBody() *ReleaseMem0PublicConnectionResponseBody
}

type ReleaseMem0PublicConnectionResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ReleaseMem0PublicConnectionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ReleaseMem0PublicConnectionResponse) String() string {
	return dara.Prettify(s)
}

func (s ReleaseMem0PublicConnectionResponse) GoString() string {
	return s.String()
}

func (s *ReleaseMem0PublicConnectionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ReleaseMem0PublicConnectionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ReleaseMem0PublicConnectionResponse) GetBody() *ReleaseMem0PublicConnectionResponseBody {
	return s.Body
}

func (s *ReleaseMem0PublicConnectionResponse) SetHeaders(v map[string]*string) *ReleaseMem0PublicConnectionResponse {
	s.Headers = v
	return s
}

func (s *ReleaseMem0PublicConnectionResponse) SetStatusCode(v int32) *ReleaseMem0PublicConnectionResponse {
	s.StatusCode = &v
	return s
}

func (s *ReleaseMem0PublicConnectionResponse) SetBody(v *ReleaseMem0PublicConnectionResponseBody) *ReleaseMem0PublicConnectionResponse {
	s.Body = v
	return s
}

func (s *ReleaseMem0PublicConnectionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
