// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveDataSet2Request interface {
	dara.Model
	String() string
	GoString() string
	SetClientId(v int64) *RemoveDataSet2Request
	GetClientId() *int64
	SetDataSetId(v int64) *RemoveDataSet2Request
	GetDataSetId() *int64
}

type RemoveDataSet2Request struct {
	ClientId  *int64 `json:"clientId,omitempty" xml:"clientId,omitempty"`
	DataSetId *int64 `json:"dataSetId,omitempty" xml:"dataSetId,omitempty"`
}

func (s RemoveDataSet2Request) String() string {
	return dara.Prettify(s)
}

func (s RemoveDataSet2Request) GoString() string {
	return s.String()
}

func (s *RemoveDataSet2Request) GetClientId() *int64 {
	return s.ClientId
}

func (s *RemoveDataSet2Request) GetDataSetId() *int64 {
	return s.DataSetId
}

func (s *RemoveDataSet2Request) SetClientId(v int64) *RemoveDataSet2Request {
	s.ClientId = &v
	return s
}

func (s *RemoveDataSet2Request) SetDataSetId(v int64) *RemoveDataSet2Request {
	s.DataSetId = &v
	return s
}

func (s *RemoveDataSet2Request) Validate() error {
	return dara.Validate(s)
}
