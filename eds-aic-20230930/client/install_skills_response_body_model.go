// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInstallSkillsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *InstallSkillsResponseBody
	GetCode() *string
	SetInstallResults(v []*InstallSkillsResponseBodyInstallResults) *InstallSkillsResponseBody
	GetInstallResults() []*InstallSkillsResponseBodyInstallResults
	SetMessage(v string) *InstallSkillsResponseBody
	GetMessage() *string
	SetRequestId(v string) *InstallSkillsResponseBody
	GetRequestId() *string
}

type InstallSkillsResponseBody struct {
	// The response code. A value of 200 indicates success.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The installation results.
	InstallResults []*InstallSkillsResponseBodyInstallResults `json:"InstallResults,omitempty" xml:"InstallResults,omitempty" type:"Repeated"`
	// The response message.
	//
	// example:
	//
	// Success.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 1CBAFFAB-B697-4049-A9B1-67E1FC****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s InstallSkillsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s InstallSkillsResponseBody) GoString() string {
	return s.String()
}

func (s *InstallSkillsResponseBody) GetCode() *string {
	return s.Code
}

func (s *InstallSkillsResponseBody) GetInstallResults() []*InstallSkillsResponseBodyInstallResults {
	return s.InstallResults
}

func (s *InstallSkillsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *InstallSkillsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *InstallSkillsResponseBody) SetCode(v string) *InstallSkillsResponseBody {
	s.Code = &v
	return s
}

func (s *InstallSkillsResponseBody) SetInstallResults(v []*InstallSkillsResponseBodyInstallResults) *InstallSkillsResponseBody {
	s.InstallResults = v
	return s
}

func (s *InstallSkillsResponseBody) SetMessage(v string) *InstallSkillsResponseBody {
	s.Message = &v
	return s
}

func (s *InstallSkillsResponseBody) SetRequestId(v string) *InstallSkillsResponseBody {
	s.RequestId = &v
	return s
}

func (s *InstallSkillsResponseBody) Validate() error {
	if s.InstallResults != nil {
		for _, item := range s.InstallResults {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type InstallSkillsResponseBodyInstallResults struct {
	// The cloud phone instance ID.
	//
	// example:
	//
	// acp-6rnonvrkf59ac****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The installation status.
	//
	// example:
	//
	// INSTALLING
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s InstallSkillsResponseBodyInstallResults) String() string {
	return dara.Prettify(s)
}

func (s InstallSkillsResponseBodyInstallResults) GoString() string {
	return s.String()
}

func (s *InstallSkillsResponseBodyInstallResults) GetInstanceId() *string {
	return s.InstanceId
}

func (s *InstallSkillsResponseBodyInstallResults) GetStatus() *string {
	return s.Status
}

func (s *InstallSkillsResponseBodyInstallResults) SetInstanceId(v string) *InstallSkillsResponseBodyInstallResults {
	s.InstanceId = &v
	return s
}

func (s *InstallSkillsResponseBodyInstallResults) SetStatus(v string) *InstallSkillsResponseBodyInstallResults {
	s.Status = &v
	return s
}

func (s *InstallSkillsResponseBodyInstallResults) Validate() error {
	return dara.Validate(s)
}
