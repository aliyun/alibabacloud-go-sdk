// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateLivePullToPushResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *UpdateLivePullToPushResponseBody
	GetDescription() *string
	SetRequestId(v string) *UpdateLivePullToPushResponseBody
	GetRequestId() *string
	SetRetCode(v int32) *UpdateLivePullToPushResponseBody
	GetRetCode() *int32
}

type UpdateLivePullToPushResponseBody struct {
	// The error description.
	//
	// example:
	//
	// OK
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 16A96B9A-F203-4EC5-8E43-CB92E68*****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The code that is returned for the request.
	//
	// >
	//
	// 	- 0 is returned if the request is normal.
	//
	// 	- For information about codes that are returned when exceptions occur, see the following Error codes table.
	//
	// example:
	//
	// 0
	RetCode *int32 `json:"RetCode,omitempty" xml:"RetCode,omitempty"`
}

func (s UpdateLivePullToPushResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateLivePullToPushResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateLivePullToPushResponseBody) GetDescription() *string {
	return s.Description
}

func (s *UpdateLivePullToPushResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateLivePullToPushResponseBody) GetRetCode() *int32 {
	return s.RetCode
}

func (s *UpdateLivePullToPushResponseBody) SetDescription(v string) *UpdateLivePullToPushResponseBody {
	s.Description = &v
	return s
}

func (s *UpdateLivePullToPushResponseBody) SetRequestId(v string) *UpdateLivePullToPushResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateLivePullToPushResponseBody) SetRetCode(v int32) *UpdateLivePullToPushResponseBody {
	s.RetCode = &v
	return s
}

func (s *UpdateLivePullToPushResponseBody) Validate() error {
	return dara.Validate(s)
}
