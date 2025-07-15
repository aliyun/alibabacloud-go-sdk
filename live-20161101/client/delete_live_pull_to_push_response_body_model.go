// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteLivePullToPushResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *DeleteLivePullToPushResponseBody
	GetDescription() *string
	SetRequestId(v string) *DeleteLivePullToPushResponseBody
	GetRequestId() *string
	SetRetCode(v int32) *DeleteLivePullToPushResponseBody
	GetRetCode() *int32
}

type DeleteLivePullToPushResponseBody struct {
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

func (s DeleteLivePullToPushResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteLivePullToPushResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteLivePullToPushResponseBody) GetDescription() *string {
	return s.Description
}

func (s *DeleteLivePullToPushResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteLivePullToPushResponseBody) GetRetCode() *int32 {
	return s.RetCode
}

func (s *DeleteLivePullToPushResponseBody) SetDescription(v string) *DeleteLivePullToPushResponseBody {
	s.Description = &v
	return s
}

func (s *DeleteLivePullToPushResponseBody) SetRequestId(v string) *DeleteLivePullToPushResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteLivePullToPushResponseBody) SetRetCode(v int32) *DeleteLivePullToPushResponseBody {
	s.RetCode = &v
	return s
}

func (s *DeleteLivePullToPushResponseBody) Validate() error {
	return dara.Validate(s)
}
