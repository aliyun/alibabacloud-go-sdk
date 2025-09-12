// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBuildStsAK2Request interface {
	dara.Model
	String() string
	GoString() string
	SetClientId(v int64) *BuildStsAK2Request
	GetClientId() *int64
	SetDataSetId(v int64) *BuildStsAK2Request
	GetDataSetId() *int64
}

type BuildStsAK2Request struct {
	ClientId  *int64 `json:"clientId,omitempty" xml:"clientId,omitempty"`
	DataSetId *int64 `json:"dataSetId,omitempty" xml:"dataSetId,omitempty"`
}

func (s BuildStsAK2Request) String() string {
	return dara.Prettify(s)
}

func (s BuildStsAK2Request) GoString() string {
	return s.String()
}

func (s *BuildStsAK2Request) GetClientId() *int64 {
	return s.ClientId
}

func (s *BuildStsAK2Request) GetDataSetId() *int64 {
	return s.DataSetId
}

func (s *BuildStsAK2Request) SetClientId(v int64) *BuildStsAK2Request {
	s.ClientId = &v
	return s
}

func (s *BuildStsAK2Request) SetDataSetId(v int64) *BuildStsAK2Request {
	s.DataSetId = &v
	return s
}

func (s *BuildStsAK2Request) Validate() error {
	return dara.Validate(s)
}
