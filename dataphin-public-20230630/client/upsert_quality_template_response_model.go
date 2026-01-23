// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpsertQualityTemplateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpsertQualityTemplateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpsertQualityTemplateResponse
	GetStatusCode() *int32
	SetBody(v *UpsertQualityTemplateResponseBody) *UpsertQualityTemplateResponse
	GetBody() *UpsertQualityTemplateResponseBody
}

type UpsertQualityTemplateResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpsertQualityTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpsertQualityTemplateResponse) String() string {
	return dara.Prettify(s)
}

func (s UpsertQualityTemplateResponse) GoString() string {
	return s.String()
}

func (s *UpsertQualityTemplateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpsertQualityTemplateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpsertQualityTemplateResponse) GetBody() *UpsertQualityTemplateResponseBody {
	return s.Body
}

func (s *UpsertQualityTemplateResponse) SetHeaders(v map[string]*string) *UpsertQualityTemplateResponse {
	s.Headers = v
	return s
}

func (s *UpsertQualityTemplateResponse) SetStatusCode(v int32) *UpsertQualityTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *UpsertQualityTemplateResponse) SetBody(v *UpsertQualityTemplateResponseBody) *UpsertQualityTemplateResponse {
	s.Body = v
	return s
}

func (s *UpsertQualityTemplateResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
