// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetShowListBackgroundResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SetShowListBackgroundResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SetShowListBackgroundResponse
	GetStatusCode() *int32
	SetBody(v *SetShowListBackgroundResponseBody) *SetShowListBackgroundResponse
	GetBody() *SetShowListBackgroundResponseBody
}

type SetShowListBackgroundResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SetShowListBackgroundResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SetShowListBackgroundResponse) String() string {
	return dara.Prettify(s)
}

func (s SetShowListBackgroundResponse) GoString() string {
	return s.String()
}

func (s *SetShowListBackgroundResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SetShowListBackgroundResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SetShowListBackgroundResponse) GetBody() *SetShowListBackgroundResponseBody {
	return s.Body
}

func (s *SetShowListBackgroundResponse) SetHeaders(v map[string]*string) *SetShowListBackgroundResponse {
	s.Headers = v
	return s
}

func (s *SetShowListBackgroundResponse) SetStatusCode(v int32) *SetShowListBackgroundResponse {
	s.StatusCode = &v
	return s
}

func (s *SetShowListBackgroundResponse) SetBody(v *SetShowListBackgroundResponseBody) *SetShowListBackgroundResponse {
	s.Body = v
	return s
}

func (s *SetShowListBackgroundResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
