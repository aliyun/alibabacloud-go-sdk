// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTimedResidentResourcePoolApplication interface {
	dara.Model
	String() string
	GoString() string
	SetAccountId(v string) *TimedResidentResourcePoolApplication
	GetAccountId() *string
	SetContent(v string) *TimedResidentResourcePoolApplication
	GetContent() *string
	SetCreatedTime(v string) *TimedResidentResourcePoolApplication
	GetCreatedTime() *string
	SetLastModifiedTime(v string) *TimedResidentResourcePoolApplication
	GetLastModifiedTime() *string
	SetOperationType(v string) *TimedResidentResourcePoolApplication
	GetOperationType() *string
	SetStatus(v string) *TimedResidentResourcePoolApplication
	GetStatus() *string
	SetTimedPoolId(v string) *TimedResidentResourcePoolApplication
	GetTimedPoolId() *string
}

type TimedResidentResourcePoolApplication struct {
	AccountId        *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
	Content          *string `json:"content,omitempty" xml:"content,omitempty"`
	CreatedTime      *string `json:"createdTime,omitempty" xml:"createdTime,omitempty"`
	LastModifiedTime *string `json:"lastModifiedTime,omitempty" xml:"lastModifiedTime,omitempty"`
	OperationType    *string `json:"operationType,omitempty" xml:"operationType,omitempty"`
	Status           *string `json:"status,omitempty" xml:"status,omitempty"`
	TimedPoolId      *string `json:"timedPoolId,omitempty" xml:"timedPoolId,omitempty"`
}

func (s TimedResidentResourcePoolApplication) String() string {
	return dara.Prettify(s)
}

func (s TimedResidentResourcePoolApplication) GoString() string {
	return s.String()
}

func (s *TimedResidentResourcePoolApplication) GetAccountId() *string {
	return s.AccountId
}

func (s *TimedResidentResourcePoolApplication) GetContent() *string {
	return s.Content
}

func (s *TimedResidentResourcePoolApplication) GetCreatedTime() *string {
	return s.CreatedTime
}

func (s *TimedResidentResourcePoolApplication) GetLastModifiedTime() *string {
	return s.LastModifiedTime
}

func (s *TimedResidentResourcePoolApplication) GetOperationType() *string {
	return s.OperationType
}

func (s *TimedResidentResourcePoolApplication) GetStatus() *string {
	return s.Status
}

func (s *TimedResidentResourcePoolApplication) GetTimedPoolId() *string {
	return s.TimedPoolId
}

func (s *TimedResidentResourcePoolApplication) SetAccountId(v string) *TimedResidentResourcePoolApplication {
	s.AccountId = &v
	return s
}

func (s *TimedResidentResourcePoolApplication) SetContent(v string) *TimedResidentResourcePoolApplication {
	s.Content = &v
	return s
}

func (s *TimedResidentResourcePoolApplication) SetCreatedTime(v string) *TimedResidentResourcePoolApplication {
	s.CreatedTime = &v
	return s
}

func (s *TimedResidentResourcePoolApplication) SetLastModifiedTime(v string) *TimedResidentResourcePoolApplication {
	s.LastModifiedTime = &v
	return s
}

func (s *TimedResidentResourcePoolApplication) SetOperationType(v string) *TimedResidentResourcePoolApplication {
	s.OperationType = &v
	return s
}

func (s *TimedResidentResourcePoolApplication) SetStatus(v string) *TimedResidentResourcePoolApplication {
	s.Status = &v
	return s
}

func (s *TimedResidentResourcePoolApplication) SetTimedPoolId(v string) *TimedResidentResourcePoolApplication {
	s.TimedPoolId = &v
	return s
}

func (s *TimedResidentResourcePoolApplication) Validate() error {
	return dara.Validate(s)
}
