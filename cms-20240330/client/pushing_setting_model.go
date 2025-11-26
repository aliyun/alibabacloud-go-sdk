// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPushingSetting interface {
	dara.Model
	String() string
	GoString() string
	SetAlertActionIds(v []*string) *PushingSetting
	GetAlertActionIds() []*string
	SetRestoreActionIds(v []*string) *PushingSetting
	GetRestoreActionIds() []*string
	SetTemplateUuid(v string) *PushingSetting
	GetTemplateUuid() *string
}

type PushingSetting struct {
	AlertActionIds   []*string `json:"alertActionIds,omitempty" xml:"alertActionIds,omitempty" type:"Repeated"`
	RestoreActionIds []*string `json:"restoreActionIds,omitempty" xml:"restoreActionIds,omitempty" type:"Repeated"`
	TemplateUuid     *string   `json:"templateUuid,omitempty" xml:"templateUuid,omitempty"`
}

func (s PushingSetting) String() string {
	return dara.Prettify(s)
}

func (s PushingSetting) GoString() string {
	return s.String()
}

func (s *PushingSetting) GetAlertActionIds() []*string {
	return s.AlertActionIds
}

func (s *PushingSetting) GetRestoreActionIds() []*string {
	return s.RestoreActionIds
}

func (s *PushingSetting) GetTemplateUuid() *string {
	return s.TemplateUuid
}

func (s *PushingSetting) SetAlertActionIds(v []*string) *PushingSetting {
	s.AlertActionIds = v
	return s
}

func (s *PushingSetting) SetRestoreActionIds(v []*string) *PushingSetting {
	s.RestoreActionIds = v
	return s
}

func (s *PushingSetting) SetTemplateUuid(v string) *PushingSetting {
	s.TemplateUuid = &v
	return s
}

func (s *PushingSetting) Validate() error {
	return dara.Validate(s)
}
