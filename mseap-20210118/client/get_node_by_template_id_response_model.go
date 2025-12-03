// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetNodeByTemplateIdResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetNodeByTemplateIdResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetNodeByTemplateIdResponse
	GetStatusCode() *int32
	SetBody(v *GetNodeByTemplateIdResponseBody) *GetNodeByTemplateIdResponse
	GetBody() *GetNodeByTemplateIdResponseBody
}

type GetNodeByTemplateIdResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetNodeByTemplateIdResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetNodeByTemplateIdResponse) String() string {
	return dara.Prettify(s)
}

func (s GetNodeByTemplateIdResponse) GoString() string {
	return s.String()
}

func (s *GetNodeByTemplateIdResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetNodeByTemplateIdResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetNodeByTemplateIdResponse) GetBody() *GetNodeByTemplateIdResponseBody {
	return s.Body
}

func (s *GetNodeByTemplateIdResponse) SetHeaders(v map[string]*string) *GetNodeByTemplateIdResponse {
	s.Headers = v
	return s
}

func (s *GetNodeByTemplateIdResponse) SetStatusCode(v int32) *GetNodeByTemplateIdResponse {
	s.StatusCode = &v
	return s
}

func (s *GetNodeByTemplateIdResponse) SetBody(v *GetNodeByTemplateIdResponseBody) *GetNodeByTemplateIdResponse {
	s.Body = v
	return s
}

func (s *GetNodeByTemplateIdResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
