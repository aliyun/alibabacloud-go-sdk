// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateMultimodalLabelStudioServiceWhiteListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateMultimodalLabelStudioServiceWhiteListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateMultimodalLabelStudioServiceWhiteListResponse
	GetStatusCode() *int32
	SetBody(v *UpdateMultimodalLabelStudioServiceWhiteListResponseBody) *UpdateMultimodalLabelStudioServiceWhiteListResponse
	GetBody() *UpdateMultimodalLabelStudioServiceWhiteListResponseBody
}

type UpdateMultimodalLabelStudioServiceWhiteListResponse struct {
	Headers    map[string]*string                                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateMultimodalLabelStudioServiceWhiteListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateMultimodalLabelStudioServiceWhiteListResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateMultimodalLabelStudioServiceWhiteListResponse) GoString() string {
	return s.String()
}

func (s *UpdateMultimodalLabelStudioServiceWhiteListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateMultimodalLabelStudioServiceWhiteListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateMultimodalLabelStudioServiceWhiteListResponse) GetBody() *UpdateMultimodalLabelStudioServiceWhiteListResponseBody {
	return s.Body
}

func (s *UpdateMultimodalLabelStudioServiceWhiteListResponse) SetHeaders(v map[string]*string) *UpdateMultimodalLabelStudioServiceWhiteListResponse {
	s.Headers = v
	return s
}

func (s *UpdateMultimodalLabelStudioServiceWhiteListResponse) SetStatusCode(v int32) *UpdateMultimodalLabelStudioServiceWhiteListResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateMultimodalLabelStudioServiceWhiteListResponse) SetBody(v *UpdateMultimodalLabelStudioServiceWhiteListResponseBody) *UpdateMultimodalLabelStudioServiceWhiteListResponse {
	s.Body = v
	return s
}

func (s *UpdateMultimodalLabelStudioServiceWhiteListResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
