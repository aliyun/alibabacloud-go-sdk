// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListHotlineTransferRegisterFileResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListHotlineTransferRegisterFileResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListHotlineTransferRegisterFileResponse
	GetStatusCode() *int32
	SetBody(v *ListHotlineTransferRegisterFileResponseBody) *ListHotlineTransferRegisterFileResponse
	GetBody() *ListHotlineTransferRegisterFileResponseBody
}

type ListHotlineTransferRegisterFileResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListHotlineTransferRegisterFileResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListHotlineTransferRegisterFileResponse) String() string {
	return dara.Prettify(s)
}

func (s ListHotlineTransferRegisterFileResponse) GoString() string {
	return s.String()
}

func (s *ListHotlineTransferRegisterFileResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListHotlineTransferRegisterFileResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListHotlineTransferRegisterFileResponse) GetBody() *ListHotlineTransferRegisterFileResponseBody {
	return s.Body
}

func (s *ListHotlineTransferRegisterFileResponse) SetHeaders(v map[string]*string) *ListHotlineTransferRegisterFileResponse {
	s.Headers = v
	return s
}

func (s *ListHotlineTransferRegisterFileResponse) SetStatusCode(v int32) *ListHotlineTransferRegisterFileResponse {
	s.StatusCode = &v
	return s
}

func (s *ListHotlineTransferRegisterFileResponse) SetBody(v *ListHotlineTransferRegisterFileResponseBody) *ListHotlineTransferRegisterFileResponse {
	s.Body = v
	return s
}

func (s *ListHotlineTransferRegisterFileResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
