// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSampleConsistencyJobResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetSampleConsistencyJobResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetSampleConsistencyJobResponse
	GetStatusCode() *int32
	SetBody(v *GetSampleConsistencyJobResponseBody) *GetSampleConsistencyJobResponse
	GetBody() *GetSampleConsistencyJobResponseBody
}

type GetSampleConsistencyJobResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetSampleConsistencyJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetSampleConsistencyJobResponse) String() string {
	return dara.Prettify(s)
}

func (s GetSampleConsistencyJobResponse) GoString() string {
	return s.String()
}

func (s *GetSampleConsistencyJobResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetSampleConsistencyJobResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetSampleConsistencyJobResponse) GetBody() *GetSampleConsistencyJobResponseBody {
	return s.Body
}

func (s *GetSampleConsistencyJobResponse) SetHeaders(v map[string]*string) *GetSampleConsistencyJobResponse {
	s.Headers = v
	return s
}

func (s *GetSampleConsistencyJobResponse) SetStatusCode(v int32) *GetSampleConsistencyJobResponse {
	s.StatusCode = &v
	return s
}

func (s *GetSampleConsistencyJobResponse) SetBody(v *GetSampleConsistencyJobResponseBody) *GetSampleConsistencyJobResponse {
	s.Body = v
	return s
}

func (s *GetSampleConsistencyJobResponse) Validate() error {
	return dara.Validate(s)
}
