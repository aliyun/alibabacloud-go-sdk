// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateVolumeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v *CreateVolumeInput) *CreateVolumeRequest
	GetBody() *CreateVolumeInput
}

type CreateVolumeRequest struct {
	// The form parameters.
	Body *CreateVolumeInput `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateVolumeRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateVolumeRequest) GoString() string {
	return s.String()
}

func (s *CreateVolumeRequest) GetBody() *CreateVolumeInput {
	return s.Body
}

func (s *CreateVolumeRequest) SetBody(v *CreateVolumeInput) *CreateVolumeRequest {
	s.Body = v
	return s
}

func (s *CreateVolumeRequest) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
