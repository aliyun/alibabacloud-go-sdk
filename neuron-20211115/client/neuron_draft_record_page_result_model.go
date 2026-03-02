// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iNeuronDraftRecordPageResult interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*NeuronDraftRecord) *NeuronDraftRecordPageResult
	GetData() []*NeuronDraftRecord
	SetRequestId(v string) *NeuronDraftRecordPageResult
	GetRequestId() *string
	SetTotal(v int32) *NeuronDraftRecordPageResult
	GetTotal() *int32
}

type NeuronDraftRecordPageResult struct {
	Data []*NeuronDraftRecord `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	// example:
	//
	// 3239281273464326823
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// 24
	Total *int32 `json:"total,omitempty" xml:"total,omitempty"`
}

func (s NeuronDraftRecordPageResult) String() string {
	return dara.Prettify(s)
}

func (s NeuronDraftRecordPageResult) GoString() string {
	return s.String()
}

func (s *NeuronDraftRecordPageResult) GetData() []*NeuronDraftRecord {
	return s.Data
}

func (s *NeuronDraftRecordPageResult) GetRequestId() *string {
	return s.RequestId
}

func (s *NeuronDraftRecordPageResult) GetTotal() *int32 {
	return s.Total
}

func (s *NeuronDraftRecordPageResult) SetData(v []*NeuronDraftRecord) *NeuronDraftRecordPageResult {
	s.Data = v
	return s
}

func (s *NeuronDraftRecordPageResult) SetRequestId(v string) *NeuronDraftRecordPageResult {
	s.RequestId = &v
	return s
}

func (s *NeuronDraftRecordPageResult) SetTotal(v int32) *NeuronDraftRecordPageResult {
	s.Total = &v
	return s
}

func (s *NeuronDraftRecordPageResult) Validate() error {
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
