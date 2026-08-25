// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetVolumeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetVolumeResponseBody
	GetCode() *string
	SetMessage(v string) *GetVolumeResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetVolumeResponseBody
	GetRequestId() *string
	SetVolume(v *E2BVolume) *GetVolumeResponseBody
	GetVolume() *E2BVolume
}

type GetVolumeResponseBody struct {
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
	// 7ADFF8D8-D4BA-5F79-AD49-DDABFEA59B6C
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The storage configuration.
	Volume *E2BVolume `json:"volume,omitempty" xml:"volume,omitempty"`
}

func (s GetVolumeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetVolumeResponseBody) GoString() string {
	return s.String()
}

func (s *GetVolumeResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetVolumeResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetVolumeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetVolumeResponseBody) GetVolume() *E2BVolume {
	return s.Volume
}

func (s *GetVolumeResponseBody) SetCode(v string) *GetVolumeResponseBody {
	s.Code = &v
	return s
}

func (s *GetVolumeResponseBody) SetMessage(v string) *GetVolumeResponseBody {
	s.Message = &v
	return s
}

func (s *GetVolumeResponseBody) SetRequestId(v string) *GetVolumeResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetVolumeResponseBody) SetVolume(v *E2BVolume) *GetVolumeResponseBody {
	s.Volume = v
	return s
}

func (s *GetVolumeResponseBody) Validate() error {
	if s.Volume != nil {
		if err := s.Volume.Validate(); err != nil {
			return err
		}
	}
	return nil
}
