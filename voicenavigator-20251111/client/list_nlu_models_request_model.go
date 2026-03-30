// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListNluModelsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *ListNluModelsRequest
	GetInstanceId() *string
}

type ListNluModelsRequest struct {
	// example:
	//
	// af81a389-91f0-4157-8d82-720edd02b66a
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s ListNluModelsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListNluModelsRequest) GoString() string {
	return s.String()
}

func (s *ListNluModelsRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListNluModelsRequest) SetInstanceId(v string) *ListNluModelsRequest {
	s.InstanceId = &v
	return s
}

func (s *ListNluModelsRequest) Validate() error {
	return dara.Validate(s)
}
