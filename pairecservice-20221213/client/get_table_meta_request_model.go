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
	// This parameter is required.
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
