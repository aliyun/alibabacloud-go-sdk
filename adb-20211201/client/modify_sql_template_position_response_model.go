// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifySqlTemplatePositionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifySqlTemplatePositionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifySqlTemplatePositionResponse
	GetStatusCode() *int32
	SetBody(v *ModifySqlTemplatePositionResponseBody) *ModifySqlTemplatePositionResponse
	GetBody() *ModifySqlTemplatePositionResponseBody
}

type ModifySqlTemplatePositionResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifySqlTemplatePositionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifySqlTemplatePositionResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifySqlTemplatePositionResponse) GoString() string {
	return s.String()
}

func (s *ModifySqlTemplatePositionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifySqlTemplatePositionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifySqlTemplatePositionResponse) GetBody() *ModifySqlTemplatePositionResponseBody {
	return s.Body
}

func (s *ModifySqlTemplatePositionResponse) SetHeaders(v map[string]*string) *ModifySqlTemplatePositionResponse {
	s.Headers = v
	return s
}

func (s *ModifySqlTemplatePositionResponse) SetStatusCode(v int32) *ModifySqlTemplatePositionResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifySqlTemplatePositionResponse) SetBody(v *ModifySqlTemplatePositionResponseBody) *ModifySqlTemplatePositionResponse {
	s.Body = v
	return s
}

func (s *ModifySqlTemplatePositionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
