// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddLargeModelResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AddLargeModelResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AddLargeModelResponse
	GetStatusCode() *int32
	SetBody(v *AddLargeModelResponseBody) *AddLargeModelResponse
	GetBody() *AddLargeModelResponseBody
}

type AddLargeModelResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddLargeModelResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddLargeModelResponse) String() string {
	return dara.Prettify(s)
}

func (s AddLargeModelResponse) GoString() string {
	return s.String()
}

func (s *AddLargeModelResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AddLargeModelResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AddLargeModelResponse) GetBody() *AddLargeModelResponseBody {
	return s.Body
}

func (s *AddLargeModelResponse) SetHeaders(v map[string]*string) *AddLargeModelResponse {
	s.Headers = v
	return s
}

func (s *AddLargeModelResponse) SetStatusCode(v int32) *AddLargeModelResponse {
	s.StatusCode = &v
	return s
}

func (s *AddLargeModelResponse) SetBody(v *AddLargeModelResponseBody) *AddLargeModelResponse {
	s.Body = v
	return s
}

func (s *AddLargeModelResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
