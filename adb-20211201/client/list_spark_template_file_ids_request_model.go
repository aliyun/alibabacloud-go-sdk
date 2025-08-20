// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSparkTemplateFileIdsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *ListSparkTemplateFileIdsRequest
	GetDBClusterId() *string
}

type ListSparkTemplateFileIdsRequest struct {
	// The ID of the AnalyticDB for MySQL Data Lakehouse Edition cluster.
	//
	// example:
	//
	// amv-bp11q28kvl688****
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
}

func (s ListSparkTemplateFileIdsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListSparkTemplateFileIdsRequest) GoString() string {
	return s.String()
}

func (s *ListSparkTemplateFileIdsRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *ListSparkTemplateFileIdsRequest) SetDBClusterId(v string) *ListSparkTemplateFileIdsRequest {
	s.DBClusterId = &v
	return s
}

func (s *ListSparkTemplateFileIdsRequest) Validate() error {
	return dara.Validate(s)
}
