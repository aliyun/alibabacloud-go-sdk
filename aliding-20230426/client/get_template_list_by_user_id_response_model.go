// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTemplateListByUserIdResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetTemplateListByUserIdResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetTemplateListByUserIdResponse
	GetStatusCode() *int32
	SetBody(v *GetTemplateListByUserIdResponseBody) *GetTemplateListByUserIdResponse
	GetBody() *GetTemplateListByUserIdResponseBody
}

type GetTemplateListByUserIdResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetTemplateListByUserIdResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetTemplateListByUserIdResponse) String() string {
	return dara.Prettify(s)
}

func (s GetTemplateListByUserIdResponse) GoString() string {
	return s.String()
}

func (s *GetTemplateListByUserIdResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetTemplateListByUserIdResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetTemplateListByUserIdResponse) GetBody() *GetTemplateListByUserIdResponseBody {
	return s.Body
}

func (s *GetTemplateListByUserIdResponse) SetHeaders(v map[string]*string) *GetTemplateListByUserIdResponse {
	s.Headers = v
	return s
}

func (s *GetTemplateListByUserIdResponse) SetStatusCode(v int32) *GetTemplateListByUserIdResponse {
	s.StatusCode = &v
	return s
}

func (s *GetTemplateListByUserIdResponse) SetBody(v *GetTemplateListByUserIdResponseBody) *GetTemplateListByUserIdResponse {
	s.Body = v
	return s
}

func (s *GetTemplateListByUserIdResponse) Validate() error {
	return dara.Validate(s)
}
