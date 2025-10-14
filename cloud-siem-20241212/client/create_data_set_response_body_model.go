// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDataSetResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDataSetRecordStatistic(v *CreateDataSetResponseBodyDataSetRecordStatistic) *CreateDataSetResponseBody
	GetDataSetRecordStatistic() *CreateDataSetResponseBodyDataSetRecordStatistic
	SetRequestId(v string) *CreateDataSetResponseBody
	GetRequestId() *string
}

type CreateDataSetResponseBody struct {
	DataSetRecordStatistic *CreateDataSetResponseBodyDataSetRecordStatistic `json:"DataSetRecordStatistic,omitempty" xml:"DataSetRecordStatistic,omitempty" type:"Struct"`
	// example:
	//
	// 9AAA9ED9-78F4-5021-86DC-D51C7511****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateDataSetResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateDataSetResponseBody) GoString() string {
	return s.String()
}

func (s *CreateDataSetResponseBody) GetDataSetRecordStatistic() *CreateDataSetResponseBodyDataSetRecordStatistic {
	return s.DataSetRecordStatistic
}

func (s *CreateDataSetResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateDataSetResponseBody) SetDataSetRecordStatistic(v *CreateDataSetResponseBodyDataSetRecordStatistic) *CreateDataSetResponseBody {
	s.DataSetRecordStatistic = v
	return s
}

func (s *CreateDataSetResponseBody) SetRequestId(v string) *CreateDataSetResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateDataSetResponseBody) Validate() error {
	if s.DataSetRecordStatistic != nil {
		if err := s.DataSetRecordStatistic.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateDataSetResponseBodyDataSetRecordStatistic struct {
	// example:
	//
	// dataset-qt0n8246gs9nackg****
	DataSetId *string `json:"DataSetId,omitempty" xml:"DataSetId,omitempty"`
	// example:
	//
	// 6
	NewDataSetRecordCount *int32 `json:"NewDataSetRecordCount,omitempty" xml:"NewDataSetRecordCount,omitempty"`
}

func (s CreateDataSetResponseBodyDataSetRecordStatistic) String() string {
	return dara.Prettify(s)
}

func (s CreateDataSetResponseBodyDataSetRecordStatistic) GoString() string {
	return s.String()
}

func (s *CreateDataSetResponseBodyDataSetRecordStatistic) GetDataSetId() *string {
	return s.DataSetId
}

func (s *CreateDataSetResponseBodyDataSetRecordStatistic) GetNewDataSetRecordCount() *int32 {
	return s.NewDataSetRecordCount
}

func (s *CreateDataSetResponseBodyDataSetRecordStatistic) SetDataSetId(v string) *CreateDataSetResponseBodyDataSetRecordStatistic {
	s.DataSetId = &v
	return s
}

func (s *CreateDataSetResponseBodyDataSetRecordStatistic) SetNewDataSetRecordCount(v int32) *CreateDataSetResponseBodyDataSetRecordStatistic {
	s.NewDataSetRecordCount = &v
	return s
}

func (s *CreateDataSetResponseBodyDataSetRecordStatistic) Validate() error {
	return dara.Validate(s)
}
