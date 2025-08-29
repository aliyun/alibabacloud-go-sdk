// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSampleConsistencyJobResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateSampleConsistencyJobResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateSampleConsistencyJobResponse
	GetStatusCode() *int32
	SetBody(v *CreateSampleConsistencyJobResponseBody) *CreateSampleConsistencyJobResponse
	GetBody() *CreateSampleConsistencyJobResponseBody
}

type CreateSampleConsistencyJobResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateSampleConsistencyJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateSampleConsistencyJobResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateSampleConsistencyJobResponse) GoString() string {
	return s.String()
}

func (s *CreateSampleConsistencyJobResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateSampleConsistencyJobResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateSampleConsistencyJobResponse) GetBody() *CreateSampleConsistencyJobResponseBody {
	return s.Body
}

func (s *CreateSampleConsistencyJobResponse) SetHeaders(v map[string]*string) *CreateSampleConsistencyJobResponse {
	s.Headers = v
	return s
}

func (s *CreateSampleConsistencyJobResponse) SetStatusCode(v int32) *CreateSampleConsistencyJobResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateSampleConsistencyJobResponse) SetBody(v *CreateSampleConsistencyJobResponseBody) *CreateSampleConsistencyJobResponse {
	s.Body = v
	return s
}

func (s *CreateSampleConsistencyJobResponse) Validate() error {
	return dara.Validate(s)
}
