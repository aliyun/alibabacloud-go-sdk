// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMultimodalLabelStudioServiceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListMultimodalLabelStudioServiceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListMultimodalLabelStudioServiceResponse
	GetStatusCode() *int32
	SetBody(v *ListMultimodalLabelStudioServiceResponseBody) *ListMultimodalLabelStudioServiceResponse
	GetBody() *ListMultimodalLabelStudioServiceResponseBody
}

type ListMultimodalLabelStudioServiceResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListMultimodalLabelStudioServiceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListMultimodalLabelStudioServiceResponse) String() string {
	return dara.Prettify(s)
}

func (s ListMultimodalLabelStudioServiceResponse) GoString() string {
	return s.String()
}

func (s *ListMultimodalLabelStudioServiceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListMultimodalLabelStudioServiceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListMultimodalLabelStudioServiceResponse) GetBody() *ListMultimodalLabelStudioServiceResponseBody {
	return s.Body
}

func (s *ListMultimodalLabelStudioServiceResponse) SetHeaders(v map[string]*string) *ListMultimodalLabelStudioServiceResponse {
	s.Headers = v
	return s
}

func (s *ListMultimodalLabelStudioServiceResponse) SetStatusCode(v int32) *ListMultimodalLabelStudioServiceResponse {
	s.StatusCode = &v
	return s
}

func (s *ListMultimodalLabelStudioServiceResponse) SetBody(v *ListMultimodalLabelStudioServiceResponseBody) *ListMultimodalLabelStudioServiceResponse {
	s.Body = v
	return s
}

func (s *ListMultimodalLabelStudioServiceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
