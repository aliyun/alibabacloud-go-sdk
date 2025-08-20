// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetStackDriftDetectionStatusResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetStackDriftDetectionStatusResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetStackDriftDetectionStatusResponse
	GetStatusCode() *int32
	SetBody(v *GetStackDriftDetectionStatusResponseBody) *GetStackDriftDetectionStatusResponse
	GetBody() *GetStackDriftDetectionStatusResponseBody
}

type GetStackDriftDetectionStatusResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetStackDriftDetectionStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetStackDriftDetectionStatusResponse) String() string {
	return dara.Prettify(s)
}

func (s GetStackDriftDetectionStatusResponse) GoString() string {
	return s.String()
}

func (s *GetStackDriftDetectionStatusResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetStackDriftDetectionStatusResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetStackDriftDetectionStatusResponse) GetBody() *GetStackDriftDetectionStatusResponseBody {
	return s.Body
}

func (s *GetStackDriftDetectionStatusResponse) SetHeaders(v map[string]*string) *GetStackDriftDetectionStatusResponse {
	s.Headers = v
	return s
}

func (s *GetStackDriftDetectionStatusResponse) SetStatusCode(v int32) *GetStackDriftDetectionStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *GetStackDriftDetectionStatusResponse) SetBody(v *GetStackDriftDetectionStatusResponseBody) *GetStackDriftDetectionStatusResponse {
	s.Body = v
	return s
}

func (s *GetStackDriftDetectionStatusResponse) Validate() error {
	return dara.Validate(s)
}
