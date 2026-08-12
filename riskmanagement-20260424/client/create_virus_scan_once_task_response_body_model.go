// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateVirusScanOnceTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateVirusScanOnceTaskResponseBody
	GetCode() *string
	SetData(v *CreateVirusScanOnceTaskResponseBodyData) *CreateVirusScanOnceTaskResponseBody
	GetData() *CreateVirusScanOnceTaskResponseBodyData
	SetMessage(v string) *CreateVirusScanOnceTaskResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateVirusScanOnceTaskResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateVirusScanOnceTaskResponseBody
	GetSuccess() *bool
}

type CreateVirusScanOnceTaskResponseBody struct {
	// The error code returned if the call fails. For more information, refer to error codes.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned data.
	Data *CreateVirusScanOnceTaskResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The message information.
	//
	// example:
	//
	// successful‌
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 739705BB-B0EF-554B-B3A8-383F4F93E067
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the call is successful. Valid values:
	//
	// - **true**: The call is successful.
	//
	// - **false**: The call fails.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateVirusScanOnceTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateVirusScanOnceTaskResponseBody) GoString() string {
	return s.String()
}

func (s *CreateVirusScanOnceTaskResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateVirusScanOnceTaskResponseBody) GetData() *CreateVirusScanOnceTaskResponseBodyData {
	return s.Data
}

func (s *CreateVirusScanOnceTaskResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateVirusScanOnceTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateVirusScanOnceTaskResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateVirusScanOnceTaskResponseBody) SetCode(v string) *CreateVirusScanOnceTaskResponseBody {
	s.Code = &v
	return s
}

func (s *CreateVirusScanOnceTaskResponseBody) SetData(v *CreateVirusScanOnceTaskResponseBodyData) *CreateVirusScanOnceTaskResponseBody {
	s.Data = v
	return s
}

func (s *CreateVirusScanOnceTaskResponseBody) SetMessage(v string) *CreateVirusScanOnceTaskResponseBody {
	s.Message = &v
	return s
}

func (s *CreateVirusScanOnceTaskResponseBody) SetRequestId(v string) *CreateVirusScanOnceTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateVirusScanOnceTaskResponseBody) SetSuccess(v bool) *CreateVirusScanOnceTaskResponseBody {
	s.Success = &v
	return s
}

func (s *CreateVirusScanOnceTaskResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateVirusScanOnceTaskResponseBodyData struct {
	// The asset selection business type. Valid values:
	//
	// - **VIRUS_SCAN_CYCLE_CONFIG**: virus scan configuration
	//
	// - **VIRUS_SCAN_ONCE_TASK**: virus scan one-time task
	//
	// example:
	//
	// VIRUS_SCAN_CYCLE_CONFIG
	BusinessType *string `json:"BusinessType,omitempty" xml:"BusinessType,omitempty"`
	// The operating system of the target asset. Valid values:
	//
	// - **windows**: Windows operating system
	//
	// - **linux**: Linux operating system
	//
	// example:
	//
	// windows
	Platform *string `json:"Platform,omitempty" xml:"Platform,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 1D345A09-5ABD-593C-9C26-5C2B28632CD6
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The unique identifier of this asset selection, which can be used to query or modify the assets corresponding to this selection.
	//
	// example:
	//
	// 87af4d19-38fc-408d-9549-2bf7b6c2a4b9
	SelectionKey *int32 `json:"SelectionKey,omitempty" xml:"SelectionKey,omitempty"`
	// The target asset type. Valid values:
	//
	// - **all_instance**: all servers
	//
	// - **instance**: select by server
	//
	// - **group**: select by group
	//
	// - **vpc**: select by VPC
	//
	// example:
	//
	// all_instance
	TargetType *string `json:"TargetType,omitempty" xml:"TargetType,omitempty"`
	// The server ID.
	//
	// example:
	//
	// 9ef1a02e1de695cb7f9fea2c6c145853eklEsP2JP0Z
	Uuid *string `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
}

func (s CreateVirusScanOnceTaskResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CreateVirusScanOnceTaskResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateVirusScanOnceTaskResponseBodyData) GetBusinessType() *string {
	return s.BusinessType
}

func (s *CreateVirusScanOnceTaskResponseBodyData) GetPlatform() *string {
	return s.Platform
}

func (s *CreateVirusScanOnceTaskResponseBodyData) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateVirusScanOnceTaskResponseBodyData) GetSelectionKey() *int32 {
	return s.SelectionKey
}

func (s *CreateVirusScanOnceTaskResponseBodyData) GetTargetType() *string {
	return s.TargetType
}

func (s *CreateVirusScanOnceTaskResponseBodyData) GetUuid() *string {
	return s.Uuid
}

func (s *CreateVirusScanOnceTaskResponseBodyData) SetBusinessType(v string) *CreateVirusScanOnceTaskResponseBodyData {
	s.BusinessType = &v
	return s
}

func (s *CreateVirusScanOnceTaskResponseBodyData) SetPlatform(v string) *CreateVirusScanOnceTaskResponseBodyData {
	s.Platform = &v
	return s
}

func (s *CreateVirusScanOnceTaskResponseBodyData) SetRequestId(v string) *CreateVirusScanOnceTaskResponseBodyData {
	s.RequestId = &v
	return s
}

func (s *CreateVirusScanOnceTaskResponseBodyData) SetSelectionKey(v int32) *CreateVirusScanOnceTaskResponseBodyData {
	s.SelectionKey = &v
	return s
}

func (s *CreateVirusScanOnceTaskResponseBodyData) SetTargetType(v string) *CreateVirusScanOnceTaskResponseBodyData {
	s.TargetType = &v
	return s
}

func (s *CreateVirusScanOnceTaskResponseBodyData) SetUuid(v string) *CreateVirusScanOnceTaskResponseBodyData {
	s.Uuid = &v
	return s
}

func (s *CreateVirusScanOnceTaskResponseBodyData) Validate() error {
	return dara.Validate(s)
}
