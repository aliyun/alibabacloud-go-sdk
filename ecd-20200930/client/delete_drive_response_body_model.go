// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDriveResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DeleteDriveResponseBody
	GetCode() *string
	SetData(v bool) *DeleteDriveResponseBody
	GetData() *bool
	SetMessage(v string) *DeleteDriveResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteDriveResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteDriveResponseBody
	GetSuccess() *bool
}

type DeleteDriveResponseBody struct {
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// true
	Data *bool `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// B7AA****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteDriveResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteDriveResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteDriveResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeleteDriveResponseBody) GetData() *bool {
	return s.Data
}

func (s *DeleteDriveResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteDriveResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteDriveResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteDriveResponseBody) SetCode(v string) *DeleteDriveResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteDriveResponseBody) SetData(v bool) *DeleteDriveResponseBody {
	s.Data = &v
	return s
}

func (s *DeleteDriveResponseBody) SetMessage(v string) *DeleteDriveResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteDriveResponseBody) SetRequestId(v string) *DeleteDriveResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteDriveResponseBody) SetSuccess(v bool) *DeleteDriveResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteDriveResponseBody) Validate() error {
	return dara.Validate(s)
}
