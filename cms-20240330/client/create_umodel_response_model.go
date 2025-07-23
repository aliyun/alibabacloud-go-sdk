// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateUmodelResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateUmodelResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateUmodelResponse
	GetStatusCode() *int32
	SetBody(v *CreateUmodelResponseBody) *CreateUmodelResponse
	GetBody() *CreateUmodelResponseBody
}

type CreateUmodelResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateUmodelResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateUmodelResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateUmodelResponse) GoString() string {
	return s.String()
}

func (s *CreateUmodelResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateUmodelResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateUmodelResponse) GetBody() *CreateUmodelResponseBody {
	return s.Body
}

func (s *CreateUmodelResponse) SetHeaders(v map[string]*string) *CreateUmodelResponse {
	s.Headers = v
	return s
}

func (s *CreateUmodelResponse) SetStatusCode(v int32) *CreateUmodelResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateUmodelResponse) SetBody(v *CreateUmodelResponseBody) *CreateUmodelResponse {
	s.Body = v
	return s
}

func (s *CreateUmodelResponse) Validate() error {
	return dara.Validate(s)
}
