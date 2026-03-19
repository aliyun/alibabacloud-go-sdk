// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartRestoreTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *StartRestoreTaskRequest
	GetClientToken() *string
	SetOwnerId(v string) *StartRestoreTaskRequest
	GetOwnerId() *string
	SetRestoreTaskId(v string) *StartRestoreTaskRequest
	GetRestoreTaskId() *string
}

type StartRestoreTaskRequest struct {
	// A client token. It ensures the idempotence of the request and prevents the same request from being submitted multiple times.
	//
	// example:
	//
	// ETnLKlblzczshOTUbOCzxxxxxxx
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	OwnerId     *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The ID of the restore job.
	//
	// This parameter is required.
	//
	// example:
	//
	// s102h7rfXXXX
	RestoreTaskId *string `json:"RestoreTaskId,omitempty" xml:"RestoreTaskId,omitempty"`
}

func (s StartRestoreTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s StartRestoreTaskRequest) GoString() string {
	return s.String()
}

func (s *StartRestoreTaskRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *StartRestoreTaskRequest) GetOwnerId() *string {
	return s.OwnerId
}

func (s *StartRestoreTaskRequest) GetRestoreTaskId() *string {
	return s.RestoreTaskId
}

func (s *StartRestoreTaskRequest) SetClientToken(v string) *StartRestoreTaskRequest {
	s.ClientToken = &v
	return s
}

func (s *StartRestoreTaskRequest) SetOwnerId(v string) *StartRestoreTaskRequest {
	s.OwnerId = &v
	return s
}

func (s *StartRestoreTaskRequest) SetRestoreTaskId(v string) *StartRestoreTaskRequest {
	s.RestoreTaskId = &v
	return s
}

func (s *StartRestoreTaskRequest) Validate() error {
	return dara.Validate(s)
}
