// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMonitorContactUpdateCmd interface {
	dara.Model
	String() string
	GoString() string
	SetEmail(v string) *MonitorContactUpdateCmd
	GetEmail() *string
	SetId(v int64) *MonitorContactUpdateCmd
	GetId() *int64
	SetName(v string) *MonitorContactUpdateCmd
	GetName() *string
	SetPhone(v string) *MonitorContactUpdateCmd
	GetPhone() *string
}

type MonitorContactUpdateCmd struct {
	// This parameter is required.
	//
	// example:
	//
	// xxx@alibaba.com
	Email *string `json:"email,omitempty" xml:"email,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 尚仁
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 15113456789
	Phone *string `json:"phone,omitempty" xml:"phone,omitempty"`
}

func (s MonitorContactUpdateCmd) String() string {
	return dara.Prettify(s)
}

func (s MonitorContactUpdateCmd) GoString() string {
	return s.String()
}

func (s *MonitorContactUpdateCmd) GetEmail() *string {
	return s.Email
}

func (s *MonitorContactUpdateCmd) GetId() *int64 {
	return s.Id
}

func (s *MonitorContactUpdateCmd) GetName() *string {
	return s.Name
}

func (s *MonitorContactUpdateCmd) GetPhone() *string {
	return s.Phone
}

func (s *MonitorContactUpdateCmd) SetEmail(v string) *MonitorContactUpdateCmd {
	s.Email = &v
	return s
}

func (s *MonitorContactUpdateCmd) SetId(v int64) *MonitorContactUpdateCmd {
	s.Id = &v
	return s
}

func (s *MonitorContactUpdateCmd) SetName(v string) *MonitorContactUpdateCmd {
	s.Name = &v
	return s
}

func (s *MonitorContactUpdateCmd) SetPhone(v string) *MonitorContactUpdateCmd {
	s.Phone = &v
	return s
}

func (s *MonitorContactUpdateCmd) Validate() error {
	return dara.Validate(s)
}
