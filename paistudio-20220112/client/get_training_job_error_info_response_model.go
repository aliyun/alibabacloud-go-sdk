// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTrainingJobErrorInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetTrainingJobErrorInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetTrainingJobErrorInfoResponse
	GetStatusCode() *int32
	SetBody(v *GetTrainingJobErrorInfoResponseBody) *GetTrainingJobErrorInfoResponse
	GetBody() *GetTrainingJobErrorInfoResponseBody
}

type GetTrainingJobErrorInfoResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetTrainingJobErrorInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetTrainingJobErrorInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s GetTrainingJobErrorInfoResponse) GoString() string {
	return s.String()
}

func (s *GetTrainingJobErrorInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetTrainingJobErrorInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetTrainingJobErrorInfoResponse) GetBody() *GetTrainingJobErrorInfoResponseBody {
	return s.Body
}

func (s *GetTrainingJobErrorInfoResponse) SetHeaders(v map[string]*string) *GetTrainingJobErrorInfoResponse {
	s.Headers = v
	return s
}

func (s *GetTrainingJobErrorInfoResponse) SetStatusCode(v int32) *GetTrainingJobErrorInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *GetTrainingJobErrorInfoResponse) SetBody(v *GetTrainingJobErrorInfoResponseBody) *GetTrainingJobErrorInfoResponse {
	s.Body = v
	return s
}

func (s *GetTrainingJobErrorInfoResponse) Validate() error {
	return dara.Validate(s)
}
