// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetBatchTemplateVersionsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetBatchTemplateVersionsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetBatchTemplateVersionsResponse
	GetStatusCode() *int32
	SetBody(v *GetBatchTemplateVersionsResponseBody) *GetBatchTemplateVersionsResponse
	GetBody() *GetBatchTemplateVersionsResponseBody
}

type GetBatchTemplateVersionsResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetBatchTemplateVersionsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetBatchTemplateVersionsResponse) String() string {
	return dara.Prettify(s)
}

func (s GetBatchTemplateVersionsResponse) GoString() string {
	return s.String()
}

func (s *GetBatchTemplateVersionsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetBatchTemplateVersionsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetBatchTemplateVersionsResponse) GetBody() *GetBatchTemplateVersionsResponseBody {
	return s.Body
}

func (s *GetBatchTemplateVersionsResponse) SetHeaders(v map[string]*string) *GetBatchTemplateVersionsResponse {
	s.Headers = v
	return s
}

func (s *GetBatchTemplateVersionsResponse) SetStatusCode(v int32) *GetBatchTemplateVersionsResponse {
	s.StatusCode = &v
	return s
}

func (s *GetBatchTemplateVersionsResponse) SetBody(v *GetBatchTemplateVersionsResponseBody) *GetBatchTemplateVersionsResponse {
	s.Body = v
	return s
}

func (s *GetBatchTemplateVersionsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
