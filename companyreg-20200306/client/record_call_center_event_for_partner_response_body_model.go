// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRecordCallCenterEventForPartnerResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v bool) *RecordCallCenterEventForPartnerResponseBody
	GetData() *bool
	SetErrorCode(v string) *RecordCallCenterEventForPartnerResponseBody
	GetErrorCode() *string
	SetErrorMsg(v string) *RecordCallCenterEventForPartnerResponseBody
	GetErrorMsg() *string
	SetRequestId(v string) *RecordCallCenterEventForPartnerResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *RecordCallCenterEventForPartnerResponseBody
	GetSuccess() *bool
}

type RecordCallCenterEventForPartnerResponseBody struct {
	// example:
	//
	// True
	Data *bool `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// NoPermission
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMsg  *string `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 1A13ABB5-7649-5031-B55C-D2E38F9F189D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s RecordCallCenterEventForPartnerResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RecordCallCenterEventForPartnerResponseBody) GoString() string {
	return s.String()
}

func (s *RecordCallCenterEventForPartnerResponseBody) GetData() *bool {
	return s.Data
}

func (s *RecordCallCenterEventForPartnerResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *RecordCallCenterEventForPartnerResponseBody) GetErrorMsg() *string {
	return s.ErrorMsg
}

func (s *RecordCallCenterEventForPartnerResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RecordCallCenterEventForPartnerResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *RecordCallCenterEventForPartnerResponseBody) SetData(v bool) *RecordCallCenterEventForPartnerResponseBody {
	s.Data = &v
	return s
}

func (s *RecordCallCenterEventForPartnerResponseBody) SetErrorCode(v string) *RecordCallCenterEventForPartnerResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *RecordCallCenterEventForPartnerResponseBody) SetErrorMsg(v string) *RecordCallCenterEventForPartnerResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *RecordCallCenterEventForPartnerResponseBody) SetRequestId(v string) *RecordCallCenterEventForPartnerResponseBody {
	s.RequestId = &v
	return s
}

func (s *RecordCallCenterEventForPartnerResponseBody) SetSuccess(v bool) *RecordCallCenterEventForPartnerResponseBody {
	s.Success = &v
	return s
}

func (s *RecordCallCenterEventForPartnerResponseBody) Validate() error {
	return dara.Validate(s)
}
