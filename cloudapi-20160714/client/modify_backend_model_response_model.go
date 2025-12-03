// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyBackendModelResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyBackendModelResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyBackendModelResponse
	GetStatusCode() *int32
	SetBody(v *ModifyBackendModelResponseBody) *ModifyBackendModelResponse
	GetBody() *ModifyBackendModelResponseBody
}

type ModifyBackendModelResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyBackendModelResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyBackendModelResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyBackendModelResponse) GoString() string {
	return s.String()
}

func (s *ModifyBackendModelResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyBackendModelResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyBackendModelResponse) GetBody() *ModifyBackendModelResponseBody {
	return s.Body
}

func (s *ModifyBackendModelResponse) SetHeaders(v map[string]*string) *ModifyBackendModelResponse {
	s.Headers = v
	return s
}

func (s *ModifyBackendModelResponse) SetStatusCode(v int32) *ModifyBackendModelResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyBackendModelResponse) SetBody(v *ModifyBackendModelResponseBody) *ModifyBackendModelResponse {
	s.Body = v
	return s
}

func (s *ModifyBackendModelResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
