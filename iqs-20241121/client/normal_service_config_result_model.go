// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iNormalServiceConfigResult interface {
	dara.Model
	String() string
	GoString() string
	SetAvailableTime(v string) *NormalServiceConfigResult
	GetAvailableTime() *string
	SetGmtCreate(v string) *NormalServiceConfigResult
	GetGmtCreate() *string
	SetStatus(v string) *NormalServiceConfigResult
	GetStatus() *string
	SetUserAPICountInfos(v []*UserAPICountInfo) *NormalServiceConfigResult
	GetUserAPICountInfos() []*UserAPICountInfo
}

type NormalServiceConfigResult struct {
	AvailableTime     *string             `json:"availableTime,omitempty" xml:"availableTime,omitempty"`
	GmtCreate         *string             `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	Status            *string             `json:"status,omitempty" xml:"status,omitempty"`
	UserAPICountInfos []*UserAPICountInfo `json:"userAPICountInfos,omitempty" xml:"userAPICountInfos,omitempty" type:"Repeated"`
}

func (s NormalServiceConfigResult) String() string {
	return dara.Prettify(s)
}

func (s NormalServiceConfigResult) GoString() string {
	return s.String()
}

func (s *NormalServiceConfigResult) GetAvailableTime() *string {
	return s.AvailableTime
}

func (s *NormalServiceConfigResult) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *NormalServiceConfigResult) GetStatus() *string {
	return s.Status
}

func (s *NormalServiceConfigResult) GetUserAPICountInfos() []*UserAPICountInfo {
	return s.UserAPICountInfos
}

func (s *NormalServiceConfigResult) SetAvailableTime(v string) *NormalServiceConfigResult {
	s.AvailableTime = &v
	return s
}

func (s *NormalServiceConfigResult) SetGmtCreate(v string) *NormalServiceConfigResult {
	s.GmtCreate = &v
	return s
}

func (s *NormalServiceConfigResult) SetStatus(v string) *NormalServiceConfigResult {
	s.Status = &v
	return s
}

func (s *NormalServiceConfigResult) SetUserAPICountInfos(v []*UserAPICountInfo) *NormalServiceConfigResult {
	s.UserAPICountInfos = v
	return s
}

func (s *NormalServiceConfigResult) Validate() error {
	if s.UserAPICountInfos != nil {
		for _, item := range s.UserAPICountInfos {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
