// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iContinuouslyPushRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppKey(v int64) *ContinuouslyPushRequest
	GetAppKey() *int64
	SetMessageId(v string) *ContinuouslyPushRequest
	GetMessageId() *string
	SetTarget(v string) *ContinuouslyPushRequest
	GetTarget() *string
	SetTargetValue(v string) *ContinuouslyPushRequest
	GetTargetValue() *string
}

type ContinuouslyPushRequest struct {
	// The AppKey information.
	//
	// This parameter is required.
	//
	// example:
	//
	// 23267207
	AppKey *int64 `json:"AppKey,omitempty" xml:"AppKey,omitempty"`
	// The message ID.
	//
	// The MessageId obtained after calling the Push operation and setting Target to TBD. It represents a message that has been saved to the push system.
	//
	// This parameter is required.
	//
	// example:
	//
	// 500131
	MessageId *string `json:"MessageId,omitempty" xml:"MessageId,omitempty"`
	// The push target:
	//
	// - **DEVICE**: Push based on devices.
	//
	// - **ACCOUNT**: Push based on accounts.
	//
	// - **ALIAS**: Push based on aliases.
	//
	// Continuous push supports only these three target types.
	//
	// This parameter is required.
	//
	// example:
	//
	// DEVICE
	Target *string `json:"Target,omitempty" xml:"Target,omitempty"`
	// Set this parameter based on Target. Separate multiple values with commas. If the limit is exceeded, push the message in multiple batches.
	//
	// - If Target is set to DEVICE, the value is in the format of `deviceid1,deviceid2`. A maximum of 1,000 devices are supported.
	//
	// - If Target is set to ACCOUNT, the value is in the format of `account1,account2`. A maximum of 1,000 accounts are supported.
	//
	// - If Target is set to ALIAS, the value is in the format of `alias1,alias2`. A maximum of 1,000 aliases are supported.
	//
	// This parameter is required.
	//
	// example:
	//
	// a692961a92534047ad3625****
	TargetValue *string `json:"TargetValue,omitempty" xml:"TargetValue,omitempty"`
}

func (s ContinuouslyPushRequest) String() string {
	return dara.Prettify(s)
}

func (s ContinuouslyPushRequest) GoString() string {
	return s.String()
}

func (s *ContinuouslyPushRequest) GetAppKey() *int64 {
	return s.AppKey
}

func (s *ContinuouslyPushRequest) GetMessageId() *string {
	return s.MessageId
}

func (s *ContinuouslyPushRequest) GetTarget() *string {
	return s.Target
}

func (s *ContinuouslyPushRequest) GetTargetValue() *string {
	return s.TargetValue
}

func (s *ContinuouslyPushRequest) SetAppKey(v int64) *ContinuouslyPushRequest {
	s.AppKey = &v
	return s
}

func (s *ContinuouslyPushRequest) SetMessageId(v string) *ContinuouslyPushRequest {
	s.MessageId = &v
	return s
}

func (s *ContinuouslyPushRequest) SetTarget(v string) *ContinuouslyPushRequest {
	s.Target = &v
	return s
}

func (s *ContinuouslyPushRequest) SetTargetValue(v string) *ContinuouslyPushRequest {
	s.TargetValue = &v
	return s
}

func (s *ContinuouslyPushRequest) Validate() error {
	return dara.Validate(s)
}
