// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteMultimodalLabelStudioServiceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteMultimodalLabelStudioServiceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteMultimodalLabelStudioServiceResponse
	GetStatusCode() *int32
	SetBody(v *DeleteMultimodalLabelStudioServiceResponseBody) *DeleteMultimodalLabelStudioServiceResponse
	GetBody() *DeleteMultimodalLabelStudioServiceResponseBody
}

type DeleteMultimodalLabelStudioServiceResponse struct {
	Headers    map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteMultimodalLabelStudioServiceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteMultimodalLabelStudioServiceResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteMultimodalLabelStudioServiceResponse) GoString() string {
	return s.String()
}

func (s *DeleteMultimodalLabelStudioServiceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteMultimodalLabelStudioServiceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteMultimodalLabelStudioServiceResponse) GetBody() *DeleteMultimodalLabelStudioServiceResponseBody {
	return s.Body
}

func (s *DeleteMultimodalLabelStudioServiceResponse) SetHeaders(v map[string]*string) *DeleteMultimodalLabelStudioServiceResponse {
	s.Headers = v
	return s
}

func (s *DeleteMultimodalLabelStudioServiceResponse) SetStatusCode(v int32) *DeleteMultimodalLabelStudioServiceResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteMultimodalLabelStudioServiceResponse) SetBody(v *DeleteMultimodalLabelStudioServiceResponseBody) *DeleteMultimodalLabelStudioServiceResponse {
	s.Body = v
	return s
}

func (s *DeleteMultimodalLabelStudioServiceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
