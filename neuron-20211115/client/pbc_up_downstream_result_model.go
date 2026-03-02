// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPbcUpDownstreamResult interface {
	dara.Model
	String() string
	GoString() string
	SetPbcItems(v []*PbcRelationItem) *PbcUpDownstreamResult
	GetPbcItems() []*PbcRelationItem
	SetRequestId(v string) *PbcUpDownstreamResult
	GetRequestId() *string
}

type PbcUpDownstreamResult struct {
	PbcItems  []*PbcRelationItem `json:"pbcItems,omitempty" xml:"pbcItems,omitempty" type:"Repeated"`
	RequestId *string            `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s PbcUpDownstreamResult) String() string {
	return dara.Prettify(s)
}

func (s PbcUpDownstreamResult) GoString() string {
	return s.String()
}

func (s *PbcUpDownstreamResult) GetPbcItems() []*PbcRelationItem {
	return s.PbcItems
}

func (s *PbcUpDownstreamResult) GetRequestId() *string {
	return s.RequestId
}

func (s *PbcUpDownstreamResult) SetPbcItems(v []*PbcRelationItem) *PbcUpDownstreamResult {
	s.PbcItems = v
	return s
}

func (s *PbcUpDownstreamResult) SetRequestId(v string) *PbcUpDownstreamResult {
	s.RequestId = &v
	return s
}

func (s *PbcUpDownstreamResult) Validate() error {
	if s.PbcItems != nil {
		for _, item := range s.PbcItems {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
