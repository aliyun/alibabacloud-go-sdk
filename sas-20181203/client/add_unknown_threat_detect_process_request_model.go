// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddUnknownThreatDetectProcessRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEventIdList(v []*int64) *AddUnknownThreatDetectProcessRequest
	GetEventIdList() []*int64
	SetProcessList(v []*AddUnknownThreatDetectProcessRequestProcessList) *AddUnknownThreatDetectProcessRequest
	GetProcessList() []*AddUnknownThreatDetectProcessRequestProcessList
	SetUuidList(v []*string) *AddUnknownThreatDetectProcessRequest
	GetUuidList() []*string
}

type AddUnknownThreatDetectProcessRequest struct {
	EventIdList []*int64                                           `json:"EventIdList,omitempty" xml:"EventIdList,omitempty" type:"Repeated"`
	ProcessList []*AddUnknownThreatDetectProcessRequestProcessList `json:"ProcessList,omitempty" xml:"ProcessList,omitempty" type:"Repeated"`
	UuidList    []*string                                          `json:"UuidList,omitempty" xml:"UuidList,omitempty" type:"Repeated"`
}

func (s AddUnknownThreatDetectProcessRequest) String() string {
	return dara.Prettify(s)
}

func (s AddUnknownThreatDetectProcessRequest) GoString() string {
	return s.String()
}

func (s *AddUnknownThreatDetectProcessRequest) GetEventIdList() []*int64 {
	return s.EventIdList
}

func (s *AddUnknownThreatDetectProcessRequest) GetProcessList() []*AddUnknownThreatDetectProcessRequestProcessList {
	return s.ProcessList
}

func (s *AddUnknownThreatDetectProcessRequest) GetUuidList() []*string {
	return s.UuidList
}

func (s *AddUnknownThreatDetectProcessRequest) SetEventIdList(v []*int64) *AddUnknownThreatDetectProcessRequest {
	s.EventIdList = v
	return s
}

func (s *AddUnknownThreatDetectProcessRequest) SetProcessList(v []*AddUnknownThreatDetectProcessRequestProcessList) *AddUnknownThreatDetectProcessRequest {
	s.ProcessList = v
	return s
}

func (s *AddUnknownThreatDetectProcessRequest) SetUuidList(v []*string) *AddUnknownThreatDetectProcessRequest {
	s.UuidList = v
	return s
}

func (s *AddUnknownThreatDetectProcessRequest) Validate() error {
	if s.ProcessList != nil {
		for _, item := range s.ProcessList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type AddUnknownThreatDetectProcessRequestProcessList struct {
	// example:
	//
	// e59b63ae983377f131ab20ec0d******
	Md5 *string `json:"Md5,omitempty" xml:"Md5,omitempty"`
	// example:
	//
	// /bin/rm
	ProcessPath *string `json:"ProcessPath,omitempty" xml:"ProcessPath,omitempty"`
	// example:
	//
	// 1330
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	// example:
	//
	// f204693a7d2ce99d6c4434e550d985ee1c7be7cb5dd9a76094369af0d2******
	Sha256 *string `json:"Sha256,omitempty" xml:"Sha256,omitempty"`
}

func (s AddUnknownThreatDetectProcessRequestProcessList) String() string {
	return dara.Prettify(s)
}

func (s AddUnknownThreatDetectProcessRequestProcessList) GoString() string {
	return s.String()
}

func (s *AddUnknownThreatDetectProcessRequestProcessList) GetMd5() *string {
	return s.Md5
}

func (s *AddUnknownThreatDetectProcessRequestProcessList) GetProcessPath() *string {
	return s.ProcessPath
}

func (s *AddUnknownThreatDetectProcessRequestProcessList) GetRemark() *string {
	return s.Remark
}

func (s *AddUnknownThreatDetectProcessRequestProcessList) GetSha256() *string {
	return s.Sha256
}

func (s *AddUnknownThreatDetectProcessRequestProcessList) SetMd5(v string) *AddUnknownThreatDetectProcessRequestProcessList {
	s.Md5 = &v
	return s
}

func (s *AddUnknownThreatDetectProcessRequestProcessList) SetProcessPath(v string) *AddUnknownThreatDetectProcessRequestProcessList {
	s.ProcessPath = &v
	return s
}

func (s *AddUnknownThreatDetectProcessRequestProcessList) SetRemark(v string) *AddUnknownThreatDetectProcessRequestProcessList {
	s.Remark = &v
	return s
}

func (s *AddUnknownThreatDetectProcessRequestProcessList) SetSha256(v string) *AddUnknownThreatDetectProcessRequestProcessList {
	s.Sha256 = &v
	return s
}

func (s *AddUnknownThreatDetectProcessRequestProcessList) Validate() error {
	return dara.Validate(s)
}
