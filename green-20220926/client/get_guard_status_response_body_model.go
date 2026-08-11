// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetGuardStatusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetLogScanStatus(v []*GetGuardStatusResponseBodyLogScanStatus) *GetGuardStatusResponseBody
	GetLogScanStatus() []*GetGuardStatusResponseBodyLogScanStatus
	SetProtectionStatus(v []*GetGuardStatusResponseBodyProtectionStatus) *GetGuardStatusResponseBody
	GetProtectionStatus() []*GetGuardStatusResponseBodyProtectionStatus
	SetRealTimeStatus(v []*GetGuardStatusResponseBodyRealTimeStatus) *GetGuardStatusResponseBody
	GetRealTimeStatus() []*GetGuardStatusResponseBodyRealTimeStatus
	SetRequestId(v string) *GetGuardStatusResponseBody
	GetRequestId() *string
}

type GetGuardStatusResponseBody struct {
	// The list of log scan statistics.
	LogScanStatus []*GetGuardStatusResponseBodyLogScanStatus `json:"LogScanStatus,omitempty" xml:"LogScanStatus,omitempty" type:"Repeated"`
	// The list of protection status statistics.
	ProtectionStatus []*GetGuardStatusResponseBodyProtectionStatus `json:"ProtectionStatus,omitempty" xml:"ProtectionStatus,omitempty" type:"Repeated"`
	// The list of real-time protection statistics.
	RealTimeStatus []*GetGuardStatusResponseBodyRealTimeStatus `json:"RealTimeStatus,omitempty" xml:"RealTimeStatus,omitempty" type:"Repeated"`
	// The ID assigned by the backend to uniquely identify a request. This ID can be used to troubleshoot issues.
	//
	// example:
	//
	// AAAAAA-BBBB-CCCCC-DDDD-EEEEEEEE****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetGuardStatusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetGuardStatusResponseBody) GoString() string {
	return s.String()
}

func (s *GetGuardStatusResponseBody) GetLogScanStatus() []*GetGuardStatusResponseBodyLogScanStatus {
	return s.LogScanStatus
}

func (s *GetGuardStatusResponseBody) GetProtectionStatus() []*GetGuardStatusResponseBodyProtectionStatus {
	return s.ProtectionStatus
}

func (s *GetGuardStatusResponseBody) GetRealTimeStatus() []*GetGuardStatusResponseBodyRealTimeStatus {
	return s.RealTimeStatus
}

func (s *GetGuardStatusResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetGuardStatusResponseBody) SetLogScanStatus(v []*GetGuardStatusResponseBodyLogScanStatus) *GetGuardStatusResponseBody {
	s.LogScanStatus = v
	return s
}

func (s *GetGuardStatusResponseBody) SetProtectionStatus(v []*GetGuardStatusResponseBodyProtectionStatus) *GetGuardStatusResponseBody {
	s.ProtectionStatus = v
	return s
}

func (s *GetGuardStatusResponseBody) SetRealTimeStatus(v []*GetGuardStatusResponseBodyRealTimeStatus) *GetGuardStatusResponseBody {
	s.RealTimeStatus = v
	return s
}

func (s *GetGuardStatusResponseBody) SetRequestId(v string) *GetGuardStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetGuardStatusResponseBody) Validate() error {
	if s.LogScanStatus != nil {
		for _, item := range s.LogScanStatus {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.ProtectionStatus != nil {
		for _, item := range s.ProtectionStatus {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.RealTimeStatus != nil {
		for _, item := range s.RealTimeStatus {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetGuardStatusResponseBodyLogScanStatus struct {
	// The status. Valid values:
	//
	// - enabled: Running.
	//
	// - disabled: Not accessed.
	//
	// example:
	//
	// enabled
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The type.
	//
	// example:
	//
	// ai_app_scan_bailian
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GetGuardStatusResponseBodyLogScanStatus) String() string {
	return dara.Prettify(s)
}

func (s GetGuardStatusResponseBodyLogScanStatus) GoString() string {
	return s.String()
}

func (s *GetGuardStatusResponseBodyLogScanStatus) GetStatus() *string {
	return s.Status
}

func (s *GetGuardStatusResponseBodyLogScanStatus) GetType() *string {
	return s.Type
}

func (s *GetGuardStatusResponseBodyLogScanStatus) SetStatus(v string) *GetGuardStatusResponseBodyLogScanStatus {
	s.Status = &v
	return s
}

func (s *GetGuardStatusResponseBodyLogScanStatus) SetType(v string) *GetGuardStatusResponseBodyLogScanStatus {
	s.Type = &v
	return s
}

func (s *GetGuardStatusResponseBodyLogScanStatus) Validate() error {
	return dara.Validate(s)
}

type GetGuardStatusResponseBodyProtectionStatus struct {
	// The status. Valid values:
	//
	// - enabled: Running.
	//
	// - disabled: Not accessed.
	//
	// example:
	//
	// enabled
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The type.
	//
	// example:
	//
	// api
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GetGuardStatusResponseBodyProtectionStatus) String() string {
	return dara.Prettify(s)
}

func (s GetGuardStatusResponseBodyProtectionStatus) GoString() string {
	return s.String()
}

func (s *GetGuardStatusResponseBodyProtectionStatus) GetStatus() *string {
	return s.Status
}

func (s *GetGuardStatusResponseBodyProtectionStatus) GetType() *string {
	return s.Type
}

func (s *GetGuardStatusResponseBodyProtectionStatus) SetStatus(v string) *GetGuardStatusResponseBodyProtectionStatus {
	s.Status = &v
	return s
}

func (s *GetGuardStatusResponseBodyProtectionStatus) SetType(v string) *GetGuardStatusResponseBodyProtectionStatus {
	s.Type = &v
	return s
}

func (s *GetGuardStatusResponseBodyProtectionStatus) Validate() error {
	return dara.Validate(s)
}

type GetGuardStatusResponseBodyRealTimeStatus struct {
	// The status. Valid values:
	//
	// - enabled: Running.
	//
	// - disabled: Not accessed.
	//
	// example:
	//
	// enabled
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The type.
	//
	// example:
	//
	// api
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GetGuardStatusResponseBodyRealTimeStatus) String() string {
	return dara.Prettify(s)
}

func (s GetGuardStatusResponseBodyRealTimeStatus) GoString() string {
	return s.String()
}

func (s *GetGuardStatusResponseBodyRealTimeStatus) GetStatus() *string {
	return s.Status
}

func (s *GetGuardStatusResponseBodyRealTimeStatus) GetType() *string {
	return s.Type
}

func (s *GetGuardStatusResponseBodyRealTimeStatus) SetStatus(v string) *GetGuardStatusResponseBodyRealTimeStatus {
	s.Status = &v
	return s
}

func (s *GetGuardStatusResponseBodyRealTimeStatus) SetType(v string) *GetGuardStatusResponseBodyRealTimeStatus {
	s.Type = &v
	return s
}

func (s *GetGuardStatusResponseBodyRealTimeStatus) Validate() error {
	return dara.Validate(s)
}
