// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetOrderAttachmentFileResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetOrderAttachmentFileResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetOrderAttachmentFileResponse
	GetStatusCode() *int32
	SetBody(v *GetOrderAttachmentFileResponseBody) *GetOrderAttachmentFileResponse
	GetBody() *GetOrderAttachmentFileResponseBody
}

type GetOrderAttachmentFileResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetOrderAttachmentFileResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetOrderAttachmentFileResponse) String() string {
	return dara.Prettify(s)
}

func (s GetOrderAttachmentFileResponse) GoString() string {
	return s.String()
}

func (s *GetOrderAttachmentFileResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetOrderAttachmentFileResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetOrderAttachmentFileResponse) GetBody() *GetOrderAttachmentFileResponseBody {
	return s.Body
}

func (s *GetOrderAttachmentFileResponse) SetHeaders(v map[string]*string) *GetOrderAttachmentFileResponse {
	s.Headers = v
	return s
}

func (s *GetOrderAttachmentFileResponse) SetStatusCode(v int32) *GetOrderAttachmentFileResponse {
	s.StatusCode = &v
	return s
}

func (s *GetOrderAttachmentFileResponse) SetBody(v *GetOrderAttachmentFileResponseBody) *GetOrderAttachmentFileResponse {
	s.Body = v
	return s
}

func (s *GetOrderAttachmentFileResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
