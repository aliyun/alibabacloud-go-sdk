// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iScanSensitiveDataRequest interface {
	dara.Model
	String() string
	GoString() string
	SetData(v string) *ScanSensitiveDataRequest
	GetData() *string
}

type ScanSensitiveDataRequest struct {
	// The data that you want to check.
	//
	// This parameter is required.
	//
	// example:
	//
	// 13900001234
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
}

func (s ScanSensitiveDataRequest) String() string {
	return dara.Prettify(s)
}

func (s ScanSensitiveDataRequest) GoString() string {
	return s.String()
}

func (s *ScanSensitiveDataRequest) GetData() *string {
	return s.Data
}

func (s *ScanSensitiveDataRequest) SetData(v string) *ScanSensitiveDataRequest {
	s.Data = &v
	return s
}

func (s *ScanSensitiveDataRequest) Validate() error {
	return dara.Validate(s)
}
