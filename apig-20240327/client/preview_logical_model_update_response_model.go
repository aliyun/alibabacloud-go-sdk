// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPreviewLogicalModelUpdateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *PreviewLogicalModelUpdateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *PreviewLogicalModelUpdateResponse
	GetStatusCode() *int32
	SetBody(v *PreviewLogicalModelUpdateResponseBody) *PreviewLogicalModelUpdateResponse
	GetBody() *PreviewLogicalModelUpdateResponseBody
}

type PreviewLogicalModelUpdateResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PreviewLogicalModelUpdateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PreviewLogicalModelUpdateResponse) String() string {
	return dara.Prettify(s)
}

func (s PreviewLogicalModelUpdateResponse) GoString() string {
	return s.String()
}

func (s *PreviewLogicalModelUpdateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *PreviewLogicalModelUpdateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *PreviewLogicalModelUpdateResponse) GetBody() *PreviewLogicalModelUpdateResponseBody {
	return s.Body
}

func (s *PreviewLogicalModelUpdateResponse) SetHeaders(v map[string]*string) *PreviewLogicalModelUpdateResponse {
	s.Headers = v
	return s
}

func (s *PreviewLogicalModelUpdateResponse) SetStatusCode(v int32) *PreviewLogicalModelUpdateResponse {
	s.StatusCode = &v
	return s
}

func (s *PreviewLogicalModelUpdateResponse) SetBody(v *PreviewLogicalModelUpdateResponseBody) *PreviewLogicalModelUpdateResponse {
	s.Body = v
	return s
}

func (s *PreviewLogicalModelUpdateResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
