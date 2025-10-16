// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyTemplateBaseInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyTemplateBaseInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyTemplateBaseInfoResponse
	GetStatusCode() *int32
	SetBody(v *ModifyTemplateBaseInfoResponseBody) *ModifyTemplateBaseInfoResponse
	GetBody() *ModifyTemplateBaseInfoResponseBody
}

type ModifyTemplateBaseInfoResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyTemplateBaseInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyTemplateBaseInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyTemplateBaseInfoResponse) GoString() string {
	return s.String()
}

func (s *ModifyTemplateBaseInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyTemplateBaseInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyTemplateBaseInfoResponse) GetBody() *ModifyTemplateBaseInfoResponseBody {
	return s.Body
}

func (s *ModifyTemplateBaseInfoResponse) SetHeaders(v map[string]*string) *ModifyTemplateBaseInfoResponse {
	s.Headers = v
	return s
}

func (s *ModifyTemplateBaseInfoResponse) SetStatusCode(v int32) *ModifyTemplateBaseInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyTemplateBaseInfoResponse) SetBody(v *ModifyTemplateBaseInfoResponseBody) *ModifyTemplateBaseInfoResponse {
	s.Body = v
	return s
}

func (s *ModifyTemplateBaseInfoResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
