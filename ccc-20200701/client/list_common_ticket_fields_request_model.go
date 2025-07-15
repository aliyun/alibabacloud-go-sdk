// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCommonTicketFieldsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *ListCommonTicketFieldsRequest
	GetInstanceId() *string
}

type ListCommonTicketFieldsRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// ccc-test
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s ListCommonTicketFieldsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListCommonTicketFieldsRequest) GoString() string {
	return s.String()
}

func (s *ListCommonTicketFieldsRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListCommonTicketFieldsRequest) SetInstanceId(v string) *ListCommonTicketFieldsRequest {
	s.InstanceId = &v
	return s
}

func (s *ListCommonTicketFieldsRequest) Validate() error {
	return dara.Validate(s)
}
