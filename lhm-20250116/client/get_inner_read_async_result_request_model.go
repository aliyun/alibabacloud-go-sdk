// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetInnerReadAsyncResultRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDataSourceName(v string) *GetInnerReadAsyncResultRequest
	GetDataSourceName() *string
}

type GetInnerReadAsyncResultRequest struct {
	// The data source name. The probe task uses this field as its dimension identifier.
	//
	// example:
	//
	// ds_dolphin_prod
	DataSourceName *string `json:"dataSourceName,omitempty" xml:"dataSourceName,omitempty"`
}

func (s GetInnerReadAsyncResultRequest) String() string {
	return dara.Prettify(s)
}

func (s GetInnerReadAsyncResultRequest) GoString() string {
	return s.String()
}

func (s *GetInnerReadAsyncResultRequest) GetDataSourceName() *string {
	return s.DataSourceName
}

func (s *GetInnerReadAsyncResultRequest) SetDataSourceName(v string) *GetInnerReadAsyncResultRequest {
	s.DataSourceName = &v
	return s
}

func (s *GetInnerReadAsyncResultRequest) Validate() error {
	return dara.Validate(s)
}
