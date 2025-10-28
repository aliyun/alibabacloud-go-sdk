// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteEcuResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *DeleteEcuResponseBody
	GetCode() *int32
	SetData(v string) *DeleteEcuResponseBody
	GetData() *string
	SetMessage(v string) *DeleteEcuResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteEcuResponseBody
	GetRequestId() *string
}

type DeleteEcuResponseBody struct {
	// The HTTP status code that is returned.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned data that indicates whether the ECU is deleted. A value of `OK` indicates that the ECU is deleted.
	//
	// example:
	//
	// OK
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// The additional information that is returned.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// fb08fd29-b197-40ab-****-****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteEcuResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteEcuResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteEcuResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *DeleteEcuResponseBody) GetData() *string {
	return s.Data
}

func (s *DeleteEcuResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteEcuResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteEcuResponseBody) SetCode(v int32) *DeleteEcuResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteEcuResponseBody) SetData(v string) *DeleteEcuResponseBody {
	s.Data = &v
	return s
}

func (s *DeleteEcuResponseBody) SetMessage(v string) *DeleteEcuResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteEcuResponseBody) SetRequestId(v string) *DeleteEcuResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteEcuResponseBody) Validate() error {
	return dara.Validate(s)
}
