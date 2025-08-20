// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSparkTemplateFullTreeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *GetSparkTemplateFullTreeRequest
	GetDBClusterId() *string
}

type GetSparkTemplateFullTreeRequest struct {
	// The ID of the AnalyticDB for MySQL Data Lakehouse Edition cluster.
	//
	// This parameter is required.
	//
	// example:
	//
	// amv-bp11q28kvl688****
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
}

func (s GetSparkTemplateFullTreeRequest) String() string {
	return dara.Prettify(s)
}

func (s GetSparkTemplateFullTreeRequest) GoString() string {
	return s.String()
}

func (s *GetSparkTemplateFullTreeRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *GetSparkTemplateFullTreeRequest) SetDBClusterId(v string) *GetSparkTemplateFullTreeRequest {
	s.DBClusterId = &v
	return s
}

func (s *GetSparkTemplateFullTreeRequest) Validate() error {
	return dara.Validate(s)
}
