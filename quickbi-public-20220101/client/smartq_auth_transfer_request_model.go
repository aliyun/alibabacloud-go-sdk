// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSmartqAuthTransferRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOriginUserId(v string) *SmartqAuthTransferRequest
	GetOriginUserId() *string
	SetTargetUserIds(v string) *SmartqAuthTransferRequest
	GetTargetUserIds() *string
}

type SmartqAuthTransferRequest struct {
	// Source user ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// ASDHASD*************12EASDA
	OriginUserId *string `json:"OriginUserId,omitempty" xml:"OriginUserId,omitempty"`
	// Target user ID array, separated by English commas.
	//
	// 	Warning: The number of user IDs cannot exceed 100.
	//
	// This parameter is required.
	//
	// example:
	//
	// 12313********dasfa,ASDASF*****SDAFEEG
	TargetUserIds *string `json:"TargetUserIds,omitempty" xml:"TargetUserIds,omitempty"`
}

func (s SmartqAuthTransferRequest) String() string {
	return dara.Prettify(s)
}

func (s SmartqAuthTransferRequest) GoString() string {
	return s.String()
}

func (s *SmartqAuthTransferRequest) GetOriginUserId() *string {
	return s.OriginUserId
}

func (s *SmartqAuthTransferRequest) GetTargetUserIds() *string {
	return s.TargetUserIds
}

func (s *SmartqAuthTransferRequest) SetOriginUserId(v string) *SmartqAuthTransferRequest {
	s.OriginUserId = &v
	return s
}

func (s *SmartqAuthTransferRequest) SetTargetUserIds(v string) *SmartqAuthTransferRequest {
	s.TargetUserIds = &v
	return s
}

func (s *SmartqAuthTransferRequest) Validate() error {
	return dara.Validate(s)
}
