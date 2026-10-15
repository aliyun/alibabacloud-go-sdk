// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iImageRemovalProAsyncResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ImageRemovalProAsyncResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ImageRemovalProAsyncResponse
	GetStatusCode() *int32
	SetBody(v *ImageRemovalProAsyncResponseBody) *ImageRemovalProAsyncResponse
	GetBody() *ImageRemovalProAsyncResponseBody
}

type ImageRemovalProAsyncResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ImageRemovalProAsyncResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ImageRemovalProAsyncResponse) String() string {
	return dara.Prettify(s)
}

func (s ImageRemovalProAsyncResponse) GoString() string {
	return s.String()
}

func (s *ImageRemovalProAsyncResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ImageRemovalProAsyncResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ImageRemovalProAsyncResponse) GetBody() *ImageRemovalProAsyncResponseBody {
	return s.Body
}

func (s *ImageRemovalProAsyncResponse) SetHeaders(v map[string]*string) *ImageRemovalProAsyncResponse {
	s.Headers = v
	return s
}

func (s *ImageRemovalProAsyncResponse) SetStatusCode(v int32) *ImageRemovalProAsyncResponse {
	s.StatusCode = &v
	return s
}

func (s *ImageRemovalProAsyncResponse) SetBody(v *ImageRemovalProAsyncResponseBody) *ImageRemovalProAsyncResponse {
	s.Body = v
	return s
}

func (s *ImageRemovalProAsyncResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
