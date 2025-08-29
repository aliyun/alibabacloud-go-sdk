// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOfflineExperimentGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *OfflineExperimentGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *OfflineExperimentGroupResponse
	GetStatusCode() *int32
	SetBody(v *OfflineExperimentGroupResponseBody) *OfflineExperimentGroupResponse
	GetBody() *OfflineExperimentGroupResponseBody
}

type OfflineExperimentGroupResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *OfflineExperimentGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s OfflineExperimentGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s OfflineExperimentGroupResponse) GoString() string {
	return s.String()
}

func (s *OfflineExperimentGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *OfflineExperimentGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *OfflineExperimentGroupResponse) GetBody() *OfflineExperimentGroupResponseBody {
	return s.Body
}

func (s *OfflineExperimentGroupResponse) SetHeaders(v map[string]*string) *OfflineExperimentGroupResponse {
	s.Headers = v
	return s
}

func (s *OfflineExperimentGroupResponse) SetStatusCode(v int32) *OfflineExperimentGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *OfflineExperimentGroupResponse) SetBody(v *OfflineExperimentGroupResponseBody) *OfflineExperimentGroupResponse {
	s.Body = v
	return s
}

func (s *OfflineExperimentGroupResponse) Validate() error {
	return dara.Validate(s)
}
