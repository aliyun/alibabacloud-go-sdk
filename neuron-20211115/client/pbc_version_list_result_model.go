// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPbcVersionListResult interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*PbcVersion) *PbcVersionListResult
	GetData() []*PbcVersion
	SetRequestId(v string) *PbcVersionListResult
	GetRequestId() *string
	SetTotal(v int32) *PbcVersionListResult
	GetTotal() *int32
}

type PbcVersionListResult struct {
	Data      []*PbcVersion `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	RequestId *string       `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// 24
	Total *int32 `json:"total,omitempty" xml:"total,omitempty"`
}

func (s PbcVersionListResult) String() string {
	return dara.Prettify(s)
}

func (s PbcVersionListResult) GoString() string {
	return s.String()
}

func (s *PbcVersionListResult) GetData() []*PbcVersion {
	return s.Data
}

func (s *PbcVersionListResult) GetRequestId() *string {
	return s.RequestId
}

func (s *PbcVersionListResult) GetTotal() *int32 {
	return s.Total
}

func (s *PbcVersionListResult) SetData(v []*PbcVersion) *PbcVersionListResult {
	s.Data = v
	return s
}

func (s *PbcVersionListResult) SetRequestId(v string) *PbcVersionListResult {
	s.RequestId = &v
	return s
}

func (s *PbcVersionListResult) SetTotal(v int32) *PbcVersionListResult {
	s.Total = &v
	return s
}

func (s *PbcVersionListResult) Validate() error {
	if s.Data != nil {
		for _, item := range s.Data {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
