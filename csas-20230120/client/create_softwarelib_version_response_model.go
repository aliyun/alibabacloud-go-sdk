// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSoftwarelibVersionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateSoftwarelibVersionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateSoftwarelibVersionResponse
	GetStatusCode() *int32
	SetBody(v *CreateSoftwarelibVersionResponseBody) *CreateSoftwarelibVersionResponse
	GetBody() *CreateSoftwarelibVersionResponseBody
}

type CreateSoftwarelibVersionResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateSoftwarelibVersionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateSoftwarelibVersionResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateSoftwarelibVersionResponse) GoString() string {
	return s.String()
}

func (s *CreateSoftwarelibVersionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateSoftwarelibVersionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateSoftwarelibVersionResponse) GetBody() *CreateSoftwarelibVersionResponseBody {
	return s.Body
}

func (s *CreateSoftwarelibVersionResponse) SetHeaders(v map[string]*string) *CreateSoftwarelibVersionResponse {
	s.Headers = v
	return s
}

func (s *CreateSoftwarelibVersionResponse) SetStatusCode(v int32) *CreateSoftwarelibVersionResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateSoftwarelibVersionResponse) SetBody(v *CreateSoftwarelibVersionResponseBody) *CreateSoftwarelibVersionResponse {
	s.Body = v
	return s
}

func (s *CreateSoftwarelibVersionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
