// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpgradeInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *UpgradeInstanceResponseBodyData) *UpgradeInstanceResponseBody
	GetData() *UpgradeInstanceResponseBodyData
	SetErrorCode(v string) *UpgradeInstanceResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *UpgradeInstanceResponseBody
	GetErrorMessage() *string
	SetHttpStatusCode(v string) *UpgradeInstanceResponseBody
	GetHttpStatusCode() *string
	SetRequestId(v string) *UpgradeInstanceResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpgradeInstanceResponseBody
	GetSuccess() *bool
}

type UpgradeInstanceResponseBody struct {
	// The returned data.
	Data *UpgradeInstanceResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
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
	// EEB0A71E-7AC7-572F-990F-EE51D3FD35B9
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether an exception occurred during the request.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpgradeInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpgradeInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *UpgradeInstanceResponseBody) GetData() *UpgradeInstanceResponseBodyData {
	return s.Data
}

func (s *UpgradeInstanceResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *UpgradeInstanceResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *UpgradeInstanceResponseBody) GetHttpStatusCode() *string {
	return s.HttpStatusCode
}

func (s *UpgradeInstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpgradeInstanceResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpgradeInstanceResponseBody) SetData(v *UpgradeInstanceResponseBodyData) *UpgradeInstanceResponseBody {
	s.Data = v
	return s
}

func (s *UpgradeInstanceResponseBody) SetErrorCode(v string) *UpgradeInstanceResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *UpgradeInstanceResponseBody) SetErrorMessage(v string) *UpgradeInstanceResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *UpgradeInstanceResponseBody) SetHttpStatusCode(v string) *UpgradeInstanceResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *UpgradeInstanceResponseBody) SetRequestId(v string) *UpgradeInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpgradeInstanceResponseBody) SetSuccess(v bool) *UpgradeInstanceResponseBody {
	s.Success = &v
	return s
}

func (s *UpgradeInstanceResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpgradeInstanceResponseBodyData struct {
	// The reason for the upgrade failure.
	//
	// example:
	//
	// null
	FailReason *string `json:"FailReason,omitempty" xml:"FailReason,omitempty"`
	// The reason key.
	//
	// example:
	//
	// null
	ReasonKey *string `json:"ReasonKey,omitempty" xml:"ReasonKey,omitempty"`
	// Indicates whether the operation is successful. Valid values:
	//
	// - true: Successful.
	//
	// - false: Failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpgradeInstanceResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s UpgradeInstanceResponseBodyData) GoString() string {
	return s.String()
}

func (s *UpgradeInstanceResponseBodyData) GetFailReason() *string {
	return s.FailReason
}

func (s *UpgradeInstanceResponseBodyData) GetReasonKey() *string {
	return s.ReasonKey
}

func (s *UpgradeInstanceResponseBodyData) GetSuccess() *bool {
	return s.Success
}

func (s *UpgradeInstanceResponseBodyData) SetFailReason(v string) *UpgradeInstanceResponseBodyData {
	s.FailReason = &v
	return s
}

func (s *UpgradeInstanceResponseBodyData) SetReasonKey(v string) *UpgradeInstanceResponseBodyData {
	s.ReasonKey = &v
	return s
}

func (s *UpgradeInstanceResponseBodyData) SetSuccess(v bool) *UpgradeInstanceResponseBodyData {
	s.Success = &v
	return s
}

func (s *UpgradeInstanceResponseBodyData) Validate() error {
	return dara.Validate(s)
}
