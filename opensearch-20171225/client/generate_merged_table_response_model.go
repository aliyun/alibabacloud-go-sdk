// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGenerateMergedTableResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GenerateMergedTableResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GenerateMergedTableResponse
	GetStatusCode() *int32
	SetBody(v *GenerateMergedTableResponseBody) *GenerateMergedTableResponse
	GetBody() *GenerateMergedTableResponseBody
}

type GenerateMergedTableResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GenerateMergedTableResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GenerateMergedTableResponse) String() string {
	return dara.Prettify(s)
}

func (s GenerateMergedTableResponse) GoString() string {
	return s.String()
}

func (s *GenerateMergedTableResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GenerateMergedTableResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GenerateMergedTableResponse) GetBody() *GenerateMergedTableResponseBody {
	return s.Body
}

func (s *GenerateMergedTableResponse) SetHeaders(v map[string]*string) *GenerateMergedTableResponse {
	s.Headers = v
	return s
}

func (s *GenerateMergedTableResponse) SetStatusCode(v int32) *GenerateMergedTableResponse {
	s.StatusCode = &v
	return s
}

func (s *GenerateMergedTableResponse) SetBody(v *GenerateMergedTableResponseBody) *GenerateMergedTableResponse {
	s.Body = v
	return s
}

func (s *GenerateMergedTableResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
