// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTableMetaRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *GetTableMetaRequest
	GetInstanceId() *string
}

type GetTableMetaRequest struct {
	// The ID of the instance. You can call the [ListInstances](https://help.aliyun.com/document_detail/2411819.html) operation to obtain the instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// Instance ID	pairec-cn-test1
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s GetTableMetaRequest) String() string {
	return dara.Prettify(s)
}

func (s GetTableMetaRequest) GoString() string {
	return s.String()
}

func (s *GetTableMetaRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetTableMetaRequest) SetInstanceId(v string) *GetTableMetaRequest {
	s.InstanceId = &v
	return s
}

func (s *GetTableMetaRequest) Validate() error {
	return dara.Validate(s)
}
