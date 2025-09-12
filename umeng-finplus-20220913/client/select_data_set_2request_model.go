// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSelectDataSet2Request interface {
	dara.Model
	String() string
	GoString() string
	SetClientId(v int64) *SelectDataSet2Request
	GetClientId() *int64
	SetDataSetId(v int64) *SelectDataSet2Request
	GetDataSetId() *int64
}

type SelectDataSet2Request struct {
	ClientId  *int64 `json:"clientId,omitempty" xml:"clientId,omitempty"`
	DataSetId *int64 `json:"dataSetId,omitempty" xml:"dataSetId,omitempty"`
}

func (s SelectDataSet2Request) String() string {
	return dara.Prettify(s)
}

func (s SelectDataSet2Request) GoString() string {
	return s.String()
}

func (s *SelectDataSet2Request) GetClientId() *int64 {
	return s.ClientId
}

func (s *SelectDataSet2Request) GetDataSetId() *int64 {
	return s.DataSetId
}

func (s *SelectDataSet2Request) SetClientId(v int64) *SelectDataSet2Request {
	s.ClientId = &v
	return s
}

func (s *SelectDataSet2Request) SetDataSetId(v int64) *SelectDataSet2Request {
	s.DataSetId = &v
	return s
}

func (s *SelectDataSet2Request) Validate() error {
	return dara.Validate(s)
}
