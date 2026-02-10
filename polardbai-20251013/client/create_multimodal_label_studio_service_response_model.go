// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateMultimodalLabelStudioServiceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateMultimodalLabelStudioServiceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateMultimodalLabelStudioServiceResponse
	GetStatusCode() *int32
	SetBody(v *CreateMultimodalLabelStudioServiceResponseBody) *CreateMultimodalLabelStudioServiceResponse
	GetBody() *CreateMultimodalLabelStudioServiceResponseBody
}

type CreateMultimodalLabelStudioServiceResponse struct {
	Headers    map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateMultimodalLabelStudioServiceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateMultimodalLabelStudioServiceResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateMultimodalLabelStudioServiceResponse) GoString() string {
	return s.String()
}

func (s *CreateMultimodalLabelStudioServiceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateMultimodalLabelStudioServiceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateMultimodalLabelStudioServiceResponse) GetBody() *CreateMultimodalLabelStudioServiceResponseBody {
	return s.Body
}

func (s *CreateMultimodalLabelStudioServiceResponse) SetHeaders(v map[string]*string) *CreateMultimodalLabelStudioServiceResponse {
	s.Headers = v
	return s
}

func (s *CreateMultimodalLabelStudioServiceResponse) SetStatusCode(v int32) *CreateMultimodalLabelStudioServiceResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateMultimodalLabelStudioServiceResponse) SetBody(v *CreateMultimodalLabelStudioServiceResponseBody) *CreateMultimodalLabelStudioServiceResponse {
	s.Body = v
	return s
}

func (s *CreateMultimodalLabelStudioServiceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
