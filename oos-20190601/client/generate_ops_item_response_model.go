// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGenerateOpsItemResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GenerateOpsItemResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GenerateOpsItemResponse
	GetStatusCode() *int32
	SetBody(v *GenerateOpsItemResponseBody) *GenerateOpsItemResponse
	GetBody() *GenerateOpsItemResponseBody
}

type GenerateOpsItemResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GenerateOpsItemResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GenerateOpsItemResponse) String() string {
	return dara.Prettify(s)
}

func (s GenerateOpsItemResponse) GoString() string {
	return s.String()
}

func (s *GenerateOpsItemResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GenerateOpsItemResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GenerateOpsItemResponse) GetBody() *GenerateOpsItemResponseBody {
	return s.Body
}

func (s *GenerateOpsItemResponse) SetHeaders(v map[string]*string) *GenerateOpsItemResponse {
	s.Headers = v
	return s
}

func (s *GenerateOpsItemResponse) SetStatusCode(v int32) *GenerateOpsItemResponse {
	s.StatusCode = &v
	return s
}

func (s *GenerateOpsItemResponse) SetBody(v *GenerateOpsItemResponseBody) *GenerateOpsItemResponse {
	s.Body = v
	return s
}

func (s *GenerateOpsItemResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
