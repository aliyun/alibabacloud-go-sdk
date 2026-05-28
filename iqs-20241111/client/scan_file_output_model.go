// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iScanFileOutput interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ScanFileOutput
	GetRequestId() *string
	SetScanFileInfoList(v []*ScanFileInfo) *ScanFileOutput
	GetScanFileInfoList() []*ScanFileInfo
	SetSearchInformation(v *UnifiedSearchInformation) *ScanFileOutput
	GetSearchInformation() *UnifiedSearchInformation
}

type ScanFileOutput struct {
	// example:
	//
	// ECB2144C-E277-5434-80E6-12D26678D364
	RequestId         *string                   `json:"requestId,omitempty" xml:"requestId,omitempty"`
	ScanFileInfoList  []*ScanFileInfo           `json:"scanFileInfoList,omitempty" xml:"scanFileInfoList,omitempty" type:"Repeated"`
	SearchInformation *UnifiedSearchInformation `json:"searchInformation,omitempty" xml:"searchInformation,omitempty"`
}

func (s ScanFileOutput) String() string {
	return dara.Prettify(s)
}

func (s ScanFileOutput) GoString() string {
	return s.String()
}

func (s *ScanFileOutput) GetRequestId() *string {
	return s.RequestId
}

func (s *ScanFileOutput) GetScanFileInfoList() []*ScanFileInfo {
	return s.ScanFileInfoList
}

func (s *ScanFileOutput) GetSearchInformation() *UnifiedSearchInformation {
	return s.SearchInformation
}

func (s *ScanFileOutput) SetRequestId(v string) *ScanFileOutput {
	s.RequestId = &v
	return s
}

func (s *ScanFileOutput) SetScanFileInfoList(v []*ScanFileInfo) *ScanFileOutput {
	s.ScanFileInfoList = v
	return s
}

func (s *ScanFileOutput) SetSearchInformation(v *UnifiedSearchInformation) *ScanFileOutput {
	s.SearchInformation = v
	return s
}

func (s *ScanFileOutput) Validate() error {
	if s.ScanFileInfoList != nil {
		for _, item := range s.ScanFileInfoList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.SearchInformation != nil {
		if err := s.SearchInformation.Validate(); err != nil {
			return err
		}
	}
	return nil
}
