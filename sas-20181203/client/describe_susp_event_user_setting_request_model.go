// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSuspEventUserSettingRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFrom(v string) *DescribeSuspEventUserSettingRequest
	GetFrom() *string
	SetId(v int32) *DescribeSuspEventUserSettingRequest
	GetId() *int32
	SetSourceIp(v string) *DescribeSuspEventUserSettingRequest
	GetSourceIp() *string
}

type DescribeSuspEventUserSettingRequest struct {
	// The ID of the request source. Set the value to **sas**.
	//
	// example:
	//
	// sas
	From *string `json:"From,omitempty" xml:"From,omitempty"`
	// The ID. You do not need to specify this parameter.
	//
	// example:
	//
	// 123
	Id *int32 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The IP address of the request. You do not need to specify this parameter.
	//
	// example:
	//
	// 127.0.XX.XX
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
}

func (s DescribeSuspEventUserSettingRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeSuspEventUserSettingRequest) GoString() string {
	return s.String()
}

func (s *DescribeSuspEventUserSettingRequest) GetFrom() *string {
	return s.From
}

func (s *DescribeSuspEventUserSettingRequest) GetId() *int32 {
	return s.Id
}

func (s *DescribeSuspEventUserSettingRequest) GetSourceIp() *string {
	return s.SourceIp
}

func (s *DescribeSuspEventUserSettingRequest) SetFrom(v string) *DescribeSuspEventUserSettingRequest {
	s.From = &v
	return s
}

func (s *DescribeSuspEventUserSettingRequest) SetId(v int32) *DescribeSuspEventUserSettingRequest {
	s.Id = &v
	return s
}

func (s *DescribeSuspEventUserSettingRequest) SetSourceIp(v string) *DescribeSuspEventUserSettingRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeSuspEventUserSettingRequest) Validate() error {
	return dara.Validate(s)
}
