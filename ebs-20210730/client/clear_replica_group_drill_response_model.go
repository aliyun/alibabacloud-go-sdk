// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iClearReplicaGroupDrillResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ClearReplicaGroupDrillResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ClearReplicaGroupDrillResponse
	GetStatusCode() *int32
	SetBody(v *ClearReplicaGroupDrillResponseBody) *ClearReplicaGroupDrillResponse
	GetBody() *ClearReplicaGroupDrillResponseBody
}

type ClearReplicaGroupDrillResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ClearReplicaGroupDrillResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ClearReplicaGroupDrillResponse) String() string {
	return dara.Prettify(s)
}

func (s ClearReplicaGroupDrillResponse) GoString() string {
	return s.String()
}

func (s *ClearReplicaGroupDrillResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ClearReplicaGroupDrillResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ClearReplicaGroupDrillResponse) GetBody() *ClearReplicaGroupDrillResponseBody {
	return s.Body
}

func (s *ClearReplicaGroupDrillResponse) SetHeaders(v map[string]*string) *ClearReplicaGroupDrillResponse {
	s.Headers = v
	return s
}

func (s *ClearReplicaGroupDrillResponse) SetStatusCode(v int32) *ClearReplicaGroupDrillResponse {
	s.StatusCode = &v
	return s
}

func (s *ClearReplicaGroupDrillResponse) SetBody(v *ClearReplicaGroupDrillResponseBody) *ClearReplicaGroupDrillResponse {
	s.Body = v
	return s
}

func (s *ClearReplicaGroupDrillResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
