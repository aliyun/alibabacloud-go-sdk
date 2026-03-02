// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMicroServicePageResult interface {
	dara.Model
	String() string
	GoString() string
	SetMicroServiceDigests(v []*MicroServiceDigest) *MicroServicePageResult
	GetMicroServiceDigests() []*MicroServiceDigest
	SetIntTotal(v int32) *MicroServicePageResult
	GetIntTotal() *int32
	SetRequestId(v string) *MicroServicePageResult
	GetRequestId() *string
	SetTotal(v int64) *MicroServicePageResult
	GetTotal() *int64
}

type MicroServicePageResult struct {
	MicroServiceDigests []*MicroServiceDigest `json:"MicroServiceDigests,omitempty" xml:"MicroServiceDigests,omitempty" type:"Repeated"`
	IntTotal            *int32                `json:"intTotal,omitempty" xml:"intTotal,omitempty"`
	RequestId           *string               `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// 24
	Total *int64 `json:"total,omitempty" xml:"total,omitempty"`
}

func (s MicroServicePageResult) String() string {
	return dara.Prettify(s)
}

func (s MicroServicePageResult) GoString() string {
	return s.String()
}

func (s *MicroServicePageResult) GetMicroServiceDigests() []*MicroServiceDigest {
	return s.MicroServiceDigests
}

func (s *MicroServicePageResult) GetIntTotal() *int32 {
	return s.IntTotal
}

func (s *MicroServicePageResult) GetRequestId() *string {
	return s.RequestId
}

func (s *MicroServicePageResult) GetTotal() *int64 {
	return s.Total
}

func (s *MicroServicePageResult) SetMicroServiceDigests(v []*MicroServiceDigest) *MicroServicePageResult {
	s.MicroServiceDigests = v
	return s
}

func (s *MicroServicePageResult) SetIntTotal(v int32) *MicroServicePageResult {
	s.IntTotal = &v
	return s
}

func (s *MicroServicePageResult) SetRequestId(v string) *MicroServicePageResult {
	s.RequestId = &v
	return s
}

func (s *MicroServicePageResult) SetTotal(v int64) *MicroServicePageResult {
	s.Total = &v
	return s
}

func (s *MicroServicePageResult) Validate() error {
	if s.MicroServiceDigests != nil {
		for _, item := range s.MicroServiceDigests {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
