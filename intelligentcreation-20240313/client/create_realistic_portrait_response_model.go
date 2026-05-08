// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateRealisticPortraitResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateRealisticPortraitResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateRealisticPortraitResponse
	GetStatusCode() *int32
	SetBody(v *CreateRealisticPortraitResponseBody) *CreateRealisticPortraitResponse
	GetBody() *CreateRealisticPortraitResponseBody
}

type CreateRealisticPortraitResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateRealisticPortraitResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateRealisticPortraitResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateRealisticPortraitResponse) GoString() string {
	return s.String()
}

func (s *CreateRealisticPortraitResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateRealisticPortraitResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateRealisticPortraitResponse) GetBody() *CreateRealisticPortraitResponseBody {
	return s.Body
}

func (s *CreateRealisticPortraitResponse) SetHeaders(v map[string]*string) *CreateRealisticPortraitResponse {
	s.Headers = v
	return s
}

func (s *CreateRealisticPortraitResponse) SetStatusCode(v int32) *CreateRealisticPortraitResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateRealisticPortraitResponse) SetBody(v *CreateRealisticPortraitResponseBody) *CreateRealisticPortraitResponse {
	s.Body = v
	return s
}

func (s *CreateRealisticPortraitResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
