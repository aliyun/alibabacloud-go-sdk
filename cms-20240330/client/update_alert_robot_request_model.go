// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAlertRobotRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDigitalEmployeeName(v string) *UpdateAlertRobotRequest
	GetDigitalEmployeeName() *string
	SetLang(v string) *UpdateAlertRobotRequest
	GetLang() *string
	SetName(v string) *UpdateAlertRobotRequest
	GetName() *string
	SetRobotSignKey(v string) *UpdateAlertRobotRequest
	GetRobotSignKey() *string
	SetType(v string) *UpdateAlertRobotRequest
	GetType() *string
	SetUrl(v string) *UpdateAlertRobotRequest
	GetUrl() *string
}

type UpdateAlertRobotRequest struct {
	// The name of the digital employee.
	//
	// example:
	//
	// apsara-ops
	DigitalEmployeeName *string `json:"digitalEmployeeName,omitempty" xml:"digitalEmployeeName,omitempty"`
	// The language.
	//
	// example:
	//
	// zh_CN
	Lang *string `json:"lang,omitempty" xml:"lang,omitempty"`
	// The name of the robot.
	//
	// example:
	//
	// test
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The signature key of the robot.
	//
	// example:
	//
	// abc123
	RobotSignKey *string `json:"robotSignKey,omitempty" xml:"robotSignKey,omitempty"`
	// The type of the robot.
	//
	// example:
	//
	// DING
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// The webhook URL of the robot.
	//
	// example:
	//
	// https://oapi.dingtalk.com/robot/send?access_token=**************
	Url *string `json:"url,omitempty" xml:"url,omitempty"`
}

func (s UpdateAlertRobotRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateAlertRobotRequest) GoString() string {
	return s.String()
}

func (s *UpdateAlertRobotRequest) GetDigitalEmployeeName() *string {
	return s.DigitalEmployeeName
}

func (s *UpdateAlertRobotRequest) GetLang() *string {
	return s.Lang
}

func (s *UpdateAlertRobotRequest) GetName() *string {
	return s.Name
}

func (s *UpdateAlertRobotRequest) GetRobotSignKey() *string {
	return s.RobotSignKey
}

func (s *UpdateAlertRobotRequest) GetType() *string {
	return s.Type
}

func (s *UpdateAlertRobotRequest) GetUrl() *string {
	return s.Url
}

func (s *UpdateAlertRobotRequest) SetDigitalEmployeeName(v string) *UpdateAlertRobotRequest {
	s.DigitalEmployeeName = &v
	return s
}

func (s *UpdateAlertRobotRequest) SetLang(v string) *UpdateAlertRobotRequest {
	s.Lang = &v
	return s
}

func (s *UpdateAlertRobotRequest) SetName(v string) *UpdateAlertRobotRequest {
	s.Name = &v
	return s
}

func (s *UpdateAlertRobotRequest) SetRobotSignKey(v string) *UpdateAlertRobotRequest {
	s.RobotSignKey = &v
	return s
}

func (s *UpdateAlertRobotRequest) SetType(v string) *UpdateAlertRobotRequest {
	s.Type = &v
	return s
}

func (s *UpdateAlertRobotRequest) SetUrl(v string) *UpdateAlertRobotRequest {
	s.Url = &v
	return s
}

func (s *UpdateAlertRobotRequest) Validate() error {
	return dara.Validate(s)
}
