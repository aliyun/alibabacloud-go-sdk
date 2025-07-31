// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetProjectProduceUserResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetProjectProduceUserResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetProjectProduceUserResponse
	GetStatusCode() *int32
	SetBody(v *GetProjectProduceUserResponseBody) *GetProjectProduceUserResponse
	GetBody() *GetProjectProduceUserResponseBody
}

type GetProjectProduceUserResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetProjectProduceUserResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetProjectProduceUserResponse) String() string {
	return dara.Prettify(s)
}

func (s GetProjectProduceUserResponse) GoString() string {
	return s.String()
}

func (s *GetProjectProduceUserResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetProjectProduceUserResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetProjectProduceUserResponse) GetBody() *GetProjectProduceUserResponseBody {
	return s.Body
}

func (s *GetProjectProduceUserResponse) SetHeaders(v map[string]*string) *GetProjectProduceUserResponse {
	s.Headers = v
	return s
}

func (s *GetProjectProduceUserResponse) SetStatusCode(v int32) *GetProjectProduceUserResponse {
	s.StatusCode = &v
	return s
}

func (s *GetProjectProduceUserResponse) SetBody(v *GetProjectProduceUserResponseBody) *GetProjectProduceUserResponse {
	s.Body = v
	return s
}

func (s *GetProjectProduceUserResponse) Validate() error {
	return dara.Validate(s)
}
