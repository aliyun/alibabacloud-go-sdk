// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReplaceCloudGtmInstanceConfigAddressPoolResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ReplaceCloudGtmInstanceConfigAddressPoolResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ReplaceCloudGtmInstanceConfigAddressPoolResponseBody
	GetSuccess() *bool
}

type ReplaceCloudGtmInstanceConfigAddressPoolResponseBody struct {
	// Unique request identification code.
	//
	// example:
	//
	// 536E9CAD-DB30-4647-AC87-AA5CC38C5382
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the operation was successful, with values:
	//
	// - true: Success.
	//
	// - false: Failure.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ReplaceCloudGtmInstanceConfigAddressPoolResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ReplaceCloudGtmInstanceConfigAddressPoolResponseBody) GoString() string {
	return s.String()
}

func (s *ReplaceCloudGtmInstanceConfigAddressPoolResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ReplaceCloudGtmInstanceConfigAddressPoolResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ReplaceCloudGtmInstanceConfigAddressPoolResponseBody) SetRequestId(v string) *ReplaceCloudGtmInstanceConfigAddressPoolResponseBody {
	s.RequestId = &v
	return s
}

func (s *ReplaceCloudGtmInstanceConfigAddressPoolResponseBody) SetSuccess(v bool) *ReplaceCloudGtmInstanceConfigAddressPoolResponseBody {
	s.Success = &v
	return s
}

func (s *ReplaceCloudGtmInstanceConfigAddressPoolResponseBody) Validate() error {
	return dara.Validate(s)
}
