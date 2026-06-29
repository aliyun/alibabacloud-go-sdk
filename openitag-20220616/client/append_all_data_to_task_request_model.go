// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAppendAllDataToTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v *OpenDatasetProxyAppendDataRequest) *AppendAllDataToTaskRequest
	GetBody() *OpenDatasetProxyAppendDataRequest
}

type AppendAllDataToTaskRequest struct {
	// Parameters.
	Body *OpenDatasetProxyAppendDataRequest `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AppendAllDataToTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s AppendAllDataToTaskRequest) GoString() string {
	return s.String()
}

func (s *AppendAllDataToTaskRequest) GetBody() *OpenDatasetProxyAppendDataRequest {
	return s.Body
}

func (s *AppendAllDataToTaskRequest) SetBody(v *OpenDatasetProxyAppendDataRequest) *AppendAllDataToTaskRequest {
	s.Body = v
	return s
}

func (s *AppendAllDataToTaskRequest) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
