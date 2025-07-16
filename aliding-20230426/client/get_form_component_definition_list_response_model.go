// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetFormComponentDefinitionListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetFormComponentDefinitionListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetFormComponentDefinitionListResponse
	GetStatusCode() *int32
	SetBody(v *GetFormComponentDefinitionListResponseBody) *GetFormComponentDefinitionListResponse
	GetBody() *GetFormComponentDefinitionListResponseBody
}

type GetFormComponentDefinitionListResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetFormComponentDefinitionListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetFormComponentDefinitionListResponse) String() string {
	return dara.Prettify(s)
}

func (s GetFormComponentDefinitionListResponse) GoString() string {
	return s.String()
}

func (s *GetFormComponentDefinitionListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetFormComponentDefinitionListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetFormComponentDefinitionListResponse) GetBody() *GetFormComponentDefinitionListResponseBody {
	return s.Body
}

func (s *GetFormComponentDefinitionListResponse) SetHeaders(v map[string]*string) *GetFormComponentDefinitionListResponse {
	s.Headers = v
	return s
}

func (s *GetFormComponentDefinitionListResponse) SetStatusCode(v int32) *GetFormComponentDefinitionListResponse {
	s.StatusCode = &v
	return s
}

func (s *GetFormComponentDefinitionListResponse) SetBody(v *GetFormComponentDefinitionListResponseBody) *GetFormComponentDefinitionListResponse {
	s.Body = v
	return s
}

func (s *GetFormComponentDefinitionListResponse) Validate() error {
	return dara.Validate(s)
}
