// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyAppRecordTemplateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyAppRecordTemplateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyAppRecordTemplateResponse
	GetStatusCode() *int32
	SetBody(v *ModifyAppRecordTemplateResponseBody) *ModifyAppRecordTemplateResponse
	GetBody() *ModifyAppRecordTemplateResponseBody
}

type ModifyAppRecordTemplateResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyAppRecordTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyAppRecordTemplateResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyAppRecordTemplateResponse) GoString() string {
	return s.String()
}

func (s *ModifyAppRecordTemplateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyAppRecordTemplateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyAppRecordTemplateResponse) GetBody() *ModifyAppRecordTemplateResponseBody {
	return s.Body
}

func (s *ModifyAppRecordTemplateResponse) SetHeaders(v map[string]*string) *ModifyAppRecordTemplateResponse {
	s.Headers = v
	return s
}

func (s *ModifyAppRecordTemplateResponse) SetStatusCode(v int32) *ModifyAppRecordTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyAppRecordTemplateResponse) SetBody(v *ModifyAppRecordTemplateResponseBody) *ModifyAppRecordTemplateResponse {
	s.Body = v
	return s
}

func (s *ModifyAppRecordTemplateResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
