// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveDataSourceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDataSourceId(v string) *RemoveDataSourceRequest
	GetDataSourceId() *string
}

type RemoveDataSourceRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// ds-000bz3***myv90r
	DataSourceId *string `json:"DataSourceId,omitempty" xml:"DataSourceId,omitempty"`
}

func (s RemoveDataSourceRequest) String() string {
	return dara.Prettify(s)
}

func (s RemoveDataSourceRequest) GoString() string {
	return s.String()
}

func (s *RemoveDataSourceRequest) GetDataSourceId() *string {
	return s.DataSourceId
}

func (s *RemoveDataSourceRequest) SetDataSourceId(v string) *RemoveDataSourceRequest {
	s.DataSourceId = &v
	return s
}

func (s *RemoveDataSourceRequest) Validate() error {
	return dara.Validate(s)
}
