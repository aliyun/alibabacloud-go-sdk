// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteVolumeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DeleteVolumeResponseBody
	GetCode() *string
	SetMessage(v string) *DeleteVolumeResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteVolumeResponseBody
	GetRequestId() *string
}

type DeleteVolumeResponseBody struct {
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
	// B5AD8B54-4358-5F5B-ACAA-52F2016459C6
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s DeleteVolumeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteVolumeResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteVolumeResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeleteVolumeResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteVolumeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteVolumeResponseBody) SetCode(v string) *DeleteVolumeResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteVolumeResponseBody) SetMessage(v string) *DeleteVolumeResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteVolumeResponseBody) SetRequestId(v string) *DeleteVolumeResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteVolumeResponseBody) Validate() error {
	return dara.Validate(s)
}
