// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGeneralAnalyzeImageResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GeneralAnalyzeImageResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GeneralAnalyzeImageResponse
	GetStatusCode() *int32
	SetBody(v *GeneralAnalyzeImageResponseBody) *GeneralAnalyzeImageResponse
	GetBody() *GeneralAnalyzeImageResponseBody
}

type GeneralAnalyzeImageResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GeneralAnalyzeImageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GeneralAnalyzeImageResponse) String() string {
	return dara.Prettify(s)
}

func (s GeneralAnalyzeImageResponse) GoString() string {
	return s.String()
}

func (s *GeneralAnalyzeImageResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GeneralAnalyzeImageResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GeneralAnalyzeImageResponse) GetBody() *GeneralAnalyzeImageResponseBody {
	return s.Body
}

func (s *GeneralAnalyzeImageResponse) SetHeaders(v map[string]*string) *GeneralAnalyzeImageResponse {
	s.Headers = v
	return s
}

func (s *GeneralAnalyzeImageResponse) SetStatusCode(v int32) *GeneralAnalyzeImageResponse {
	s.StatusCode = &v
	return s
}

func (s *GeneralAnalyzeImageResponse) SetBody(v *GeneralAnalyzeImageResponseBody) *GeneralAnalyzeImageResponse {
	s.Body = v
	return s
}

func (s *GeneralAnalyzeImageResponse) Validate() error {
	return dara.Validate(s)
}
