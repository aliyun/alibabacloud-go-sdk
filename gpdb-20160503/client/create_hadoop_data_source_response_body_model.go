// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateHadoopDataSourceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDataSourceId(v int32) *CreateHadoopDataSourceResponseBody
	GetDataSourceId() *int32
	SetRequestId(v string) *CreateHadoopDataSourceResponseBody
	GetRequestId() *string
}

type CreateHadoopDataSourceResponseBody struct {
	// Data source ID.
	//
	// example:
	//
	// 123
	DataSourceId *int32 `json:"DataSourceId,omitempty" xml:"DataSourceId,omitempty"`
	// Request ID.
	//
	// example:
	//
	// 2C125605-266F-41CA-8AC5-3A643D4F42C5
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateHadoopDataSourceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateHadoopDataSourceResponseBody) GoString() string {
	return s.String()
}

func (s *CreateHadoopDataSourceResponseBody) GetDataSourceId() *int32 {
	return s.DataSourceId
}

func (s *CreateHadoopDataSourceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateHadoopDataSourceResponseBody) SetDataSourceId(v int32) *CreateHadoopDataSourceResponseBody {
	s.DataSourceId = &v
	return s
}

func (s *CreateHadoopDataSourceResponseBody) SetRequestId(v string) *CreateHadoopDataSourceResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateHadoopDataSourceResponseBody) Validate() error {
	return dara.Validate(s)
}
