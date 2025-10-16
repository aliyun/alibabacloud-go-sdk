// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateImageByInstanceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateImageByInstanceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateImageByInstanceResponse
	GetStatusCode() *int32
	SetBody(v *CreateImageByInstanceResponseBody) *CreateImageByInstanceResponse
	GetBody() *CreateImageByInstanceResponseBody
}

type CreateImageByInstanceResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateImageByInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateImageByInstanceResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateImageByInstanceResponse) GoString() string {
	return s.String()
}

func (s *CreateImageByInstanceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateImageByInstanceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateImageByInstanceResponse) GetBody() *CreateImageByInstanceResponseBody {
	return s.Body
}

func (s *CreateImageByInstanceResponse) SetHeaders(v map[string]*string) *CreateImageByInstanceResponse {
	s.Headers = v
	return s
}

func (s *CreateImageByInstanceResponse) SetStatusCode(v int32) *CreateImageByInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateImageByInstanceResponse) SetBody(v *CreateImageByInstanceResponseBody) *CreateImageByInstanceResponse {
	s.Body = v
	return s
}

func (s *CreateImageByInstanceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
