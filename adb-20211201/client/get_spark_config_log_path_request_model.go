// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSparkConfigLogPathRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *GetSparkConfigLogPathRequest
	GetDBClusterId() *string
}

type GetSparkConfigLogPathRequest struct {
	// The ID of the AnalyticDB for MySQL Data Lakehouse Edition cluster.
	//
	// This parameter is required.
	//
	// example:
	//
	// am-adsdxxxx
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
}

func (s GetSparkConfigLogPathRequest) String() string {
	return dara.Prettify(s)
}

func (s GetSparkConfigLogPathRequest) GoString() string {
	return s.String()
}

func (s *GetSparkConfigLogPathRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *GetSparkConfigLogPathRequest) SetDBClusterId(v string) *GetSparkConfigLogPathRequest {
	s.DBClusterId = &v
	return s
}

func (s *GetSparkConfigLogPathRequest) Validate() error {
	return dara.Validate(s)
}
