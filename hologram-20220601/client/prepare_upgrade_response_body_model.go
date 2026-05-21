// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPrepareUpgradeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *PrepareUpgradeResponseBodyData) *PrepareUpgradeResponseBody
	GetData() *PrepareUpgradeResponseBodyData
	SetErrorCode(v string) *PrepareUpgradeResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *PrepareUpgradeResponseBody
	GetErrorMessage() *string
	SetHttpStatusCode(v string) *PrepareUpgradeResponseBody
	GetHttpStatusCode() *string
	SetRequestId(v string) *PrepareUpgradeResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *PrepareUpgradeResponseBody
	GetSuccess() *bool
}

type PrepareUpgradeResponseBody struct {
	Data *PrepareUpgradeResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
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
	// 77B97AFB-7C9D-50FF-A72D-F13FD73E49D8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s PrepareUpgradeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s PrepareUpgradeResponseBody) GoString() string {
	return s.String()
}

func (s *PrepareUpgradeResponseBody) GetData() *PrepareUpgradeResponseBodyData {
	return s.Data
}

func (s *PrepareUpgradeResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *PrepareUpgradeResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *PrepareUpgradeResponseBody) GetHttpStatusCode() *string {
	return s.HttpStatusCode
}

func (s *PrepareUpgradeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *PrepareUpgradeResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *PrepareUpgradeResponseBody) SetData(v *PrepareUpgradeResponseBodyData) *PrepareUpgradeResponseBody {
	s.Data = v
	return s
}

func (s *PrepareUpgradeResponseBody) SetErrorCode(v string) *PrepareUpgradeResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *PrepareUpgradeResponseBody) SetErrorMessage(v string) *PrepareUpgradeResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *PrepareUpgradeResponseBody) SetHttpStatusCode(v string) *PrepareUpgradeResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *PrepareUpgradeResponseBody) SetRequestId(v string) *PrepareUpgradeResponseBody {
	s.RequestId = &v
	return s
}

func (s *PrepareUpgradeResponseBody) SetSuccess(v bool) *PrepareUpgradeResponseBody {
	s.Success = &v
	return s
}

func (s *PrepareUpgradeResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type PrepareUpgradeResponseBodyData struct {
	// example:
	//
	// null
	FailReason *string `json:"FailReason,omitempty" xml:"FailReason,omitempty"`
	// example:
	//
	// null
	ReasonKey *string `json:"ReasonKey,omitempty" xml:"ReasonKey,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s PrepareUpgradeResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s PrepareUpgradeResponseBodyData) GoString() string {
	return s.String()
}

func (s *PrepareUpgradeResponseBodyData) GetFailReason() *string {
	return s.FailReason
}

func (s *PrepareUpgradeResponseBodyData) GetReasonKey() *string {
	return s.ReasonKey
}

func (s *PrepareUpgradeResponseBodyData) GetSuccess() *bool {
	return s.Success
}

func (s *PrepareUpgradeResponseBodyData) SetFailReason(v string) *PrepareUpgradeResponseBodyData {
	s.FailReason = &v
	return s
}

func (s *PrepareUpgradeResponseBodyData) SetReasonKey(v string) *PrepareUpgradeResponseBodyData {
	s.ReasonKey = &v
	return s
}

func (s *PrepareUpgradeResponseBodyData) SetSuccess(v bool) *PrepareUpgradeResponseBodyData {
	s.Success = &v
	return s
}

func (s *PrepareUpgradeResponseBodyData) Validate() error {
	return dara.Validate(s)
}
