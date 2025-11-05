// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartReplicaGroupDrillResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *StartReplicaGroupDrillResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *StartReplicaGroupDrillResponse
	GetStatusCode() *int32
	SetBody(v *StartReplicaGroupDrillResponseBody) *StartReplicaGroupDrillResponse
	GetBody() *StartReplicaGroupDrillResponseBody
}

type StartReplicaGroupDrillResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StartReplicaGroupDrillResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StartReplicaGroupDrillResponse) String() string {
	return dara.Prettify(s)
}

func (s StartReplicaGroupDrillResponse) GoString() string {
	return s.String()
}

func (s *StartReplicaGroupDrillResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *StartReplicaGroupDrillResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *StartReplicaGroupDrillResponse) GetBody() *StartReplicaGroupDrillResponseBody {
	return s.Body
}

func (s *StartReplicaGroupDrillResponse) SetHeaders(v map[string]*string) *StartReplicaGroupDrillResponse {
	s.Headers = v
	return s
}

func (s *StartReplicaGroupDrillResponse) SetStatusCode(v int32) *StartReplicaGroupDrillResponse {
	s.StatusCode = &v
	return s
}

func (s *StartReplicaGroupDrillResponse) SetBody(v *StartReplicaGroupDrillResponseBody) *StartReplicaGroupDrillResponse {
	s.Body = v
	return s
}

func (s *StartReplicaGroupDrillResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
