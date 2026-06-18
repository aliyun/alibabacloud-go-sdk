// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPreCheckSqlFlashbackTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *PreCheckSqlFlashbackTaskResponseBodyData) *PreCheckSqlFlashbackTaskResponseBody
	GetData() *PreCheckSqlFlashbackTaskResponseBodyData
	SetMessage(v string) *PreCheckSqlFlashbackTaskResponseBody
	GetMessage() *string
	SetRequestId(v string) *PreCheckSqlFlashbackTaskResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *PreCheckSqlFlashbackTaskResponseBody
	GetSuccess() *bool
}

type PreCheckSqlFlashbackTaskResponseBody struct {
	// The result set.
	Data *PreCheckSqlFlashbackTaskResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The response message. This parameter is empty when the request succeeds. If the request fails, an exception message is returned, such as an error code.
	//
	// example:
	//
	// successs
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 14036EBE-***-44DB-ACE9-AC6157D3A6FC
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The request result.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s PreCheckSqlFlashbackTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s PreCheckSqlFlashbackTaskResponseBody) GoString() string {
	return s.String()
}

func (s *PreCheckSqlFlashbackTaskResponseBody) GetData() *PreCheckSqlFlashbackTaskResponseBodyData {
	return s.Data
}

func (s *PreCheckSqlFlashbackTaskResponseBody) GetMessage() *string {
	return s.Message
}

func (s *PreCheckSqlFlashbackTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *PreCheckSqlFlashbackTaskResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *PreCheckSqlFlashbackTaskResponseBody) SetData(v *PreCheckSqlFlashbackTaskResponseBodyData) *PreCheckSqlFlashbackTaskResponseBody {
	s.Data = v
	return s
}

func (s *PreCheckSqlFlashbackTaskResponseBody) SetMessage(v string) *PreCheckSqlFlashbackTaskResponseBody {
	s.Message = &v
	return s
}

func (s *PreCheckSqlFlashbackTaskResponseBody) SetRequestId(v string) *PreCheckSqlFlashbackTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *PreCheckSqlFlashbackTaskResponseBody) SetSuccess(v bool) *PreCheckSqlFlashbackTaskResponseBody {
	s.Success = &v
	return s
}

func (s *PreCheckSqlFlashbackTaskResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type PreCheckSqlFlashbackTaskResponseBodyData struct {
	// The check result.
	CheckResult *PreCheckSqlFlashbackTaskResponseBodyDataCheckResult `json:"CheckResult,omitempty" xml:"CheckResult,omitempty" type:"Struct"`
}

func (s PreCheckSqlFlashbackTaskResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s PreCheckSqlFlashbackTaskResponseBodyData) GoString() string {
	return s.String()
}

func (s *PreCheckSqlFlashbackTaskResponseBodyData) GetCheckResult() *PreCheckSqlFlashbackTaskResponseBodyDataCheckResult {
	return s.CheckResult
}

func (s *PreCheckSqlFlashbackTaskResponseBodyData) SetCheckResult(v *PreCheckSqlFlashbackTaskResponseBodyDataCheckResult) *PreCheckSqlFlashbackTaskResponseBodyData {
	s.CheckResult = v
	return s
}

func (s *PreCheckSqlFlashbackTaskResponseBodyData) Validate() error {
	if s.CheckResult != nil {
		if err := s.CheckResult.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type PreCheckSqlFlashbackTaskResponseBodyDataCheckResult struct {
	// Indicates whether valid binary log files exist. This value is used to determine whether operations such as flashback and synchronization can be performed.
	BinlogExists *bool `json:"BinlogExists,omitempty" xml:"BinlogExists,omitempty"`
	// Indicates whether the recording of original SQL query events is enabled. Valid values:
	//
	// - **true**: Enabled.
	//
	// - **false**: Disabled.
	BinlogRowQueryEventEnabled *bool `json:"BinlogRowQueryEventEnabled,omitempty" xml:"BinlogRowQueryEventEnabled,omitempty"`
	// Indicates whether the upgrade can succeed.
	CanUpgradeSupportBinlogRowQueryEvents *bool `json:"CanUpgradeSupportBinlogRowQueryEvents,omitempty" xml:"CanUpgradeSupportBinlogRowQueryEvents,omitempty"`
	// Indicates whether the table exists. Valid values: true and false.
	HasTable *bool `json:"HasTable,omitempty" xml:"HasTable,omitempty"`
	// Indicates whether log recording is supported. Valid values:
	//
	// - **true**: Supported.
	//
	// - **false**: Not supported.
	SupportBinlogRowQueryEvents *bool `json:"SupportBinlogRowQueryEvents,omitempty" xml:"SupportBinlogRowQueryEvents,omitempty"`
}

func (s PreCheckSqlFlashbackTaskResponseBodyDataCheckResult) String() string {
	return dara.Prettify(s)
}

func (s PreCheckSqlFlashbackTaskResponseBodyDataCheckResult) GoString() string {
	return s.String()
}

func (s *PreCheckSqlFlashbackTaskResponseBodyDataCheckResult) GetBinlogExists() *bool {
	return s.BinlogExists
}

func (s *PreCheckSqlFlashbackTaskResponseBodyDataCheckResult) GetBinlogRowQueryEventEnabled() *bool {
	return s.BinlogRowQueryEventEnabled
}

func (s *PreCheckSqlFlashbackTaskResponseBodyDataCheckResult) GetCanUpgradeSupportBinlogRowQueryEvents() *bool {
	return s.CanUpgradeSupportBinlogRowQueryEvents
}

func (s *PreCheckSqlFlashbackTaskResponseBodyDataCheckResult) GetHasTable() *bool {
	return s.HasTable
}

func (s *PreCheckSqlFlashbackTaskResponseBodyDataCheckResult) GetSupportBinlogRowQueryEvents() *bool {
	return s.SupportBinlogRowQueryEvents
}

func (s *PreCheckSqlFlashbackTaskResponseBodyDataCheckResult) SetBinlogExists(v bool) *PreCheckSqlFlashbackTaskResponseBodyDataCheckResult {
	s.BinlogExists = &v
	return s
}

func (s *PreCheckSqlFlashbackTaskResponseBodyDataCheckResult) SetBinlogRowQueryEventEnabled(v bool) *PreCheckSqlFlashbackTaskResponseBodyDataCheckResult {
	s.BinlogRowQueryEventEnabled = &v
	return s
}

func (s *PreCheckSqlFlashbackTaskResponseBodyDataCheckResult) SetCanUpgradeSupportBinlogRowQueryEvents(v bool) *PreCheckSqlFlashbackTaskResponseBodyDataCheckResult {
	s.CanUpgradeSupportBinlogRowQueryEvents = &v
	return s
}

func (s *PreCheckSqlFlashbackTaskResponseBodyDataCheckResult) SetHasTable(v bool) *PreCheckSqlFlashbackTaskResponseBodyDataCheckResult {
	s.HasTable = &v
	return s
}

func (s *PreCheckSqlFlashbackTaskResponseBodyDataCheckResult) SetSupportBinlogRowQueryEvents(v bool) *PreCheckSqlFlashbackTaskResponseBodyDataCheckResult {
	s.SupportBinlogRowQueryEvents = &v
	return s
}

func (s *PreCheckSqlFlashbackTaskResponseBodyDataCheckResult) Validate() error {
	return dara.Validate(s)
}
