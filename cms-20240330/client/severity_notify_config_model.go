// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSeverityNotifyConfig interface {
	dara.Model
	String() string
	GoString() string
	SetReceivers(v []*DirectNotifyReceiver) *SeverityNotifyConfig
	GetReceivers() []*DirectNotifyReceiver
	SetSendRecoverNotification(v bool) *SeverityNotifyConfig
	GetSendRecoverNotification() *bool
}

type SeverityNotifyConfig struct {
	Receivers               []*DirectNotifyReceiver `json:"receivers,omitempty" xml:"receivers,omitempty" type:"Repeated"`
	SendRecoverNotification *bool                   `json:"sendRecoverNotification,omitempty" xml:"sendRecoverNotification,omitempty"`
}

func (s SeverityNotifyConfig) String() string {
	return dara.Prettify(s)
}

func (s SeverityNotifyConfig) GoString() string {
	return s.String()
}

func (s *SeverityNotifyConfig) GetReceivers() []*DirectNotifyReceiver {
	return s.Receivers
}

func (s *SeverityNotifyConfig) GetSendRecoverNotification() *bool {
	return s.SendRecoverNotification
}

func (s *SeverityNotifyConfig) SetReceivers(v []*DirectNotifyReceiver) *SeverityNotifyConfig {
	s.Receivers = v
	return s
}

func (s *SeverityNotifyConfig) SetSendRecoverNotification(v bool) *SeverityNotifyConfig {
	s.SendRecoverNotification = &v
	return s
}

func (s *SeverityNotifyConfig) Validate() error {
	if s.Receivers != nil {
		for _, item := range s.Receivers {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
