// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpgradeClusterResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *UpgradeClusterResponseBody
	GetErrorCode() *string
	SetHttpCode(v string) *UpgradeClusterResponseBody
	GetHttpCode() *string
	SetMessage(v string) *UpgradeClusterResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpgradeClusterResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpgradeClusterResponseBody
	GetSuccess() *bool
}

type UpgradeClusterResponseBody struct {
	// The error code.
	//
	// example:
	//
	// mse-100-000
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 202
	HttpCode *string `json:"HttpCode,omitempty" xml:"HttpCode,omitempty"`
	// The returned message.
	//
	// - If the request succeeds, a success message is returned.
	//
	// - If the request fails, the cause of the failure is returned.
	//
	// example:
	//
	// The request is processed successfully.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 8D855418-177B-5FF8-A021-75B930AD890A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request succeeded. Valid values:
	//
	// - `true`: The request succeeded.
	//
	// - `false`: The request failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpgradeClusterResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpgradeClusterResponseBody) GoString() string {
	return s.String()
}

func (s *UpgradeClusterResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *UpgradeClusterResponseBody) GetHttpCode() *string {
	return s.HttpCode
}

func (s *UpgradeClusterResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpgradeClusterResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpgradeClusterResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpgradeClusterResponseBody) SetErrorCode(v string) *UpgradeClusterResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *UpgradeClusterResponseBody) SetHttpCode(v string) *UpgradeClusterResponseBody {
	s.HttpCode = &v
	return s
}

func (s *UpgradeClusterResponseBody) SetMessage(v string) *UpgradeClusterResponseBody {
	s.Message = &v
	return s
}

func (s *UpgradeClusterResponseBody) SetRequestId(v string) *UpgradeClusterResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpgradeClusterResponseBody) SetSuccess(v bool) *UpgradeClusterResponseBody {
	s.Success = &v
	return s
}

func (s *UpgradeClusterResponseBody) Validate() error {
	return dara.Validate(s)
}
