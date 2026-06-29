// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTaskTemplateViewsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetTaskTemplateViewsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetTaskTemplateViewsResponse
	GetStatusCode() *int32
	SetBody(v *GetTaskTemplateViewsResponseBody) *GetTaskTemplateViewsResponse
	GetBody() *GetTaskTemplateViewsResponseBody
}

type GetTaskTemplateViewsResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetTaskTemplateViewsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetTaskTemplateViewsResponse) String() string {
	return dara.Prettify(s)
}

func (s GetTaskTemplateViewsResponse) GoString() string {
	return s.String()
}

func (s *GetTaskTemplateViewsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetTaskTemplateViewsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetTaskTemplateViewsResponse) GetBody() *GetTaskTemplateViewsResponseBody {
	return s.Body
}

func (s *GetTaskTemplateViewsResponse) SetHeaders(v map[string]*string) *GetTaskTemplateViewsResponse {
	s.Headers = v
	return s
}

func (s *GetTaskTemplateViewsResponse) SetStatusCode(v int32) *GetTaskTemplateViewsResponse {
	s.StatusCode = &v
	return s
}

func (s *GetTaskTemplateViewsResponse) SetBody(v *GetTaskTemplateViewsResponseBody) *GetTaskTemplateViewsResponse {
	s.Body = v
	return s
}

func (s *GetTaskTemplateViewsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
