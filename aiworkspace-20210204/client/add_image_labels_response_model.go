// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddImageLabelsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AddImageLabelsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AddImageLabelsResponse
	GetStatusCode() *int32
	SetBody(v *AddImageLabelsResponseBody) *AddImageLabelsResponse
	GetBody() *AddImageLabelsResponseBody
}

type AddImageLabelsResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddImageLabelsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddImageLabelsResponse) String() string {
	return dara.Prettify(s)
}

func (s AddImageLabelsResponse) GoString() string {
	return s.String()
}

func (s *AddImageLabelsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AddImageLabelsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AddImageLabelsResponse) GetBody() *AddImageLabelsResponseBody {
	return s.Body
}

func (s *AddImageLabelsResponse) SetHeaders(v map[string]*string) *AddImageLabelsResponse {
	s.Headers = v
	return s
}

func (s *AddImageLabelsResponse) SetStatusCode(v int32) *AddImageLabelsResponse {
	s.StatusCode = &v
	return s
}

func (s *AddImageLabelsResponse) SetBody(v *AddImageLabelsResponseBody) *AddImageLabelsResponse {
	s.Body = v
	return s
}

func (s *AddImageLabelsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
