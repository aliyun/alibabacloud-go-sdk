// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteCustomCallTaggingRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *DeleteCustomCallTaggingRequest
	GetInstanceId() *string
	SetNumber(v string) *DeleteCustomCallTaggingRequest
	GetNumber() *string
}

type DeleteCustomCallTaggingRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// ccc-test
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1312121****
	Number *string `json:"Number,omitempty" xml:"Number,omitempty"`
}

func (s DeleteCustomCallTaggingRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteCustomCallTaggingRequest) GoString() string {
	return s.String()
}

func (s *DeleteCustomCallTaggingRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DeleteCustomCallTaggingRequest) GetNumber() *string {
	return s.Number
}

func (s *DeleteCustomCallTaggingRequest) SetInstanceId(v string) *DeleteCustomCallTaggingRequest {
	s.InstanceId = &v
	return s
}

func (s *DeleteCustomCallTaggingRequest) SetNumber(v string) *DeleteCustomCallTaggingRequest {
	s.Number = &v
	return s
}

func (s *DeleteCustomCallTaggingRequest) Validate() error {
	return dara.Validate(s)
}
