// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSaveReceiverDetailResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SaveReceiverDetailResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SaveReceiverDetailResponse
	GetStatusCode() *int32
	SetBody(v *SaveReceiverDetailResponseBody) *SaveReceiverDetailResponse
	GetBody() *SaveReceiverDetailResponseBody
}

type SaveReceiverDetailResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SaveReceiverDetailResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SaveReceiverDetailResponse) String() string {
	return dara.Prettify(s)
}

func (s SaveReceiverDetailResponse) GoString() string {
	return s.String()
}

func (s *SaveReceiverDetailResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SaveReceiverDetailResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SaveReceiverDetailResponse) GetBody() *SaveReceiverDetailResponseBody {
	return s.Body
}

func (s *SaveReceiverDetailResponse) SetHeaders(v map[string]*string) *SaveReceiverDetailResponse {
	s.Headers = v
	return s
}

func (s *SaveReceiverDetailResponse) SetStatusCode(v int32) *SaveReceiverDetailResponse {
	s.StatusCode = &v
	return s
}

func (s *SaveReceiverDetailResponse) SetBody(v *SaveReceiverDetailResponseBody) *SaveReceiverDetailResponse {
	s.Body = v
	return s
}

func (s *SaveReceiverDetailResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
