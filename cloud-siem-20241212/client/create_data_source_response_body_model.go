// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDataSourceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDataSourceId(v string) *CreateDataSourceResponseBody
	GetDataSourceId() *string
	SetRequestId(v string) *CreateDataSourceResponseBody
	GetRequestId() *string
}

type CreateDataSourceResponseBody struct {
	// example:
	//
	// ds-jl67vixpe1scwysgyu3x。
	DataSourceId *string `json:"DataSourceId,omitempty" xml:"DataSourceId,omitempty"`
	// example:
	//
	// 6276D891-*****-55B2-87B9-74D413F7****。
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateDataSourceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateDataSourceResponseBody) GoString() string {
	return s.String()
}

func (s *CreateDataSourceResponseBody) GetDataSourceId() *string {
	return s.DataSourceId
}

func (s *CreateDataSourceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateDataSourceResponseBody) SetDataSourceId(v string) *CreateDataSourceResponseBody {
	s.DataSourceId = &v
	return s
}

func (s *CreateDataSourceResponseBody) SetRequestId(v string) *CreateDataSourceResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateDataSourceResponseBody) Validate() error {
	return dara.Validate(s)
}
