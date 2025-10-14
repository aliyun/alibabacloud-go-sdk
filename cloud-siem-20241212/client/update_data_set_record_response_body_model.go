// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateDataSetRecordResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDataSetRecordStatistic(v *UpdateDataSetRecordResponseBodyDataSetRecordStatistic) *UpdateDataSetRecordResponseBody
	GetDataSetRecordStatistic() *UpdateDataSetRecordResponseBodyDataSetRecordStatistic
	SetRequestId(v string) *UpdateDataSetRecordResponseBody
	GetRequestId() *string
}

type UpdateDataSetRecordResponseBody struct {
	DataSetRecordStatistic *UpdateDataSetRecordResponseBodyDataSetRecordStatistic `json:"DataSetRecordStatistic,omitempty" xml:"DataSetRecordStatistic,omitempty" type:"Struct"`
	// example:
	//
	// 9AAA9ED9-78F4-5021-86DC-D51C7511****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateDataSetRecordResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateDataSetRecordResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateDataSetRecordResponseBody) GetDataSetRecordStatistic() *UpdateDataSetRecordResponseBodyDataSetRecordStatistic {
	return s.DataSetRecordStatistic
}

func (s *UpdateDataSetRecordResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateDataSetRecordResponseBody) SetDataSetRecordStatistic(v *UpdateDataSetRecordResponseBodyDataSetRecordStatistic) *UpdateDataSetRecordResponseBody {
	s.DataSetRecordStatistic = v
	return s
}

func (s *UpdateDataSetRecordResponseBody) SetRequestId(v string) *UpdateDataSetRecordResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateDataSetRecordResponseBody) Validate() error {
	if s.DataSetRecordStatistic != nil {
		if err := s.DataSetRecordStatistic.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateDataSetRecordResponseBodyDataSetRecordStatistic struct {
	// example:
	//
	// 12
	NewDataSetRecordCount *int32 `json:"NewDataSetRecordCount,omitempty" xml:"NewDataSetRecordCount,omitempty"`
	// example:
	//
	// 4
	UpdateDataSetRecordCount *int32 `json:"UpdateDataSetRecordCount,omitempty" xml:"UpdateDataSetRecordCount,omitempty"`
}

func (s UpdateDataSetRecordResponseBodyDataSetRecordStatistic) String() string {
	return dara.Prettify(s)
}

func (s UpdateDataSetRecordResponseBodyDataSetRecordStatistic) GoString() string {
	return s.String()
}

func (s *UpdateDataSetRecordResponseBodyDataSetRecordStatistic) GetNewDataSetRecordCount() *int32 {
	return s.NewDataSetRecordCount
}

func (s *UpdateDataSetRecordResponseBodyDataSetRecordStatistic) GetUpdateDataSetRecordCount() *int32 {
	return s.UpdateDataSetRecordCount
}

func (s *UpdateDataSetRecordResponseBodyDataSetRecordStatistic) SetNewDataSetRecordCount(v int32) *UpdateDataSetRecordResponseBodyDataSetRecordStatistic {
	s.NewDataSetRecordCount = &v
	return s
}

func (s *UpdateDataSetRecordResponseBodyDataSetRecordStatistic) SetUpdateDataSetRecordCount(v int32) *UpdateDataSetRecordResponseBodyDataSetRecordStatistic {
	s.UpdateDataSetRecordCount = &v
	return s
}

func (s *UpdateDataSetRecordResponseBodyDataSetRecordStatistic) Validate() error {
	return dara.Validate(s)
}
