// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddLabelsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AddLabelsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AddLabelsResponse
	GetStatusCode() *int32
	SetBody(v *AddLabelsResponseBody) *AddLabelsResponse
	GetBody() *AddLabelsResponseBody
}

type AddLabelsResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddLabelsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddLabelsResponse) String() string {
	return dara.Prettify(s)
}

func (s AddLabelsResponse) GoString() string {
	return s.String()
}

func (s *AddLabelsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AddLabelsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AddLabelsResponse) GetBody() *AddLabelsResponseBody {
	return s.Body
}

func (s *AddLabelsResponse) SetHeaders(v map[string]*string) *AddLabelsResponse {
	s.Headers = v
	return s
}

func (s *AddLabelsResponse) SetStatusCode(v int32) *AddLabelsResponse {
	s.StatusCode = &v
	return s
}

func (s *AddLabelsResponse) SetBody(v *AddLabelsResponseBody) *AddLabelsResponse {
	s.Body = v
	return s
}

func (s *AddLabelsResponse) Validate() error {
	return dara.Validate(s)
}
