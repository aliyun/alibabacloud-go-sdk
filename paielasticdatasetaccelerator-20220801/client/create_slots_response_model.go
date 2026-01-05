// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSlotsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateSlotsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateSlotsResponse
	GetStatusCode() *int32
	SetBody(v *CreateSlotsResponseBody) *CreateSlotsResponse
	GetBody() *CreateSlotsResponseBody
}

type CreateSlotsResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateSlotsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateSlotsResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateSlotsResponse) GoString() string {
	return s.String()
}

func (s *CreateSlotsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateSlotsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateSlotsResponse) GetBody() *CreateSlotsResponseBody {
	return s.Body
}

func (s *CreateSlotsResponse) SetHeaders(v map[string]*string) *CreateSlotsResponse {
	s.Headers = v
	return s
}

func (s *CreateSlotsResponse) SetStatusCode(v int32) *CreateSlotsResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateSlotsResponse) SetBody(v *CreateSlotsResponseBody) *CreateSlotsResponse {
	s.Body = v
	return s
}

func (s *CreateSlotsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
