// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCustomCallTaggingRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCustomNumberList(v string) *CreateCustomCallTaggingRequest
	GetCustomNumberList() *string
	SetInstanceId(v string) *CreateCustomCallTaggingRequest
	GetInstanceId() *string
}

type CreateCustomCallTaggingRequest struct {
	// This parameter is required.
	CustomNumberList *string `json:"CustomNumberList,omitempty" xml:"CustomNumberList,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ccc-test
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s CreateCustomCallTaggingRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateCustomCallTaggingRequest) GoString() string {
	return s.String()
}

func (s *CreateCustomCallTaggingRequest) GetCustomNumberList() *string {
	return s.CustomNumberList
}

func (s *CreateCustomCallTaggingRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CreateCustomCallTaggingRequest) SetCustomNumberList(v string) *CreateCustomCallTaggingRequest {
	s.CustomNumberList = &v
	return s
}

func (s *CreateCustomCallTaggingRequest) SetInstanceId(v string) *CreateCustomCallTaggingRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateCustomCallTaggingRequest) Validate() error {
	return dara.Validate(s)
}
