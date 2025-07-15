// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpgradeInstanceVersionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *UpgradeInstanceVersionResponseBody
	GetCode() *int32
	SetMessage(v string) *UpgradeInstanceVersionResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpgradeInstanceVersionResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpgradeInstanceVersionResponseBody
	GetSuccess() *bool
}

type UpgradeInstanceVersionResponseBody struct {
	// The HTTP status code that is returned. The HTTP status code 200 indicates that the request is successful.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The error message returned.
	//
	// example:
	//
	// operation success.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// ABA4A7FD-E10F-45C7-9774-A5236015***
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpgradeInstanceVersionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpgradeInstanceVersionResponseBody) GoString() string {
	return s.String()
}

func (s *UpgradeInstanceVersionResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *UpgradeInstanceVersionResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpgradeInstanceVersionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpgradeInstanceVersionResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpgradeInstanceVersionResponseBody) SetCode(v int32) *UpgradeInstanceVersionResponseBody {
	s.Code = &v
	return s
}

func (s *UpgradeInstanceVersionResponseBody) SetMessage(v string) *UpgradeInstanceVersionResponseBody {
	s.Message = &v
	return s
}

func (s *UpgradeInstanceVersionResponseBody) SetRequestId(v string) *UpgradeInstanceVersionResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpgradeInstanceVersionResponseBody) SetSuccess(v bool) *UpgradeInstanceVersionResponseBody {
	s.Success = &v
	return s
}

func (s *UpgradeInstanceVersionResponseBody) Validate() error {
	return dara.Validate(s)
}
