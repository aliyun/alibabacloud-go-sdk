// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iWebInstanceInfo interface {
	dara.Model
	String() string
	GoString() string
	SetImageUrl(v string) *WebInstanceInfo
	GetImageUrl() *string
	SetInstanceId(v string) *WebInstanceInfo
	GetInstanceId() *string
	SetStatus(v string) *WebInstanceInfo
	GetStatus() *string
	SetVersionId(v string) *WebInstanceInfo
	GetVersionId() *string
}

type WebInstanceInfo struct {
	ImageUrl *string `json:"ImageUrl,omitempty" xml:"ImageUrl,omitempty"`
	// example:
	//
	// c-66691780-1522405d-3021e147e0c3
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Status     *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// 1
	VersionId *string `json:"VersionId,omitempty" xml:"VersionId,omitempty"`
}

func (s WebInstanceInfo) String() string {
	return dara.Prettify(s)
}

func (s WebInstanceInfo) GoString() string {
	return s.String()
}

func (s *WebInstanceInfo) GetImageUrl() *string {
	return s.ImageUrl
}

func (s *WebInstanceInfo) GetInstanceId() *string {
	return s.InstanceId
}

func (s *WebInstanceInfo) GetStatus() *string {
	return s.Status
}

func (s *WebInstanceInfo) GetVersionId() *string {
	return s.VersionId
}

func (s *WebInstanceInfo) SetImageUrl(v string) *WebInstanceInfo {
	s.ImageUrl = &v
	return s
}

func (s *WebInstanceInfo) SetInstanceId(v string) *WebInstanceInfo {
	s.InstanceId = &v
	return s
}

func (s *WebInstanceInfo) SetStatus(v string) *WebInstanceInfo {
	s.Status = &v
	return s
}

func (s *WebInstanceInfo) SetVersionId(v string) *WebInstanceInfo {
	s.VersionId = &v
	return s
}

func (s *WebInstanceInfo) Validate() error {
	return dara.Validate(s)
}
