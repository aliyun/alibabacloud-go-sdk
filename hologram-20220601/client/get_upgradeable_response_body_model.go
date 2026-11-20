// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetUpgradeableResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*GetUpgradeableResponseBodyData) *GetUpgradeableResponseBody
	GetData() []*GetUpgradeableResponseBodyData
	SetErrorCode(v string) *GetUpgradeableResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *GetUpgradeableResponseBody
	GetErrorMessage() *string
	SetHttpStatusCode(v string) *GetUpgradeableResponseBody
	GetHttpStatusCode() *string
	SetRequestId(v string) *GetUpgradeableResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetUpgradeableResponseBody
	GetSuccess() *bool
}

type GetUpgradeableResponseBody struct {
	// The response data.
	Data []*GetUpgradeableResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The error code.
	//
	// example:
	//
	// null
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message.
	//
	// example:
	//
	// null
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *string `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// Id of the request
	//
	// example:
	//
	// B8E7613D-E09F-59D3-B35A-FBB7680DCE8F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. This parameter indicates whether an exception occurred during the request and is independent of the business logic.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetUpgradeableResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetUpgradeableResponseBody) GoString() string {
	return s.String()
}

func (s *GetUpgradeableResponseBody) GetData() []*GetUpgradeableResponseBodyData {
	return s.Data
}

func (s *GetUpgradeableResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *GetUpgradeableResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *GetUpgradeableResponseBody) GetHttpStatusCode() *string {
	return s.HttpStatusCode
}

func (s *GetUpgradeableResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetUpgradeableResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetUpgradeableResponseBody) SetData(v []*GetUpgradeableResponseBodyData) *GetUpgradeableResponseBody {
	s.Data = v
	return s
}

func (s *GetUpgradeableResponseBody) SetErrorCode(v string) *GetUpgradeableResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetUpgradeableResponseBody) SetErrorMessage(v string) *GetUpgradeableResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *GetUpgradeableResponseBody) SetHttpStatusCode(v string) *GetUpgradeableResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetUpgradeableResponseBody) SetRequestId(v string) *GetUpgradeableResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetUpgradeableResponseBody) SetSuccess(v bool) *GetUpgradeableResponseBody {
	s.Success = &v
	return s
}

func (s *GetUpgradeableResponseBody) Validate() error {
	if s.Data != nil {
		for _, item := range s.Data {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetUpgradeableResponseBodyData struct {
	// The reason for the upgrade failure.
	//
	// example:
	//
	// null
	FailReason *string `json:"FailReason,omitempty" xml:"FailReason,omitempty"`
	// The key of the upgrade reason.
	//
	// example:
	//
	// null
	ReasonKey *string `json:"ReasonKey,omitempty" xml:"ReasonKey,omitempty"`
	// Indicates whether the instance can be upgraded. Valid values:
	//
	// - true: The instance can be upgraded.
	//
	// - false: The instance cannot be upgraded because it is already running the latest version.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetUpgradeableResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetUpgradeableResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetUpgradeableResponseBodyData) GetFailReason() *string {
	return s.FailReason
}

func (s *GetUpgradeableResponseBodyData) GetReasonKey() *string {
	return s.ReasonKey
}

func (s *GetUpgradeableResponseBodyData) GetSuccess() *bool {
	return s.Success
}

func (s *GetUpgradeableResponseBodyData) SetFailReason(v string) *GetUpgradeableResponseBodyData {
	s.FailReason = &v
	return s
}

func (s *GetUpgradeableResponseBodyData) SetReasonKey(v string) *GetUpgradeableResponseBodyData {
	s.ReasonKey = &v
	return s
}

func (s *GetUpgradeableResponseBodyData) SetSuccess(v bool) *GetUpgradeableResponseBodyData {
	s.Success = &v
	return s
}

func (s *GetUpgradeableResponseBodyData) Validate() error {
	return dara.Validate(s)
}
