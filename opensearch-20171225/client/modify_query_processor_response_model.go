// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyQueryProcessorResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyQueryProcessorResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyQueryProcessorResponse
	GetStatusCode() *int32
	SetBody(v *ModifyQueryProcessorResponseBody) *ModifyQueryProcessorResponse
	GetBody() *ModifyQueryProcessorResponseBody
}

type ModifyQueryProcessorResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyQueryProcessorResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyQueryProcessorResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyQueryProcessorResponse) GoString() string {
	return s.String()
}

func (s *ModifyQueryProcessorResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyQueryProcessorResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyQueryProcessorResponse) GetBody() *ModifyQueryProcessorResponseBody {
	return s.Body
}

func (s *ModifyQueryProcessorResponse) SetHeaders(v map[string]*string) *ModifyQueryProcessorResponse {
	s.Headers = v
	return s
}

func (s *ModifyQueryProcessorResponse) SetStatusCode(v int32) *ModifyQueryProcessorResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyQueryProcessorResponse) SetBody(v *ModifyQueryProcessorResponseBody) *ModifyQueryProcessorResponse {
	s.Body = v
	return s
}

func (s *ModifyQueryProcessorResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
