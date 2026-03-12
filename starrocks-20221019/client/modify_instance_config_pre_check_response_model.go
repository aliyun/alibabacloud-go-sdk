// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyInstanceConfigPreCheckResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyInstanceConfigPreCheckResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyInstanceConfigPreCheckResponse
	GetStatusCode() *int32
	SetBody(v *ModifyInstanceConfigPreCheckResponseBody) *ModifyInstanceConfigPreCheckResponse
	GetBody() *ModifyInstanceConfigPreCheckResponseBody
}

type ModifyInstanceConfigPreCheckResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyInstanceConfigPreCheckResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyInstanceConfigPreCheckResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyInstanceConfigPreCheckResponse) GoString() string {
	return s.String()
}

func (s *ModifyInstanceConfigPreCheckResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyInstanceConfigPreCheckResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyInstanceConfigPreCheckResponse) GetBody() *ModifyInstanceConfigPreCheckResponseBody {
	return s.Body
}

func (s *ModifyInstanceConfigPreCheckResponse) SetHeaders(v map[string]*string) *ModifyInstanceConfigPreCheckResponse {
	s.Headers = v
	return s
}

func (s *ModifyInstanceConfigPreCheckResponse) SetStatusCode(v int32) *ModifyInstanceConfigPreCheckResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyInstanceConfigPreCheckResponse) SetBody(v *ModifyInstanceConfigPreCheckResponseBody) *ModifyInstanceConfigPreCheckResponse {
	s.Body = v
	return s
}

func (s *ModifyInstanceConfigPreCheckResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
