// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetRegistryScanDayNumRequest interface {
	dara.Model
	String() string
	GoString() string
	SetScanDayNum(v int32) *SetRegistryScanDayNumRequest
	GetScanDayNum() *int32
}

type SetRegistryScanDayNumRequest struct {
	// The cycle at which you want to scan your images. Unit: days.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	ScanDayNum *int32 `json:"ScanDayNum,omitempty" xml:"ScanDayNum,omitempty"`
}

func (s SetRegistryScanDayNumRequest) String() string {
	return dara.Prettify(s)
}

func (s SetRegistryScanDayNumRequest) GoString() string {
	return s.String()
}

func (s *SetRegistryScanDayNumRequest) GetScanDayNum() *int32 {
	return s.ScanDayNum
}

func (s *SetRegistryScanDayNumRequest) SetScanDayNum(v int32) *SetRegistryScanDayNumRequest {
	s.ScanDayNum = &v
	return s
}

func (s *SetRegistryScanDayNumRequest) Validate() error {
	return dara.Validate(s)
}
