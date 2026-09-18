// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAlertRobotRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDigitalEmployeeName(v string) *CreateAlertRobotRequest
	GetDigitalEmployeeName() *string
	SetLang(v string) *CreateAlertRobotRequest
	GetLang() *string
	SetName(v string) *CreateAlertRobotRequest
	GetName() *string
	SetRobotId(v string) *CreateAlertRobotRequest
	GetRobotId() *string
	SetRobotSignKey(v string) *CreateAlertRobotRequest
	GetRobotSignKey() *string
	SetType(v string) *CreateAlertRobotRequest
	GetType() *string
	SetUrl(v string) *CreateAlertRobotRequest
	GetUrl() *string
	SetWorkspace(v string) *CreateAlertRobotRequest
	GetWorkspace() *string
}

type CreateAlertRobotRequest struct {
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
	// The unique ID of the robot.
	//
	// example:
	//
	// test
	RobotId *string `json:"robotId,omitempty" xml:"robotId,omitempty"`
	// The signature key of the robot.
	//
	// example:
	//
	// abc123
	RobotSignKey *string `json:"robotSignKey,omitempty" xml:"robotSignKey,omitempty"`
	// The type of the robot.
	//
	// This parameter is required.
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
	// The workspace name.
	//
	// example:
	//
	// default-cms-1423134313712421-cn-shanghai
	Workspace *string `json:"workspace,omitempty" xml:"workspace,omitempty"`
}

func (s CreateAlertRobotRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateAlertRobotRequest) GoString() string {
	return s.String()
}

func (s *CreateAlertRobotRequest) GetDigitalEmployeeName() *string {
	return s.DigitalEmployeeName
}

func (s *CreateAlertRobotRequest) GetLang() *string {
	return s.Lang
}

func (s *CreateAlertRobotRequest) GetName() *string {
	return s.Name
}

func (s *CreateAlertRobotRequest) GetRobotId() *string {
	return s.RobotId
}

func (s *CreateAlertRobotRequest) GetRobotSignKey() *string {
	return s.RobotSignKey
}

func (s *CreateAlertRobotRequest) GetType() *string {
	return s.Type
}

func (s *CreateAlertRobotRequest) GetUrl() *string {
	return s.Url
}

func (s *CreateAlertRobotRequest) GetWorkspace() *string {
	return s.Workspace
}

func (s *CreateAlertRobotRequest) SetDigitalEmployeeName(v string) *CreateAlertRobotRequest {
	s.DigitalEmployeeName = &v
	return s
}

func (s *CreateAlertRobotRequest) SetLang(v string) *CreateAlertRobotRequest {
	s.Lang = &v
	return s
}

func (s *CreateAlertRobotRequest) SetName(v string) *CreateAlertRobotRequest {
	s.Name = &v
	return s
}

func (s *CreateAlertRobotRequest) SetRobotId(v string) *CreateAlertRobotRequest {
	s.RobotId = &v
	return s
}

func (s *CreateAlertRobotRequest) SetRobotSignKey(v string) *CreateAlertRobotRequest {
	s.RobotSignKey = &v
	return s
}

func (s *CreateAlertRobotRequest) SetType(v string) *CreateAlertRobotRequest {
	s.Type = &v
	return s
}

func (s *CreateAlertRobotRequest) SetUrl(v string) *CreateAlertRobotRequest {
	s.Url = &v
	return s
}

func (s *CreateAlertRobotRequest) SetWorkspace(v string) *CreateAlertRobotRequest {
	s.Workspace = &v
	return s
}

func (s *CreateAlertRobotRequest) Validate() error {
	return dara.Validate(s)
}
