// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateIDCImportCommandRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *CreateIDCImportCommandRequest
	GetClusterId() *string
}

type CreateIDCImportCommandRequest struct {
	// The cluster ID. You can call the ListCluster operation to query the cluster ID. For more information, see [ListCluster](https://help.aliyun.com/document_detail/154995.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// 7246cxxx-53xx-xxxx-xxxx-xxxxxxxx
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
}

func (s CreateIDCImportCommandRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateIDCImportCommandRequest) GoString() string {
	return s.String()
}

func (s *CreateIDCImportCommandRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *CreateIDCImportCommandRequest) SetClusterId(v string) *CreateIDCImportCommandRequest {
	s.ClusterId = &v
	return s
}

func (s *CreateIDCImportCommandRequest) Validate() error {
	return dara.Validate(s)
}
