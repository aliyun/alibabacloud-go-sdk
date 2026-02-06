// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteContactWhiteListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetContactWhiteListId(v string) *DeleteContactWhiteListRequest
	GetContactWhiteListId() *string
	SetInstanceId(v string) *DeleteContactWhiteListRequest
	GetInstanceId() *string
	SetOperator(v string) *DeleteContactWhiteListRequest
	GetOperator() *string
}

type DeleteContactWhiteListRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 111111
	ContactWhiteListId *string `json:"ContactWhiteListId,omitempty" xml:"ContactWhiteListId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// af81a389-91f0-4157-8d82-720edd02b66a
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// 160131
	Operator *string `json:"Operator,omitempty" xml:"Operator,omitempty"`
}

func (s DeleteContactWhiteListRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteContactWhiteListRequest) GoString() string {
	return s.String()
}

func (s *DeleteContactWhiteListRequest) GetContactWhiteListId() *string {
	return s.ContactWhiteListId
}

func (s *DeleteContactWhiteListRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DeleteContactWhiteListRequest) GetOperator() *string {
	return s.Operator
}

func (s *DeleteContactWhiteListRequest) SetContactWhiteListId(v string) *DeleteContactWhiteListRequest {
	s.ContactWhiteListId = &v
	return s
}

func (s *DeleteContactWhiteListRequest) SetInstanceId(v string) *DeleteContactWhiteListRequest {
	s.InstanceId = &v
	return s
}

func (s *DeleteContactWhiteListRequest) SetOperator(v string) *DeleteContactWhiteListRequest {
	s.Operator = &v
	return s
}

func (s *DeleteContactWhiteListRequest) Validate() error {
	return dara.Validate(s)
}
