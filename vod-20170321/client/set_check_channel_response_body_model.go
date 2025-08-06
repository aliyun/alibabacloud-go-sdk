// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetCheckChannelResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAudit(v *SetCheckChannelResponseBodyAudit) *SetCheckChannelResponseBody
	GetAudit() *SetCheckChannelResponseBodyAudit
	SetRequestId(v string) *SetCheckChannelResponseBody
	GetRequestId() *string
}

type SetCheckChannelResponseBody struct {
	Audit     *SetCheckChannelResponseBodyAudit `json:"Audit,omitempty" xml:"Audit,omitempty" type:"Struct"`
	RequestId *string                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SetCheckChannelResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SetCheckChannelResponseBody) GoString() string {
	return s.String()
}

func (s *SetCheckChannelResponseBody) GetAudit() *SetCheckChannelResponseBodyAudit {
	return s.Audit
}

func (s *SetCheckChannelResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SetCheckChannelResponseBody) SetAudit(v *SetCheckChannelResponseBodyAudit) *SetCheckChannelResponseBody {
	s.Audit = v
	return s
}

func (s *SetCheckChannelResponseBody) SetRequestId(v string) *SetCheckChannelResponseBody {
	s.RequestId = &v
	return s
}

func (s *SetCheckChannelResponseBody) Validate() error {
	return dara.Validate(s)
}

type SetCheckChannelResponseBodyAudit struct {
	Channel     *string `json:"Channel,omitempty" xml:"Channel,omitempty"`
	CreateTime  *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	CustomerId  *string `json:"CustomerId,omitempty" xml:"CustomerId,omitempty"`
	LegalSwitch *string `json:"LegalSwitch,omitempty" xml:"LegalSwitch,omitempty"`
	UpdateTime  *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s SetCheckChannelResponseBodyAudit) String() string {
	return dara.Prettify(s)
}

func (s SetCheckChannelResponseBodyAudit) GoString() string {
	return s.String()
}

func (s *SetCheckChannelResponseBodyAudit) GetChannel() *string {
	return s.Channel
}

func (s *SetCheckChannelResponseBodyAudit) GetCreateTime() *string {
	return s.CreateTime
}

func (s *SetCheckChannelResponseBodyAudit) GetCustomerId() *string {
	return s.CustomerId
}

func (s *SetCheckChannelResponseBodyAudit) GetLegalSwitch() *string {
	return s.LegalSwitch
}

func (s *SetCheckChannelResponseBodyAudit) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *SetCheckChannelResponseBodyAudit) SetChannel(v string) *SetCheckChannelResponseBodyAudit {
	s.Channel = &v
	return s
}

func (s *SetCheckChannelResponseBodyAudit) SetCreateTime(v string) *SetCheckChannelResponseBodyAudit {
	s.CreateTime = &v
	return s
}

func (s *SetCheckChannelResponseBodyAudit) SetCustomerId(v string) *SetCheckChannelResponseBodyAudit {
	s.CustomerId = &v
	return s
}

func (s *SetCheckChannelResponseBodyAudit) SetLegalSwitch(v string) *SetCheckChannelResponseBodyAudit {
	s.LegalSwitch = &v
	return s
}

func (s *SetCheckChannelResponseBodyAudit) SetUpdateTime(v string) *SetCheckChannelResponseBodyAudit {
	s.UpdateTime = &v
	return s
}

func (s *SetCheckChannelResponseBodyAudit) Validate() error {
	return dara.Validate(s)
}
