// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckAccountNameAvailableResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CheckAccountNameAvailableResponseBody
	GetRequestId() *string
	SetResourceGroupId(v string) *CheckAccountNameAvailableResponseBody
	GetResourceGroupId() *string
}

type CheckAccountNameAvailableResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 5E4AA101-1EE5-41C0-AE6D-0F066331AC1C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ID of the resource group.
	//
	// example:
	//
	// rg-acfmy****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s CheckAccountNameAvailableResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CheckAccountNameAvailableResponseBody) GoString() string {
	return s.String()
}

func (s *CheckAccountNameAvailableResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CheckAccountNameAvailableResponseBody) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *CheckAccountNameAvailableResponseBody) SetRequestId(v string) *CheckAccountNameAvailableResponseBody {
	s.RequestId = &v
	return s
}

func (s *CheckAccountNameAvailableResponseBody) SetResourceGroupId(v string) *CheckAccountNameAvailableResponseBody {
	s.ResourceGroupId = &v
	return s
}

func (s *CheckAccountNameAvailableResponseBody) Validate() error {
	return dara.Validate(s)
}
