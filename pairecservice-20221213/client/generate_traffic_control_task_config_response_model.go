// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGenerateTrafficControlTaskConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GenerateTrafficControlTaskConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GenerateTrafficControlTaskConfigResponse
	GetStatusCode() *int32
	SetBody(v *GenerateTrafficControlTaskConfigResponseBody) *GenerateTrafficControlTaskConfigResponse
	GetBody() *GenerateTrafficControlTaskConfigResponseBody
}

type GenerateTrafficControlTaskConfigResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GenerateTrafficControlTaskConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GenerateTrafficControlTaskConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s GenerateTrafficControlTaskConfigResponse) GoString() string {
	return s.String()
}

func (s *GenerateTrafficControlTaskConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GenerateTrafficControlTaskConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GenerateTrafficControlTaskConfigResponse) GetBody() *GenerateTrafficControlTaskConfigResponseBody {
	return s.Body
}

func (s *GenerateTrafficControlTaskConfigResponse) SetHeaders(v map[string]*string) *GenerateTrafficControlTaskConfigResponse {
	s.Headers = v
	return s
}

func (s *GenerateTrafficControlTaskConfigResponse) SetStatusCode(v int32) *GenerateTrafficControlTaskConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *GenerateTrafficControlTaskConfigResponse) SetBody(v *GenerateTrafficControlTaskConfigResponseBody) *GenerateTrafficControlTaskConfigResponse {
	s.Body = v
	return s
}

func (s *GenerateTrafficControlTaskConfigResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
