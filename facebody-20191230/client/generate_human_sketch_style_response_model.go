// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGenerateHumanSketchStyleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GenerateHumanSketchStyleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GenerateHumanSketchStyleResponse
	GetStatusCode() *int32
	SetBody(v *GenerateHumanSketchStyleResponseBody) *GenerateHumanSketchStyleResponse
	GetBody() *GenerateHumanSketchStyleResponseBody
}

type GenerateHumanSketchStyleResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GenerateHumanSketchStyleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GenerateHumanSketchStyleResponse) String() string {
	return dara.Prettify(s)
}

func (s GenerateHumanSketchStyleResponse) GoString() string {
	return s.String()
}

func (s *GenerateHumanSketchStyleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GenerateHumanSketchStyleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GenerateHumanSketchStyleResponse) GetBody() *GenerateHumanSketchStyleResponseBody {
	return s.Body
}

func (s *GenerateHumanSketchStyleResponse) SetHeaders(v map[string]*string) *GenerateHumanSketchStyleResponse {
	s.Headers = v
	return s
}

func (s *GenerateHumanSketchStyleResponse) SetStatusCode(v int32) *GenerateHumanSketchStyleResponse {
	s.StatusCode = &v
	return s
}

func (s *GenerateHumanSketchStyleResponse) SetBody(v *GenerateHumanSketchStyleResponseBody) *GenerateHumanSketchStyleResponse {
	s.Body = v
	return s
}

func (s *GenerateHumanSketchStyleResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
