// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetDataShareInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrMessage(v string) *SetDataShareInstanceResponseBody
	GetErrMessage() *string
	SetRequestId(v string) *SetDataShareInstanceResponseBody
	GetRequestId() *string
	SetStatus(v string) *SetDataShareInstanceResponseBody
	GetStatus() *string
}

type SetDataShareInstanceResponseBody struct {
	// The error message returned if the operation fails.
	//
	// example:
	//
	// error message
	ErrMessage *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// B4CAF581-2AC7-41AD-8940-D5**********
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The state of the operation. Valid values:
	//
	// 	- **success**: The operation is successful.
	//
	// 	- **failed**: The operation fails.
	//
	// example:
	//
	// success
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s SetDataShareInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SetDataShareInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *SetDataShareInstanceResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *SetDataShareInstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SetDataShareInstanceResponseBody) GetStatus() *string {
	return s.Status
}

func (s *SetDataShareInstanceResponseBody) SetErrMessage(v string) *SetDataShareInstanceResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *SetDataShareInstanceResponseBody) SetRequestId(v string) *SetDataShareInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *SetDataShareInstanceResponseBody) SetStatus(v string) *SetDataShareInstanceResponseBody {
	s.Status = &v
	return s
}

func (s *SetDataShareInstanceResponseBody) Validate() error {
	return dara.Validate(s)
}
