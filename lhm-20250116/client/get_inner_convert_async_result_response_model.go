// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetInnerConvertAsyncResultResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetInnerConvertAsyncResultResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetInnerConvertAsyncResultResponse
	GetStatusCode() *int32
	SetBody(v *GetInnerConvertAsyncResultResponseBody) *GetInnerConvertAsyncResultResponse
	GetBody() *GetInnerConvertAsyncResultResponseBody
}

type GetInnerConvertAsyncResultResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetInnerConvertAsyncResultResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetInnerConvertAsyncResultResponse) String() string {
	return dara.Prettify(s)
}

func (s GetInnerConvertAsyncResultResponse) GoString() string {
	return s.String()
}

func (s *GetInnerConvertAsyncResultResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetInnerConvertAsyncResultResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetInnerConvertAsyncResultResponse) GetBody() *GetInnerConvertAsyncResultResponseBody {
	return s.Body
}

func (s *GetInnerConvertAsyncResultResponse) SetHeaders(v map[string]*string) *GetInnerConvertAsyncResultResponse {
	s.Headers = v
	return s
}

func (s *GetInnerConvertAsyncResultResponse) SetStatusCode(v int32) *GetInnerConvertAsyncResultResponse {
	s.StatusCode = &v
	return s
}

func (s *GetInnerConvertAsyncResultResponse) SetBody(v *GetInnerConvertAsyncResultResponseBody) *GetInnerConvertAsyncResultResponse {
	s.Body = v
	return s
}

func (s *GetInnerConvertAsyncResultResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
