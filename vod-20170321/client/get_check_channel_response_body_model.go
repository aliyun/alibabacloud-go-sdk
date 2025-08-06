// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCheckChannelResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAudit(v *GetCheckChannelResponseBodyAudit) *GetCheckChannelResponseBody
	GetAudit() *GetCheckChannelResponseBodyAudit
	SetRequestId(v string) *GetCheckChannelResponseBody
	GetRequestId() *string
}

type GetCheckChannelResponseBody struct {
	Audit     *GetCheckChannelResponseBodyAudit `json:"Audit,omitempty" xml:"Audit,omitempty" type:"Struct"`
	RequestId *string                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetCheckChannelResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetCheckChannelResponseBody) GoString() string {
	return s.String()
}

func (s *GetCheckChannelResponseBody) GetAudit() *GetCheckChannelResponseBodyAudit {
	return s.Audit
}

func (s *GetCheckChannelResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetCheckChannelResponseBody) SetAudit(v *GetCheckChannelResponseBodyAudit) *GetCheckChannelResponseBody {
	s.Audit = v
	return s
}

func (s *GetCheckChannelResponseBody) SetRequestId(v string) *GetCheckChannelResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetCheckChannelResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetCheckChannelResponseBodyAudit struct {
	Channel     *string `json:"Channel,omitempty" xml:"Channel,omitempty"`
	CreateTime  *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	CustomerId  *string `json:"CustomerId,omitempty" xml:"CustomerId,omitempty"`
	LegalSwitch *string `json:"LegalSwitch,omitempty" xml:"LegalSwitch,omitempty"`
	UpdateTime  *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s GetCheckChannelResponseBodyAudit) String() string {
	return dara.Prettify(s)
}

func (s GetCheckChannelResponseBodyAudit) GoString() string {
	return s.String()
}

func (s *GetCheckChannelResponseBodyAudit) GetChannel() *string {
	return s.Channel
}

func (s *GetCheckChannelResponseBodyAudit) GetCreateTime() *string {
	return s.CreateTime
}

func (s *GetCheckChannelResponseBodyAudit) GetCustomerId() *string {
	return s.CustomerId
}

func (s *GetCheckChannelResponseBodyAudit) GetLegalSwitch() *string {
	return s.LegalSwitch
}

func (s *GetCheckChannelResponseBodyAudit) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *GetCheckChannelResponseBodyAudit) SetChannel(v string) *GetCheckChannelResponseBodyAudit {
	s.Channel = &v
	return s
}

func (s *GetCheckChannelResponseBodyAudit) SetCreateTime(v string) *GetCheckChannelResponseBodyAudit {
	s.CreateTime = &v
	return s
}

func (s *GetCheckChannelResponseBodyAudit) SetCustomerId(v string) *GetCheckChannelResponseBodyAudit {
	s.CustomerId = &v
	return s
}

func (s *GetCheckChannelResponseBodyAudit) SetLegalSwitch(v string) *GetCheckChannelResponseBodyAudit {
	s.LegalSwitch = &v
	return s
}

func (s *GetCheckChannelResponseBodyAudit) SetUpdateTime(v string) *GetCheckChannelResponseBodyAudit {
	s.UpdateTime = &v
	return s
}

func (s *GetCheckChannelResponseBodyAudit) Validate() error {
	return dara.Validate(s)
}
