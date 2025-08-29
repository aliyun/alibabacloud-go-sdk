// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOfflineLaboratoryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *OfflineLaboratoryRequest
	GetInstanceId() *string
}

type OfflineLaboratoryRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// pairec-cn-abcdefg1234
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s OfflineLaboratoryRequest) String() string {
	return dara.Prettify(s)
}

func (s OfflineLaboratoryRequest) GoString() string {
	return s.String()
}

func (s *OfflineLaboratoryRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *OfflineLaboratoryRequest) SetInstanceId(v string) *OfflineLaboratoryRequest {
	s.InstanceId = &v
	return s
}

func (s *OfflineLaboratoryRequest) Validate() error {
	return dara.Validate(s)
}
