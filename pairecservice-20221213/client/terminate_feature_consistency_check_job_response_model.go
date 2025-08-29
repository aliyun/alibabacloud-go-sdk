// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTerminateFeatureConsistencyCheckJobResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *TerminateFeatureConsistencyCheckJobResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *TerminateFeatureConsistencyCheckJobResponse
	GetStatusCode() *int32
	SetBody(v *TerminateFeatureConsistencyCheckJobResponseBody) *TerminateFeatureConsistencyCheckJobResponse
	GetBody() *TerminateFeatureConsistencyCheckJobResponseBody
}

type TerminateFeatureConsistencyCheckJobResponse struct {
	Headers    map[string]*string                               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *TerminateFeatureConsistencyCheckJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s TerminateFeatureConsistencyCheckJobResponse) String() string {
	return dara.Prettify(s)
}

func (s TerminateFeatureConsistencyCheckJobResponse) GoString() string {
	return s.String()
}

func (s *TerminateFeatureConsistencyCheckJobResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *TerminateFeatureConsistencyCheckJobResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *TerminateFeatureConsistencyCheckJobResponse) GetBody() *TerminateFeatureConsistencyCheckJobResponseBody {
	return s.Body
}

func (s *TerminateFeatureConsistencyCheckJobResponse) SetHeaders(v map[string]*string) *TerminateFeatureConsistencyCheckJobResponse {
	s.Headers = v
	return s
}

func (s *TerminateFeatureConsistencyCheckJobResponse) SetStatusCode(v int32) *TerminateFeatureConsistencyCheckJobResponse {
	s.StatusCode = &v
	return s
}

func (s *TerminateFeatureConsistencyCheckJobResponse) SetBody(v *TerminateFeatureConsistencyCheckJobResponseBody) *TerminateFeatureConsistencyCheckJobResponse {
	s.Body = v
	return s
}

func (s *TerminateFeatureConsistencyCheckJobResponse) Validate() error {
	return dara.Validate(s)
}
