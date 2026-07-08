// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddCloudAccessResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AddCloudAccessResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AddCloudAccessResponse
	GetStatusCode() *int32
	SetBody(v *AddCloudAccessResponseBody) *AddCloudAccessResponse
	GetBody() *AddCloudAccessResponseBody
}

type AddCloudAccessResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddCloudAccessResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddCloudAccessResponse) String() string {
	return dara.Prettify(s)
}

func (s AddCloudAccessResponse) GoString() string {
	return s.String()
}

func (s *AddCloudAccessResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AddCloudAccessResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AddCloudAccessResponse) GetBody() *AddCloudAccessResponseBody {
	return s.Body
}

func (s *AddCloudAccessResponse) SetHeaders(v map[string]*string) *AddCloudAccessResponse {
	s.Headers = v
	return s
}

func (s *AddCloudAccessResponse) SetStatusCode(v int32) *AddCloudAccessResponse {
	s.StatusCode = &v
	return s
}

func (s *AddCloudAccessResponse) SetBody(v *AddCloudAccessResponseBody) *AddCloudAccessResponse {
	s.Body = v
	return s
}

func (s *AddCloudAccessResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
