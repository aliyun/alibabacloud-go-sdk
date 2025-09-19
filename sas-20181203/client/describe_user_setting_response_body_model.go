// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeUserSettingResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAlertLevels(v []*string) *DescribeUserSettingResponseBody
	GetAlertLevels() []*string
	SetInvalidWarningKeepDays(v int32) *DescribeUserSettingResponseBody
	GetInvalidWarningKeepDays() *int32
	SetRequestId(v string) *DescribeUserSettingResponseBody
	GetRequestId() *string
}

type DescribeUserSettingResponseBody struct {
	// The severities of alerts. If this parameter is empty, no custom alerts are generated.
	AlertLevels []*string `json:"AlertLevels,omitempty" xml:"AlertLevels,omitempty" type:"Repeated"`
	// The number of days during which you want to retain invalid alerts.
	//
	// example:
	//
	// 30
	InvalidWarningKeepDays *int32 `json:"InvalidWarningKeepDays,omitempty" xml:"InvalidWarningKeepDays,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 09969D2C-4FAD-429E-BFBF-9A60DEF8****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeUserSettingResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeUserSettingResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeUserSettingResponseBody) GetAlertLevels() []*string {
	return s.AlertLevels
}

func (s *DescribeUserSettingResponseBody) GetInvalidWarningKeepDays() *int32 {
	return s.InvalidWarningKeepDays
}

func (s *DescribeUserSettingResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeUserSettingResponseBody) SetAlertLevels(v []*string) *DescribeUserSettingResponseBody {
	s.AlertLevels = v
	return s
}

func (s *DescribeUserSettingResponseBody) SetInvalidWarningKeepDays(v int32) *DescribeUserSettingResponseBody {
	s.InvalidWarningKeepDays = &v
	return s
}

func (s *DescribeUserSettingResponseBody) SetRequestId(v string) *DescribeUserSettingResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeUserSettingResponseBody) Validate() error {
	return dara.Validate(s)
}
