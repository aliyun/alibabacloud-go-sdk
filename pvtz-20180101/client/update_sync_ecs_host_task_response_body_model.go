// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateSyncEcsHostTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateSyncEcsHostTaskResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateSyncEcsHostTaskResponseBody
	GetSuccess() *bool
}

type UpdateSyncEcsHostTaskResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// test-FC9A-4595-8D96-089D73D7A63D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateSyncEcsHostTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateSyncEcsHostTaskResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateSyncEcsHostTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateSyncEcsHostTaskResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateSyncEcsHostTaskResponseBody) SetRequestId(v string) *UpdateSyncEcsHostTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateSyncEcsHostTaskResponseBody) SetSuccess(v bool) *UpdateSyncEcsHostTaskResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateSyncEcsHostTaskResponseBody) Validate() error {
	return dara.Validate(s)
}
