// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyCallbackMetaResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyCallbackMetaResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyCallbackMetaResponse
	GetStatusCode() *int32
	SetBody(v *ModifyCallbackMetaResponseBody) *ModifyCallbackMetaResponse
	GetBody() *ModifyCallbackMetaResponseBody
}

type ModifyCallbackMetaResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyCallbackMetaResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyCallbackMetaResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyCallbackMetaResponse) GoString() string {
	return s.String()
}

func (s *ModifyCallbackMetaResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyCallbackMetaResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyCallbackMetaResponse) GetBody() *ModifyCallbackMetaResponseBody {
	return s.Body
}

func (s *ModifyCallbackMetaResponse) SetHeaders(v map[string]*string) *ModifyCallbackMetaResponse {
	s.Headers = v
	return s
}

func (s *ModifyCallbackMetaResponse) SetStatusCode(v int32) *ModifyCallbackMetaResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyCallbackMetaResponse) SetBody(v *ModifyCallbackMetaResponseBody) *ModifyCallbackMetaResponse {
	s.Body = v
	return s
}

func (s *ModifyCallbackMetaResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
