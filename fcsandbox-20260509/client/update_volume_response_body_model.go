// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateVolumeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *UpdateVolumeResponseBody
	GetCode() *string
	SetMessage(v string) *UpdateVolumeResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateVolumeResponseBody
	GetRequestId() *string
	SetVolume(v *E2BVolume) *UpdateVolumeResponseBody
	GetVolume() *E2BVolume
}

type UpdateVolumeResponseBody struct {
	// example:
	//
	// 200
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// example:
	//
	// 7ADFF8D8-D4BA-5F79-AD49-DDABFEA59B6C
	RequestId *string    `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Volume    *E2BVolume `json:"volume,omitempty" xml:"volume,omitempty"`
}

func (s UpdateVolumeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateVolumeResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateVolumeResponseBody) GetCode() *string {
	return s.Code
}

func (s *UpdateVolumeResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateVolumeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateVolumeResponseBody) GetVolume() *E2BVolume {
	return s.Volume
}

func (s *UpdateVolumeResponseBody) SetCode(v string) *UpdateVolumeResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateVolumeResponseBody) SetMessage(v string) *UpdateVolumeResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateVolumeResponseBody) SetRequestId(v string) *UpdateVolumeResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateVolumeResponseBody) SetVolume(v *E2BVolume) *UpdateVolumeResponseBody {
	s.Volume = v
	return s
}

func (s *UpdateVolumeResponseBody) Validate() error {
	if s.Volume != nil {
		if err := s.Volume.Validate(); err != nil {
			return err
		}
	}
	return nil
}
