// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAdHocTaskResultResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetAdHocTaskResultResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetAdHocTaskResultResponse
	GetStatusCode() *int32
	SetBody(v *GetAdHocTaskResultResponseBody) *GetAdHocTaskResultResponse
	GetBody() *GetAdHocTaskResultResponseBody
}

type GetAdHocTaskResultResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAdHocTaskResultResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAdHocTaskResultResponse) String() string {
	return dara.Prettify(s)
}

func (s GetAdHocTaskResultResponse) GoString() string {
	return s.String()
}

func (s *GetAdHocTaskResultResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetAdHocTaskResultResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetAdHocTaskResultResponse) GetBody() *GetAdHocTaskResultResponseBody {
	return s.Body
}

func (s *GetAdHocTaskResultResponse) SetHeaders(v map[string]*string) *GetAdHocTaskResultResponse {
	s.Headers = v
	return s
}

func (s *GetAdHocTaskResultResponse) SetStatusCode(v int32) *GetAdHocTaskResultResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAdHocTaskResultResponse) SetBody(v *GetAdHocTaskResultResponseBody) *GetAdHocTaskResultResponse {
	s.Body = v
	return s
}

func (s *GetAdHocTaskResultResponse) Validate() error {
	return dara.Validate(s)
}
