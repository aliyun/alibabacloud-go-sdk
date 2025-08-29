// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteSampleConsistencyJobResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteSampleConsistencyJobResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteSampleConsistencyJobResponse
	GetStatusCode() *int32
	SetBody(v *DeleteSampleConsistencyJobResponseBody) *DeleteSampleConsistencyJobResponse
	GetBody() *DeleteSampleConsistencyJobResponseBody
}

type DeleteSampleConsistencyJobResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteSampleConsistencyJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteSampleConsistencyJobResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteSampleConsistencyJobResponse) GoString() string {
	return s.String()
}

func (s *DeleteSampleConsistencyJobResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteSampleConsistencyJobResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteSampleConsistencyJobResponse) GetBody() *DeleteSampleConsistencyJobResponseBody {
	return s.Body
}

func (s *DeleteSampleConsistencyJobResponse) SetHeaders(v map[string]*string) *DeleteSampleConsistencyJobResponse {
	s.Headers = v
	return s
}

func (s *DeleteSampleConsistencyJobResponse) SetStatusCode(v int32) *DeleteSampleConsistencyJobResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteSampleConsistencyJobResponse) SetBody(v *DeleteSampleConsistencyJobResponseBody) *DeleteSampleConsistencyJobResponse {
	s.Body = v
	return s
}

func (s *DeleteSampleConsistencyJobResponse) Validate() error {
	return dara.Validate(s)
}
