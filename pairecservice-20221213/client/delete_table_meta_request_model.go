// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteTableMetaRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *DeleteTableMetaRequest
	GetInstanceId() *string
}

type DeleteTableMetaRequest struct {
	// This parameter is required.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s DeleteTableMetaRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteTableMetaRequest) GoString() string {
	return s.String()
}

func (s *DeleteTableMetaRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DeleteTableMetaRequest) SetInstanceId(v string) *DeleteTableMetaRequest {
	s.InstanceId = &v
	return s
}

func (s *DeleteTableMetaRequest) Validate() error {
	return dara.Validate(s)
}
