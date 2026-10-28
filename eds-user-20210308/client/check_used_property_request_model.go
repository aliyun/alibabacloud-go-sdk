// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckUsedPropertyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBusinessChannel(v string) *CheckUsedPropertyRequest
	GetBusinessChannel() *string
	SetPropertyId(v int64) *CheckUsedPropertyRequest
	GetPropertyId() *int64
}

type CheckUsedPropertyRequest struct {
	// The channel.
	//
	// example:
	//
	// ENTERPRISE
	BusinessChannel *string `json:"BusinessChannel,omitempty" xml:"BusinessChannel,omitempty"`
	// The property ID. You can call the [ListProperty](~~ListProperty~~) operation to query property IDs.
	//
	// This parameter is required.
	//
	// example:
	//
	// 390
	PropertyId *int64 `json:"PropertyId,omitempty" xml:"PropertyId,omitempty"`
}

func (s CheckUsedPropertyRequest) String() string {
	return dara.Prettify(s)
}

func (s CheckUsedPropertyRequest) GoString() string {
	return s.String()
}

func (s *CheckUsedPropertyRequest) GetBusinessChannel() *string {
	return s.BusinessChannel
}

func (s *CheckUsedPropertyRequest) GetPropertyId() *int64 {
	return s.PropertyId
}

func (s *CheckUsedPropertyRequest) SetBusinessChannel(v string) *CheckUsedPropertyRequest {
	s.BusinessChannel = &v
	return s
}

func (s *CheckUsedPropertyRequest) SetPropertyId(v int64) *CheckUsedPropertyRequest {
	s.PropertyId = &v
	return s
}

func (s *CheckUsedPropertyRequest) Validate() error {
	return dara.Validate(s)
}
