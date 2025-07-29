// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetEssayCorrectionTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetEssayCorrectionTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetEssayCorrectionTaskResponse
	GetStatusCode() *int32
	SetBody(v *GetEssayCorrectionTaskResponseBody) *GetEssayCorrectionTaskResponse
	GetBody() *GetEssayCorrectionTaskResponseBody
}

type GetEssayCorrectionTaskResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetEssayCorrectionTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetEssayCorrectionTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s GetEssayCorrectionTaskResponse) GoString() string {
	return s.String()
}

func (s *GetEssayCorrectionTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetEssayCorrectionTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetEssayCorrectionTaskResponse) GetBody() *GetEssayCorrectionTaskResponseBody {
	return s.Body
}

func (s *GetEssayCorrectionTaskResponse) SetHeaders(v map[string]*string) *GetEssayCorrectionTaskResponse {
	s.Headers = v
	return s
}

func (s *GetEssayCorrectionTaskResponse) SetStatusCode(v int32) *GetEssayCorrectionTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *GetEssayCorrectionTaskResponse) SetBody(v *GetEssayCorrectionTaskResponseBody) *GetEssayCorrectionTaskResponse {
	s.Body = v
	return s
}

func (s *GetEssayCorrectionTaskResponse) Validate() error {
	return dara.Validate(s)
}
