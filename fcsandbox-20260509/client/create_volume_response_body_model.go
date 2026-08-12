// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateVolumeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateVolumeResponseBody
	GetCode() *string
	SetMessage(v string) *CreateVolumeResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateVolumeResponseBody
	GetRequestId() *string
	SetVolume(v *E2BVolume) *CreateVolumeResponseBody
	GetVolume() *E2BVolume
}

type CreateVolumeResponseBody struct {
	// The error code.
	//
	// example:
	//
	// 200
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The response message.
	//
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// B5AD8B54-4358-5F5B-ACAA-52F2016459C6
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The storage configuration.
	Volume *E2BVolume `json:"volume,omitempty" xml:"volume,omitempty"`
}

func (s CreateVolumeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateVolumeResponseBody) GoString() string {
	return s.String()
}

func (s *CreateVolumeResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateVolumeResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateVolumeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateVolumeResponseBody) GetVolume() *E2BVolume {
	return s.Volume
}

func (s *CreateVolumeResponseBody) SetCode(v string) *CreateVolumeResponseBody {
	s.Code = &v
	return s
}

func (s *CreateVolumeResponseBody) SetMessage(v string) *CreateVolumeResponseBody {
	s.Message = &v
	return s
}

func (s *CreateVolumeResponseBody) SetRequestId(v string) *CreateVolumeResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateVolumeResponseBody) SetVolume(v *E2BVolume) *CreateVolumeResponseBody {
	s.Volume = v
	return s
}

func (s *CreateVolumeResponseBody) Validate() error {
	if s.Volume != nil {
		if err := s.Volume.Validate(); err != nil {
			return err
		}
	}
	return nil
}
