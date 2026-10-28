// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPropertyValueRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBusinessChannel(v string) *ListPropertyValueRequest
	GetBusinessChannel() *string
	SetPropertyId(v int64) *ListPropertyValueRequest
	GetPropertyId() *int64
}

type ListPropertyValueRequest struct {
	// The channel.
	//
	// example:
	//
	// ENTERPRISE
	BusinessChannel *string `json:"BusinessChannel,omitempty" xml:"BusinessChannel,omitempty"`
	// The property ID. You can call [ListProperty](~~ListProperty~~) to query the property ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 390
	PropertyId *int64 `json:"PropertyId,omitempty" xml:"PropertyId,omitempty"`
}

func (s ListPropertyValueRequest) String() string {
	return dara.Prettify(s)
}

func (s ListPropertyValueRequest) GoString() string {
	return s.String()
}

func (s *ListPropertyValueRequest) GetBusinessChannel() *string {
	return s.BusinessChannel
}

func (s *ListPropertyValueRequest) GetPropertyId() *int64 {
	return s.PropertyId
}

func (s *ListPropertyValueRequest) SetBusinessChannel(v string) *ListPropertyValueRequest {
	s.BusinessChannel = &v
	return s
}

func (s *ListPropertyValueRequest) SetPropertyId(v int64) *ListPropertyValueRequest {
	s.PropertyId = &v
	return s
}

func (s *ListPropertyValueRequest) Validate() error {
	return dara.Validate(s)
}
