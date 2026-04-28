// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePolarFsObjectResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreatePolarFsObjectResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreatePolarFsObjectResponse
	GetStatusCode() *int32
	SetBody(v *CreatePolarFsObjectResponseBody) *CreatePolarFsObjectResponse
	GetBody() *CreatePolarFsObjectResponseBody
}

type CreatePolarFsObjectResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreatePolarFsObjectResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreatePolarFsObjectResponse) String() string {
	return dara.Prettify(s)
}

func (s CreatePolarFsObjectResponse) GoString() string {
	return s.String()
}

func (s *CreatePolarFsObjectResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreatePolarFsObjectResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreatePolarFsObjectResponse) GetBody() *CreatePolarFsObjectResponseBody {
	return s.Body
}

func (s *CreatePolarFsObjectResponse) SetHeaders(v map[string]*string) *CreatePolarFsObjectResponse {
	s.Headers = v
	return s
}

func (s *CreatePolarFsObjectResponse) SetStatusCode(v int32) *CreatePolarFsObjectResponse {
	s.StatusCode = &v
	return s
}

func (s *CreatePolarFsObjectResponse) SetBody(v *CreatePolarFsObjectResponseBody) *CreatePolarFsObjectResponse {
	s.Body = v
	return s
}

func (s *CreatePolarFsObjectResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
