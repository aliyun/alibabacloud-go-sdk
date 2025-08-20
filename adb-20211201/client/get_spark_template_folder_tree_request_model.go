// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSparkTemplateFolderTreeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *GetSparkTemplateFolderTreeRequest
	GetDBClusterId() *string
}

type GetSparkTemplateFolderTreeRequest struct {
	// The ID of the AnalyticDB for MySQL Data Lakehouse Edition cluster.
	//
	// This parameter is required.
	//
	// example:
	//
	// amv-bp11q28kvl688****
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
}

func (s GetSparkTemplateFolderTreeRequest) String() string {
	return dara.Prettify(s)
}

func (s GetSparkTemplateFolderTreeRequest) GoString() string {
	return s.String()
}

func (s *GetSparkTemplateFolderTreeRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *GetSparkTemplateFolderTreeRequest) SetDBClusterId(v string) *GetSparkTemplateFolderTreeRequest {
	s.DBClusterId = &v
	return s
}

func (s *GetSparkTemplateFolderTreeRequest) Validate() error {
	return dara.Validate(s)
}
