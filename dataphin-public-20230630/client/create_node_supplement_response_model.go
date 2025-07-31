// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateNodeSupplementResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateNodeSupplementResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateNodeSupplementResponse
	GetStatusCode() *int32
	SetBody(v *CreateNodeSupplementResponseBody) *CreateNodeSupplementResponse
	GetBody() *CreateNodeSupplementResponseBody
}

type CreateNodeSupplementResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateNodeSupplementResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateNodeSupplementResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateNodeSupplementResponse) GoString() string {
	return s.String()
}

func (s *CreateNodeSupplementResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateNodeSupplementResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateNodeSupplementResponse) GetBody() *CreateNodeSupplementResponseBody {
	return s.Body
}

func (s *CreateNodeSupplementResponse) SetHeaders(v map[string]*string) *CreateNodeSupplementResponse {
	s.Headers = v
	return s
}

func (s *CreateNodeSupplementResponse) SetStatusCode(v int32) *CreateNodeSupplementResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateNodeSupplementResponse) SetBody(v *CreateNodeSupplementResponseBody) *CreateNodeSupplementResponse {
	s.Body = v
	return s
}

func (s *CreateNodeSupplementResponse) Validate() error {
	return dara.Validate(s)
}
