// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAutoUpgradeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v bool) *UpdateAutoUpgradeResponseBody
	GetData() *bool
	SetErrorCode(v string) *UpdateAutoUpgradeResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *UpdateAutoUpgradeResponseBody
	GetErrorMessage() *string
	SetHttpStatusCode(v string) *UpdateAutoUpgradeResponseBody
	GetHttpStatusCode() *string
	SetRequestId(v string) *UpdateAutoUpgradeResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateAutoUpgradeResponseBody
	GetSuccess() *bool
}

type UpdateAutoUpgradeResponseBody struct {
	// example:
	//
	// true
	Data *bool `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// null
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// example:
	//
	// null
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *string `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 2A8DEF6E-067E-5DB0-BAE1-2894266E6C6A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateAutoUpgradeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateAutoUpgradeResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateAutoUpgradeResponseBody) GetData() *bool {
	return s.Data
}

func (s *UpdateAutoUpgradeResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *UpdateAutoUpgradeResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *UpdateAutoUpgradeResponseBody) GetHttpStatusCode() *string {
	return s.HttpStatusCode
}

func (s *UpdateAutoUpgradeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateAutoUpgradeResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateAutoUpgradeResponseBody) SetData(v bool) *UpdateAutoUpgradeResponseBody {
	s.Data = &v
	return s
}

func (s *UpdateAutoUpgradeResponseBody) SetErrorCode(v string) *UpdateAutoUpgradeResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *UpdateAutoUpgradeResponseBody) SetErrorMessage(v string) *UpdateAutoUpgradeResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *UpdateAutoUpgradeResponseBody) SetHttpStatusCode(v string) *UpdateAutoUpgradeResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *UpdateAutoUpgradeResponseBody) SetRequestId(v string) *UpdateAutoUpgradeResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateAutoUpgradeResponseBody) SetSuccess(v bool) *UpdateAutoUpgradeResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateAutoUpgradeResponseBody) Validate() error {
	return dara.Validate(s)
}
