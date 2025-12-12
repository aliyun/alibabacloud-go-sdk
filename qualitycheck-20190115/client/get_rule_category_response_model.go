// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetRuleCategoryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetRuleCategoryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetRuleCategoryResponse
	GetStatusCode() *int32
	SetBody(v *GetRuleCategoryResponseBody) *GetRuleCategoryResponse
	GetBody() *GetRuleCategoryResponseBody
}

type GetRuleCategoryResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetRuleCategoryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetRuleCategoryResponse) String() string {
	return dara.Prettify(s)
}

func (s GetRuleCategoryResponse) GoString() string {
	return s.String()
}

func (s *GetRuleCategoryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetRuleCategoryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetRuleCategoryResponse) GetBody() *GetRuleCategoryResponseBody {
	return s.Body
}

func (s *GetRuleCategoryResponse) SetHeaders(v map[string]*string) *GetRuleCategoryResponse {
	s.Headers = v
	return s
}

func (s *GetRuleCategoryResponse) SetStatusCode(v int32) *GetRuleCategoryResponse {
	s.StatusCode = &v
	return s
}

func (s *GetRuleCategoryResponse) SetBody(v *GetRuleCategoryResponseBody) *GetRuleCategoryResponse {
	s.Body = v
	return s
}

func (s *GetRuleCategoryResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
