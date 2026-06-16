// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchGroup interface {
	dara.Model
	String() string
	GoString() string
	SetBatchId(v string) *BatchGroup
	GetBatchId() *string
	SetRecords(v []*ExperimentRecord) *BatchGroup
	GetRecords() []*ExperimentRecord
}

type BatchGroup struct {
	BatchId *string             `json:"batchId,omitempty" xml:"batchId,omitempty"`
	Records []*ExperimentRecord `json:"records,omitempty" xml:"records,omitempty" type:"Repeated"`
}

func (s BatchGroup) String() string {
	return dara.Prettify(s)
}

func (s BatchGroup) GoString() string {
	return s.String()
}

func (s *BatchGroup) GetBatchId() *string {
	return s.BatchId
}

func (s *BatchGroup) GetRecords() []*ExperimentRecord {
	return s.Records
}

func (s *BatchGroup) SetBatchId(v string) *BatchGroup {
	s.BatchId = &v
	return s
}

func (s *BatchGroup) SetRecords(v []*ExperimentRecord) *BatchGroup {
	s.Records = v
	return s
}

func (s *BatchGroup) Validate() error {
	if s.Records != nil {
		for _, item := range s.Records {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
