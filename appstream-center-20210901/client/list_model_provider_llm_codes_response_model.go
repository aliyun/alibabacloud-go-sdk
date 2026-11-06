// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListModelProviderLlmCodesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListModelProviderLlmCodesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListModelProviderLlmCodesResponse
	GetStatusCode() *int32
	SetBody(v *ListModelProviderLlmCodesResponseBody) *ListModelProviderLlmCodesResponse
	GetBody() *ListModelProviderLlmCodesResponseBody
}

type ListModelProviderLlmCodesResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListModelProviderLlmCodesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListModelProviderLlmCodesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListModelProviderLlmCodesResponse) GoString() string {
	return s.String()
}

func (s *ListModelProviderLlmCodesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListModelProviderLlmCodesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListModelProviderLlmCodesResponse) GetBody() *ListModelProviderLlmCodesResponseBody {
	return s.Body
}

func (s *ListModelProviderLlmCodesResponse) SetHeaders(v map[string]*string) *ListModelProviderLlmCodesResponse {
	s.Headers = v
	return s
}

func (s *ListModelProviderLlmCodesResponse) SetStatusCode(v int32) *ListModelProviderLlmCodesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListModelProviderLlmCodesResponse) SetBody(v *ListModelProviderLlmCodesResponseBody) *ListModelProviderLlmCodesResponse {
	s.Body = v
	return s
}

func (s *ListModelProviderLlmCodesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
