// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyOpenLogShipperResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyOpenLogShipperResponseBody
	GetRequestId() *string
}

type ModifyOpenLogShipperResponseBody struct {
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// 25EC270F-5783-4416-AD7C-1EDF063A039C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyOpenLogShipperResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyOpenLogShipperResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyOpenLogShipperResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyOpenLogShipperResponseBody) SetRequestId(v string) *ModifyOpenLogShipperResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyOpenLogShipperResponseBody) Validate() error {
	return dara.Validate(s)
}
