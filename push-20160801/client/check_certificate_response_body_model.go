// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckCertificateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAndroid(v bool) *CheckCertificateResponseBody
	GetAndroid() *bool
	SetDevelopmentCertInfo(v *CheckCertificateResponseBodyDevelopmentCertInfo) *CheckCertificateResponseBody
	GetDevelopmentCertInfo() *CheckCertificateResponseBodyDevelopmentCertInfo
	SetIOS(v bool) *CheckCertificateResponseBody
	GetIOS() *bool
	SetProductionCertInfo(v *CheckCertificateResponseBodyProductionCertInfo) *CheckCertificateResponseBody
	GetProductionCertInfo() *CheckCertificateResponseBodyProductionCertInfo
	SetRequestId(v string) *CheckCertificateResponseBody
	GetRequestId() *string
}

type CheckCertificateResponseBody struct {
	// example:
	//
	// false
	Android             *bool                                            `json:"Android,omitempty" xml:"Android,omitempty"`
	DevelopmentCertInfo *CheckCertificateResponseBodyDevelopmentCertInfo `json:"DevelopmentCertInfo,omitempty" xml:"DevelopmentCertInfo,omitempty" type:"Struct"`
	// example:
	//
	// true
	IOS                *bool                                           `json:"IOS,omitempty" xml:"IOS,omitempty"`
	ProductionCertInfo *CheckCertificateResponseBodyProductionCertInfo `json:"ProductionCertInfo,omitempty" xml:"ProductionCertInfo,omitempty" type:"Struct"`
	// example:
	//
	// 9998B3CC-ED9E-4CB3-A8FB-DCC61296BFBC
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CheckCertificateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CheckCertificateResponseBody) GoString() string {
	return s.String()
}

func (s *CheckCertificateResponseBody) GetAndroid() *bool {
	return s.Android
}

func (s *CheckCertificateResponseBody) GetDevelopmentCertInfo() *CheckCertificateResponseBodyDevelopmentCertInfo {
	return s.DevelopmentCertInfo
}

func (s *CheckCertificateResponseBody) GetIOS() *bool {
	return s.IOS
}

func (s *CheckCertificateResponseBody) GetProductionCertInfo() *CheckCertificateResponseBodyProductionCertInfo {
	return s.ProductionCertInfo
}

func (s *CheckCertificateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CheckCertificateResponseBody) SetAndroid(v bool) *CheckCertificateResponseBody {
	s.Android = &v
	return s
}

func (s *CheckCertificateResponseBody) SetDevelopmentCertInfo(v *CheckCertificateResponseBodyDevelopmentCertInfo) *CheckCertificateResponseBody {
	s.DevelopmentCertInfo = v
	return s
}

func (s *CheckCertificateResponseBody) SetIOS(v bool) *CheckCertificateResponseBody {
	s.IOS = &v
	return s
}

func (s *CheckCertificateResponseBody) SetProductionCertInfo(v *CheckCertificateResponseBodyProductionCertInfo) *CheckCertificateResponseBody {
	s.ProductionCertInfo = v
	return s
}

func (s *CheckCertificateResponseBody) SetRequestId(v string) *CheckCertificateResponseBody {
	s.RequestId = &v
	return s
}

func (s *CheckCertificateResponseBody) Validate() error {
	if s.DevelopmentCertInfo != nil {
		if err := s.DevelopmentCertInfo.Validate(); err != nil {
			return err
		}
	}
	if s.ProductionCertInfo != nil {
		if err := s.ProductionCertInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CheckCertificateResponseBodyDevelopmentCertInfo struct {
	// example:
	//
	// 1470024000000
	ExipreTime *int64 `json:"ExipreTime,omitempty" xml:"ExipreTime,omitempty"`
	// example:
	//
	// EXPIRED
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s CheckCertificateResponseBodyDevelopmentCertInfo) String() string {
	return dara.Prettify(s)
}

func (s CheckCertificateResponseBodyDevelopmentCertInfo) GoString() string {
	return s.String()
}

func (s *CheckCertificateResponseBodyDevelopmentCertInfo) GetExipreTime() *int64 {
	return s.ExipreTime
}

func (s *CheckCertificateResponseBodyDevelopmentCertInfo) GetStatus() *string {
	return s.Status
}

func (s *CheckCertificateResponseBodyDevelopmentCertInfo) SetExipreTime(v int64) *CheckCertificateResponseBodyDevelopmentCertInfo {
	s.ExipreTime = &v
	return s
}

func (s *CheckCertificateResponseBodyDevelopmentCertInfo) SetStatus(v string) *CheckCertificateResponseBodyDevelopmentCertInfo {
	s.Status = &v
	return s
}

func (s *CheckCertificateResponseBodyDevelopmentCertInfo) Validate() error {
	return dara.Validate(s)
}

type CheckCertificateResponseBodyProductionCertInfo struct {
	// example:
	//
	// 1764561600000
	ExipreTime *int64 `json:"ExipreTime,omitempty" xml:"ExipreTime,omitempty"`
	// example:
	//
	// OK
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s CheckCertificateResponseBodyProductionCertInfo) String() string {
	return dara.Prettify(s)
}

func (s CheckCertificateResponseBodyProductionCertInfo) GoString() string {
	return s.String()
}

func (s *CheckCertificateResponseBodyProductionCertInfo) GetExipreTime() *int64 {
	return s.ExipreTime
}

func (s *CheckCertificateResponseBodyProductionCertInfo) GetStatus() *string {
	return s.Status
}

func (s *CheckCertificateResponseBodyProductionCertInfo) SetExipreTime(v int64) *CheckCertificateResponseBodyProductionCertInfo {
	s.ExipreTime = &v
	return s
}

func (s *CheckCertificateResponseBodyProductionCertInfo) SetStatus(v string) *CheckCertificateResponseBodyProductionCertInfo {
	s.Status = &v
	return s
}

func (s *CheckCertificateResponseBodyProductionCertInfo) Validate() error {
	return dara.Validate(s)
}
