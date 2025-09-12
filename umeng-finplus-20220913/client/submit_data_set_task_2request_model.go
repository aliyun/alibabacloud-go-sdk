// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitDataSetTask2Request interface {
	dara.Model
	String() string
	GoString() string
	SetClientId(v int64) *SubmitDataSetTask2Request
	GetClientId() *int64
	SetDataSetId(v int64) *SubmitDataSetTask2Request
	GetDataSetId() *int64
}

type SubmitDataSetTask2Request struct {
	ClientId  *int64 `json:"clientId,omitempty" xml:"clientId,omitempty"`
	DataSetId *int64 `json:"dataSetId,omitempty" xml:"dataSetId,omitempty"`
}

func (s SubmitDataSetTask2Request) String() string {
	return dara.Prettify(s)
}

func (s SubmitDataSetTask2Request) GoString() string {
	return s.String()
}

func (s *SubmitDataSetTask2Request) GetClientId() *int64 {
	return s.ClientId
}

func (s *SubmitDataSetTask2Request) GetDataSetId() *int64 {
	return s.DataSetId
}

func (s *SubmitDataSetTask2Request) SetClientId(v int64) *SubmitDataSetTask2Request {
	s.ClientId = &v
	return s
}

func (s *SubmitDataSetTask2Request) SetDataSetId(v int64) *SubmitDataSetTask2Request {
	s.DataSetId = &v
	return s
}

func (s *SubmitDataSetTask2Request) Validate() error {
	return dara.Validate(s)
}
