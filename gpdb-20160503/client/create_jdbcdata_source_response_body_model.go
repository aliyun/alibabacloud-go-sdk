// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateJDBCDataSourceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDataSourceId(v string) *CreateJDBCDataSourceResponseBody
	GetDataSourceId() *string
	SetRequestId(v string) *CreateJDBCDataSourceResponseBody
	GetRequestId() *string
}

type CreateJDBCDataSourceResponseBody struct {
	// Data source ID.
	//
	// example:
	//
	// 123
	DataSourceId *string `json:"DataSourceId,omitempty" xml:"DataSourceId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 2C125605-266F-41CA-8AC5-3A643D4F42C5
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateJDBCDataSourceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateJDBCDataSourceResponseBody) GoString() string {
	return s.String()
}

func (s *CreateJDBCDataSourceResponseBody) GetDataSourceId() *string {
	return s.DataSourceId
}

func (s *CreateJDBCDataSourceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateJDBCDataSourceResponseBody) SetDataSourceId(v string) *CreateJDBCDataSourceResponseBody {
	s.DataSourceId = &v
	return s
}

func (s *CreateJDBCDataSourceResponseBody) SetRequestId(v string) *CreateJDBCDataSourceResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateJDBCDataSourceResponseBody) Validate() error {
	return dara.Validate(s)
}
