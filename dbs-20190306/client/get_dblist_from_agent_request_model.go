// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDBListFromAgentRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBackupGatewayId(v int64) *GetDBListFromAgentRequest
	GetBackupGatewayId() *int64
	SetClientToken(v string) *GetDBListFromAgentRequest
	GetClientToken() *string
	SetOwnerId(v string) *GetDBListFromAgentRequest
	GetOwnerId() *string
	SetSourceEndpointRegion(v string) *GetDBListFromAgentRequest
	GetSourceEndpointRegion() *string
	SetTaskId(v int64) *GetDBListFromAgentRequest
	GetTaskId() *int64
}

type GetDBListFromAgentRequest struct {
	// The ID of the backup gateway.
	//
	// example:
	//
	// 160813
	BackupGatewayId *int64 `json:"BackupGatewayId,omitempty" xml:"BackupGatewayId,omitempty"`
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the value, but you must make sure that it is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length.
	//
	// example:
	//
	// ETnLKlblzczshOTUbOCzxxxxxxxxxx
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	OwnerId     *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region in which the backup gateway resides.
	//
	// example:
	//
	// cn-hangzhou
	SourceEndpointRegion *string `json:"SourceEndpointRegion,omitempty" xml:"SourceEndpointRegion,omitempty"`
	// The ID of the asynchronous task.
	//
	// example:
	//
	// 123456
	TaskId *int64 `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s GetDBListFromAgentRequest) String() string {
	return dara.Prettify(s)
}

func (s GetDBListFromAgentRequest) GoString() string {
	return s.String()
}

func (s *GetDBListFromAgentRequest) GetBackupGatewayId() *int64 {
	return s.BackupGatewayId
}

func (s *GetDBListFromAgentRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *GetDBListFromAgentRequest) GetOwnerId() *string {
	return s.OwnerId
}

func (s *GetDBListFromAgentRequest) GetSourceEndpointRegion() *string {
	return s.SourceEndpointRegion
}

func (s *GetDBListFromAgentRequest) GetTaskId() *int64 {
	return s.TaskId
}

func (s *GetDBListFromAgentRequest) SetBackupGatewayId(v int64) *GetDBListFromAgentRequest {
	s.BackupGatewayId = &v
	return s
}

func (s *GetDBListFromAgentRequest) SetClientToken(v string) *GetDBListFromAgentRequest {
	s.ClientToken = &v
	return s
}

func (s *GetDBListFromAgentRequest) SetOwnerId(v string) *GetDBListFromAgentRequest {
	s.OwnerId = &v
	return s
}

func (s *GetDBListFromAgentRequest) SetSourceEndpointRegion(v string) *GetDBListFromAgentRequest {
	s.SourceEndpointRegion = &v
	return s
}

func (s *GetDBListFromAgentRequest) SetTaskId(v int64) *GetDBListFromAgentRequest {
	s.TaskId = &v
	return s
}

func (s *GetDBListFromAgentRequest) Validate() error {
	return dara.Validate(s)
}
