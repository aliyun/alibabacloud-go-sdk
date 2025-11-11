// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSmartClipTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetSmartClipTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetSmartClipTaskResponse
	GetStatusCode() *int32
	SetBody(v *GetSmartClipTaskResponseBody) *GetSmartClipTaskResponse
	GetBody() *GetSmartClipTaskResponseBody
}

type GetSmartClipTaskResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetSmartClipTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetSmartClipTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s GetSmartClipTaskResponse) GoString() string {
	return s.String()
}

func (s *GetSmartClipTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetSmartClipTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetSmartClipTaskResponse) GetBody() *GetSmartClipTaskResponseBody {
	return s.Body
}

func (s *GetSmartClipTaskResponse) SetHeaders(v map[string]*string) *GetSmartClipTaskResponse {
	s.Headers = v
	return s
}

func (s *GetSmartClipTaskResponse) SetStatusCode(v int32) *GetSmartClipTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *GetSmartClipTaskResponse) SetBody(v *GetSmartClipTaskResponseBody) *GetSmartClipTaskResponse {
	s.Body = v
	return s
}

func (s *GetSmartClipTaskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
