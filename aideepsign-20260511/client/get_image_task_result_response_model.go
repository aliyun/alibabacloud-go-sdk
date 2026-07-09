// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetImageTaskResultResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetImageTaskResultResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetImageTaskResultResponse
	GetStatusCode() *int32
	SetBody(v *GetImageTaskResultResponseBody) *GetImageTaskResultResponse
	GetBody() *GetImageTaskResultResponseBody
}

type GetImageTaskResultResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetImageTaskResultResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetImageTaskResultResponse) String() string {
	return dara.Prettify(s)
}

func (s GetImageTaskResultResponse) GoString() string {
	return s.String()
}

func (s *GetImageTaskResultResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetImageTaskResultResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetImageTaskResultResponse) GetBody() *GetImageTaskResultResponseBody {
	return s.Body
}

func (s *GetImageTaskResultResponse) SetHeaders(v map[string]*string) *GetImageTaskResultResponse {
	s.Headers = v
	return s
}

func (s *GetImageTaskResultResponse) SetStatusCode(v int32) *GetImageTaskResultResponse {
	s.StatusCode = &v
	return s
}

func (s *GetImageTaskResultResponse) SetBody(v *GetImageTaskResultResponseBody) *GetImageTaskResultResponse {
	s.Body = v
	return s
}

func (s *GetImageTaskResultResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
