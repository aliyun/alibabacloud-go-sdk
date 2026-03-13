// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iResourceOperation interface {
	dara.Model
	String() string
	GoString() string
	SetCreatorId(v string) *ResourceOperation
	GetCreatorId() *string
	SetGmtCreatedTime(v string) *ResourceOperation
	GetGmtCreatedTime() *string
	SetGmtEndTime(v string) *ResourceOperation
	GetGmtEndTime() *string
	SetGmtModifiedTime(v string) *ResourceOperation
	GetGmtModifiedTime() *string
	SetGmtStartTime(v string) *ResourceOperation
	GetGmtStartTime() *string
	SetObjectId(v string) *ResourceOperation
	GetObjectId() *string
	SetObjectType(v string) *ResourceOperation
	GetObjectType() *string
	SetOperationDescription(v string) *ResourceOperation
	GetOperationDescription() *string
	SetOperationId(v string) *ResourceOperation
	GetOperationId() *string
	SetOperationSpecJson(v string) *ResourceOperation
	GetOperationSpecJson() *string
	SetOperationType(v string) *ResourceOperation
	GetOperationType() *string
	SetReasonCode(v string) *ResourceOperation
	GetReasonCode() *string
	SetReasonMessage(v string) *ResourceOperation
	GetReasonMessage() *string
	SetStatus(v string) *ResourceOperation
	GetStatus() *string
}

type ResourceOperation struct {
	CreatorId            *string `json:"CreatorId,omitempty" xml:"CreatorId,omitempty"`
	GmtCreatedTime       *string `json:"GmtCreatedTime,omitempty" xml:"GmtCreatedTime,omitempty"`
	GmtEndTime           *string `json:"GmtEndTime,omitempty" xml:"GmtEndTime,omitempty"`
	GmtModifiedTime      *string `json:"GmtModifiedTime,omitempty" xml:"GmtModifiedTime,omitempty"`
	GmtStartTime         *string `json:"GmtStartTime,omitempty" xml:"GmtStartTime,omitempty"`
	ObjectId             *string `json:"ObjectId,omitempty" xml:"ObjectId,omitempty"`
	ObjectType           *string `json:"ObjectType,omitempty" xml:"ObjectType,omitempty"`
	OperationDescription *string `json:"OperationDescription,omitempty" xml:"OperationDescription,omitempty"`
	OperationId          *string `json:"OperationId,omitempty" xml:"OperationId,omitempty"`
	OperationSpecJson    *string `json:"OperationSpecJson,omitempty" xml:"OperationSpecJson,omitempty"`
	OperationType        *string `json:"OperationType,omitempty" xml:"OperationType,omitempty"`
	ReasonCode           *string `json:"ReasonCode,omitempty" xml:"ReasonCode,omitempty"`
	ReasonMessage        *string `json:"ReasonMessage,omitempty" xml:"ReasonMessage,omitempty"`
	Status               *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ResourceOperation) String() string {
	return dara.Prettify(s)
}

func (s ResourceOperation) GoString() string {
	return s.String()
}

func (s *ResourceOperation) GetCreatorId() *string {
	return s.CreatorId
}

func (s *ResourceOperation) GetGmtCreatedTime() *string {
	return s.GmtCreatedTime
}

func (s *ResourceOperation) GetGmtEndTime() *string {
	return s.GmtEndTime
}

func (s *ResourceOperation) GetGmtModifiedTime() *string {
	return s.GmtModifiedTime
}

func (s *ResourceOperation) GetGmtStartTime() *string {
	return s.GmtStartTime
}

func (s *ResourceOperation) GetObjectId() *string {
	return s.ObjectId
}

func (s *ResourceOperation) GetObjectType() *string {
	return s.ObjectType
}

func (s *ResourceOperation) GetOperationDescription() *string {
	return s.OperationDescription
}

func (s *ResourceOperation) GetOperationId() *string {
	return s.OperationId
}

func (s *ResourceOperation) GetOperationSpecJson() *string {
	return s.OperationSpecJson
}

func (s *ResourceOperation) GetOperationType() *string {
	return s.OperationType
}

func (s *ResourceOperation) GetReasonCode() *string {
	return s.ReasonCode
}

func (s *ResourceOperation) GetReasonMessage() *string {
	return s.ReasonMessage
}

func (s *ResourceOperation) GetStatus() *string {
	return s.Status
}

func (s *ResourceOperation) SetCreatorId(v string) *ResourceOperation {
	s.CreatorId = &v
	return s
}

func (s *ResourceOperation) SetGmtCreatedTime(v string) *ResourceOperation {
	s.GmtCreatedTime = &v
	return s
}

func (s *ResourceOperation) SetGmtEndTime(v string) *ResourceOperation {
	s.GmtEndTime = &v
	return s
}

func (s *ResourceOperation) SetGmtModifiedTime(v string) *ResourceOperation {
	s.GmtModifiedTime = &v
	return s
}

func (s *ResourceOperation) SetGmtStartTime(v string) *ResourceOperation {
	s.GmtStartTime = &v
	return s
}

func (s *ResourceOperation) SetObjectId(v string) *ResourceOperation {
	s.ObjectId = &v
	return s
}

func (s *ResourceOperation) SetObjectType(v string) *ResourceOperation {
	s.ObjectType = &v
	return s
}

func (s *ResourceOperation) SetOperationDescription(v string) *ResourceOperation {
	s.OperationDescription = &v
	return s
}

func (s *ResourceOperation) SetOperationId(v string) *ResourceOperation {
	s.OperationId = &v
	return s
}

func (s *ResourceOperation) SetOperationSpecJson(v string) *ResourceOperation {
	s.OperationSpecJson = &v
	return s
}

func (s *ResourceOperation) SetOperationType(v string) *ResourceOperation {
	s.OperationType = &v
	return s
}

func (s *ResourceOperation) SetReasonCode(v string) *ResourceOperation {
	s.ReasonCode = &v
	return s
}

func (s *ResourceOperation) SetReasonMessage(v string) *ResourceOperation {
	s.ReasonMessage = &v
	return s
}

func (s *ResourceOperation) SetStatus(v string) *ResourceOperation {
	s.Status = &v
	return s
}

func (s *ResourceOperation) Validate() error {
	return dara.Validate(s)
}
