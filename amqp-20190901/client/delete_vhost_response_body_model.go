// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteVhostResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *DeleteVhostResponseBody
	GetCode() *int32
	SetData(v string) *DeleteVhostResponseBody
	GetData() *string
	SetMessage(v string) *DeleteVhostResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteVhostResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteVhostResponseBody
	GetSuccess() *bool
}

type DeleteVhostResponseBody struct {
	Code      *int32  `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteVhostResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteVhostResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteVhostResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *DeleteVhostResponseBody) GetData() *string {
	return s.Data
}

func (s *DeleteVhostResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteVhostResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteVhostResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteVhostResponseBody) SetCode(v int32) *DeleteVhostResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteVhostResponseBody) SetData(v string) *DeleteVhostResponseBody {
	s.Data = &v
	return s
}

func (s *DeleteVhostResponseBody) SetMessage(v string) *DeleteVhostResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteVhostResponseBody) SetRequestId(v string) *DeleteVhostResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteVhostResponseBody) SetSuccess(v bool) *DeleteVhostResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteVhostResponseBody) Validate() error {
	return dara.Validate(s)
}
