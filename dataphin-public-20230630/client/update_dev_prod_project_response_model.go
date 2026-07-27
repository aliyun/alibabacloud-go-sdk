// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateDevProdProjectResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateDevProdProjectResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateDevProdProjectResponse
	GetStatusCode() *int32
	SetBody(v *UpdateDevProdProjectResponseBody) *UpdateDevProdProjectResponse
	GetBody() *UpdateDevProdProjectResponseBody
}

type UpdateDevProdProjectResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateDevProdProjectResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateDevProdProjectResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateDevProdProjectResponse) GoString() string {
	return s.String()
}

func (s *UpdateDevProdProjectResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateDevProdProjectResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateDevProdProjectResponse) GetBody() *UpdateDevProdProjectResponseBody {
	return s.Body
}

func (s *UpdateDevProdProjectResponse) SetHeaders(v map[string]*string) *UpdateDevProdProjectResponse {
	s.Headers = v
	return s
}

func (s *UpdateDevProdProjectResponse) SetStatusCode(v int32) *UpdateDevProdProjectResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateDevProdProjectResponse) SetBody(v *UpdateDevProdProjectResponseBody) *UpdateDevProdProjectResponse {
	s.Body = v
	return s
}

func (s *UpdateDevProdProjectResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
